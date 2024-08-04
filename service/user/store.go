package user

import (
	"database/sql"

	"github.com/xyugen/realtime-chat-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
}
