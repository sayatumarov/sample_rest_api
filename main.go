package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers.GetUsers).Methods("GET")
	http.ListenAndServe(":8000", router)
}
