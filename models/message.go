package models

import "github.com/google/uuid"

type Message struct {
	ID         uuid.UUID `gorm:"primary_key"`
	SenderID   uuid.UUID `gorm:"not null"`
	ReceiverID uuid.UUID `gorm:"not null"`
	Content    string `gorm:"not null"`

	Sender   User `gorm:"foreignkey:SenderID"`   // Association with the sender user
	Receiver User `gorm:"foreignkey:ReceiverID"` 
}



