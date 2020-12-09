package handlers

import (
	"html/template"
	"net/http"

	"github.com/vlasove/course/Lec8/cookies"
)

//LoginPageHandler ...
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	//Показываем юзеру форму для логина
	parsedTemplate, _ := template.ParseFiles("templates/login.html")
	parsedTemplate.Execute(w, nil)
}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//Извлечем имя пользователя из сессии и подадим на вход шаблону home.html
	UserName := cookies.GetUserName(r)
	if UserName != "" {
		data := map[string]interface{}{
			"UserName": UserName,
		}
		parsedTemplate, _ := template.ParseFiles("templates/home.html")
		parsedTemplate.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

//LoginFormPageHandler ...
func LoginFormPageHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	target := "/"
	if username != "" && password != "" {
		cookies.SetSession(username, w)
		target = "/home"
	}
	http.Redirect(w, r, target, 302)
}

//LogoutFormPageHandler ...
func LogoutFormPageHandler(w http.ResponseWriter, r *http.Request) {
	//Очищаем сессию и редиректим на LoginPage
	cookies.ClearSession(w)
	http.Redirect(w, r, "/", 302)
}
