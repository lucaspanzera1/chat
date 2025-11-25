package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Anonymous"
	}

	user, err := wsh.userRepo.Create(context.Background(), username)
	if err != nil {
		log.Printf("Erro ao criar usuário: %v", err)
		http.Error(w, "Erro ao autenticar", http.StatusInternalServerError)
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
		Username: username,
		UserID:   user.ID,
	}

	wsh.hub.Register <- c

	joinMsg := models.Message{
		ID:        uuid.New().String(),
		Username:  "Sistema",
		Content:   username + " entrou no chat",
		Timestamp: time.Now(),
		Type:      "join",
	}
	wsh.messageRepo.Create(context.Background(), &joinMsg, user.ID)
	wsh.hub.Broadcast <- joinMsg

	go c.WritePump()
	go c.ReadPump(wsh.hub.Broadcast, wsh.hub.Unregister, wsh.messageRepo)
}
