package db

import (
	"database/sql"
	"fmt"
)

func AddDBTransfer(data Transfer) bool {
	db := connectDB()
	addDb, err := db.Prepare(`INSERT INTO "transfer" ( "imageid" ,"uuid","status") VALUES (?,?)`)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if _, err := addDb.Exec(data.ImageID, data.UUID, data.Status); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func RemoveDBTransfer(uuid string) bool {
	db := connectDB()
	deletedb := "DELETE FROM transfer WHERE uuid = ?"
	_, err := db.Exec(deletedb, uuid)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetAllDBTransfer() []Transfer {
	db := connectDB()
	rows, err := db.Query("SELECT * FROM transfer")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var bg []Transfer
	for rows.Next() {
		var b Transfer
		err := rows.Scan(&b.ID, &b.ImageID, &b.UUID, &b.Status)
		if err != nil {
			fmt.Println(err)
		}
		bg = append(bg, b)
	}
	return bg
}

func GetDBTransfer(id int) (Transfer, bool) {
	db := connectDB()

	rows := db.QueryRow("SELECT * FROM transfer WHERE id = ?", id)
	var b Transfer
	err := rows.Scan(&b.ID, &b.ImageID, &b.UUID, &b.Status)

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

func ChangeDBTransferStatus(id, status int) bool {
	db := connectDB()

	_, err := db.Exec("UPDATE image SET status = ? WHERE id = ?", status, id)
	if err != nil {
		fmt.Println("Error: DBUpdate Error (Group Name)")
		return false
	}

	return true
}
