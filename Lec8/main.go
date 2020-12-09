package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vlasove/course/Lec8/handlers"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.LoginPageHandler)
	router.HandleFunc("/home", handlers.HomePageHandler)
	router.HandleFunc("/login", handlers.LoginFormPageHandler).Methods("POST")
	router.HandleFunc("/logout", handlers.LogoutFormPageHandler).Methods("POST")
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}
