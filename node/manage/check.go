package manage

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/etc"
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

func FileExistsCheck(filename string) bool {
	if f, err := os.Stat(filename); os.IsNotExist(err) || f.IsDir() {
		return false
	} else {
		return true
	}
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
	count3 := 0
	socket := etc.SocketPath(name)
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
				fmt.Println("DB not found!!")
				count1++
			}
			//stop status 1
			if status == 0 {
				count1++
				fmt.Println("-->vm stop process")
				fmt.Println("VMID: " + strconv.Itoa(data.ID) + " Count: " + strconv.Itoa(count1))
			}
			if count1 == 1 {
				fmt.Println("VMID: " + strconv.Itoa(data.ID) + " stop vm process!!")
				return true
			}
			//_, err = pipeline.CombinedOutput(
			//	[]string{"ps", "axf"},
			//	[]string{"grep", "/" + name + ".sock"},
			//)
			//stop status 2
			if FileExistsCheck(socket) {
				count2 = 0
			} else {
				count2++
				fmt.Println("-->sock file not found")
				fmt.Println("VMID: " + strconv.Itoa(data.ID) + " Count: " + strconv.Itoa(count2))
				if count2 == 2 {
					db.VMDBStatusUpdate(data.ID, 0)
					fmt.Println("VMID: " + strconv.Itoa(data.ID) + " sock file error!!")
					return true
				}
			}
			_, err = pipeline.Output(
				[]string{"echo", "info status"},
				[]string{"sudo", "socat", "-", socket},
			)
			if err != nil {
				fmt.Println("already stopped")
				count3++
				fmt.Println("-->sock file not open")
				fmt.Println("VMID: " + strconv.Itoa(data.ID) + " Count: " + strconv.Itoa(count3))
				if count3 == 2 {
					db.VMDBStatusUpdate(data.ID, 0)
					fmt.Println("VMID: " + strconv.Itoa(data.ID) + " sock file not open!!")
					return true
				}
			} else {
				count3 = 0
			}
		}
	}
	return true
}
