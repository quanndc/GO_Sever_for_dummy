package core

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo"

	"google.golang.org/api/option"
)

type Server struct {
	Echo       *echo.Echo
	Db         *Database
	Firebase   *firebase.App
	AuthClient *auth.Client
}

func NewServer() *Server {

	db, err := NewDatabase()
	if err != nil {
		panic(err)
	}

	opt := option.WithCredentialsFile("admin-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	server := &Server{
		Echo:       echo.New(),
		Db:         db,
		Firebase:   app,
		AuthClient: authClient,
	}

	return server
}

func (s Server) Start() error {
	return s.Echo.Start(os.Getenv("HOST"))
}
