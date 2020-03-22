package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yoneyan/vm_mgr/ggate/data"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(CORS)

	//token
	router.POST("/api/v1/token", data.GenerateToken)
	router.POST("/api/v1/token/check", data.CheckToken)
	router.DELETE("/api/v1/token", data.DeleteToken)
	//router.HandleFunc("/api/v1/token",data.GetUserVM).Methods("GET")
	//vm
	router.PUT("/api/v1/vm", data.CreateVM)
	router.DELETE("/api/v1/vm/:id", data.DeleteVM)
	router.GET("/api/v1/vm", data.GetUserVM)
	router.GET("/api/v1/vm/:id", data.GetVM)
	router.PUT("/api/v1/vm/:id/power", data.StartVM)
	router.DELETE("/api/v1/vm/:id/power", data.StopVM)
	router.PUT("/api/v1/vm/:id/reset", data.ResetVM)
	router.PUT("/api/v1/vm/:id/pause", data.ResumeVM)
	router.DELETE("/api/v1/vm/:id/pause", data.PauseVM)
	//image
	router.GET("/api/v1/image", data.GetAllImage)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func CORS(c *gin.Context) {

	//c.Header("Access-Control-Allow-Headers", "Accept, Content-ID, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-ID", "application/json")
	c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
