package vnc

import (
	"fmt"
	"github.com/koding/websocketproxy"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func WebSocketProxy() func(echo.Context) error {
	return func(c echo.Context) error {
		port := c.Param("port")
		p, _ := strconv.Atoi(port)

		u := &url.URL{
			Scheme: "ws",
			Host:   fmt.Sprintf("%s:%d", "172.16.190.101", p),
			Path:   "/",
		}
		fmt.Println(u)

		ws := &websocketproxy.WebsocketProxy{
			Backend: func(r *http.Request) *url.URL {
				return u
			},
		}
		delete(c.Request().Header, "Origin")
		log.Printf("[DEBUG] websocket proxy requesting to backend '%s'", ws.Backend(c.Request()))
		ws.ServeHTTP(c.Response(), c.Request())

		return nil
	}
}
