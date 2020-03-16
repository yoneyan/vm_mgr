package data

import (
	"github.com/gin-gonic/gin"
	"github.com/yoneyan/vm_mgr/ggate/client"
	"log"
)

func GetUserVM(c *gin.Context) {
	log.Println("------GetUserVM------")

	token := GetToken(c.Request.Header.Get("Authorization"))
	data := client.GetUserVMClient(token)

	c.JSON(200, data)
}
