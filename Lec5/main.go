package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/schema"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	Username string
	Password string
	Age      int
	Phone    string
	Link     string
}

//LoginPageHandler ....
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("error while executing template:", err)
			return
		}
	} else {
		user := ReadUserForm(r)
		fmt.Fprintf(w, "Hello "+user.Username+" !!")
	}

}

//ReadUserForm ...
func ReadUserForm(r *http.Request) *User {
	r.ParseForm()                           //Получить все данные из запроса, которые касаются форм запроса
	user := new(User)                       //Пустышка пользователя
	decoder := schema.NewDecoder()          // Стандартный декодер для форм
	err := decoder.Decode(user, r.PostForm) // Перенесем в поинтер на User все, что было в теле POST запроса касаемо формы.
	if err != nil {
		log.Println("error mapping user from Post form:", err)
	}
	return user
}

func main() {
	http.HandleFunc("/login", LoginPageHandler)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
