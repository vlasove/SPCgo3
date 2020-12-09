package main

import (
	"log"
	"net/http"
	"text/template"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	Username string
	Age      int
	Phone    string
	Link     string
}

//LoginPageHandler ....
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/login.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error while executing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/login", LoginPageHandler)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
