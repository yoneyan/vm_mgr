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
	router.GET("/api/v1/vm", data.GetUserVM)
	//router.StrictSlash(true)
	//router.HandleFunc("/api/v1/vm/group", data.GetUserVM).Methods("GET")
	//router.HandleFunc("/api/v1/vm", data.GetUserVM).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CORS(c *gin.Context) {

	//c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
