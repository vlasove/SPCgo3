package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Перевая функция-отображение в виде промежуточной переменной
var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world from GetRequestHandler with variable creation!"))
	},
)

//PostRequestHandler ...
var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PostRequestHandler working with variable creation!"))
	},
)

//MultyRequestHandler ...
var MultyRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		w.Write([]byte("Hello, " + name))
	},
)
