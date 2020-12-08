## Лекция 1. Создание простейшего HTTP сервера
***Задача*** : каким образом можно определить простейший ```http``` -сервера.

### Шаг 1. Использование стандартной библиотеки
Создадим файл ```main.go``` с следующим содержанием:
```
package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
)

//helloWolrd - функция отображения
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	//TODO: создать соотношение вида ```url``` => handlerFunc (функция отображения)
	http.HandleFunc("/", helloWorld)
	//1-ый вариант - явная обработка ошибки
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

	//2-ой вариант - классическое логирование
	//log.Fatal(http.ListenAndServe(CONN_HOST + ":" + CONN_PORT , nil))

}

```
Основная идея - предложить 2 варианта обработчиков для прослушивания запросов, а так же ввести понятие ***функция отображения***.
***Функция отображения*** - это функциональная сущность, которая ```может быть вызвана``` в результате осуществления запроса.

Определим ```makefile```:
```
.PHONY: run 
run:
	go run main.go 

.PHONY: build
build:
	go build main.go 

.PHONY: exec
exec:
	./main 

DEFAULT_GOAL := run
```

### Шаг 2. Сборка HTTP сервера с простейшей аутентификацией
***Аутентификация*** - это процесс узнавания (```свой-чужой```). 
***Основная идея*** реализации аутентификации состоит в ***оборачивании*** функции-отображения в обертку (```декоратор```), который будет проверять соответствие необходимой переданной информации с уже существующей в локальной БД.
Создадим файл ```main.go```:
```
package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
	//Известныйй пользователь
	USER     = "admin"
	PASSWORD = "admin"
)

//helloWolrd - функция отображения
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

//Обертка (Wrapper) == декоратор
//Данный декоратор будет проверять, что тот, кто вызвал функцию-отображение - пользователь
//присутсвующий в нашей бд
func AuthWrapper(handler http.HandlerFunc, realm string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//Блок проверки
		user, pass, ok := r.BasicAuth() // возвращает все переданные данные о текущем пользователе
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(PASSWORD)) != 1 {
			w.Header().Set("www-my-authservice", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are unauthorized to. Access denied\n"))
			return
		}
		//Блок вызова
		handler(w, r)
	}
}

func main() {
	//Связь url с функцией отображения
	http.HandleFunc("/", AuthWrapper(helloWorld, "Please enter your username and password"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}


```

***Что делает AuthWrapper?*** Функция ```func AuthWrapper(handler http.HandlerFunc, realm string) http.HandlerFunc``` выполняет следующие действия:
* Принимает на вход стандартную функцию-отображение (```handler```) 
* Перед ее вызовом проверяет, что пользователь передал необходимую информацию при осуществлении данного запроса:
```
user, pass, ok := r.BasicAuth() // возвращает все переданные данные о текущем пользователе
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(PASSWORD)) != 1 {
			w.Header().Set("www-my-authservice", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are unauthorized to. Access denied\n"))
			return
		}
``` 
* Если пользователь ***НЕ ПЕРЕДАЛ НИЧЕГО*** или ***ПЕРЕДАЛ НЕПРАВИЛЬНЫЙ ЛОГИН*** или ***ПЕРЕДАЛ НЕПРАВИЛЬНЫЙ ПАРОЛЬ*** (Неправильный == несовпадающий с той информацией, которая есть у нас про этого пользователя) - выводим сообщение про ```www-my-authservice....```
* В случае, если все совпало - вызываем беспрепятственно ```handler(w, r)```.

Такой подход называется ***декорирование***. P.S. ```Большинство аутентификационных операций проводятся при помощи использования декорирования. В целом, любой middleware работает в качестве декоратора```


### Шаг 3. Оптимизация HTTP responses с использованием GZIP оптимайзера
Хотим, чтобы наш сервер отвечал сжатыми в ```.gzip``` объектами, с целью разгрузки потока байт при общении с клиентами.
Воспользуемся готовым рещением в виде ```go get github.com/gorilla/handlers``` (выполнить команду в терминале).
Поскольку теперь проект несет сторонние зависимости - созаддим файл ```go.mod``` (```go mod init```).
Для решения данной задачи создадим ```main.go```
```
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
```
***Что было сделано?*** Теперь мы ввели мультиплексер(промежуточная сущность, через которую будут проводиться все операции по взаимодейсвтие с функциями-отображениями). Данный мультиплексер в связке с ```handlers.CompressHandler``` позволяет сжимать ```.gzip``` все ```response```'s в gzip с целью оптимизации потокового обращения к клиенту. 

```P.S.``` После иницилазации ```go.mod``` рекомендуутся не выполнять ```install gopls``` , т.к. данное расширение требует достаточно глубокой настройки.


***Определение:*** ```multyplexer``` - промежуточная сущность, которая позволяет регулировать соотношения ```url - ACTIVITY - handleFunc```. 