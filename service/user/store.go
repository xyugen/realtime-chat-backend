package user

import (
	"github.com/xyugen/realtime-chat-backend/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
	result := s.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	var user types.User
	result := s.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	var user types.User
	result := s.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *Store) SearchUser(username string) ([]types.User, error) {
	var users []types.User
	result := s.db.Where("username LIKE ?", "%"+username+"%").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
