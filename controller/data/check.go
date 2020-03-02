package data

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/db"
	"github.com/yoneyan/vm_mgr/proto/proto-go"
	"strconv"
)

func CreateVMCheck(d *grpc.VMData) bool {
	if d.Vcpu < 0 || d.Vmem < 0 || d.Option.Vnc < 0 || d.Storage < 0 {
		return false
	}
	if d.Vmname == "" || d.Option.StoragePath == "" {
		return false
	}
	return true
}

func ExistGroupCheck(name string) bool {
	data := db.GetDBAllGroup()
	for _, a := range data {
		if a.Name == name {
			return true
		}
	}
	return false
}

func ExistUserCheck(name string) bool {
	data := db.GetDBAllUser()
	fmt.Println("UserCount: " + strconv.Itoa(len(data)))
	for _, a := range data {
		if a.Name == name {
			fmt.Println("Check OK(Exists User)")
			return true
		}
	}
	return false
}

func GroupAllUserCheck(name string) bool {
	data := db.GetDBAllGroup()
	for i, _ := range data {
		if SearchAllGroupUser(data[i].Admin, name) {
			if SearchAllGroupUser(data[i].User, name) {
				fmt.Println("Error: Exists group user.(Admin,User)")
				return false
			}
		}
	}
	return true
}
