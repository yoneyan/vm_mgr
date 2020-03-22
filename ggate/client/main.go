package client

import "github.com/yoneyan/vm_mgr/ggate/etc"

type AuthResult struct {
	Result   bool   `json:"result"`
	Token    string `json:"token"`
	UserName string `json:"username"`
	UserID   int    `json:"userid"`
}

type Result struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
}

func GetgRPCServerAddress() string {
	return etc.GetControllerIP()
}
