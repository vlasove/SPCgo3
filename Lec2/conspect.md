## Лекция 2. Продолжение Лекции 1. Реализация эффективного routing'а

***Задача***: когда в проекте 1-2 ```url``` ссылки - здесь все однозначно. Но когда появляется достаточно большео количество 
сопоставлений ```читаемость``` кода стремительно падает. Хочется поддерживать много сопоставлений без потери читаемости.
```
/hello -> helloHandleFunc
/exit -> exitHandleFunc
```

### Шаг 1. Использование стандартной библиотеки и базовая структуризация сопоставлений
Создадим ```main.go``` и  ```handlers/handlers.go``` и введем поддержку 3-ёх различных сопоставлений:
* ```"/"```
* ```"/login```
* ```"/logout```

```
//main.go
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

```

```
//handlers/handlers.go
package handlers

import (
	"fmt"
	"net/http"
)

//Описание отображений
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home page view!</h1>")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login page view!")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout page view!")
}

```
***Основная польза*** - явно разделить отображения от основного блока с ***сопоставлением***. Отображения находятся в пакете ```handlers```, блок сопотсавлений пристуствует в входной точке программы.



***Дома*** ОБЯЗАТЕЛЬНО почитать, как создаются многофайловые проекты на ```GO``` и как определяются правила ```import```.

### Шаг 2. Использование мультиплексера и стандартной структуризации
```go get github.com/gorilla/mux``` - сторонний пакет для работы с мультиплексером. Соответственно, сразу инициализируем ```go mod init```. 
***Задача*** как структуризуется набор соотношений в присутсвии мультиплексера?
При использовании мультиплексера подход к структурированию чаще всего не изменяется, но сами функции-отображения выступают в качестве промежуточных переменных.

Реализуем набор следующих сопоставлений:
* ```"/"``` -> GetRequest
* ```"/post"``` -> PostRequest
* ```"/hello/{name}"``` -> MultyRequest (поддерживает и Get и Put)

Теперь реализация ```main.go``` выглядит так
```
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	//Иинициализация мультиплексера
	router := mux.NewRouter()

	//Если используем gorilla мультиплексер необходимо указывать поддерживаемые методы
	router.Handle().Methods("GET")
	router.Handle().Methods("POST")
	router.Handle().Methods("GET", "PUT")

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}
```

А в файле ```handlers/handlers.go``` распишем наши функции-отображения в виде набор переменных
```
//handlers/handlers.go
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


```

***Проверим, что все работает***. Не будем конфигурировать ```Postman``` , а воспользуемся стандартной утилитой ```curl```.
* ```GET /``` : ```curl -X GET -i http://localhost:8080/```
* ```GET /hello/<name>``` : ```curl -X GET -i http://localhost:8080/hello/<name>```
* ```POST /post``` : ```curl -X POST -i http://localhost:8080/post```
* ```PUT /hello/<name>``` : ...

***Основаная идея*** показать, как декомпозируется сервер при использовании сторонних мультиплексеров (```gorilla/mux```).


### Шаг 3. Логирование , силами gorilla
Добавим возможность логировать события на нашем сервере. Для этого нам достаточно внести небольшое количество изменений внутри функции ```main```. 
Для добавления функционала логирования необходимо завести ```server.log``` файл:
```
//main.go
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

```

Данная реализация удовлетворяет следующим требованиям:
* Запросы по ```"/"``` логируются через ```os.Stdout```
* Запросы ```/post, /hello/{name}``` логируются через файл ```server.log```
