package data

import (
	"github.com/gin-gonic/gin"
	"github.com/yoneyan/vm_mgr/ggate/client"
	"log"
)

type ShutdownData struct {
	Force bool `json:"force"`
}

func GetUserVM(c *gin.Context) {
	log.Println("------GetUserVM------")

	token := GetToken(c.Request.Header.Get("Authorization"))
	data := client.GetUserVMClient(token)

	c.JSON(200, data)
}

/*
func GetVM(c *gin.Context){
	log.Println("------GetVM------")

	id := c.Param("id")
	token := GetToken(c.Request.Header.Get("Authorization"))

}

*/

func StartVM(c *gin.Context) {
	log.Println("------GetUserVM------")

	id := c.Param("id")
	token := GetToken(c.Request.Header.Get("Authorization"))

	r := client.StartVMClient(token, id)

	c.JSON(200, r)

}

func StopVM(c *gin.Context) {
	log.Println("------StopAndShutdownVM------")

	var data ShutdownData
	c.BindJSON(&data)

	token := GetToken(c.Request.Header.Get("Authorization"))

	var r client.Result

	id := c.Param("id")
	if data.Force {
		r = client.StopVMClient(token, id)
	} else {
		r = client.ShutdownVMClient(token, id)
	}
	c.JSON(200, r)

}

func ResetVM(c *gin.Context) {
	log.Println("------ResetVM------")

	id := c.Param("id")
	token := GetToken(c.Request.Header.Get("Authorization"))

	r := client.ResetVMClient(token, id)

	c.JSON(200, r)
}

func PauseVM(c *gin.Context) {
	log.Println("------ResetVM------")

	id := c.Param("id")
	token := GetToken(c.Request.Header.Get("Authorization"))

	r := client.PauseVMClient(token, id)

	c.JSON(200, r)
}

func ResumeVM(c *gin.Context) {
	log.Println("------ResetVM------")

	id := c.Param("id")
	token := GetToken(c.Request.Header.Get("Authorization"))

	r := client.ResumeVMClient(token, id)

	c.JSON(200, r)
}
