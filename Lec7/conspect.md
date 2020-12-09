## Лекция 7. Создание собственных Cookie-файлов

Первым делом, для создания собственных ```cookie``` выполинм команду : ```go get github.com/gorilla/securecookie```

### Шаг 1. Конфигурируем собственный cookieHandler
```
//Определение handler'a для создания собственных cookie
var cookieHandler *securecookie.SecureCookie

func init(){
	cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

```

### Шаг 2. Создание собственных cookie
Для начала определим стартовый шаблон ```main.go```
```
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

const (
	connPort = "8080"
	connHost = "localhost"
)

//Определение handler'a для создания собственных cookie
var cookieHandler *securecookie.SecureCookie

func init() {
	cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

//Блок создания cookie
func createCookie(w http.ResponseWriter, r *http.Request) {}

//Блок чтения cookie
func readCookie(w http.ResponseWriter, r *http.Request) {}
func main() {
	http.HandleFunc("/create", createCookie)
	http.HandleFunc("/read", readCookie)

	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
	}

}

```

### Шаг 3. Функция createCookie
Для создания cookie необходимо определить объект, оторый будет выступать в качестве парного сопоставления ```key:value```. Затем данный объект нужно будет передать в пользовательскую сессию.

```
//Блок создания cookie
func createCookie(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{"username": "Alex"}
	//Для передачи данных через cookie необходимо сериализовать
	base64Encoded, err := cookieHandler.Encode("key", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "first-cookie",
			Value: base64Encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	w.Write([]byte("Cookie created!"))
}
```

### Шаг 4. Функция readCookie
Теперь получим информацию из созданного ```cookie```
```
//Блок чтения cookie
func readCookie(w http.ResponseWriter, r *http.Request) {
	log.Println("Now reading cookie proces....")
	cookie, err := r.Cookie("first-cookie")
	if cookie != nil && err == nil {
		value := make(map[string]string)
		if err = cookieHandler.Decode("key", cookie.Value, &value); err == nil {
			w.Write([]byte(fmt.Sprintf("Hello, %v !\n", value["username"])))
		}
	} else {
		log.Println("Cookie not found in this request....")
		w.Write([]byte("Hello !"))
	}
}

```