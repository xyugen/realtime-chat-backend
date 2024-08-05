package types

import (
	"gorm.io/gorm"
)

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

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
