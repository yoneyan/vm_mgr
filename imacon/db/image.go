package db

import (
	"database/sql"
	"fmt"
)

func AddDBImage(data Image) bool {
	db := connectDB()
	addDb, err := db.Prepare(`INSERT INTO "image" ("filename" , "name" ,"tag" ,"type" ,"capacity" , "addtime" , "minmem" ,"authority" ,"status") VALUES (?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if _, err := addDb.Exec(data.FileName, data.Name, data.Tag, data.Type, data.Capacity, data.AddTime, data.MinMem, data.Authority, data.Status); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func RemoveDBImage(id int) bool {
	db := connectDB()
	deletedb := "DELETE FROM image WHERE id = ?"
	_, err := db.Exec(deletedb, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetNameDBImage(name string) []Image {
	db := connectDB()
	rows, err := db.Query("SELECT * FROM image")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var bg []Image
	for rows.Next() {
		var b Image
		err := rows.Scan(&b.ID, &b.FileName, &b.Name, &b.Tag, &b.Type, &b.Capacity, &b.AddTime, &b.MinMem, &b.Authority, &b.Status)
		if err != nil {
			fmt.Println(err)
		}
		bg = append(bg, b)
	}
	return bg
}

func GetDBImage(name, tag string) (Image, bool) {
	db := connectDB()

	rows := db.QueryRow("SELECT * FROM image WHERE name = ? AND tag = ?", name, tag)
	var b Image
	err := rows.Scan(&b.ID, &b.FileName, &b.Name, &b.Tag, &b.Type, &b.Capacity, &b.AddTime, &b.MinMem, &b.Authority, &b.Status)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("Not found")
		return b, false
	case err != nil:
		fmt.Println(err)
		fmt.Println("Error: DBError")
		return b, false
	default:
		return b, true
	}
}

func ChangeDBImageName(id int, data string) bool {
	db := connectDB()

	_, err := db.Exec("UPDATE image SET name = ? WHERE id = ?", data, id)
	if err != nil {
		fmt.Println("Error: DBUpdate Error (Group Name)")
		return false
	}

	return true
}

func ChangeDBImageTag(id int, tag string) bool {
	db := connectDB()

	_, err := db.Exec("UPDATE image SET tag = ? WHERE id = ?", tag, id)
	if err != nil {
		fmt.Println("Error: DBUpdate Error (Group Name)")
		return false
	}

	return true
}

func ChangeDBImageAuthority(id, authority int) bool {
	db := connectDB()

	_, err := db.Exec("UPDATE image SET authority = ? WHERE id = ?", authority, id)
	if err != nil {
		fmt.Println("Error: DBUpdate Error (Group Name)")
		return false
	}

	return true
}

func ChangeDBImageStatus(id, status int) bool {
	db := connectDB()

	_, err := db.Exec("UPDATE image SET status = ? WHERE id = ?", status, id)
	if err != nil {
		fmt.Println("Error: DBUpdate Error (Group Name)")
		return false
	}

	return true
}

func ChangeDBImageFileName(id int, filename string) bool {
	db := connectDB()

	_, err := db.Exec("UPDATE image SET filename = ? WHERE id = ?", filename, id)
	if err != nil {
		fmt.Println("Error: DBUpdate Error (Group Name)")
		return false
	}

	return true
}
