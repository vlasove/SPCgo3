package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	//Зададим ограничения на уровне структуры
	Username string `valid:"alpha, required"`
	Password string `valid:"alpha, required"`
	Age      int
	Phone    string
	Link     string
}

//VaildateUser ...
func ValidateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validateError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validateError, "Username")
		passwordError := govalidator.ErrorByField(validateError, "Password")
		if usernameError != "" {
			log.Println("username validation error:", usernameError)
			return valid, "Validation error with Username field"
		}

		if passwordError != "" {
			log.Println("password validation error:", passwordError)
			return valid, "Validation error with Password field"
		}
	}
	return valid, "Validation Error"
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
		valid, validationError := ValidateUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, validationError)
			return
		}
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
