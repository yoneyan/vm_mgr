package data

import (
	"fmt"
	"strings"
)

func GetToken(token string) string {
	d := strings.Split(token, " ")

	if d[0] == "Bearer" {
		fmt.Println("RestAPI Token: " + d[1])
		return d[1]
	}
	return ""
}
