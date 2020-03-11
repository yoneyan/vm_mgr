package vm

import "C"
import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/etc"
	"github.com/yoneyan/vm_mgr/node/manage"
	"log"
	"strings"
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

func CreateVMProcess(c *CreateVMInformation) (string, bool) {
	fmt.Println("----VMNewCreate")

	if manage.VMVncExistsCheck(c.VNC) {
		fmt.Println("A VM with the same vnc port exists. So, change vnc port of the VM.")
	} else {
		if manage.VMExistsName(c.Name) {
			fmt.Println("A VM with the same name exists. So, change the name of the VM.")
		} else {
			if len(c.Net) != 0 {
				d := strings.Split(c.Net, ",")
				fmt.Println(d)
				var net []string
				fmt.Println(net)
				net = append(net, "0")
				for _, a := range d {
					net = append(net, a)
					net = append(net, manage.GenerateMacAddresss())
				}
				fmt.Println(net)
				c.Net = strings.Join(net, ",")
				fmt.Println(c.Net)
			} else {
				c.Net = ""
			}
			CreateVMDBProcess(c)
			err := RunQEMUCmd(CreateGenerateCmd(c))
			if err != nil {
				fmt.Println(err)
				log.Fatal("VMNewCreate Error!!")
				return "Error: RunQEMUCmd", false
			} else {
				db.VMDBStatusUpdate(c.ID, 1)
			}
		}
	}

	return "ok", true
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
	if db.AddDBVM(dbdata) == false {
		fmt.Println("Error: Failed add vm database")
	}
}
