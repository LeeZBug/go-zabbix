package models

type ZabbixAuthInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type ZabbixAuthPlayload struct {
	Jsonrpc string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  ZabbixAuthInfo `json:"params"`
	Id      int            `json:"id"`
}
