package routes

import "github.com/kataras/iris/v12"
import "go-zabbix/controllers/zabbix"

func RegisterZabbixRoutes(app *iris.Application) {
	zabbixAPI := app.Party("/zabbix")
	{
		zabbixAPI.Use(iris.Compression)
		zabbixAPI.Get("/", zabbix.GetInfo)
	}
}
