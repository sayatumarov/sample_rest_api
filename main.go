package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sayatumarov/sample_rest_api/getRequest"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getRequest.GetUsers).Methods("GET")
	http.ListenAndServe(":8000", router)
}
