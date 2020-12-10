## Лекция 12. Интеграция баз данных и модели данных

***Модель данных*** - способ взаимодействия серверных объектов с базой данных.
***Подготовка миграции*** - создание SQL запросов по изменению структуры таблицы.
***Миграция*** - применение подготовленных запросов к таблице.

### Шаг 1. Инциализация первичной миграции
Создадим скрипт ```migrations/v1.sql```
```
use mydb;

CREATE TABLE IF NOT EXISTS employee
(uid INT(10) NOT NULL AUTO_INCREMENT, name VARCHAR(100) NULL DEFAULT NULL, PRIMARY KEY (uid));

commit;
```
Первая миграция - выполним вручную (через ```MySql admin``` или ```PostgreSQL admin```).

### Шаг 2. Получение адаптера
В данном проекте используется ```MySql``` , для такой СУБД существует адаптер ```go get github.com/go-sql-driver/mysql```

### Шаг 3. Установка первичного подключения
***Рекомендуется*** запустить ```SQL сервер``` в качестве демона.
В файле ```main.go``` изложим все необходимые данные для подключения к таблице.
```
const (
	connHost = "8080"
	connPort = "localhost"
	//Конфиги СУБД
	driverName     = "mysql"               //Поменять в случае необходимости использования сторонних драйверов
	dataSourceName = "root:password@/mydb" //Поменять в случае необходимости использования сторонних драйверов
)

//Определение объекта БД
var db *sql.DB

//Определение ошибки этапа подключения
var connectionError error

//Инициалзируем подключение
func init() {
	//Сообщаем, через какой драйвер и к какой базе хотим подключиться
	db, connectionError = sql.Open(driverName, dataSourceName)
	if connectionError != nil {
		log.Fatal("error while connectiong to database:", connectionError)
	}
}
```

### Шаг 4. Стандартная функция GetCUrrentDB
Функция, возвращающая имя текущей базы данных
```
//GetCurrentDB ...
func GetCurrentDB(w http.ResponseWriter, r *http.Request) {
	//Сформируем запрос, который вернет имя базы данных
	var query = "SELECT DATABASE() as db"
	//Выполняем запрос
	rows, err := db.Query(query)
	if err != nil {
		log.Println("error executing DATABASE() query:", err)
		return
	}
	var currentDB string
	for rows.Next() {
		rows.Scan(&currentDB)
	}
	fmt.Fprintf(w, "Current Database: %s", currentDB)
}

func main() {
	//сРАЗУ ЖЕ НЕ ЗАБЫВАЕМ
	defer db.Close()
	http.HandleFunc("/", GetCurrentDB)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting server", err)
		return
	}
}

```

### Шаг 5. Добавление записи
Создадим функцию, которая добавляет новую запись в нашу таблицу.
```
//CreateEmployee ...
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	//Получение очереди
	vals := r.URL.Query()
	name, ok := vals["name"] // предеаем в качестве параметров на этапе /employee/create/name=bob
	if ok {
		log.Println("ready to insert new employee....")
		stmt, err := db.Prepare("INSERT INTO employee (uid, name) VALUES(NULL, ?)")
		if err != nil {
			log.Println("error while query INSERT preparing:", err)
			return
		}
		result, err := stmt.Exec(name[0])
		if err != nil {
			log.Println("error while executing INSERT:", err)
			return
		}
		//Получим id последнего вставленного пользователя
		id, err := result.LastInsertId()
		fmt.Fprintf(w, "Last inseted id:", id)
	} else {
		fmt.Fprintf(w, "Can not parse arguments in this request")
	}
}
```

### Шаг 6. Получение всех записей из бд
Получим абсолютно всех из таблицы ```employees``` : ```SELECT * FROM employee```. Для этого создадим функцию ```ReadAllEmployees``` ,а также определим структуру ```Employee```.
```
//ReadAllEmployees ...
func ReadAllEmployees(w http.ResponseWriter, r *http.Request) {
	var query = "SELECT * FROM employee"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("error while query SELECT executing:", err)
		return
	}
	employees := []Employee{}
	for rows.Next() {
		var uid uint
		var name string
		err = rows.Scan(&uid, &name)
		if err != nil {
			log.Println("error while scanning values from SELECT query:", err)
			continue
		}
		employee := Employee{ID: uid, Name: name}
		employees = append(employees, employee)
	}
	json.NewEncoder(w).Encode(employees)
}
```

### Шаг 7. Реализация обновления записи
Реализуем простейший функционал обновления информации про сотрудника.
```
//UpdateEmployee ...
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	//Считываем id из запроса
	vars := mux.Vars(r)
	id := vars["id"]
	//получим имя для обновления
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Println("ready to update name...")
		stmt, err := db.Prepare("UPDATE employee SET name=? WHERE uid=?")
		if err != nil {
			log.Println("error while praparing UPDATE query:", err)
			return
		}
		_, err = stmt.Exec(name[0], id)
		if err != nil {
			log.Println("error while executing UPDATE query:", err)
			return
		}
		fmt.Fprintf(w, "Info about emplyee updated successfully!")
	} else {
		fmt.Fprintf(w, "can not find parameters in put request")
	}
}
```

### Шаг 8. Удаление записи из таблицы
```
//DeleteEmployee ...
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Println("ready to delete from table employee")
		stmt, err := db.Prepare("DELETE FROM employee WHERE name=?")
		if err != nil {
			log.Println("error while preparing DELETE query:", err)
			return
		}
		_, err = stmt.Exec(name[0])
		if err != nil {
			log.Println("error while executing DELETE query:", err)
			return
		}
		fmt.Fprintf(w, "employee was deleted successfully")

	} else {
		fmt.Fprintf(w, "can not find parameters in delete request")
	}
}
```

### Шаг 9. Резюме
Теперь у нас реализован полноценный функционал ```CrUD```:
* ```Cr``` - ```CREATE```
* ```U``` - ```Update```
* ```D``` - ```Delete```

Говорят, что на ***сервере определена модель данных*** , если существует как минимум 1 объект , имеющий связь с ***базой данных***, а так же имеющий правила ***взаимодействия с базой данных (реализован CrUD)***

### Шаг 10. А где адаптер?
```
_ "github.com/go-sql-driver/mysql" // Явное определение адаптера
```
