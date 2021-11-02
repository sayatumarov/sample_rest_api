package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/simpleDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers.GetUsers).Methods("GET")
	http.ListenAndServe(":8000", router)
}
