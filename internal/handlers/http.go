package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lucaspanzera1/chat/internal/models"
	"github.com/lucaspanzera1/chat/internal/repository"
)

type HTTPHandler struct {
	messageRepo *repository.MessageRepository
}

func NewHTTPHandler(messageRepo *repository.MessageRepository) *HTTPHandler {
	return &HTTPHandler{messageRepo: messageRepo}
}

func (h *HTTPHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	messages, err := h.messageRepo.GetRecent(r.Context(), limit)
	if err != nil {
		http.Error(w, "Erro ao buscar mensagens", http.StatusInternalServerError)
		return
	}

	if messages == nil {
		messages = []models.Message{}
	}

	json.NewEncoder(w).Encode(messages)
}
