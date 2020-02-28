package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yoneyan/vm_mgr/controller/etc"
	"log"
)

const db_name = "./main.sql"

type Node struct {
	ID       int
	HostName string
	IP       string
	Port     int
	Auth     int
	MaxCPU   int
	MaxMem   int
	Status   int
}

type VmUser struct {
	ID   int
	Name string
	Pass string
	Auth int
}

type VmGroup struct {
	ID         int
	Name       string
	Admin      string
	User       string
	MaxCPU     int
	MaxMem     int
	MaxStorage int
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
	//Node data
	createdb(`CREATE TABLE IF NOT EXISTS "node" ("id" INTEGER PRIMARY KEY, "hostname" VARCHAR(255), "ip" VARCHAR(255), "port" INT, "auth" INT,"maxcpu" INT "maxmem" INT "status" INT`)
	//user data
	createdb(`CREATE TABLE IF NOT EXISTS "user" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255), "pass" VARCHAR(255),"auth" INT`)
	//group data
	createdb(`CREATE TABLE IF NOT EXISTS "group" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255),"user" VARCHAR(10000),"admin" VARCHAR(1000),"maxcpu" INT,"maxmem" INT,"maxstorage" INT`)
	return true
}

//Node

func AddDBNode(data Node) bool {
	db := *connectdb()
	addDb, err := db.Prepare(`INSERT INTO "node" ("id","hostname","ip","port","auth","maxcpu","maxmem","status") VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println("DBError!!")
		return false
	}

	if _, err := addDb.Exec(data.ID, data.HostName, data.IP, data.Port, data.Auth, data.MaxCPU, data.MaxMem, data.Status); err != nil {
		fmt.Println("Add Error!!")
		return false
	}
	return true
}

func DeleteDBNode(id int) bool {
	db := *connectdb()
	deleteDb := "DELETE FROM node WHERE id = ?"
	_, err := db.Exec(deleteDb, id)
	if err != nil {
		fmt.Println("Delete Failed!!")
		return false
	}
	return true
}

func NodeDBStatusUpdate(id, status int) bool {
	db := *connectdb()

	cmd := "UPDATE node SET status = ? WHERE id = ?"
	_, err := db.Exec(cmd, status, id)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func GetDBNodeID(id int) (Node, bool) {
	db := *connectdb()

	rows, err := db.Query("SELECT * FROM node WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var b Node
	err = rows.Scan(b.ID, b.HostName, b.IP, b.Port, b.Auth, b.MaxCPU, b.MaxMem, b.Status)

	if err != nil {
		fmt.Println(err)
		return b, false
	}

	return b, true
}

func GetDBAllNode() []Node {

	db := *connectdb()

	rows, err := db.Query("SELECT * FROM node")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var bg []Node
	for rows.Next() {
		var b Node
		err := rows.Scan(b.ID, b.HostName, b.IP, b.Port, b.Auth, b.MaxCPU, b.MaxMem, b.Status)
		if err != nil {
			log.Println(err)
		}
		bg = append(bg, b)
	}
	return bg

}

//User

func AddDBUser(data VmUser) bool {
	db := connectdb()
	addDb, err := db.Prepare(`INSERT INTO "user" ("name","pass") VALUES (?,?)`)
	if err != nil {
		panic(err)
		return false
	}

	if _, err := addDb.Exec(data.Name, etc.Hashgenerate(data.Pass)); err != nil {
		panic(err)
		return false
	}

	return true
}

func DeleteDBUser(name string) bool {
	db := connectdb()
	deleteDb := "DELETE FROM user WHERE name = ?"
	_, err := db.Exec(deleteDb, name)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func TestPassDBUser(name, pass string) bool {
	db := connectdb()
	var hash string
	if err := db.QueryRow("SELECT pass FROM user WHERE name = ?", name).Scan(&hash); err != nil {
		log.Fatal(err)
	}

	return etc.Verifyhashdata(pass, hash)
}

//group
func AddDBGroup(data VmGroup) bool {
	db := connectdb()
	addDb, err := db.Prepare(`INSERT INTO "group" ("name","user","admin","maxcpu","maxmem","maxstorage") VALUES (?,?,?,?,?,?)`)
	if err != nil {
		panic(err)
		return false
	}

	if _, err := addDb.Exec(data.Name, data.User, data.Admin, data.MaxCPU, data.MaxMem, data.MaxStorage); err != nil {
		panic(err)
		return false
	}

	return true
}

func DeleteDBGroup(name string) bool {
	db := connectdb()
	deleteDb := "DELETE FROM group WHERE name = ?"
	_, err := db.Exec(deleteDb, name)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}
