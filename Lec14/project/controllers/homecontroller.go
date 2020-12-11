package controllers

import beego "github.com/astaxie/beego/server/web"

//Аналог []Route из прошлый лекций
type HomeController struct {
	beego.Controller
}

func (this *HomeController) HomePage() {
	this.TplName = "home.html"
}
