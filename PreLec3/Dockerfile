# Какую версию компилятора нужно использовать для работы приложения
FROM golang:1.12.0-alpine3.9
# Создать директорию app внутри контейнера
RUN mkdir /app


# Загрузка git
RUN apk add --no-cache git 


# Перенести все содержимое из текущего места запуска Dockerfile'a внутрь /app
ADD . /app
# cd app
WORKDIR /app 
# Выполнимая команда

# Установка зависимостей
RUN go mod download

RUN go build -o main . 
# Каким образом будем запускать наше приложение?
CMD ["/app/main"]