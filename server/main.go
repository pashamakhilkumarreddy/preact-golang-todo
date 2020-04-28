package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to React Golang App")
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
