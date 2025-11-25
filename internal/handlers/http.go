package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/lucaspanzera1/chat/internal/auth"
	"github.com/lucaspanzera1/chat/internal/models"
	"github.com/lucaspanzera1/chat/internal/repository"
)

type HTTPHandler struct {
	messageRepo *repository.MessageRepository
	roomRepo    *repository.RoomRepository
	userRepo    *repository.UserRepository
}

func NewHTTPHandler(messageRepo *repository.MessageRepository, roomRepo *repository.RoomRepository, userRepo *repository.UserRepository) *HTTPHandler {
	return &HTTPHandler{
		messageRepo: messageRepo,
		roomRepo:    roomRepo,
		userRepo:    userRepo,
	}
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
		log.Printf("Erro ao buscar mensagens: %v", err)
		http.Error(w, "Erro ao buscar mensagens", http.StatusInternalServerError)
		return
	}

	if messages == nil {
		messages = []models.Message{}
	}

	json.NewEncoder(w).Encode(messages)
}

func (h *HTTPHandler) GetRoomHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	roomID := r.URL.Query().Get("roomId")
	if roomID == "" {
		roomID = "00000000-0000-0000-0000-000000000001"
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	messages, err := h.messageRepo.GetRecentByRoom(r.Context(), roomID, limit)
	if err != nil {
		log.Printf("Erro ao buscar mensagens: %v", err)
		http.Error(w, "Erro ao buscar mensagens", http.StatusInternalServerError)
		return
	}

	if messages == nil {
		messages = []models.Message{}
	}

	json.NewEncoder(w).Encode(messages)
}

func (h *HTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Token não fornecido", http.StatusUnauthorized)
		return
	}

	claims, err := auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	users, err := h.roomRepo.GetAllUsers(r.Context(), claims.UserID)
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *HTTPHandler) CreatePrivateRoom(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		log.Println("CreatePrivateRoom: Token não fornecido")
		http.Error(w, "Token não fornecido", http.StatusUnauthorized)
		return
	}

	claims, err := auth.ValidateToken(token)
	if err != nil {
		log.Printf("CreatePrivateRoom: Erro ao validar token: %v", err)
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	var req struct {
		OtherUserID string `json:"otherUserId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("CreatePrivateRoom: Erro ao decodificar body: %v", err)
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	log.Printf("GetOrCreatePrivateRoom: Buscando/criando sala entre %s e %s", claims.UserID, req.OtherUserID)

	room, err := h.roomRepo.GetOrCreatePrivateRoom(r.Context(), claims.UserID, req.OtherUserID)
	if err != nil {
		log.Printf("GetOrCreatePrivateRoom: Erro: %v", err)
		http.Error(w, "Erro ao buscar/criar sala privada: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("GetOrCreatePrivateRoom: Retornando sala: %s", room.ID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}
