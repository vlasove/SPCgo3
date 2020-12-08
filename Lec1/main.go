package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
	//Известныйй пользователь
	USER     = "admin"
	PASSWORD = "admin"
)

//helloWolrd - функция отображения
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

//Обертка (Wrapper) == декоратор
//Данный декоратор будет проверять, что тот, кто вызвал функцию-отображение - пользователь
//присутсвующий в нашей бд
func AuthWrapper(handler http.HandlerFunc, realm string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//Блок проверки
		user, pass, ok := r.BasicAuth() // возвращает все переданные данные о текущем пользователе
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(PASSWORD)) != 1 {
			w.Header().Set("www-my-authservice", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are unauthorized to. Access denied\n"))
			return
		}
		//Блок вызова
		handler(w, r)
	}
}

func main() {
	//Связь url с функцией отображения
	http.HandleFunc("/", AuthWrapper(helloWorld, "Please enter your username and password"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
