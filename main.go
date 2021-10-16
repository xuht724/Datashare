package main

import (
	"goProject/models/user"
	_ "goProject/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*", "X-Requested-With", "Content-Type", "Referer", "User-Agent"},
		ExposeHeaders:    []string{"*", "Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Referer", "User-Agent", "X-Requested-With", "Referer", "User-Agent"},
		AllowCredentials: true,
	}))

	// log.Init()
	user.Init()
	beego.Run()
}
