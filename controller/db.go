package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const db_name = "./main.sql"

type controller struct {
	HostName string
	ip       string
	port     int
	user     string
	password string
}

type vmUser struct {
	Name      string
	Pass      string
	Authority int
	MaxCPU    int
	MaxMemory int
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

func initdb() bool {
	createdb(`CREATE TABLE IF NOT EXISTS "controller" ("id" INTEGER PRIMARY KEY, "hostname" VARCHAR(255), "ip" VARCHAR(255), "port" INT, "user" VARCHAR(255),"password" VARCHAR(255))`)

	createdb(`CREATE TABLE IF NOT EXISTS "user" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(255), "pass" VARCHAR(255), "authority" INT)`)
	/*

		/*
			//create vm data table
			_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "vm"("id" INTEGER PRIMARY KEY,"name" varchar,"ip" varchar,"port" integer,"cpu" integer,"memory" integer,"storage" integer)`,)
			if err != nil{
				log.Fatalln(err)
				panic(err)
				value = false
			}
			//create vm data table
			_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "vm"("id" INTEGER PRIMARY KEY,"name" varchar,"vcpu" integer,"vmemory" integer,"vstorage" integer,"vnc" integer,)`,)
			if err != nil{
				log.Fatalln(err)
				panic(err)
				value = false
			}
	*/
	return true
}

//Controller

func addDBController(data controller) bool {
	db := *connectdb()
	addDb, err := db.Prepare(`INSERT INTO "controller" ("hostname","ip","port","user","password") VALUES (?,?,?,?,?)`)
	if err != nil {
		panic(err)
		return false
	}

	if _, err := addDb.Exec(data.HostName, data.ip, data.port, data.user, data.password); err != nil {
		panic(err)
		return false
	}
	return true
}

func deleteDBController(hostname string) bool {
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

func addDBUser(data vmUser) bool {
	db := connectdb()
	addDb, err := db.Prepare(`INSERT INTO "user" ("name","pass","authority") VALUES (?,?,?)`)
	if err != nil {
		panic(err)
		return false
	}

	if _, err := addDb.Exec(data.Name, hashgenerate(data.Pass), data.Authority); err != nil {
		panic(err)
		return false
	}

	return true
}

func deleteDBUser(name string) bool {
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

	return verifyhashdata(pass, hash)
}
