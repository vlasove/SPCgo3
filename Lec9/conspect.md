## Лекция 9. Использование REST API в качестве промежуточного сервера

***Хотелка:*** посмотреть, как проектируется ```API``` для сервера? Как прикрутить к этому ```API``` клиента?

### Шаг 1. Простейший GET запрос
Для реализации ```GET``` запроса создадим файл ```get.go```
```
//get.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter) // Наполнение набор соотношений
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting API server:", err)
		return
	}
}

```

***С целью*** обеспечения быстрого и единообразного масштабирвования определим функцию ```AddRoutes```:
```
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"getEmplyee",
		"GET",
		"/employee/{id}",
		getEmplyee,
	},
}

//Функция для наполнения набора внутри мультиплексера
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
```


***Теперь*** определим структуры ```Employee```
```
//Employee ...
type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee
```

И добавим пару новых сотрудников в бд.
```
func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Bob", LastName: "Jack"},
		Employee{Id: "2", FirstName: "Alice", LastName: "Tompson"},
		Employee{Id: "3", FirstName: "George", LastName: "Lighter"},
	}
}

```

Теперь реализуем 2 функции : ```getEmployees``` и ```getEmployee```
```
func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, employee := range employees {
		if employee.Id == id {
			if err := json.NewEncoder(w).Encode(employee); err != nil {
				log.Println("error getting employee by id::", err)
			}
		}
	}
}
```

В итоге весь файл целиком
```
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"getEmployee",
		"GET",
		"/employee/{id}",
		getEmployee,
	},
}

//Employee ...
type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

//Инициализация хранилища сотрудников
var employees Employees

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Bob", LastName: "Jack"},
		Employee{Id: "2", FirstName: "Alice", LastName: "Tompson"},
		Employee{Id: "3", FirstName: "George", LastName: "Lighter"},
	}
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

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, employee := range employees {
		if employee.Id == id {
			if err := json.NewEncoder(w).Encode(employee); err != nil {
				log.Println("error getting employee by id::", err)
			}
		}
	}
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
```