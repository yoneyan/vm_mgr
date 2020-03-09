package vm

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/manage"
	"strconv"
)

func StartVMProcess(id int) bool {
	fmt.Println("-----StartVMProcess-----")
	if manage.VMExistsID(id) == false {
		fmt.Println("VM Not Found!!")
		return false
	}
	data, err := db.VMDBGetData(id)
	if err != nil {
		fmt.Println("VM Data Not Found!!")
		return false
	}
	status, err := db.VMDBGetVMStatus(id)
	if status == 1 {
		fmt.Println("VM is power on!!")
		return false
	} else if status > 1 || status < 0 {
		fmt.Println("VM status is error!! status: " + strconv.Itoa(status))
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
