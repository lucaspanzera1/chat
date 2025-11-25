package client

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/lucaspanzera1/chat/internal/models"
	"github.com/lucaspanzera1/chat/internal/repository"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type HubInterface interface {
	BroadcastLeave(username string)
}

type Client struct {
	Hub      HubInterface
	Conn     *websocket.Conn
	Send     chan models.Message
	Username string
	UserID   string
}

func NewClient(hub HubInterface, conn *websocket.Conn, username string) *Client {
	return &Client{
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan models.Message, 256),
		Username: username,
	}
}

func (c *Client) ReadPump(broadcast chan<- models.Message, unregister chan<- *Client, messageRepo *repository.MessageRepository) {
	defer func() {
		if c.Hub != nil {
			c.Hub.BroadcastLeave(c.Username)
		}

		leaveMsg := models.Message{
			ID:        uuid.New().String(),
			Username:  "Sistema",
			Content:   c.Username + " saiu do chat",
			Timestamp: time.Now(),
			Type:      "leave",
		}
		messageRepo.Create(context.Background(), &leaveMsg, c.UserID)

		unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var incoming struct {
			Content string `json:"content"`
		}
		if err := json.Unmarshal(data, &incoming); err != nil {
			log.Printf("error unmarshaling: %v", err)
			continue
		}

		msg := models.Message{
			ID:        uuid.New().String(),
			Username:  c.Username,
			Content:   incoming.Content,
			Timestamp: time.Now(),
			Type:      "message",
		}

		if err := messageRepo.Create(context.Background(), &msg, c.UserID); err != nil {
			log.Printf("Erro ao salvar mensagem: %v", err)
		}

		broadcast <- msg
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteJSON(message); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
