package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/vlasove/course/Lec2/handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	//Иинициализация мультиплексера
	router := mux.NewRouter()

	//Если используем gorilla мультиплексер необходимо указывать поддерживаемые методы
	router.Handle("/", handlers.GetRequestHandler).Methods("GET")
	router.Handle("/post", handlers.PostRequestHandler).Methods("POST")
	router.Handle("/hello/{name}", handlers.MultyRequestHandler).Methods("GET", "PUT")

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}
