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
	SearchUser(username string) ([]User, error)
}

type User struct {
	Base
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
}

// Conversation
type ConversationStore interface {
	CreateConversation(conversation Conversation) error
	GetConversationsByUserId(userId int, username string) ([]Conversation, error)
	GetConversationByUserIds(user1Id int, user2Id int) (*Conversation, error)
	GetConversationByIDAndUserID(conversationId int, userId int) ([]Conversation, error)
	GetConversationById(conversationId int) (*Conversation, error)
}

type Conversation struct {
	Base
	User1ID int  `json:"user1Id" gorm:"index:idx_user1_user2"`
	User1   User `json:"user1" gorm:"foreignKey:User1ID;references:ID"`
	User2ID int  `json:"user2Id" gorm:"index:idx_user1_user2"`
	User2   User `json:"user2" gorm:"foreignKey:User2ID;references:ID"`
}

// Message
type MessageStore interface {
	CreateMessage(message Message) error
	// GetMessagesByConversationId(conversationId int) ([]Message, error)
}

type Message struct {
	Base
	ConversationID int    `json:"conversationId" gorm:"index,not null"`
	Content        string `json:"content"`
	SenderID       int    `json:"senderId" gorm:"index,not null"`
	Sender         User   `json:"sender" gorm:"foreignKey:SenderID;references:ID"`
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

type CreateMessagePayload struct {
	// ConversationId int    `json:"conversationId" validate:"required"`
	Content string `json:"content" validate:"required"`
	// SenderID       int    `json:"senderId" validate:"required"`
}
