## Лекция 13. Дополнение до NoSQL

Для того, чтобы взаимодействовать с ```NoSQL``` СУБД, необходимо:
* ```mongo``` запуск сервера
* ```go get gopkg.in/mgo.v2```

***Задача простая*** получить имя базы используя ```mongo```?

### Шаг 1. Определим функцию GetCurrentDB()
```
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
)

const (
	connPort = "8080"
	connHost = "localhost"
	//ПОдключение к бд
	mongoDBURL = "127.0.0.1"
)

//Определим объект сессии
var session *mgo.Session

//Определим ошибки подключения
var connectionError error

func init() {
	session, connectionError := mgo.Dial(mongoDBURL)
	if connectionError != nil {
		log.Fatal("error connecting to mongo")

	}
	session.SetMode(mgo.Monotonic, true)
}

//GetCurrentDB ...
func GetCurrentDB(w http.ResponseWriter, r *http.Request) {
	db, err := session.DatabaseNames()
	if err != nil {
		log.Println("error getting databse name:", err)
		return
	}
	fmt.Fprintf(w, "databse name is :: %s", strings.Join(db, ", "))
}

func main() {
	http.HandleFunc("/", GetCurrentDB)
	defer session.Close()

	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}

```