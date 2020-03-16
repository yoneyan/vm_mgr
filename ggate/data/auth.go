package data

import (
	"github.com/gin-gonic/gin"
	"github.com/yoneyan/vm_mgr/ggate/client"
	"log"
)

type UserAuth struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func CheckToken(c *gin.Context) {
	log.Println("------CheckToken------")

	token := GetToken(c.Request.Header.Get("Authorization"))

	result := client.CheckTokenClient(token)

	c.JSON(200, result)
}

func DeleteToken(c *gin.Context) {
	log.Println("------DeleteToken------")

	token := GetToken(c.Request.Header.Get("Authorization"))

	result := client.DeleteTokenClient(token)

	c.JSON(200, result)
}

/*
func GetAllToken(c *gin.Context) {
	log.Println("------GetAllToken------")
	var result Result

	token := GetToken(c.Request.Header.Get("Authorization"))

	ok := client.DeleteTokenClient(token)
	if ok {
		result.Result = true
		result.Info = "OK"
	} else if token == "" {
		result.Result = false
	} else {
		result.Result = false
		result.Info = "Auth NG"
	}

	c.JSON(200, result)
}

*/

func GenerateToken(c *gin.Context) {
	log.Println("------GenerateToken------")

	var auth UserAuth
	c.BindJSON(&auth)

	result := client.GenerateTokenClient(auth.User, auth.Pass)

	c.JSON(200, result)
}
