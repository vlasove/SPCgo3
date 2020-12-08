package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	custom_handlers "github.com/vlasove/course/Lec2/handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	//Иинициализация мультиплексера
	router := mux.NewRouter()

	//Если используем gorilla мультиплексер необходимо указывать поддерживаемые методы
	//Добавим возможность логирования
	router.Handle("/", handlers.LoggingHandler(os.Stdout, custom_handlers.GetRequestHandler)).Methods("GET")
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	router.Handle("/post", handlers.LoggingHandler(logFile, custom_handlers.PostRequestHandler)).Methods("POST")
	router.Handle("/hello/{name}", handlers.LoggingHandler(logFile, custom_handlers.MultyRequestHandler)).Methods("GET", "PUT")

	err = http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}
