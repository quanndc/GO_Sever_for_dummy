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

	// db.AutoMigrate(&models.Task{})

	// db.Create(&models.Task{

	// 	Title:       "Task 1",
	// 	Description: "Description 1",
	// 	Status:      "Done",
	// })

	return &Database{
		Db: db,
	}, nil
}

// func NewTask(	) {
// 	db, _ := sql.Open("sqlite", "./test.db")
// 	statement, _ := db.Prepare(create)
// 	statement.Exec()

// }
