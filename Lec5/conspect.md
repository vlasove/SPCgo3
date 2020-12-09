## Лекция 5. Продолжение Лекции 4.

***Проблема:***Как конфигурировать ```FileServer``` в условиях наличия мультиплексера?

***Решеине*** Научимся сообщать мультиплексеру откуда брать ```static``` файлы?

### Шаг 1. Перепишем main.go на мультиплексер
```
package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	Username string
	Age      int
	Phone    string
	Link     string
}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: "NewUser",
		Age:      35,
		Phone:    "+49 999 999 22 33",
		Link:     "github.com/new_user/portfolio",
	}
	parserdTemplate, _ := template.ParseFiles("templates/home.html")
	err := parserdTemplate.Execute(w, user)
	if err != nil {
		log.Println("error while parsing template with user:", err)
		return
	}

}

func main() {

	//Реконфигурация static через мультиплексер
	router := mux.NewRouter()

	router.HandleFunc("/", HomePageHandler).Methods("GET")
	//Поддержка самого файл-сервера
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	//Запуск сервера
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}

```
В данном исполнении мы изменили только функцию ```main```. 
***Основное преимущество такого подхода*** : для разных рутеров выбирать ***разные*** наборы статики.


### Шаг 2. Веб-формы и их поддержка. Пустая форма.
Каким образом можно отобразить веб-формы на странице и каким образом она коннектится к ```go application```?
Создадим форму логина : ```templates/login.html```:
```
```