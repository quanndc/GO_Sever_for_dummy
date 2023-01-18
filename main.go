package main

import (
	_ "github.com/mattn/go-sqlite3"
	"itss.edu.vn/todo/services/apis"
	"itss.edu.vn/todo/services/core"
	"itss.edu.vn/todo/services/utilities"
)

// func searchTask(db *gorm.DB, search string) []models.Task {
// 	rows, _ := db.Distinct(("SELECT id, title FROM tasks WHERE title LIKE ?", "%"+search+"%"))
// 	defer rows.Close()

// 	err := rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	var tasks []models.Task
// 	for rows.Next() {
// 		var task models.Task
// 		rows.Scan(&task.ID, &task.Title)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return tasks
// }

func main() {

	server := core.NewServer()

	server.Echo.Validator = utilities.NewValidator()

	_ = apis.NewHealthyAPI("/healthy", server)

	_ = apis.NewTaskApis("/tasks", server)
	server.Start()

}
