package models

import "time"

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`  // "general" ou "private"
	Users     []string  `json:"users"` // IDs dos usuários
	CreatedAt time.Time `json:"createdAt"`
}

type RoomUser struct {
	RoomID   string `json:"roomId"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
}
