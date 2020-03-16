package client

var address = "127.0.0.1:50200"

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

func RegistergRPCServerAddress(ip string) {
	address = ip
}

func GetgRPCServerAddress() string {
	return address
}
