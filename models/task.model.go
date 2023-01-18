package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          string    `json:"id" validate:"required,min=3,max=10" gorm:"primaryKey;"`
	Title       string    `json:"title" validate:"required,max=50" gorm:"not null;size:50;"`
	Description string    `json:"description"`
	Status      string    `json:"status" validate:"required,oneof=todo doing done" gorm:"not null;index;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	DueDate     time.Time `json:"due_date"`
	Owner       User      `json:"owner_id" gorm:"embedded;embeddedPrefix:user_"`
	AssigneeID  User      `json:"assignee_id" gorm:"embedded;embeddedPrefix:user_"`
}
