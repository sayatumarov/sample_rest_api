package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sayatumarov/sample_rest_api/getRequest"
	"github.com/sayatumarov/sample_rest_api/postRequest"
	"github.com/sayatumarov/sample_rest_api/serv"
)

func main() {
	var err error
	serv.DBCon, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/simpleDB")
	if err != nil {
		panic(err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/users", getRequest.GetUsers).Methods("GET")
	router.HandleFunc("/users", postRequest.RegUser).Methods("GET")
	http.ListenAndServe(":8000", router)
}
