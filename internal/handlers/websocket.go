package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/lucaspanzera1/chat/internal/client"
	"github.com/lucaspanzera1/chat/internal/hub"
	"github.com/lucaspanzera1/chat/internal/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Anonymous"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &client.Client{
		Conn:     conn,
		Send:     make(chan models.Message, 256),
		Username: username,
	}

	h.Register <- c

	h.Broadcast <- models.Message{
		ID:        uuid.New().String(),
		Username:  "Sistema",
		Content:   username + " entrou no chat",
		Timestamp: time.Now(),
		Type:      "join",
	}

	go c.WritePump()
	go c.ReadPump(h.Broadcast, h.Unregister)
}
