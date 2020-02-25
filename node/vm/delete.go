package vm

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/manage"
)

func DeleteVMProcess(id int) bool {
	result := manage.VMExistsID(id)
	if result == false {
		fmt.Println("VMID Not Found!!")
		return false
	}
	err := VMStop(id)
	if err != nil {
		fmt.Println("Already stopped!!")
	} else {
		fmt.Println("Stop process end!!")
	}

	db.DeleteDBVM(id)
	return true
}
