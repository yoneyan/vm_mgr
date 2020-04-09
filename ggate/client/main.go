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

var use = false

func GetgRPCServerAddress() string {
	return etc.GetControllerIP()
}

func UseChange(u bool) {
	use = u
}

func UseStatus() bool {
	return use
}
