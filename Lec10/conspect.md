## Лекция 10. Продолжение Лекции 9

***Задача*** реализовать ```PUT``` и ```DELETE``` запросы.

### Шаг 1. Реализация ```PUT``` запроса
В основной файле ```main.go``` определим новую пару отношений:
```
var routes = Routes{
	Route{
		"GetEmployees",
		"GET",
		"/employees",
		handlers.GetEmployees,
	},
	Route{
		"GetEmployee",
		"GET",
		"/employee/{id}",
		handlers.GetEmployee,
	},
	Route{
		"AddEmployee",
		"POST",
		"/employee/add",
		handlers.AddEmployee,
	},
	Route{
		"UpdateEmployee",
		"PUT",
		"/employee/update",
		handlers.UpdateEmployee,
	},
}
```

Теперь в ```handlers.go``` добавим новый обработчик.
```
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

```

### Шаг 2. Реализация ```DELETE``` запроса
Получает на вход информацию в виде ```{"Id" : "1", "firstName" : "Bob", "lastName" : "Jack"}``` если персонаж с таким ```Id``` присутствует в слайсе ```employees``` - удаляем его из слайса. В противном случае - ничего не делаем.
* В случае удаления вывести в консоль ```employee with that ID successfully deleted```
* В случае если не нашли кого удалять ```employee with that ID does not exist```

В ```main.go``` добавим новый ```Route```:
```
Route{
		"DeleteEmployee",
		"DELETE",
		"/employee/delete",
		handlers.DeleteEmployee,
	},
```
Должно получиться следующее:
* ```curl -H "Content-Type: application/json" -X DELETE -d '{"Id" : "1", "firstName" : "Bob", "lastName" : "Jack"}' http://localhost:8080/employee/delete``` -> ```employee with that ID successfully deleted```

* ```curl -H "Content-Type: application/json" -X DELETE -d '{"Id" : "200", "firstName" : "Bob", "lastName" : "Jack"}' http://localhost:8080/employee/delete``` -> ```employee with that ID does not exist```