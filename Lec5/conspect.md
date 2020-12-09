## Лекция 5. Продолжение Лекции 4.

***Проблема:***Как конфигурировать ```FileServer``` в условиях наличия мультиплексера?

***Решеине*** Научимся сообщать мультиплексеру откуда брать ```static``` файлы?

### Шаг 1. Перепишем main.go на мультиплексер
```
package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	Username string
	Age      int
	Phone    string
	Link     string
}

//HomePageHandler ...
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: "NewUser",
		Age:      35,
		Phone:    "+49 999 999 22 33",
		Link:     "github.com/new_user/portfolio",
	}
	parserdTemplate, _ := template.ParseFiles("templates/home.html")
	err := parserdTemplate.Execute(w, user)
	if err != nil {
		log.Println("error while parsing template with user:", err)
		return
	}

}

func main() {

	//Реконфигурация static через мультиплексер
	router := mux.NewRouter()

	router.HandleFunc("/", HomePageHandler).Methods("GET")
	//Поддержка самого файл-сервера
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	//Запуск сервера
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}

}

```
В данном исполнении мы изменили только функцию ```main```. 
***Основное преимущество такого подхода*** : для разных рутеров выбирать ***разные*** наборы статики.


### Шаг 2. Веб-формы и их поддержка. Пустая форма.
Каким образом можно отобразить веб-формы на странице и каким образом она коннектится к ```go application```?
Создадим форму логина : ```templates/login.html```:
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
***Важно*** , чтобы поле ```type="password"``` было указано в поле ввода пароля.

Теперь пропишем лоигку в ```main.go``` файле:
```
package main

import (
	"log"
	"net/http"
	"text/template"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	Username string
	Age      int
	Phone    string
	Link     string
}

//LoginPageHandler ....
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/login.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error while executing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/login", LoginPageHandler)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}

```

### Шаг 3. Чтение из формы.
Для взаимодействия с формой воспользуемся пакетом ```go get github.com/gorilla/schema```. Поля ```Username```  и ```Password``` уже определены в структуре ```User```. Опишем функцию, по чтению данных из формы.
```
//ReadUserForm ...
func ReadUserForm(r *http.Request) *User {
	r.ParseForm()                           //Получить все данные из запроса, которые касаются форм запроса
	user := new(User)                       //Пустышка пользователя
	decoder := schema.NewDecoder()          // Стандартный декодер для форм
	err := decoder.Decode(user, r.PostForm) // Перенесем в поинтер на User все, что было в теле POST запроса касаемо формы.
	if err != nil {
		log.Println("error mapping user from Post form:", err)
	}
	return user
}

```

Теперь воспользуемся этой функцией:
```
//LoginPageHandler ....
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("error while executing template:", err)
			return
		}
	} else {
		user := ReadUserForm(r)
		fmt.Fprintf(w, "Hello "+user.Username+" !!")
	}

}

```

Весь код ```main.go``` файла:
```
package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/schema"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	Username string
	Password string
	Age      int
	Phone    string
	Link     string
}

//LoginPageHandler ....
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("error while executing template:", err)
			return
		}
	} else {
		user := ReadUserForm(r)
		fmt.Fprintf(w, "Hello "+user.Username+" !!")
	}

}

//ReadUserForm ...
func ReadUserForm(r *http.Request) *User {
	r.ParseForm()                           //Получить все данные из запроса, которые касаются форм запроса
	user := new(User)                       //Пустышка пользователя
	decoder := schema.NewDecoder()          // Стандартный декодер для форм
	err := decoder.Decode(user, r.PostForm) // Перенесем в поинтер на User все, что было в теле POST запроса касаемо формы.
	if err != nil {
		log.Println("error mapping user from Post form:", err)
	}
	return user
}

func main() {
	http.HandleFunc("/login", LoginPageHandler)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}

```


### Шаг 4. Валидация форм.
Существует 3 уровня валидации любой формы.:
* Валидация на уровне шаблона (```.js```) 
* Валидация на уровне приема ```POST``` запроса
* Валидация на уровне модели 

Рассмотирм валидацию на уровне ```POST``` запроса. Добавим в ```main.go``` функцию ```VaildateUser```. С нуля писать валидаторы не будем. Мы воспользуемся низкоуровней заготовкой для создания валидационных правил :```go get github.com/asaskevich/govalidator```
И теперь воспользуемся им в нашем ```main.go```
```
//main.go
...
//User ...
type User struct {
	//Зададим ограничения на уровне структуры
	Username string `valid:"alpha, required"`
	Password string `valid:"alpha, required"`
	Age      int
	Phone    string
	Link     string
}
//VaildateUser
func ValidateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string){
	valid, validateError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validateError, "Username")
		passwordError := govalidator.ErrorByField(validateError, "Password")
		if usernameError != "" {
			log.Println("username validation error:", usernameError)
			return valid, "Validation error with Username field"
		}

		if passwordError != ""{
			log.Println("password validation error:", passwordError)
			return valid, "Validation error with Password field"
		}
	}
	return valid, "Validation Error"
}
...
```

Теперь воспользуемся функцией ```ValidateUser``` :
```
//LoginPageHandler ....
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("error while executing template:", err)
			return
		}
	} else {
		user := ReadUserForm(r)
		valid, validationError := ValidateUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, validationError)
			return 
		}
		fmt.Fprintf(w, "Hello "+user.Username+" !!")
	}

}
```

Теперь полный код выглядит так:
```
//main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

//User ...
type User struct {
	//Зададим ограничения на уровне структуры
	Username string `valid:"alpha, required"`
	Password string `valid:"alpha, required"`
	Age      int
	Phone    string
	Link     string
}

//VaildateUser
func ValidateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validateError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validateError, "Username")
		passwordError := govalidator.ErrorByField(validateError, "Password")
		if usernameError != "" {
			log.Println("username validation error:", usernameError)
			return valid, "Validation error with Username field"
		}

		if passwordError != "" {
			log.Println("password validation error:", passwordError)
			return valid, "Validation error with Password field"
		}
	}
	return valid, "Validation Error"
}

//LoginPageHandler ....
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login.html")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			log.Println("error while executing template:", err)
			return
		}
	} else {
		user := ReadUserForm(r)
		valid, validationError := ValidateUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, validationError)
			return
		}
		fmt.Fprintf(w, "Hello "+user.Username+" !!")
	}

}

//ReadUserForm ...
func ReadUserForm(r *http.Request) *User {
	r.ParseForm()                           //Получить все данные из запроса, которые касаются форм запроса
	user := new(User)                       //Пустышка пользователя
	decoder := schema.NewDecoder()          // Стандартный декодер для форм
	err := decoder.Decode(user, r.PostForm) // Перенесем в поинтер на User все, что было в теле POST запроса касаемо формы.
	if err != nil {
		log.Println("error mapping user from Post form:", err)
	}
	return user
}

func main() {
	http.HandleFunc("/login", LoginPageHandler)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server:", err)
		return
	}
}

```