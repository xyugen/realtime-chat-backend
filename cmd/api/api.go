package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xyugen/realtime-chat-backend/service/conversation"
	"github.com/xyugen/realtime-chat-backend/service/user"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	conversationStore := conversation.NewStore(s.db)
	conversationHandler := conversation.NewHandler(conversationStore, userStore)
	conversationHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
