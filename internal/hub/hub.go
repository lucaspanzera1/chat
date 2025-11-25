package hub

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucaspanzera1/chat/internal/client"
	"github.com/lucaspanzera1/chat/internal/models"
)

type Hub struct {
	Clients    map[*client.Client]bool
	Broadcast  chan models.Message
	Register   chan *client.Client
	Unregister chan *client.Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*client.Client]bool),
		Broadcast:  make(chan models.Message),
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
	}
}

func (h *Hub) ClientCount() int {
	return len(h.Clients)
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.broadcastCount()

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				h.broadcastCount()
			}

		case message := <-h.Broadcast:
			message.OnlineCount = len(h.Clients)
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

func (h *Hub) broadcastCount() {
	countMsg := models.Message{
		Type:        "count",
		OnlineCount: len(h.Clients),
	}
	for client := range h.Clients {
		select {
		case client.Send <- countMsg:
		default:
		}
	}
}

func (h *Hub) BroadcastLeave(username string) {
	msg := models.Message{
		ID:        uuid.New().String(),
		Username:  "Sistema",
		Content:   username + " saiu do chat",
		Timestamp: time.Now(),
		Type:      "leave",
	}
	h.Broadcast <- msg
}
