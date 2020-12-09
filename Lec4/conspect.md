## Лекция 4. Включение static/css

На прошлой лекции рассмотрели простейший пример использования шаблонов. Теперь рассмотрим простейший пример использования ```css```.

### Шаг 1. Определим css файл
Стандартно, все стили располагают по пути ```static/css/main.css```. 
С следующим содержимым:
```
body {
    color : red;
}
```

### Шаг 2. Определим шаблон с включением ```css```
В директории ```templates``` определим файл ```templates/home.html```:
```
<!--templates/home.html-->
<html>
    <head>
        <title>HomePage for Go app</title>
        <link rel="stylesheet" href="/static/css/main.css">
    </head>
    <body>
        <h2>Page with info about {{.Username}}</h2>
        <p>Age :{{ .Age}}</p>
        <p>Phone : {{ .Phone}}</p>
        <p>Github link : {{ .Link}}</p>
    </body>
</html>
```

### Шаг 4. Опишем main.go
В файле ```main.go``` определим стандартную связь с отображением , ***а также*** пропишем новый объект ***FileServer***.
При использовании static-контента в приложении важно помнить про 2 варианта размещения ```static```:
* Хостить статики на стороннем сервисе, а в шаблоны вставлять ссылки (***Плюсы:*** серверу не нужно хранить исходники, ***Минусы***: завязывание на производительности клиента и доступности стороннего сервиса)

* Хостим в внутренних сетях (Пример : ```CDN``` - ```Content Delivery Network```, размещенные внутри файловой системы).

* Гибридное размещение 

Мы же сейчас отконфигурируем свой собственный ```FileServer``` внутри приложения (создадим простейшую ```CDN``` внутри ```go-runtime```). И соответственно, будем из нее брать ```static```- файлы.

```
//main.go
package main

import (
	"log"
	"net/http"
	"text/template"
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

	// //Конфигурация FileServer
	fileServer := http.FileServer(http.Dir("static"))
	// //Обработка и перенаправление запросов через fileServer
	http.Handle("/static/", http.StripPrefix("/static/", fileServer)) //static/css/main.css
	//Описание соотношений
	http.HandleFunc("/", HomePageHandler)

	//Запуск сервера
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}

```

Основная идея состоит в том, что теперь мы все запросы шаблонов на использование ```static``` будем переадресовывать локальной директории. ```http.Handle("/static/", http.StripPrefix("/static/", fileServer)) ```
