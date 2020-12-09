## Лекция 8. Реализация механизмов login/logout через cookie
***Задача:*** реализовать механизм ```login/logout``` с использованием сохранения состояния через наши собственные cookie.

### Шаг 1. Инициализация шаблонов
Возьмем шаблон ```templates/login.html```:
```
<!DOCTYPE html>
<html>
    <head>
        <title>Login Page</title>
    </head>
    <body> 
        <form method="post" action="/login">
            <label for="username">Username</label>
            <input type="text" id="username" name="username">

            <label for="password">Password</label>
            <input type="password" id="password" name="password">

            <button type="submit">Login</button>
        </form>

    </body>
</html>
```
Создадим шаблон ```templates/home.html```:
```
<!DOCTYPE html>
<html>
    <head>
        <title>Home page</title>
    </head>
    <body>
        <h2>Welcome {{ .userName}}!</h2>
        <form method="post" action="/logout">
            <button type="submit">Logout</button>
        </form>
    </body>
</html>
```

### Шаг 2. Подготовка основного файла
Определим ```main.go``` с стандартным содержимым:
```
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//LoginPageHandler ...
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {}

//LoginFormPageHandler ...
func LoginFormPageHandler(w http.ResponseWriter, r *http.Request) {}

//LogoutFormPageHandler ...
func LogoutFormPageHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", LoginPageHandler)
	router.HandleFunc("/home", HomePageHandler)
	router.HandleFunc("/login", LoginFormPageHandler).Methods("POST")
	router.HandleFunc("/logot", LogoutFormPageHandler).Methods("POST")
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}

```

### Шаг 3. Опишем логику передачи данных через cookie
* Нужно уметь читать информацию о пользователе из запроса ```getUserName()```
* Нам нужно инициализировать и настривать сессию ```setSession()```
* Нам нужно уметь очищать пользовательскую сессию ```clearSession()```

```
var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//SetSession ...
func SetSession(userName string, response http.ResponseWriter) {
	value := map[string]string{"username": userName}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

//ClearSession ...
func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//GetUserName ...
func GetUserName(request *http.Request) (userName string) {
	cookie, err := request.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}

```
### Шаг 4. Описание взаимодействия с отображениями
Необходимо релизовать 4 функции-отображения:
* ```HomePageHandler(w http.ResponseWriter, r *http.Request) {}```
```
//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//Извлечем имя пользователя из сессии и подадим на вход шаблону home.html
	userName := GetUserName(r)
	if userName != "" {
		data := map[string]interface{}{
			"username": userName,
		}
		parsedTemplate, _ := template.ParseFiles("templates/home.html")
		parsedTemplate.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
```

* ```LoginPageHandler(w http.ResponseWriter, r *http.Request)```
```
//LoginPageHandler ...
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	//Показываем юзеру форму для логина
	parsedTemplate, _ := template.ParseFiles("templates/login.html")
	parsedTemplate.Execute(w, nil)
}
```

* ```LogoutFormPageHandler(w http.ResponseWriter, r *http.Request)```
```
//LogoutFormPageHandler ...
func LogoutFormPageHandler(w http.ResponseWriter, r *http.Request) {
	//Очищаем сессию и редиректим на LoginPage
	ClearSession(w)
	http.Redirect(w, r, "/", 302)
}
```

* ```LoginFormPageHandler(w http.ResponseWriter, r *http.Request)```
```
//LoginFormPageHandler ...
func LoginFormPageHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	target := "/"
	if username != "" && password != "" {
		SetSession(username, w)
		target = "/home"
	}
	http.Redirect(w, r, target, 302)
}
```

### Шаг 5. Исходник проекта.
```
//main.go
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//SetSession ...
func SetSession(userName string, response http.ResponseWriter) {
	value := map[string]string{"username": userName}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

//ClearSession ...
func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//GetUserName ...
func GetUserName(request *http.Request) (userName string) {
	cookie, err := request.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}

//LoginPageHandler ...
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	//Показываем юзеру форму для логина
	parsedTemplate, _ := template.ParseFiles("templates/login.html")
	parsedTemplate.Execute(w, nil)
}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//Извлечем имя пользователя из сессии и подадим на вход шаблону home.html
	userName := GetUserName(r)
	if userName != "" {
		data := map[string]interface{}{
			"username": userName,
		}
		parsedTemplate, _ := template.ParseFiles("templates/home.html")
		parsedTemplate.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

//LoginFormPageHandler ...
func LoginFormPageHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	target := "/"
	if username != "" && password != "" {
		SetSession(username, w)
		target = "/home"
	}
	http.Redirect(w, r, target, 302)
}

//LogoutFormPageHandler ...
func LogoutFormPageHandler(w http.ResponseWriter, r *http.Request) {
	//Очищаем сессию и редиректим на LoginPage
	ClearSession(w)
	http.Redirect(w, r, "/", 302)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", LoginPageHandler)
	router.HandleFunc("/home", HomePageHandler)
	router.HandleFunc("/login", LoginFormPageHandler).Methods("POST")
	router.HandleFunc("/logout", LogoutFormPageHandler).Methods("POST")
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}
```