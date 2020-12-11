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
	this.Ctx.WriteString("You successfully logged in!\n")
}

func (this *SessionController) Logout() {
	this.SetSession("authenticated", false)
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.WriteString("You successfully logged out!\n")
}

func (this *SessionController) HomeAfterLogin() {

	isAuth := this.GetSession("authenticated")
	if isAuth == nil || isAuth == false {
		this.Ctx.WriteString("You are unauthorized for this page!\n")
		return
	}

	this.TplName = "home.html"
	//this.Ctx.WriteString("Welcome to this HomePage analog!\n")
}
