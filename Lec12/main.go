package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Явное определение адаптера --- запуск init() и использование ПЕРЕМЕННЫХ ПАКЕТА

	"github.com/gorilla/mux"
)

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

//CreateEmployee ...
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	//Получение очереди
	vals := r.URL.Query()
	name, ok := vals["name"]
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
		fmt.Fprintf(w, "Last inseted id: %d", id)
	} else {
		fmt.Fprintf(w, "Can not parse arguments in this request")
	}
}

//Employee ...
type Employee struct {
	ID   uint   `json:"uid"`
	Name string `json:"name"`
}

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

func main() {
	router := mux.NewRouter()
	//Сразу выполним отложенное закрытие
	defer db.Close()
	router.HandleFunc("/", GetCurrentDB)
	// предеаем в качестве параметров на этапе /employee/create?name=bob
	router.HandleFunc("/employee/create", CreateEmployee).Methods("POST")
	// Получим абсолютно всех сотрудников из БД
	router.HandleFunc("/employees", ReadAllEmployees).Methods("GET")
	//обновляем строку в таблице про сотрудника с id
	//                  /employee/update/2?name=alice
	router.HandleFunc("/employee/update/{id}", UpdateEmployee).Methods("PUT")
	//удаление из таблицы
	//                    /employee/delete?name=bob
	router.HandleFunc("/employee/delete", DeleteEmployee).Methods("DELETE")
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting server", err)
		return
	}
}
