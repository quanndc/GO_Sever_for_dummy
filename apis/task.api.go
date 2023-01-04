package apis

import (
	"net/http"

	"github.com/labstack/echo"
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/models"
)

func NewTaskApis(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)

	tasks := make([]*models.Task, 0)

	apis.POST("/", func(c echo.Context) error {
		task := &models.Task{}

		if err := c.Bind(task); err != nil {
			return echo.ErrBadRequest
		}

		if task.Title == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "title is required",
			})
		}

		if err := c.Validate(task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}
		tasks = append(tasks, task)
		return c.NoContent(http.StatusCreated)
	})

	apis.GET("/all", func(c echo.Context) error {
		return c.JSON(http.StatusOK, tasks)
	})

	apis.DELETE("/id:task_id", func(c echo.Context) error {
		task_id := c.Param("task_id")
		for i, task := range tasks {
			if task.ID == task_id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				return c.NoContent(http.StatusNoContent)
			}
		}
		return c.NoContent(http.StatusNotFound)
	})

	return apis

}
