package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const db_name = "./main.sql"

type NodeVM struct {
	ID          int
	Name        string
	CPU         int
	Mem         int
	StoragePath string
	Net         string
	Vnc         int
	Socket      string
	Status      int
}

func connectdb() *sql.DB {
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("SQL open error")
		panic(err)
	}

	//defer db.Close()

	return db
}

func createdb(database string) bool {
	db := *connectdb()

	_, err := db.Exec(database)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func Initdb() bool {
	//nodevm data
	createdb(`CREATE TABLE IF NOT EXISTS "nodevm" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255), "cpu" INT,"memory" INT, "storagepath" VARCHAR(255),"net" VARCHAR(255),"vnc" INT, "socket" VARCHAR(255),"status" INT)`)
	return true
}

//NodeVM

func AddDBVM(data NodeVM) bool {
	fmt.Println("add database: " + data.Name)
	db := *connectdb()
	addDb, err := db.Prepare(`INSERT INTO "nodevm" ("name","cpu","memory","storagepath","net","vnc","socket","status") VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		panic(err)
		return false
	}

	if _, err := addDb.Exec(data.Name, data.CPU, data.Mem, data.StoragePath, data.Net, data.Vnc, data.Socket, data.Status); err != nil {
		panic(err)
		return false
	}
	return true
}

func DeleteDBVM(id int) bool {
	db := connectdb()
	deleteDb := "DELETE FROM nodevm WHERE id = ?"
	_, err := db.Exec(deleteDb, id)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func GetDBAll() []NodeVM {
	db := *connectdb()

	cmd := "SELECT * FROM nodevm"
	rows, _ := db.Query(cmd)

	defer rows.Close()

	var bg []NodeVM
	for rows.Next() {
		var b NodeVM
		err := rows.Scan(&b.ID, &b.Name, &b.CPU, &b.Mem, &b.StoragePath, &b.Net, &b.Vnc, &b.Socket, &b.Status)
		if err != nil {
			log.Println(err)
		}
		bg = append(bg, b)
	}
	return bg
}

func VMDBGetVMID(name string) (int, error) {
	data := GetDBAll()
	for i, _ := range data {
		if data[i].Name == name {
			return data[i].ID, nil
		}
	}

	return -1, fmt.Errorf("Not Found!!!")
}

func VMDBGetVMStatus(id int) (int, error) {
	//0: PowerOff 1: PowerOn 2:Suspend 3: TmpStop 4: busy
	data := GetDBAll()
	for i, _ := range data {
		if data[i].ID == id {
			return data[i].Status, nil
		}
	}

	return -1, fmt.Errorf("Not Found!!!")
}

func VMDBStatusUpdate(id, status int) bool {
	db := *connectdb()
	fmt.Println("status:" + string(status))

	cmd := "UPDATE nodevm SET status = ? WHERE id = ?"
	_, err := db.Exec(cmd, status, id)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	return true
}

func VMDBGetData(id int) (*NodeVM, error) {
	data := GetDBAll()
	var result NodeVM
	for i, _ := range data {
		if data[i].ID == id {
			result = data[i]
			fmt.Println(i)
			return &result, nil
		}
	}
	return &result, fmt.Errorf("Not Found")
}

func VMDBUpdate(data *NodeVM) {
}

func VMDBStatusStop(id int) {

}
