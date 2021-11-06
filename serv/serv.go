package serv

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/simpleDB")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	DB = db
}
