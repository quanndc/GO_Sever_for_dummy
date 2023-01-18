package business

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/models"
)

type AuthBusiness struct {
	server *core.Server
}

func NewAuthBusiness(s *core.Server) *AuthBusiness {
	return &AuthBusiness{
		server: s,
	}
}

func (b AuthBusiness) Register(regUser *models.UserRegistraionRequest) error {
	newUser := &auth.UserToCreate{}
	newUser.Email("")
	newUser.Password("")
	_, err := b.server.AuthClient.CreateUser(context.Background(), newUser)

	return err
}
