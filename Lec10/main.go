package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vlasove/course/Lec10/handlers"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"GetEmployees",
		"GET",
		"/employees",
		handlers.GetEmployees,
	},
	Route{
		"GetEmployee",
		"GET",
		"/employee/{id}",
		handlers.GetEmployee,
	},
	Route{
		"AddEmployee",
		"POST",
		"/employee/add",
		handlers.AddEmployee,
	},
	Route{
		"UpdateEmployee",
		"PUT",
		"/employee/update",
		handlers.UpdateEmployee,
	},
	Route{
		"DeleteEmployee",
		"DELETE",
		"/employee/delete",
		handlers.DeleteEmployee,
	},
}

//AddRoutes ...
func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter) // Наполнение набор соотношений
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting API server:", err)
		return
	}
}
