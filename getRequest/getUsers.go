package getRequest

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sayatumarov/sample_rest_api/serv"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := serv.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	var users []User
	result, err := db.Query("SELECT id, name, password from users")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Name, &user.Password)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
