package main

import (
	_ "go-pos/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

