package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const db_name = "./main.sql"

func initdb() bool {

	value := true

	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("SQL open error")
		panic(err)
		value = false
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "controller" ("id" INTEGER PRIMARY KEY, "hostname" VARCHAR(255), "ip" VARCHAR(255), "port" INT, "user" VARCHAR(255),"password" VARCHAR(255))`)
	if err != nil {
		panic(err)
	}
	/*
		//create controller data table
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "controller"{"id" INTEGER PRIMARY KEY,"name" varchar,"ip" varchar,"port" integer)`, )
		if err != nil {
			log.Fatalln(err)
			fmt.Println("failed create controller data table")
			panic(err)
			value = false
		}
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
	return value
}

func db_controller(function ,hostname, ip string, port int, user, password string) bool{
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("SQL open error")
		panic(err)
		return false
	}

	if(function == "add"){
		add_db, err := db.Prepare(`INSERT INTO "controller" ("hostname","ip","port","user","password") VALUES (?,?,?,?,?)`)
		if err != nil {
			panic(err)
			return false
		}

		if _, err := add_db.Exec(hostname, ip, port, user, password); err != nil {
			panic(err)
			return false
		}
	}else if (function == "remove") {
		delete_db := "DELETE FROM controller WHERE hostname = ?"
		_, err = db.Exec(delete_db, hostname)
		if err != nil {
			log.Fatalln(err)
			return false
		}
	}

	return true
}