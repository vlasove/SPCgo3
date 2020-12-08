package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
)

//helloWolrd - функция отображения
func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	//Определим базовый рутер (multyplexer)
	mux := http.NewServeMux()
	//Связь url с функцией отображения
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}
