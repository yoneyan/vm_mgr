package vm

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
)

func StartupProcess() {
	data := db.GetDBAll()
	var autostart []int
	for i, _ := range data {
		db.VMDBStatusUpdate(data[i].ID, 0)
		fmt.Printf("Status 0  VMID: %d", data[i].ID)
		if data[i].AutoStart {
			autostart = append(autostart, data[i].ID)
		}
	}
	fmt.Printf("AutoStartVMID: ")
	fmt.Println(autostart)

	for i, _ := range autostart {
		if StartVMProcess(autostart[i]) {
			fmt.Printf("Start VMID: %d", i)
		} else {
			fmt.Printf("Failed start VMID: %d", i)
		}

	}
	fmt.Println()

	fmt.Println("Start process is end!!")
}
