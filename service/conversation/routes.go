package conversation

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/xyugen/realtime-chat-backend/service/auth"
	"github.com/xyugen/realtime-chat-backend/types"
	"github.com/xyugen/realtime-chat-backend/utils"
)

type Handler struct {
	store        types.ConversationStore
	userStore    types.UserStore
	messageStore types.MessageStore
}

func NewHandler(store types.ConversationStore, userStore types.UserStore, messageStore types.MessageStore) *Handler {
	return &Handler{store: store, userStore: userStore, messageStore: messageStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/conversations", auth.WithJWTAuth(h.handleGetConversations, h.userStore)).Methods("GET")
	router.HandleFunc("/conversation/new", auth.WithJWTAuth(h.handleCreateConversation, h.userStore)).Methods("POST")

	router.HandleFunc("/conversation/{id}", auth.WithJWTAuth(h.handleGetConversation, h.userStore)).Methods("GET")

	router.HandleFunc("/conversation/{id}/messages", auth.WithJWTAuth(h.handleGetMessages, h.userStore)).Methods("GET")
	router.HandleFunc("/conversation/{id}/messages/new", auth.WithJWTAuth(h.handleCreateMessage, h.userStore)).Methods("POST")
}

func (h *Handler) handleGetConversation(w http.ResponseWriter, r *http.Request) {
	conversationId := mux.Vars(r)["id"]
	userId := auth.GetUserIDFromContext(r.Context())

	conversationIdInt, err := strconv.Atoi(conversationId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid conversation id"))
		return
	}

	// user must be part of convo
	if _, err := h.store.GetConversationByIDAndUserID(conversationIdInt, userId); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user is not part of conversation"))
		return
	}

	conversation, err := h.store.GetConversationById(conversationIdInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, conversation)
}

func (h *Handler) handleCreateConversation(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	// get conversation payload
	var conversation types.CreateConversationPayload
	if err := utils.ParseJSON(r, &conversation); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(conversation); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	if userID == conversation.User2ID {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("cannot create conversation with self"))
		return
	}

	// check if user exists
	if _, err := h.userStore.GetUserByID(conversation.User2ID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d does not exist", conversation.User2ID))
		return
	}

	// check if convo exists
	if _, err := h.store.GetConversationByUserIds(userID, conversation.User2ID); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("conversation already exists"))
		return
	}

	// create new convo
	if err := h.store.CreateConversation(types.Conversation{
		User1ID: userID,
		User2ID: conversation.User2ID,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return success
	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleGetConversations(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	username := r.URL.Query().Get("username")

	c, err := h.store.GetConversationsByUserId(userID, username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, c)
}

func (h *Handler) handleGetMessages(w http.ResponseWriter, r *http.Request) {
	conversationId := mux.Vars(r)["id"]
	conversationIdInt, err := strconv.Atoi(conversationId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid conversation id"))
		return
	}

	fmt.Println(conversationIdInt)
}

func (h *Handler) handleCreateMessage(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	conversationId := mux.Vars(r)["id"]
	conversationIdInt, err := strconv.Atoi(conversationId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid conversation id"))
		return
	}

	// get message payload
	var message types.CreateMessagePayload
	if err := utils.ParseJSON(r, &message); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(message); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if user exists
	if _, err := h.userStore.GetUserByID(userID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d does not exist", userID))
		return
	}

	// check if convo exists
	if _, err := h.store.GetConversationById(conversationIdInt); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("conversation with id %d does not exist", conversationIdInt))
		return
	}

	// user must be part of convo
	if _, err := h.store.GetConversationByIDAndUserID(conversationIdInt, userID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user is not part of conversation"))
		return
	}

	// create new message
	if err := h.messageStore.CreateMessage(types.Message{
		ConversationID: conversationIdInt,
		SenderID:       userID,
		Content:        message.Content,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return success
	utils.WriteJSON(w, http.StatusOK, nil)
}
