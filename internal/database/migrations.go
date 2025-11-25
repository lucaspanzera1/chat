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
	}

	for _, query := range queries {
		if _, err := DB.Exec(context.Background(), query); err != nil {
			return err
		}
	}

	log.Println("✓ Migrações executadas com sucesso")
	return nil
}
