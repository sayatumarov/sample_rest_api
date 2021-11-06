package serv

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/simpleDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	DB = db
}
