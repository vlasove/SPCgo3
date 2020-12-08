## Лекция 3. Работа с шаблонами

***Проблема:*** отображение байтовых последовательностей достаточно неудобно. Для более наглядной структуризации информации на веб-страницах воспользуемся ```html```-шаблоны.

***Задача:*** прикрутить набор ```html```-шаблонов и заставить сервер работать вместе с ними.

### Шаг 1. Создание простейшего шаблона
Для того, чтобы создать простейший шаблон определим стандартное место, для расположения всех ```html``` шаблонов в проекте.
* Создадим файл ```templates/home.html```
* Определим содержимое шаблона:
```
<!--templates/home.html-->
<html>
    <head>
        <title>
            Home Page Title
        </title>
    </head>
    <body>
        <h1>
            Home Page body!!
        </h1>
        <!--{{ }} - это отбращение к встроенному шаблонизатору-->
        <p>Hello, {{ .Name }}</p>
        <p>Your age is {{ .Age }}</p>
    </body>
</html>
```

Теперь внутри файла ```main.go``` пропишу основную логику взаимодействия с шаблоном:
```
//main.go
package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	//CONN_PORT ...
	CONN_PORT = "8080"
	//CONN_HOST ...
	CONN_HOST = "localhost"
)

//User ...
type User struct {
	Age  string
	Name string
}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Age: "35", Name: "Bob"}
	// //Спарсим шаблон
	parsedTemplate, _ := template.ParseFiles("templates/home.html")
	err := parsedTemplate.Execute(w, user) // В writer поместим шаблон и передадим шаблонизатору нашего пользователя
	if err != nil {
		log.Println("An error happend while parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", HomePageHandler)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}

```