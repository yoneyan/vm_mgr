package vm

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/manage"
)

func StartVMProcess(id int) bool {
	if manage.VMExistsID(id) == false {
		fmt.Println("VM Not Found!!")
		return false
	}
	data, err := db.VMDBGetData(id)
	if err != nil {
		fmt.Println("VM Data Not Found!!")
		return false
	}
	var c CreateVMInformation
	c.Name = data.Name
	c.CPU = data.CPU
	c.Mem = data.Mem
	c.Net = data.Net
	c.VNC = data.Vnc
	c.StoragePath = data.StoragePath

	cmd := CreateGenerateCmd(&c)

	err = RunQEMUCmd(cmd)
	if err != nil {
		fmt.Println(err)

		return false
	}

	fmt.Println("Start End")
	result := db.VMDBStatusUpdate(id, 1)
	if result {
		return true
	} else {
		return false
	}

}
