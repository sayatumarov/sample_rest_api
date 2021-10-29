package main

import (
	"fmt"
	"net/http"
)

func main() {
	home := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	}
	http.HandleFunc("/bar", home)
	http.ListenAndServe(":8080", nil)
}
