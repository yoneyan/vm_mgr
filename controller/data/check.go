package data

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/db"
	"github.com/yoneyan/vm_mgr/proto/proto-go/node"
)

func CreateVMCheck(d *node.VMData) bool {
	if d.Vcpu < 0 || d.Vmem < 0 || d.Vnc < 0 || d.Storage < 0 {
		return false
	}
	if d.Vmname == "" || d.StoragePath == "" {
		return false
	}
	return true
}

func ExistGroupCheck(name string) bool {

	return true
}

func GroupAllUserCheck(name string) bool {
	data := db.GetDBAllGroup()
	for i, _ := range data {
		if SearchAllGroupUser(data[i].Admin, name) == false && SearchAllGroupUser(data[i].User, name) == false {
			fmt.Println("Error: Exists group user.")
			return false
		}
	}
	return true
}
