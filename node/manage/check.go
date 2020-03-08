package manage

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/yoneyan/vm_mgr/node/db"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func VMExistsName(name string) bool {
	_, err := db.VMDBGetVMID(name)
	if err != nil {
		return false
	}
	return true
}

func VMExistsID(id int) bool {
	_, err := db.VMDBGetData(id)
	if err != nil {
		return false
	}
	return true
}

func VMExistsCheck(name string, id int) bool {
	if VMExistsID(id) || VMExistsName(name) {
		return true
	} else {
		return false
	}
}

func VMVncExistsCheck(vnc int) bool {
	result := db.GetDBAll()
	for a, _ := range result {
		if result[a].Vnc == vnc {
			return true
		}
	}
	return false
}

func VMIDCheck(id int) bool {
	if id < 0 {
		fmt.Println("VMID Check NG.")
		return false
	} else {
		fmt.Println("VMID Check OK.")
		return true
	}
}

func VMLifeCheck(name string) bool {
	id, err := db.VMDBGetVMID(name)
	if err != nil {
		fmt.Println("DB Read Error")
		return false
	}
	data, err := db.VMDBGetData(id)
	if err != nil {
		fmt.Println("DB Read Error")
		return false
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	count1 := 0
	count2 := 0
loop:
	for {
		select {
		case <-sc:
			fmt.Println("interrupt")
			break loop
		case <-time.After(10 * time.Second):
			fmt.Println("-----VMLIFECheck-----")
			fmt.Println("VMID: " + strconv.Itoa(data.ID))
			status, err := db.VMDBGetVMStatus(id)
			if err != nil {
				fmt.Println(false)
			}
			//stop status 1
			if status == 0 {
				count1++
				fmt.Println("-->vm stop process")
				fmt.Println("VMID: " + strconv.Itoa(data.ID) + " Count: " + strconv.Itoa(data.Status))
				if count1 == 2 {
					fmt.Println("VMID: " + strconv.Itoa(data.ID) + "stop vm process")
					return true
				}
			}

			_, err = pipeline.CombinedOutput(
				[]string{"ps", "axf"},
				[]string{"grep", "/" + name + ".sock"},
			)
			//stop status 2
			if err != nil {
				count2++
				fmt.Println("-->grep error")
				fmt.Println("VMID: " + strconv.Itoa(data.ID) + " Count: " + strconv.Itoa(data.Status))
				if count2 == 2 {
					db.VMDBStatusUpdate(data.ID, 0)
					fmt.Println("VMID: " + strconv.Itoa(data.ID) + "grep error!!")
					return true
				}
			}

			count1 = 0
			count2 = 0
		}
	}
	return true
}
