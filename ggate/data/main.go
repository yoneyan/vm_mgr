package data

import (
	"fmt"
	"strings"
)

type Result struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
}

func GetToken(token string) string {
	d := strings.Split(token, " ")

	if d[0] == "Bearer" {
		fmt.Println("RestAPI Token: " + d[1])
		return d[1]
	}
	return ""
}
