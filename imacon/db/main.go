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

func conneectDB() *sql.DB {
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		fmt.Println("SQL open error")
		fmt.Println(err)
		//panic(err)
	}

	//defer db.Close()
	return db
}

func createDB(database string) bool {
	db := *conneectDB()

	_, err := db.Exec(database)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func InitDB() bool {
	//ImageData
	createDB(`CREATE TABLE IF NOT EXISTS "imagedata" ("id" INTEGER PRIMARY KEY, "name" VARCHAR(500), "type" INT,"capacity" INT, "addtime" INT, "minmem" INT,"onlyadmin" INT)`)
	return true
}
