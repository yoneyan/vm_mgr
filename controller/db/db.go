package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const DBPath = "./main.db"

type Node struct {
	ID        int
	HostName  string
	IP        string
	Path      string
	OnlyAdmin int
	MaxCPU    int
	MaxMem    int
	Status    int
}

type User struct {
	ID   int
	Name string
	Pass string
	Auth int
}

type Group struct {
	ID         int
	Name       string
	Admin      string
	User       string
	MaxVM      int
	MaxCPU     int
	MaxMem     int
	MaxStorage int
	Net        string
}

func connectdb() *sql.DB {
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		fmt.Println("SQL open error")
		fmt.Println(err)
		//panic(err)
	}

	//defer db.Close()
	return db
}

func createdb(database string) bool {
	db := *connectdb()

	_, err := db.Exec(database)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func InitDB() bool {
	//Node data
	createdb(`CREATE TABLE IF NOT EXISTS "node" ("id" INTEGER PRIMARY KEY, "hostname" VARCHAR(255), "ip" VARCHAR(255), "path" VARCHAR(2000), "onlyadmin" INT,"maxcpu" INT ,"maxmem" INT, "status" INT)`)
	//user data
	createdb(`CREATE TABLE IF NOT EXISTS "userdata" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255), "pass" VARCHAR(255))`)
	//group data
	createdb(`CREATE TABLE IF NOT EXISTS "groupdata" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255),"admin" VARCHAR(500),"user" VARCHAR(2000),"maxvm" INT,"maxcpu" INT,"maxmem" INT,"maxstorage" INT,"net" VARCHAR(255))`)

	return true
}
