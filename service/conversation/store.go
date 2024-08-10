package conversation

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

func (s *Store) CreateConversation(conversation types.Conversation) error {
	result := s.db.Create(&conversation)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetConversationsByUserId(userId int, username string) ([]types.Conversation, error) {
	var conversations []types.Conversation
	var result *gorm.DB
	if username == "" {
		result = s.db.
			Joins("User1").
			Joins("User2").
			Where("user1_id = ? OR user2_id = ?", userId, userId).
			Find(&conversations)
	} else {
		usernamePattern := "%" + username + "%"
		result = s.db.
			Joins("User1").
			Joins("User2").
			Where("(user1_id = ? AND user2.username LIKE ?) OR (user2_id = ? AND user1.username LIKE ?)", userId, usernamePattern, userId, usernamePattern).
			Find(&conversations)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return conversations, nil
}

func (s *Store) GetConversationByUserIds(user1Id int, user2Id int) (*types.Conversation, error) {
	var conversation types.Conversation
	result := s.db.
		Preload("User1").
		Preload("User2").
		Where("user1_id = ? AND user2_id = ?", user1Id, user2Id).Or("user1_id = ? AND user2_id = ?", user2Id, user1Id).
		First(&conversation)
	if result.Error != nil {
		return nil, result.Error
	}

	return &conversation, nil
}

func (s *Store) GetConversationById(conversationId int) (*types.Conversation, error) {
	var conversation types.Conversation
	result := s.db.
		Preload("User1").
		Preload("User2").
		Where("id = ?", conversationId).
		First(&conversation)
	if result.Error != nil {
		return nil, result.Error
	}

	return &conversation, nil
}
