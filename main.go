package main

import "github.com/kataras/iris/v12"
import "go-zabbix/routes"

func main() {
	app := iris.New()
	routes.RegisterZabbixRoutes(app)

	app.Listen(":8080")
}
