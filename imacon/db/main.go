package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const DBPath = "./imacon.db"

type Image struct {
	ID        int
	FileName  string
	Name      string
	Tag       string
	Type      int
	Capacity  int
	AddTime   int
	Authority int
	MinMem    int
	Status    int
}

type Transfer struct {
	ID      int
	ImageID int
	UUID    string
	Status  int
}

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		fmt.Println("SQL open error")
		fmt.Println(err)
	}

	return db
}

func createDB(database string) bool {
	db := *connectDB()
	defer db.Close()

	_, err := db.Exec(database)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func InitDB() bool {
	//ImageData
	createDB(`CREATE TABLE IF NOT EXISTS "image" ("id" INTEGER PRIMARY KEY, "filename" VARCHAR(500), "name" VARCHAR(500),"tag" VARCHAR(500),"type" INT,"capacity" INT, "addtime" INT, "minmem" INT,"authority" INT,"status" INT)`)
	//TransferData
	createDB(`CREATE TABLE IF NOT EXISTS "transfer" ("id" INTEGER PRIMARY KEY, "imageid" INT ,"uuid" VARCHAR(500) ,"status" INT`)

	return true
}
