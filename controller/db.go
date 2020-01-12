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

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "BOOKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}

	/*
		//create controller data table
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "controller"{"id" INTEGER PRIMARY KEY,"name" varchar,"ip" varchar,"port" integer)`,)
		if err != nil{
			log.Fatalln(err)
			fmt.Println("failed create controller data table")
			panic(err)
			value = false
		}
		//create node data table
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "node"("id" INTEGER PRIMARY KEY,"name" varchar,"ip" varchar,"port" integer,"cpu" integer,"memory" integer,"storage" integer)`,)
		if err != nil{
			log.Fatalln(err)
			panic(err)
			value = false
		}
		//create vm data table
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "node"("id" INTEGER PRIMARY KEY,"name" varchar,"vcpu" integer,"vmemory" integer,"vstorage" integer,"vnc" integer,)`,)
		if err != nil{
			log.Fatalln(err)
			panic(err)
			value = false
		}
	*/
	return value
}
