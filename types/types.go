package types

import (
	"gorm.io/gorm"
)

// User
type UserStore interface {
	CreateUser(user User) error
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// Conversation
type ConversationStore interface {
	CreateConversation(conversation Conversation) error
	GetConversationsByUserId(userId int) (*Conversation, error)
	GetConversationByUserIds(user1Id int, user2Id int) (*Conversation, error)
}

type Conversation struct {
	gorm.Model
	User1ID int `json:"user1_id"`
	User2ID int `json:"user2_id"`
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
	User1ID int `json:"user1Id" validate:"required,nefield=User2ID"`
	User2ID int `json:"user2Id" validate:"required"`
}
