package types

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

// User
type UserStore interface {
	CreateUser(user User) error
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

type User struct {
	Base
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
}

// Conversation
type ConversationStore interface {
	CreateConversation(conversation Conversation) error
	GetConversationsByUserId(userId int) ([]Conversation, error)
	GetConversationByUserIds(user1Id int, user2Id int) (*Conversation, error)
}

type Conversation struct {
	Base
	User1ID int `json:"user1Id" gorm:"index:idx_user1_user2"`
	User2ID int `json:"user2Id" gorm:"index:idx_user1_user2"`
}

// Payloads (separated for scale reasons)
type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginUserPayload struct {
	Username string `json:"username" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=6"`
}

type CreateConversationPayload struct {
	User2ID int `json:"user2Id" validate:"required"`
}
