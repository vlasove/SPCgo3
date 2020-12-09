package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/boj/redistore.v1"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//Определение хранилища пользовательской сессии
//Переопределим на случай использования Redis
var store *redistore.RediStore
var err error

//Инициализация хранилища и конфигурация Redis
func init() {
	store, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		log.Fatal("error getting redis store:", err)
	}
}

//HomePageHandler ....
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var authenticated interface{} = session.Values["authenticated"]

	if authenticated != nil {
		//Если в authenticated что-то есть(значит пользователь либо логинился, либо уже разлогинился)
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "You are unauthorized for this page", http.StatusForbidden)
			return
		}
		fmt.Fprintln(w, "Home page for authorized user!")
	} else {
		//В cookie в принципе остуствует информация про залогиненного пользователя
		http.Error(w, "You are unauthorized for this page", http.StatusForbidden)
		return
	}
}

//LoginPageHandler ...
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	if err = session.Save(r, w); err != nil {
		log.Fatalf("Error saving session : %v", err)
	}
	fmt.Fprintln(w, "You are successfully logged in!")
}

//LogoutPageHandler ...
func LogoutPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	if err = session.Save(r, w); err != nil {
		log.Fatalf("Error saving session after logout: %v", err)
	}
	fmt.Fprintf(w, "You are successfully logged out!")
}

func main() {
	http.HandleFunc("/home", HomePageHandler)
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/logout", LogoutPageHandler)

	err := http.ListenAndServe(connHost+":"+connPort, nil)
	defer store.Close()
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
