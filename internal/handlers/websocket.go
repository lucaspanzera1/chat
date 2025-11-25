package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lucaspanzera1/chat/internal/auth"
	"github.com/lucaspanzera1/chat/internal/client"
	"github.com/lucaspanzera1/chat/internal/hub"
	"github.com/lucaspanzera1/chat/internal/models"
	"github.com/lucaspanzera1/chat/internal/repository"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSHandler struct {
	hub         *hub.Hub
	userRepo    *repository.UserRepository
	messageRepo *repository.MessageRepository
}

func NewWSHandler(h *hub.Hub, userRepo *repository.UserRepository, messageRepo *repository.MessageRepository) *WSHandler {
	return &WSHandler{
		hub:         h,
		userRepo:    userRepo,
		messageRepo: messageRepo,
	}
}

func (wsh *WSHandler) ServeWS(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	roomID := r.URL.Query().Get("roomId")

	if roomID == "" {
		roomID = "00000000-0000-0000-0000-000000000001" // Sala geral
	}

	if token == "" {
		http.Error(w, "Token não fornecido", http.StatusUnauthorized)
		return
	}

	claims, err := auth.ValidateToken(token)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	user, err := wsh.userRepo.GetByID(context.Background(), claims.UserID)
	if err != nil || user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &client.Client{
		Hub:      wsh.hub,
		Conn:     conn,
		Send:     make(chan models.Message, 256),
		Username: user.Username,
		UserID:   user.ID,
		RoomID:   roomID,
	}

	wsh.hub.Register <- c

	unregisterFunc := func(client *client.Client) {
		wsh.hub.Unregister <- client
	}

	go c.WritePump()
	go c.ReadPump(wsh.hub.Broadcast, unregisterFunc, wsh.messageRepo)
}
