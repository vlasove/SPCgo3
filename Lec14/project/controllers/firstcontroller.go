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

func (this *FirstController) GetDashboard() {
	this.Data["employees"] = employees
	this.TplName = "dashboard.html"
}
