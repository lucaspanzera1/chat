package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucaspanzera1/chat/internal/models"
)

type MessageRepository struct {
	db *pgxpool.Pool
}

func NewMessageRepository(db *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(ctx context.Context, msg *models.Message, userID string) error {
	query := `INSERT INTO messages (id, user_id, username, content, type, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(ctx, query, msg.ID, userID, msg.Username, msg.Content, msg.Type, msg.Timestamp)
	return err
}

func (r *MessageRepository) GetRecent(ctx context.Context, limit int) ([]models.Message, error) {
	query := `SELECT id, username, content, type, created_at 
			  FROM messages 
			  ORDER BY created_at DESC 
			  LIMIT $1`

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.Username, &msg.Content, &msg.Type, &msg.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}
