## Лекция 6. Пользовательские сессии

***Проблема*** : как научиться контролировать данные, сохраненные в пользовательской сессии? Как создать сессию? Как подключить ```redis```? Как создать свои ```cookie```?

### Шаг 1. Подготовим базовый файл.
В качестве базового файла ```main.go``` , который будет поддерживать 3 операции:
* ```/home``` (сюда не может попасть неаутентифицированный пользователь)
* ```/login```
* ```/logout```

***Хотелки:*** хотим создать минимальный функционал аутентификации с использованием пользовательской сессии. Запрос ```/home``` перед своим выполнением должен просматривать сессию и валидировать поле ```autheticated```: 
* если ```authenticated == true``` то доступ к данному запросу разрешен.
* если ```authenticated == false``` - доступ запрещен.

***Что нужно?*** При выполнении ```/login``` проставляем пользователю ```true```, при выполнении ```/logout``` проставляем ```false```.

Создадим набор шаблонов ```templates/home.html```, ```templates/register/login.html```.
Для начала установим пакет для работы с сессиями: ```go get github.com/gorilla/sessions```
Пропишем логику на уровне ```main.go```.
```
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//Определение хранилища пользовательской сессии
var store *sessions.CookieStore

//Инициализация хранилища
func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

//HomePageHandler ....
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var authenticated interface{} = session.Values["autheticated"]
	if authenticated != nil {
		//Если в authenticated что-то есть(значит пользователь либо логинился, либо уже разлогинился)
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "You are unauthorized for this page", http.StatusForbidden)
			return
		}
		fmt.Fprintln(w, "Home page for authorized user!")
	} else {
		//В cookie в принципе остуствует информация про залогиненного пользователя
		http.Error(w, "You are unauthorized for this page", http.StatusForbidden)
		return
	}
}

//LoginPageHandler ...
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintln(w, "You are successfully logged in!")
}

//LogoutPageHandler ...
func LogoutPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintf(w, "You are successfully logged out!")
}

func main() {
	http.HandleFunc("/home", HomePageHandler)
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/logout", LogoutPageHandler)

	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}

```

***Проверим, что все работает***. Для этого выполним следующие команды:
* ```curl -X GET http://localhost:8080/home```(должны увидеть сообщение ```You are unauthorized for this page```)
* Попробуем залогиниться : ```curl -X GET -i http://localhost:8080/login``` (Из данного сообщения нас инетресует пара session-name=value```
session-name=MTYwNzUxMjE0OXxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQThBRFdGMWRHaGxiblJwWTJGMFpXUUVZbTl2YkFJQ0FBRT18qpyNAATn_-wyE290iXeL4EP64IJQRxcnlqAcOpbnvbE=;```)
* Теперь попробуем перейти на запрос ```/home``` , передав дополнительные параметры в виде
``` curl --cookie "session-name=MTYwNzUxMjE0OXxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQThBRFdGMWRHaGxiblJwWTJGMFpXUUVZbTl2YkFJQ0FBRT18qpyNAATn_-wyE290iXeL4EP64IJQRxcnlqAcOpbnvbE=;" http://localhost:8080/home ```

***Проблем:*** не смотря на то, что пользователь ```/logout``` имея на имя старой сессии пользователь может вернуть на страницу ```/home```. Позднее, исправим это через взаимодействие с клиентом.


### Шаг 2. Реализация сессии через Redis
Один из самых популярных пакетов для работы с ```Redis``` является : ```go get gopkg.in/boj/redistore.v1```
* Теперь необходимо объявить, что наше хранилище Cookie будет совмещено с ```Redis```