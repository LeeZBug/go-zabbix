package service

import (
	"encoding/json"
	"fmt"
	"go-zabbix/config"
	"go-zabbix/models"
	"io"
	"log"
	"strings"
)
import "net/http"

func ZabbixserviceGetinfo(address string) {
	zabbixAuthPlayload := &models.ZabbixAuthPlayload{
		Jsonrpc: "2.0",
		Method:  "user.login",
		Params: models.ZabbixAuthInfo{
			User:     config.ZabbixUser,
			Password: config.ZabbixPassword,
		},
		Id: 0,
	}
	buf, err := json.Marshal(zabbixAuthPlayload)
	if err != nil {
		fmt.Println(err)
	}
	reqBody := strings.NewReader(string(buf))
	res, err := http.Post(config.ZabbixApiUrl, "application/json", reqBody)
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
	}
	fmt.Printf("%s", body)
}
