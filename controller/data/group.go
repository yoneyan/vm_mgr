package data

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/db"
	"strings"
)

func AddGroupUser(name, group string) bool {
	data, result := VerifyGroup(name, group)
	if result == false {
		return false
	}
	d, result := ProcessStringToArray(data.User, name, 0)
	if result == false {
		fmt.Println("Error: User is exists this group")
		return false
	}

	if db.ChangeDBGroupUser(data.ID, d) {
		return true
	}
	return false
}

func AddGroupAdmin(name, group string) bool {
	data, result := VerifyGroup(name, group)
	if result == false {
		return false
	}
	d, result := ProcessStringToArray(data.Admin, name, 0)
	if result == false {
		fmt.Println("Error: User is exists this group")
		return false
	}

	if db.ChangeDBGroupUser(data.ID, d) {
		return true
	}
	return false
}

func DeleteGroupAdmin(name, group string) bool {
	data, result := VerifyGroup(name, group)
	if result == false {
		return false
	}
	d, result := ProcessStringToArray(data.Admin, name, 1)
	if result == false {
		fmt.Println("Error: User is exists this group")
		return false
	}

	if db.ChangeDBGroupAdmin(data.ID, d) {
		return true
	}
	return false
}
func DeleteGroupUser(name, group string) bool {
	data, result := VerifyGroup(name, group)
	if result == false {
		return false
	}
	d, result := ProcessStringToArray(data.User, name, 1)
	if result == false {
		fmt.Println("Error: User is exists this group")
		return false
	}

	if db.ChangeDBGroupUser(data.ID, d) {
		return true
	}
	return false
}

func VerifyGroup(name, group string) (*db.Group, bool) {
	var d *db.Group
	_, result := db.GetDBUserID(name)
	if result == false {
		fmt.Println("Error: Not exists this User")
		return d, false
	}
	id, result := db.GetDBGroupID(group)
	if result == false {
		fmt.Println("Error: Not exists this group")
		return d, false
	}
	data, result := db.GetDBGroup(id)
	if result == false {
		fmt.Println("Error: Failed GetGroup")
		return d, false
	}
	return &data, true
}

//mode 0: create 1:delete
func ProcessStringToArray(basedata, data string, mode int) (string, bool) {
	if mode == 0 {
		basedataarray := strings.Split(basedata, ",")
		for _, a := range basedataarray {
			if a == data {
				return "0", false
			}
		}
		basedataarray = append(basedataarray, data)
		result := strings.Join(basedataarray, ",")
		fmt.Println("stringdata: " + result)
		return result, true
	} else if mode == 1 {
		var dataarray []string
		basedataarray := strings.Split(basedata, ",")
		for _, a := range basedataarray {
			if a == data {
				return "0", false
			} else {
				dataarray = append(dataarray, a)
			}
		}
		result := strings.Join(dataarray, ",")
		fmt.Println("stringdata: " + result)
		return result, true
	}
	return "0", false
}

//basedata: group admin and user data
func SearchAllGroupUser(basedata, searchnamedata string) bool {
	dataarray := strings.Split(basedata, ",")
	for _, d := range dataarray {
		if d == searchnamedata {
			return true
		}
	}
	return false
}

//mode 0:admin 1:user
func SearchGroupUser(searchnamedata, group string, mode int) bool {
	id, result := db.GetDBGroupID(group)
	if result == false {
		fmt.Println("Error: Not found group")
	}
	data, result := db.GetDBGroup(id)
	if result == false {
		fmt.Println("Error: Not found group")
	}
	if mode == 0 {
		d := strings.Split(data.Admin, ",")
		for _, a := range d {
			if a == searchnamedata {
				return true
			}
		}
	} else if mode == 1 {
		d := strings.Split(data.User, ",")
		for _, a := range d {
			if a == searchnamedata {
				return true
			}
		}
	}
	return false
}
