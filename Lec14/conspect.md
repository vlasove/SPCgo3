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

### Шаг 3. Создание views
* ```Controllers``` - набор правил взаимодействия ***отображений*** и ***моделей***
* ```Views``` - набор отображений верхнего уровня (шаблоны, ```.js``` фреймворки и т.д.)

Создадим файл ```views/home.html```
```
<!DOCTYPE html>
<html>
    <head>
        <title>Dashboard</title>
    </head>
    <body>
        <h1>Current Employees list</h1>
        {{ range .employees }}
            <h2>Employee ID: {{ .Id}}</h2>
            <p>First Name : {{ .FirstName}}</p>
            <p>Last Name : {{ .LastName}}</p>
            <br>
        {{ end }}
    </body>
</html>
```

Восползуемся данным шаблоном. Перейдем в ```controllers/firstcontroller.go```:
```
....
func (this *FirstController) GetDashboard() {
	this.Data["employees"] = employees
	this.TplName = "dasboard.html"
}

...
```

Зарегестрируем новый ```route``` : ```routes/router.go```
```
package routers

import (
	"project/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{},
		"get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{},
		"get:GetDashboard")
}

```

### Шаг 4. Изменим view домашней страницы
Для этого создадим ```controllers/homecontroller.go```
```
package controllers

import beego "github.com/astaxie/beego/server/web"

//Аналог []Route из прошлый лекций
type HomeController struct {
	beego.Controller
}

func (this *HomeController) HomePage() {
	this.TplName = "home.html"
}

```

Теперь вместо стандартной домашней страницы будем показывать страницу шаблона. В файле ```routers/router.go```
```
func init() {
	beego.Router("/", &controllers.HomeController{},
		"get:HomePage") //Новая часть
	beego.Router("/employees", &controllers.FirstController{},
		"get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{},
		"get:GetDashboard")
}
```

В центральном блоке ```home.html``` поместим всю необходимую для отображения информацию:
```
<div class="container">
    <h1>Home page</h1>
    <p>
      To  <a href="/dashboard">dashboard</a>
      |
      To  <a href="/employees">JSON</a>
    </p>
</div>
```

Для обновления состояния ```dashboard.html``` выполним аналогичную замену и отрисуем сотрудников в виде :
```
{{ range .employees }}
<div class="card">
    <div class="card-header">
        <span class="font-weight-bold">Employee ID: {{ .Id}}</span>
    </div>
    <div class="card-body">
        <p>First Name : {{ .FirstName}}</p>
        <p>Last Name : {{ .LastName}}</p>
    </div>

    <div class="card-footer text-center text-muted">
        <a href="/">Link 1</a> | <a href="/">Link 2</a>
    </div>
        
</div>
{{ end }}

```


### Шаг 5. Создание переменной сессии в контексте ```beego```
Создадим новый контроллер сессии ```controllers/sessioncontroller.go```
```
package controllers

import beego "github.com/astaxie/beego/server/web"

//Аналог []Route из прошлый лекций
type SessionController struct {
	beego.Controller
}

//Для работы с сессией стандартно используется 3 шага
//login
//logout
//все остальные сценарии
func (this *SessionController) Login() {
	this.SetSession("authenticated", true)
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.WriteString("You successfully logged in!")
}

func (this *SessionController) Logout() {
	this.SetSession("authenticated", false)
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.WriteString("You successfully logged out!")
}

func (this *SessionController) HomeAfterLogin() {

	isAuth := this.GetSession("authenticated")
	if isAuth == nil || isAuth == false {
		this.Ctx.WriteString("You are unauthorized for this page!")
		return
	}
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.WriteString("Welcome to this HomePage analog!\n")
}

```

Теперь на уровне ```routers``` проведем регистрацию нашего нового контроллера:
```
beego.Router("/home", &controllers.SessionController{},
		"get:HomeAfterLogin")
	beego.Router("/login", &controllers.SessionController{},
		"get:Login")
	beego.Router("/logout", &controllers.SessionController{},
		"get:Logout")
}
```

Теперь необходимо подсказать приложению, через какой инструмент мы взаимодействуем с сессией?
Хотим через ```redis```.

Для обеспечения функционала сессионности приложения в ```main.go``` файле импортируем зависимость с ```regis```
```
import (
	_ "project/routers"
	beego "github.com/astaxie/beego/server/web"
	_ "github.com/astaxie/beego/server/web/session/redis"
)

```

Теперь пропишем на уровне конфигуратива какой сервис для сессий мы используем. Зайдем в ```conf/app.conf``` и допишем 3 следующие строки:
```
...
SessionOn = true
SessionProvider = "redis"
SessionProviderConfig = "127.0.0.1:6379"
```

### Шаг 6. Простейший фильтр
Создадим простейший фильтр, выполняющий роль логгера. На фильтре будет лежать следующая задача:
* Фильтр будет считывать ```IP``` адрес ,а также в какое время (локальное) произошел тот или иной запрос

***Фильтр*** - функционал, позволяющий выполнять действия ***ДО*** осуществления прямого запроса и ***ПОСЛЕ*** него.

Определим пакет ```filters``` внутри которого создадим ```simplefilter.go```
```
package filters

import (
	"fmt"
	"time"

	context "github.com/astaxie/beego/server/web/context"
)

var LogManager = func(ctx *context.Context) {
	fmt.Println("IP :: " + ctx.Request.RemoteAddr + ", Time :: " + time.Now().Format(time.RFC850))
}

```

В ```router.go``` подключим наш фильтр:
```
```