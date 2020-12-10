package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vlasove/course/Lec10/models"
)

var employees models.Employees

func init() {
	employees = models.Employees{
		models.Employee{Id: "1", FirstName: "Bob", LastName: "Jack"},
		models.Employee{Id: "2", FirstName: "Alice", LastName: "Tompson"},
		models.Employee{Id: "3", FirstName: "George", LastName: "Lighter"},
	}
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, employee := range employees {
		if employee.Id == id {
			if err := json.NewEncoder(w).Encode(employee); err != nil {
				log.Println("error getting employee by id::", err)
			}
		}
	}
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Println("error while parsing POST-body:", err)
		return
	}
	log.Println("POST body successeded parsing")
	employees = append(employees, models.Employee{
		Id:        employee.Id,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
	})
	json.NewEncoder(w).Encode(employees)

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Println("error while parsing body PUT method:", err)
		return
	}

	var isUpsert = true
	for idx, empl := range employees {
		if empl.Id == employee.Id {
			isUpsert = false
			log.Println("found employee with that ID ... updating")
			employees[idx].FirstName = employee.FirstName
			employees[idx].LastName = employee.LastName
			break
		}
	}
	if isUpsert {
		employees = append(employees, models.Employee{
			Id:        employee.Id,
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
		})
	}
	json.NewEncoder(w).Encode(employees)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	//Удаляем по Id из тела запроса.
}

// sudo apt install npm
