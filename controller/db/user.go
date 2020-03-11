package db

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/etc"
	"log"
)

//userdata
func AddDBUser(data User) bool {
	db := connectdb()
	addDb, err := db.Prepare(`INSERT INTO "userdata" ("name","pass") VALUES (?,?)`)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if _, err := addDb.Exec(data.Name, etc.Hashgenerate(data.Pass)); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func RemoveDBUser(name string) bool {
	db := connectdb()
	deleteDb := "DELETE FROM userdata WHERE name = ?"
	_, err := db.Exec(deleteDb, name)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func PassAuthDBUser(name, pass string) bool {
	db := connectdb()
	var hash string
	if err := db.QueryRow("SELECT pass FROM userdata WHERE name = ?", name).Scan(&hash); err != nil {
		fmt.Println(err)
	}
	fmt.Println("User OnlyAdmin Success")

	return etc.Verifyhashdata(pass, hash)
}

func GetDBUserID(name string) (int, bool) {
	db := connectdb()

	var id int
	if err := db.QueryRow("SELECT id FROM userdata WHERE name = ?", name).Scan(&id); err != nil {
		fmt.Println(err)
		return 0, false
	}

	return id, true
}

func GetDBAllUser() []User {

	db := *connectdb()

	rows, err := db.Query("SELECT * FROM userdata")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var bg []User
	for rows.Next() {
		var b User
		err := rows.Scan(&b.ID, &b.Name, &b.Pass)
		if err != nil {
			log.Println(err)
		}
		bg = append(bg, b)
	}
	return bg
}

func ChangeDBUserName(id int, data string) bool {
	db := connectdb()

	dbdata := "UPDATE userdata SET user = ? WHERE id = ?"
	_, err := db.Exec(dbdata, data, id)

	if err != nil {
		fmt.Println("Error: DBUpdate Error (User Name)")
		return false
	}

	return true
}

func ChangeDBUserPassword(id int, data string) bool {
	db := connectdb()

	dbdata := "UPDATE userdata SET pass = ? WHERE id = ?"
	_, err := db.Exec(dbdata, etc.Hashgenerate(data), id)

	if err != nil {
		fmt.Println("Error: DBUpdate Error (User Pass)")
		return false
	}

	return true
}
