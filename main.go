package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sayatumarov/sample_rest_api/getRequest"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getRequest.GetUsers).Methods("GET")
	http.ListenAndServe(":8000", router)
}
