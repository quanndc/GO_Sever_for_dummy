package apis

import (
	"github.com/labstack/echo"
	"itss.edu.vn/todo/services/core"
)

func NewHealthyAPI(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)

	apis.GET("/", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	return apis
}
