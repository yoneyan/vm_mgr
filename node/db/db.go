package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const dbName = "./main.db"

type VM struct {
	ID          int
	Name        string
	CPU         int
	Mem         int
	Storage     string
	StoragePath string
	Net         string
	Vnc         int
	Socket      string
	Status      int
	AutoStart   bool
}

func connectdb() *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err)
		fmt.Println("SQL open error")
	}

	//defer db.Close()

	return db
}

func Createdb() bool {
	db := *connectdb()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "vm" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255), "cpu" INT,"memory" INT, "storage" VARCHAR(500),"storagepath" VARCHAR(500),"net" VARCHAR(255),"vnc" INT, "socket" VARCHAR(255),"status" INT,"autostart" boolean)`)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error: SQL open Error")
		return false
	}
	return true
}

//VM

func AddDBVM(data VM) bool {
	fmt.Println("add database: " + data.Name)
	db := *connectdb()
	addDb, err := db.Prepare(`INSERT INTO "vm" ("name","cpu","memory","storage","storagepath","net","vnc","socket","status","autostart") VALUES (?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error: SQL Prepare Error")
		return false
	}

	if _, err := addDb.Exec(data.Name, data.CPU, data.Mem, data.Storage, data.StoragePath, data.Net, data.Vnc, data.Socket, data.Status, data.AutoStart); err != nil {
		fmt.Println(err)
		fmt.Println("Error: SQL Exec Error")
		return false
	}
	return true
}

func DeleteDBVM(id int) bool {
	db := connectdb()
	deleteDb := "DELETE FROM vm WHERE id = ?"
	_, err := db.Exec(deleteDb, id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error: SQL open Error")
		return false
	}
	return true
}

func VMDBGetAll() []VM {
	db := *connectdb()

	cmd := "SELECT * FROM vm"
	rows, _ := db.Query(cmd)

	defer rows.Close()

	var bg []VM
	for rows.Next() {
		var b VM
		err := rows.Scan(&b.ID, &b.Name, &b.CPU, &b.Mem, &b.Storage, &b.StoragePath, &b.Net, &b.Vnc, &b.Socket, &b.Status, &b.AutoStart)
		if err != nil {
			fmt.Println(err)
		}
		bg = append(bg, b)
	}
	return bg
}

func VMDBGetVMID(name string) (int, error) {
	data := VMDBGetAll()
	for i, _ := range data {
		if data[i].Name == name {
			return data[i].ID, nil
		}
	}

	return -1, fmt.Errorf("Not Found!!!")
}

func VMDBGetVMStatus(id int) (int, error) {
	//0: PowerOff 1: PowerOn 2:Suspend 3: TmpStop 4: busy
	data := VMDBGetAll()
	for i, _ := range data {
		if data[i].ID == id {
			return data[i].Status, nil
		}
	}

	return -1, fmt.Errorf("Not Found!!!")
}

func VMDBStatusUpdate(id, status int) bool {
	db := *connectdb()

	cmd := "UPDATE vm SET status = ? WHERE id = ?"
	_, err := db.Exec(cmd, status, id)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	return true
}

func VMDBGetData(id int) (*VM, error) {
	data := VMDBGetAll()
	var result VM
	for i, _ := range data {
		if data[i].ID == id {
			result = data[i]
			fmt.Println(i)
			return &result, nil
		}
	}
	return &result, fmt.Errorf("Not Found")
}

func VMDBUpdate(data *VM) {
}

func VMDBStatusStop(id int) {

}
