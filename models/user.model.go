package models

type User struct {
	ID        string `json:"id" validate:"required,min=5,max=10" gorm:"primaryKey;"`
	Username  string `json:"username" validate:"required,max=50" gorm:"uniqueIndex;not null;size:50;"`
	HPassword string `json:"password" validate:"required" gorm:"not null;"`
}

type UserRegistraionRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
