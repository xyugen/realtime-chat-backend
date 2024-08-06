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
	GetConversationsByUserId(userId int) ([]Conversation, error)
	GetConversationByUserIds(user1Id int, user2Id int) (*Conversation, error)
}

type Conversation struct {
	gorm.Model
	User1ID int   `json:"user1_id" gorm:"index:idx_user1_user2,not null"`
	User1   *User `json:"user1" gorm:"foreignKey:User1ID"`
	User2ID int   `json:"user2_id" gorm:"index:idx_user1_user2,not null"`
	User2   *User `json:"user2" gorm:"foreignKey:User2ID"`
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
