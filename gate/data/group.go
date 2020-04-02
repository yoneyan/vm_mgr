package data

import (
	"github.com/gin-gonic/gin"
	"github.com/yoneyan/vm_mgr/gate/client"
	"log"
)

func GetGroup(c *gin.Context) {
	log.Println("------GetGroup------")

	token := GetToken(c.Request.Header.Get("Authorization"))
	result := client.GetGroup(token)

	c.JSON(200, result)

}
