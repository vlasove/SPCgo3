package models

//Employee ...
type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

//Инициализация хранилища сотрудников
var employees Employees
