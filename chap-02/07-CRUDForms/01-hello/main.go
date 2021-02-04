package main

import (
	"log"
	"net/http"
	"text/template"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("main.html"))
	tpl.Execute(w, nil)
}

func main() {
	log.Println("Listening localhost@PORT:3000")

	http.HandleFunc("/", Home)

	http.ListenAndServe(":3000", nil)
}
