package data

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/db"
)

//User Certification Tool is testing now !!
func AdminUserCertification(name, pass string) bool {
	if db.PassAuthDBUser(name, pass) {
		SearchGroupUser(name, "admin", 0)
		{
			fmt.Println("Certification OK!! (Administrator)")
			return true
		}
	}
	return false
}

func UserCertification(name, pass string) bool {
	if db.PassAuthDBUser(name, pass) {
		fmt.Println("Certification OK!! (User)")
		return true
	}
	return false
}

func GroupUserCertification(name, pass, group string) bool {
	if db.PassAuthDBUser(name, pass) && SearchGroupUser(name, group, 1) {
		fmt.Println("Certification OK!! (GroupUser)")
		return true
	}
	return false
}
