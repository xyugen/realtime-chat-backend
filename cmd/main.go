package main

import (
	"fmt"
	"log"

	"github.com/xyugen/realtime-chat-backend/cmd/api"
	"github.com/xyugen/realtime-chat-backend/config"
	"github.com/xyugen/realtime-chat-backend/db"
	"github.com/xyugen/realtime-chat-backend/types"
	"gorm.io/gorm"
)

func main() {
	db, err := db.NewMySQLiteStorage(config.Envs.TursoDatabaseUrl, config.Envs.TursoAuthToken)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *gorm.DB) {
	err := db.AutoMigrate(&types.User{}, &types.Conversation{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized")
}
