package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	//CONN_PORT ...
	CONN_PORT = "8080"
	//CONN_HOST ...
	CONN_HOST = "localhost"
)

//User ...
type User struct {
	Age  string
	Name string
}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Age: "35", Name: "Bob"}
	// //Спарсим шаблон
	parsedTemplate, _ := template.ParseFiles("templates/home.html")
	err := parsedTemplate.Execute(w, user) // В writer поместим шаблон и передадим шаблонизатору нашего пользователя
	if err != nil {
		log.Println("An error happend while parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", HomePageHandler)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
