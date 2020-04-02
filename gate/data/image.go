package data

import (
	"github.com/gin-gonic/gin"
	"github.com/yoneyan/vm_mgr/gate/client"
	"log"
)

func GetAllImage(c *gin.Context) {
	log.Println("------GetImage------")

	token := GetToken(c.Request.Header.Get("Authorization"))
	result := client.GetAllImage(token)

	c.JSON(200, result)

}
