## Make и его алиасы

Для более удобной работы с проектом/контейнерами/конфигами будем пользоваться ```makefile```'ми. Данные файлы хранят в себе наборы пар ```name``` : ```shell command```. Для того, чтобы создать файл с такими командами , необходимо инициализирвоать файл ```makefile```.

### Шаг 1. Создадим базовое приложение
В файле ```main.go``` напишем код по сложению 2-ух целых чисел
```
//SuperProject
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(a + b)
}

```

Что можно делать с эти проектом?
* Выполнение на месте```go run main.go```
* Компиляция ```go build main.go```
* Запуск исполняемого файла ```./main```

### Шаг 2. Опишем makefile
Хотим через команды ```make run```, ```make build``` ```make exec``` выполнять все функции из предыдущего пункта. Для этого в файле ```makefile``` определим 3 инструкции.
```
.PHONY: run
run :
	go run main.go

.PHONY: build
build:
	go build main.go 

.PHONY: exec 
exec:
	./main 

DEFAULT_GOAL := run 
```


### Шаг 3. Сделаем проект чуть более правдоподобным
Декомпозируем наш код и введем ряд дополнительных функций , также добавим пачку простейших модульных тестов.
```
//main.go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	result := add(a, b)*sub(b, a) - mult(a, a)
	fmt.Println("Result:", result)
}
```

Теперь создадим файл с модульными тестами ```main_test.go``` и опишем там 3 простейших теста:
```
//main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	var ans int
	ans = add(2, 3)
	if ans != 5 {
		t.Error("Expected 5 but has :", ans)
	}

	ans = add(2, 0)
	if ans != 2 {
		t.Error("Expected 2 but has :", ans)
	}
}

func TestSub(t *testing.T) {
	var ans int
	ans = sub(10, 2)
	if ans != 8 {
		t.Error("Expected 8 but has :", ans)
	}
}

func TestMult(t *testing.T) {
	var ans int
	ans = mult(2, 3)
	if ans != 6 {
		t.Error("Expected 6 but has :", ans)
	}
}
```

Для запуска используем команду ```go test -v``` .

Теперь добавим эту опцию к ```makefile```:
```
.PHONY: test
test :
	go test -v

.PHONY: build
build:
	go build main.go 

.PHONY: exec 
exec:
	./main 

.PHONY: run 
run:
	go run main.go

DEFAULT_GOAL := test
```