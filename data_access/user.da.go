package dataaccess

import (
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/models"
)

type UserDataAccess struct {
	server *core.Server
}

func NewUserDataAccess(server *core.Server) *UserDataAccess {
	return &UserDataAccess{
		server: server,
	}
}

func (d UserDataAccess) GetById(id string) (*models.User, error) {
	user := &models.User{}
	if err := d.server.Db.Db.First(user, map[string]interface{}{
		"id": id,
	}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d UserDataAccess) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}
	if err := d.server.Db.Db.Where("username = ? ", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
