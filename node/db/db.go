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
	createdb(`CREATE TABLE IF NOT EXISTS "nodevm" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255), "cpu" INT,"memory" INT, "storagepath" VARCHAR(255),"net" VARCHAR(255),"vnc" INT, "socket" VARCHAR(255))`)
	return true
}

//NodeVM

func AddDBVM(data NodeVM) bool {
	db := *connectdb()
	addDb, err := db.Prepare(`INSERT INTO "nodevm" ("name","cpu","memory","storagepath","net","vnc","socket") VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		panic(err)
		return false
	}

	if _, err := addDb.Exec(data.Name, data.CPU, data.Mem, data.StoragePath, data.Net, data.Vnc, data.Socket); err != nil {
		panic(err)
		return false
	}
	return true
}

func DeleteDBVM(name string) bool {
	db := connectdb()
	deleteDb := "DELETE FROM nodevm WHERE name = ?"
	_, err := db.Exec(deleteDb, name)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}
