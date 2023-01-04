package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          string `json:"id" validate:"required,min=3,max=10"`
	Title       string `json:"title" validate:"required,max=50"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required,oneof=todo doing done"`
	CreatedAt   int64  `json:"created_at"`
	DueDate     int64  `json:"due_date"`
	OwnerID     string `json:"owner_id"`
	AssigneeID  string `json:"assignee_id"`
}
