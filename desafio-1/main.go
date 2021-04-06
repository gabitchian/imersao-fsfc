package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8000.")
	http.ListenAndServe(":8000", nil)
}
