package conversation

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/xyugen/realtime-chat-backend/service/auth"
	"github.com/xyugen/realtime-chat-backend/types"
	"github.com/xyugen/realtime-chat-backend/utils"
)

type Handler struct {
	store     types.ConversationStore
	userStore types.UserStore
}

func NewHandler(store types.ConversationStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// router.HandleFunc("/conversation", h.handleCreateConversation).Methods("POST")
	router.HandleFunc("/conversation/new", auth.WithJWTAuth(h.handleCreateConversation, h.userStore)).Methods("POST")
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
	c, err := h.store.GetConversationsByUserId(1)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, c)
}
