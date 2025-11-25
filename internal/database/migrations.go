package database

import (
	"context"
	"log"
)

func RunMigrations() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS messages (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			username VARCHAR(50) NOT NULL,
			content TEXT NOT NULL,
			type VARCHAR(20) DEFAULT 'message',
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_messages_user_id ON messages(user_id)`,
		`CREATE TABLE IF NOT EXISTS rooms (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(100),
			type VARCHAR(20) NOT NULL DEFAULT 'general',
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS room_users (
			room_id UUID REFERENCES rooms(id) ON DELETE CASCADE,
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			joined_at TIMESTAMP DEFAULT NOW(),
			PRIMARY KEY (room_id, user_id)
		)`,
		`ALTER TABLE messages ADD COLUMN IF NOT EXISTS room_id UUID REFERENCES rooms(id) ON DELETE CASCADE`,
		`CREATE INDEX IF NOT EXISTS idx_messages_room_id ON messages(room_id)`,
		`INSERT INTO rooms (id, name, type) VALUES ('00000000-0000-0000-0000-000000000001', 'General', 'general') ON CONFLICT DO NOTHING`,
	}

	for _, query := range queries {
		if _, err := DB.Exec(context.Background(), query); err != nil {
			return err
		}
	}

	log.Println("✓ Migrações executadas com sucesso")
	return nil
}
