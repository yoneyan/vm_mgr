package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rakyll/statik/fs"
	"github.com/yoneyan/vm_mgr/wgate/vnc"
	_ "github.com/yoneyan/vm_mgr/wgate/web"
	"net/http"
)

func main() {
	statikFs, err := fs.New()
	if err != nil {
		fmt.Println("Error")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/:port/vnc", vnc.WebSocketProxy())
	e.GET("/console/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(statikFs))))

	e.Start("0.0.0.0:80")

}
