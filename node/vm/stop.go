package vm

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/manage"
	"log"
)

func VMStop(id int) error {
	if manage.VMExistsID(id) == false {
		fmt.Println("VMID Error")
		return nil
	}

	result, err := db.VMDBGetData(id)
	if result.Status == 1 {
		fmt.Println("Power On State")
	} else if result.Status == 0 {
		fmt.Println("Power Off State")
	} else {
		fmt.Println("Power State Error")
		return nil
	}

	fmt.Println(result)
	if err != nil {
		log.Fatalf("Error!!")
	}

	fmt.Println(result.Name)

	//ps axf | grep test|grep qemu  | grep -v grep | awk '{print "kill -9 " $1}' | sudo sh
	fmt.Println("-----VMStop Command-----")
	out, err := pipeline.CombinedOutput(
		[]string{"ps", "axf"},
		[]string{"grep", result.Name},
		[]string{"grep", "qemu"},
		[]string{"grep", "-v", "grep"},
		[]string{"awk", "{print \"kill -9 \" $1}"},
		[]string{"sudo", "sh"},
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(out))

	if db.VMDBStatusUpdate(id, 0) {
		fmt.Println("Power Off state")
	} else {
		fmt.Println("state Error!!")
	}
	return nil
}
