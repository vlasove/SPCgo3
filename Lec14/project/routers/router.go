package routers

import (
	"project/controllers"

	beego "github.com/astaxie/beego/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{},
		"get:GetEmployees")
}
