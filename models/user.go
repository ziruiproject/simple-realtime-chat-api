package models

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID	`gorm:"primary_key;column:id"`
	Email string	`gorm:"not null;unique;column:email"`
	Name string		`gorm:"not null;column:name"`
	Password string	`gorm:"not null;column:password"`
	Profile string	`gorm:"not null;column:profile_img"`
	CreatedAt int	`gorm:"not null;column:created_at"`
	UpdatedAt int	`gorm:"not null;column:updated_at"`
}