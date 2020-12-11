package main

import (
	_ "project/routers"

	beego "github.com/astaxie/beego/server/web"
	_ "github.com/astaxie/beego/server/web/session/redis"
)

func main() {
	beego.Run()
}
