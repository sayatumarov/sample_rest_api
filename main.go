package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sayatumarov/sample_rest_api/serv"
)

var db *sql.DB
var err error

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
	db.SetConnMaxLifetime(3)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", serv.GetUsers).Methods("GET")
	http.ListenAndServe(":8000", r)
}
