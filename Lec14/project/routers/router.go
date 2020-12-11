package routers

import (
	"project/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	beego.Router("/", &controllers.HomeController{},
		"get:HomePage")
	beego.Router("/employees", &controllers.FirstController{},
		"get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{},
		"get:GetDashboard")

	beego.Router("/home", &controllers.SessionController{},
		"get:HomeAfterLogin")
	beego.Router("/login", &controllers.SessionController{},
		"get:Login")
	beego.Router("/logout", &controllers.SessionController{},
		"get:Logout")
}
