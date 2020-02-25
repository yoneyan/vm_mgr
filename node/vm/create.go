package vm

import "C"
import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/etc"
	"github.com/yoneyan/vm_mgr/node/manage"
	"log"
)

type CreateVMInformation struct {
	ID          int
	Name        string
	CPU         int
	Mem         int
	StoragePath string
	CDROM       string
	Net         string
	VNC         int
	AutoStart   bool
}

func CreateVMProcess(c *CreateVMInformation) error {
	fmt.Println("----VMNewCreate")

	if manage.VMVncExistsCheck(c.VNC) {
		fmt.Println("A VM with the same vnc port exists. So, change vnc port of the VM.")
	} else {
		if manage.VMExistsName(c.Name) {
			fmt.Println("A VM with the same name exists. So, change the name of the VM.")
		} else {
			CreateVMDBProcess(c)
			err := RunQEMUCmd(CreateGenerateCmd(c))
			if err != nil {
				log.Fatal(err)
				return fmt.Errorf("VMNewCreate Error!!")
			} else {
				db.VMDBStatusUpdate(c.ID, 1)
			}
			return nil
		}
	}

	return nil
}

func CreateVMDBProcess(c *CreateVMInformation) {
	dbdata := db.NodeVM{
		Name:        c.Name,
		CPU:         c.CPU,
		Mem:         c.Mem,
		StoragePath: c.StoragePath,
		Net:         c.Net,
		Vnc:         c.VNC,
		Socket:      etc.SocketGenerate(c.Name),
		Status:      0,
		AutoStart:   c.AutoStart,
	}
	db.AddDBVM(dbdata)
}
