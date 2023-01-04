package core

import (
	"os"

	"github.com/labstack/echo"
)

type Server struct {
	Echo *echo.Echo
	Db   *Database
}

func NewServer() *Server {

	db, err := NewDatabase()
	if err != nil {
		panic(err)
	}
	server := &Server{
		Echo: echo.New(),
		Db:   db,
	}
	return server
}

func (s Server) Start() error {
	return s.Echo.Start(os.Getenv("HOST"))
}
