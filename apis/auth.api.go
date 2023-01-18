package apis

import (
	"net/http"

	"github.com/labstack/echo"
	"itss.edu.vn/todo/services/business"
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/models"
)

func NewAuthApi(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)

	authBusiness := business.NewAuthBusiness(server)

	apis.POST("/register", func(c echo.Context) error {
		var register models.UserRegistraionRequest
		if err := c.Bind(&register); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := c.Validate(register); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := authBusiness.Register(&register); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, "OK")
	})

	return apis
}
