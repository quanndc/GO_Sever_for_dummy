package apis

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/middlewares"
	"itss.edu.vn/todo/services/models"
)

func NewUserApis(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)
	users := make([]*models.User, 0)

	apis.Use(middlewares.NewAuthMiddleware())

	apis.POST("/", func(c echo.Context) error {
		user := &models.User{}
		if err := c.Bind(user); err != nil {
			return err
		}
		if err := c.Validate(user); err != nil {
			return err
		}
		if user.Username == "" {
			return c.JSON(400, map[string]string{
				"message": "username is required",
			})
		}
		if user.HPassword == "" {
			return c.JSON(400, map[string]string{
				"message": "password is required",
			})
		}
		user.ID = uuid.New().String()
		fmt.Fprintln(c.Response().Writer, user.ID)
		users = append(users, user)
		return c.NoContent(http.StatusCreated)
	})

	apis.GET("/all", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})

	apis.GET("/id:id", func(c echo.Context) error {
		id := c.Param("id")
		fmt.Fprintln(c.Response().Writer, id)
		for _, user := range users {
			if user.ID == id {
				return c.JSON(http.StatusOK, user)
			}
		}
		return c.NoContent(http.StatusNotFound)
	})

	apis.PUT("/id:id", func(c echo.Context) error {
		id := c.Param("id")
		fmt.Fprintln(c.Response().Writer, id)
		for _, user := range users {
			if user.ID == id {
				user.Username = c.FormValue("username")
				user.HPassword = c.FormValue("password")
				return c.NoContent(http.StatusOK)
			}
		}
		return c.NoContent(http.StatusNotFound)
	})

	return apis
}
