package models

import "github.com/google/uuid"

type Message struct {
	ID uuid.UUID	`gorm:"primary_key;column:id"`
	UserId uuid.UUID `gorm:"column:user_id"`
	Content string `gorm:"not null;column:content"`
	User User `gorm:"foreignKey:user_id;references:id"`
}