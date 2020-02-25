package manage

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
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

func VMIDCheck(id int) bool {
	if id < 0 {
		fmt.Println("VMID Check NG.")
		return false
	} else {
		fmt.Println("VMID Check OK.")
		return true
	}
}
