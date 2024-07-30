package zabbix

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go-zabbix/service"
)

func GetInfo(ctx iris.Context) {
	address := ctx.URLParam("address")
	fmt.Println(address)
	service.ZabbixserviceGetinfo(address)
}
