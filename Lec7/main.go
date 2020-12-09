package main

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

//Определение handler'a для создания собственных cookie
var cookieHandler *securecookie.SecureCookie

func init() {
	cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

//Блок создания cookie
func createCookie(w http.ResponseWriter, r *http.Request) {}

//Блок чтения cookie
func readCookie(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/create", createCookie)
	http.HandleFunc("/read", readCookie)

	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
	}

}
