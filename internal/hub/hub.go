package hub

import (
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

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcast:
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
