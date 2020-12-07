## Hello Web из контейнера

***Задача***: напишем простейший ```Hello world``` используя ```net/http``` , упакуем все в контейнер , проверим работу и отправим на ```docker-hub```.

### Шаг 1. Напишем ядро
В файле ```main.go```:
```
//main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web!")
}

func hiWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from my go application!")
}

func main() {
	http.HandleFunc("/", helloWeb)
	http.HandleFunc("/hi", hiWeb)
	fmt.Println("Our application working......")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

```
Проверим что все работает: ```go run main.go``` и в адресной строке пройдем по пути ```localhost:8081/``` и ```localhost:8081/hi```


## Шаг 2. Создадим makefile для менежмента
Хотим ```make run```, ```make build```, ```make exec``` выполнять операции с нашим приложением:
```
.PHONY: run 
run:
	go run main.go

.PHONY:build
build:
	go build main.go 

.PHONY:exec
exec:
	./main

DEFAULT_GOAL := run
```

## Шаг 3. Запакуем в docker-контейнер
Наша задача состоит в том, что мы сейчас хотим создать образ, внутри которого будет запускаться наше приложение. Для этого сначала необходимо прописать инструкцию по сборке данного образа - ```Dockerfile```.
```
# Какую версию компилятора нужно использовать для работы приложения
FROM golang:1.12.0-alpine3.9
# Создать директорию app внутри контейнера
RUN mkdir /app
# Перенести все содержимое из текущего места запуска Dockerfile'a внутрь /app
ADD . /app
# cd app
WORKDIR /app 
# Выполнимая команда
RUN go build -o main . 
# Каким образом будем запускать наше приложение?
CMD ["/app/main"]
```

* ***Сначала*** нужно собрать образ : ```sudo docker build -t my-go-app .```
* Для того, чтобы убедиться, что образ существует : ```sudo docker images```
* Запусти образ с условием ***ПОРТ 8081 ВНУТРИ КОНТЕЙНЕРА СВЯЗАТЬ С НАШИМ ЛКОАЛЬНОМ ПОРТОМ 8080***: 
```sudo docker run -p 8080:8081 -it my-go-app``` (для того, чтобы установить ```Ctrl+C```)
* Если нужно образ удалить, выполним команду ```sudo docker image rm <IMAGE_ID> -f```

Теперь перенесем в ```makefile``` команды по работе с ```docker```:
```
...
.PHONY: docker_build
docker_build:
	sudo docker build -t my-go-app .

.PHONY: docker_run:
docker_run:
	sudo docker run -p 8080:8081 -it my-go-app 

.PHONY: images
images:
	sudo docker images 
...
```


### Шаг 4. Docker-hub

* Зарегестрироваться на сайте ```https://hub.docker.com/```
* Локально залогинимся в аккаунт docker-hub : ```sudo docker login```
* Подготовим образ к транспортировке в ```docker-hub``` : имена образов должны соответствовать определенному правило :
```<docherhub-username>/<dockerhub-repo> : TAG```` 
Для того, чтобы тегнуть существующий контейнер выполним команду:  ```sudo docker tag <CONTAINER-ID> vlasovevg/my-go-app:latest```
* Для отправки на ```docker-hub``` : ```sudo docker push vlasovevg/my-go-app```
* Для того, чтобы стянуть : ```sudo docker pull vlasovevg/my-go-app:latest```
