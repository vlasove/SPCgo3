package main

import (
	_ "project/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}

