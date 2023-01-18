package apis

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/models"
)

func NewTaskApis(endpoint string, server *core.Server) *echo.Group {
	apis := server.Echo.Group(endpoint)

	tasks := make([]*models.Task, 0)

	apis.POST("/", func(c echo.Context) error {
		task := &models.Task{}
		db, _ := core.NewDatabase()
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
		task.ID = uuid.New().String()
		task.CreatedAt = time.Now()

		fmt.Fprintln(c.Response().Writer, task.ID)
		tasks = append(tasks, task)
		db.Db.Create(task)
		return c.NoContent(http.StatusCreated)
	})

	apis.GET("/all", func(c echo.Context) error {
		return c.JSON(http.StatusOK, tasks)
	})

	apis.DELETE("/id:task_id", func(c echo.Context) error {
		task_id := c.Param("task_id")
		fmt.Fprintln(c.Response().Writer, task_id)
		for i, task := range tasks {
			if task.ID == task_id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				return c.NoContent(http.StatusOK)
			}
		}
		return c.NoContent(http.StatusNotFound)
	})

	return apis

}
