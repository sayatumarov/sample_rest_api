package postRequest

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Name     string
	Password string
}

func RegUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t user
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, t.Name)
}
