package serv

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Name     string
	Password string
}

var db *sql.DB
var err error

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8888)/simpleDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT id, name, password from users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var user User
		er := result.Scan(&user.Id, &user.Name, &user.Password)
		if er != nil {
			panic(er.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
