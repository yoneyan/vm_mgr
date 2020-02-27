package data

import "strconv"

func GenerateAddress(ip string, port int) string {
	return ip + ":" + strconv.Itoa(port)
}
