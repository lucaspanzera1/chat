package repository

import (
	"context"
	"log"

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
	query := `INSERT INTO messages (id, room_id, user_id, username, content, type, created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.Exec(ctx, query, msg.ID, msg.RoomID, userID, msg.Username, msg.Content, msg.Type, msg.Timestamp)
	return err
}

func (r *MessageRepository) GetRecent(ctx context.Context, limit int) ([]models.Message, error) {
	query := `SELECT id, COALESCE(room_id, '00000000-0000-0000-0000-000000000001'), username, content, type, created_at
			  FROM messages 
			  WHERE room_id = '00000000-0000-0000-0000-000000000001' OR room_id IS NULL
			  ORDER BY created_at DESC 
			  LIMIT $1`

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		log.Printf("Erro na query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.Username, &msg.Content, &msg.Type, &msg.Timestamp); err != nil {
			log.Printf("Erro ao fazer scan: %v", err)
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Erro ao iterar rows: %v", err)
		return nil, err
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

func (r *MessageRepository) GetRecentByRoom(ctx context.Context, roomID string, limit int) ([]models.Message, error) {
	query := `SELECT id, COALESCE(room_id, '00000000-0000-0000-0000-000000000001'), username, content, type, created_at
			  FROM messages 
			  WHERE room_id = $1
			  ORDER BY created_at DESC 
			  LIMIT $2`

	rows, err := r.db.Query(ctx, query, roomID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.RoomID, &msg.Username, &msg.Content, &msg.Type, &msg.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}
