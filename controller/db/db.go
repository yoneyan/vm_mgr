package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yoneyan/vm_mgr/controller/etc"
	"log"
)

const db_name = "./main.sql"

type Controller struct {
	HostName string
	IP       string
	GRPCPort int
	SSHPort  int
	User     string
	Pass     string
}

type VmUser struct {
	Name string
	Pass string
}

type VmGroup struct {
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
	//controller data
	createdb(`CREATE TABLE IF NOT EXISTS "controller" ("id" INTEGER PRIMARY KEY, "hostname" VARCHAR(255), "ip" VARCHAR(255), "grpcport" INT,"sshport" INT, "user" VARCHAR(255),"password" VARCHAR(255))`)
	//user data
	createdb(`CREATE TABLE IF NOT EXISTS "user" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255), "pass" VARCHAR(255)`)
	//group data
	createdb(`CREATE TABLE IF NOT EXISTS "group" ("id" INTEGER PRIMARY Key, "name" VARCHAR(255),"user" VARCHAR(10000),"admin" VARCHAR(1000),"maxcpu" INT,"maxmem" INT,"maxstorage" INT`)
	//vm data
	createdb(`CREATE TABLE IF NOT EXISTS "vm" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255),"vcpu" INT,"vram" INT,"vstorage" integer,"vnc" INT,"net" VARCHAR(255))`)
	return true
}

//Controller

func AddDBController(data Controller) bool {
	db := *connectdb()
	addDb, err := db.Prepare(`INSERT INTO "controller" ("hostname","ip","grpcport","sshport","user","password") VALUES (?,?,?,?,?,?)`)
	if err != nil {
		panic(err)
		return false
	}

	if _, err := addDb.Exec(data.HostName, data.IP, data.GRPCPort, data.SSHPort, data.User, data.Pass); err != nil {
		panic(err)
		return false
	}
	return true
}

func DeleteDBController(hostname string) bool {
	db := connectdb()
	deleteDb := "DELETE FROM controller WHERE hostname = ?"
	_, err := db.Exec(deleteDb, hostname)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
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
