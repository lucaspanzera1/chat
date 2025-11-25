package models

import "time"

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"` // "general", "private" ou "group"
	Users     []string  `json:"users"`
	CreatedBy string    `json:"createdBy,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type RoomUser struct {
	RoomID   string `json:"roomId"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

type CreateGroupRequest struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"userIds"` // IDs dos usuários a adicionar (mínimo 2)
}
