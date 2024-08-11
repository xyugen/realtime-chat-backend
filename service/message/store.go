package message

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

func (s *Store) CreateMessage(message types.Message) error {
	result := s.db.Create(&message)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetMessagesByConversationId(conversationId int) ([]types.Message, error) {
	var messages []types.Message
	result := s.db.
		Preload("Sender").
		Where("conversation_id = ?", conversationId).
		Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}
