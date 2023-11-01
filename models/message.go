package models

import "github.com/google/uuid"

type Message struct {
	ID uuid.UUID	`gorm:"primary_key;column:id"`
	UserId uuid.UUID `gorm:"primary_key;column:id"`
}