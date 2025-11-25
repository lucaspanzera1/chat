package hub

import (
	"github.com/lucaspanzera1/chat/internal/models"
)

type ClientInterface interface {
	GetRoomID() string
	GetSendChannel() chan models.Message
}

type Hub struct {
	Rooms      map[string]map[ClientInterface]bool
	Broadcast  chan models.Message
	Register   chan ClientInterface
	Unregister chan ClientInterface
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]map[ClientInterface]bool),
		Broadcast:  make(chan models.Message),
		Register:   make(chan ClientInterface),
		Unregister: make(chan ClientInterface),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			roomID := client.GetRoomID()
			if h.Rooms[roomID] == nil {
				h.Rooms[roomID] = make(map[ClientInterface]bool)
			}
			h.Rooms[roomID][client] = true
			h.broadcastCountToRoom(roomID)

		case client := <-h.Unregister:
			roomID := client.GetRoomID()
			if clients, ok := h.Rooms[roomID]; ok {
				if _, exists := clients[client]; exists {
					delete(clients, client)
					close(client.GetSendChannel())
					h.broadcastCountToRoom(roomID)
				}
			}

		case message := <-h.Broadcast:
			if clients, ok := h.Rooms[message.RoomID]; ok {
				message.OnlineCount = len(clients)
				for client := range clients {
					select {
					case client.GetSendChannel() <- message:
					default:
						close(client.GetSendChannel())
						delete(clients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) broadcastCountToRoom(roomID string) {
	if clients, ok := h.Rooms[roomID]; ok {
		countMsg := models.Message{
			RoomID:      roomID,
			Type:        "count",
			OnlineCount: len(clients),
		}
		for client := range clients {
			select {
			case client.GetSendChannel() <- countMsg:
			default:
			}
		}
	}
}

func (h *Hub) BroadcastLeave(username string) {}
