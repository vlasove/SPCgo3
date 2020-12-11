## Лекция 14. MVC фреймворка 

***MVC*** - паттерн проектирования веб-приложения, основанный на связи ```Model <-> View <-> Controller```
Рассмотрим ```Beego```(```Revel```, ```Utron```). 
Получение старой утилиты (необязательно, но в случае, если не видит ```bee```): ```go get github.com/beego/bee```
Получение пакета :```go get github.com/astaxie/beego```

***Хотелки***:
* Создание пустого проекта
* Как определяется первая пара ```controller-router```
* Как создаются отображения в ```beego```
* Как можно натянуть ```Bootstrap4```
* Создание переменной сессии в контексте ```beego```
* Простейший фильтр
* Панель администратора и мониторинг приложения


### Шаг 1. Создание пустого проекта
* Выполним команду ```bee new <project_name>``` (после этого в директории появится проект). Перейдем в него
* ```cd <project_name>```
* Запустим проект ```bee run```

### Шаг 2. Определение пары ```controller-router```
* В директории ```controllers```  создадим файл ```firstcontroller.go```
```
/project/controllers/firstcontroller.go
package controllers

import beego "github.com/astaxie/beego/server/web"

//Аналог []Route из прошлый лекций
type FirstController struct {
	beego.Controller
}

type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees Employees

func init() {
	employees = Employees{
		Employee{Id: 1, FirstName: "Bob", LastName: "BobLast"},
		Employee{Id: 2, FirstName: "Alice", LastName: "AliceLast"},
	}
}

func (this *FirstController) GetEmployees() {
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = employees
	this.ServeJSON()
}

```

Подключим данный контроллер к общим ```routes``` приложения:
```
/project/routers/router.go
package routers

import (
	"project/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{},
		"get:GetEmployees")
}

```