package core

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func NewDatabase() (*Database, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Database{
		Db: db,
	}, nil
}
