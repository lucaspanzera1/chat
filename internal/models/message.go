package models

import "time"

type Message struct {
	ID          string    `json:"id"`
	RoomID      string    `json:"roomId"`
	Username    string    `json:"username"`
	AvatarURL   string    `json:"avatarUrl,omitempty"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
	Type        string    `json:"type"`
	OnlineCount int       `json:"onlineCount,omitempty"`
}
