package main

import (
	"itss.edu.vn/todo/services/apis"
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/utilities"
)

func main() {
	server := core.NewServer()

	server.Echo.Validator = utilities.NewValidator()

	_ = apis.NewHealthyAPI("/healthy", server)

	_ = apis.NewTaskApis("/tasks", server)
	server.Start()

}
