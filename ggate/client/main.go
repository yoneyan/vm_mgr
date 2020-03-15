package client

var address = "127.0.0.1:50200"

func RegistergRPCServerAddress(ip string) {
	address = ip
}

func GetgRPCServerAddress() string {
	return address
}
