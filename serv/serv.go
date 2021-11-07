package serv

import "database/sql"

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() {
	db, err := sql.Open("mysql", "root:root@/simpleDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	DB = db
}
