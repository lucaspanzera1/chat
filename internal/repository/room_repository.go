package repository

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucaspanzera1/chat/internal/models"
)

type RoomRepository struct {
	db *pgxpool.Pool
}

func NewRoomRepository(db *pgxpool.Pool) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) GetOrCreatePrivateRoom(ctx context.Context, user1ID, user2ID string) (*models.Room, error) {

	query := `SELECT DISTINCT r.id, r.type, r.created_at
			  FROM rooms r
			  INNER JOIN room_users ru1 ON ru1.room_id = r.id
			  INNER JOIN room_users ru2 ON ru2.room_id = r.id
			  WHERE r.type = 'private'
			  AND (
				  (ru1.user_id = $1 AND ru2.user_id = $2)
				  OR
				  (ru1.user_id = $2 AND ru2.user_id = $1)
			  )
			  LIMIT 1`

	room := &models.Room{}
	err := r.db.QueryRow(ctx, query, user1ID, user2ID).Scan(&room.ID, &room.Type, &room.CreatedAt)

	if err == nil {
		room.Users = []string{user1ID, user2ID}
		log.Printf("✓ Sala privada existente: %s", room.ID)
		return room, nil
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		log.Printf("✗ Erro ao buscar sala: %v", err)
		return nil, err
	}

	roomID := uuid.New().String()
	insertRoom := `INSERT INTO rooms (id, type) VALUES ($1, 'private') RETURNING id, type, created_at`

	err = r.db.QueryRow(ctx, insertRoom, roomID).Scan(&room.ID, &room.Type, &room.CreatedAt)
	if err != nil {
		log.Printf("✗ Erro ao criar sala: %v", err)
		return nil, err
	}

	insertUsers := `INSERT INTO room_users (room_id, user_id) VALUES ($1, $2), ($1, $3)`
	if _, err := r.db.Exec(ctx, insertUsers, roomID, user1ID, user2ID); err != nil {
		log.Printf("✗ Erro ao adicionar usuários: %v", err)
		return nil, err
	}

	room.Users = []string{user1ID, user2ID}
	log.Printf("✓ Nova sala privada criada: %s", room.ID)

	return room, nil
}

func (r *RoomRepository) GetUserRooms(ctx context.Context, userID string) ([]models.RoomUser, error) {
	query := `SELECT r.id, r.type, u.id, u.username
			  FROM rooms r
			  INNER JOIN room_users ru ON ru.room_id = r.id
			  LEFT JOIN room_users ru2 ON ru2.room_id = r.id AND ru2.user_id != $1
			  LEFT JOIN users u ON u.id = ru2.user_id
			  WHERE ru.user_id = $1
			  ORDER BY r.created_at DESC`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.RoomUser
	for rows.Next() {
		var ru models.RoomUser
		var roomType, otherUserID, otherUsername *string
		if err := rows.Scan(&ru.RoomID, &roomType, &otherUserID, &otherUsername); err != nil {
			return nil, err
		}
		if otherUsername != nil {
			ru.Username = *otherUsername
		}
		rooms = append(rooms, ru)
	}

	return rooms, nil
}

func (r *RoomRepository) GetAllUsers(ctx context.Context, excludeUserID string) ([]models.User, error) {
	query := `SELECT id, username, email FROM users WHERE id != $1 ORDER BY username`

	rows, err := r.db.Query(ctx, query, excludeUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *RoomRepository) CreateGroup(ctx context.Context, name string, creatorID string, userIDs []string) (*models.Room, error) {
	if len(userIDs) < 2 {
		return nil, errors.New("grupo deve ter no mínimo 3 usuários (criador + 2)")
	}

	log.Printf("CreateGroup: Criando grupo '%s' com %d membros", name, len(userIDs)+1)

	roomID := uuid.New().String()
	insertRoom := `INSERT INTO rooms (id, name, type, created_by) 
				   VALUES ($1, $2, 'group', $3) 
				   RETURNING id, name, type, created_by, created_at`

	room := &models.Room{}
	err := r.db.QueryRow(ctx, insertRoom, roomID, name, creatorID).
		Scan(&room.ID, &room.Name, &room.Type, &room.CreatedBy, &room.CreatedAt)
	if err != nil {
		log.Printf("✗ Erro ao criar grupo: %v", err)
		return nil, err
	}

	// Adicionar criador ao grupo
	allUserIDs := append([]string{creatorID}, userIDs...)

	// Inserir todos os membros
	for _, userID := range allUserIDs {
		insertUser := `INSERT INTO room_users (room_id, user_id) VALUES ($1, $2)`
		if _, err := r.db.Exec(ctx, insertUser, roomID, userID); err != nil {
			log.Printf("✗ Erro ao adicionar usuário ao grupo: %v", err)
			return nil, err
		}
	}

	room.Users = allUserIDs
	log.Printf("✓ Grupo '%s' criado com sucesso: %s", name, room.ID)

	return room, nil
}

func (r *RoomRepository) GetGroupMembers(ctx context.Context, roomID string) ([]models.User, error) {
	query := `SELECT u.id, u.username, u.email
			  FROM users u
			  INNER JOIN room_users ru ON ru.user_id = u.id
			  WHERE ru.room_id = $1
			  ORDER BY u.username`

	rows, err := r.db.Query(ctx, query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *RoomRepository) GetUserGroups(ctx context.Context, userID string) ([]models.Room, error) {
	query := `SELECT DISTINCT r.id, r.name, r.type, r.created_by, r.created_at
			  FROM rooms r
			  INNER JOIN room_users ru ON ru.room_id = r.id
			  WHERE ru.user_id = $1 AND r.type = 'group'
			  ORDER BY r.created_at DESC`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.Name, &room.Type, &room.CreatedBy, &room.CreatedAt); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *RoomRepository) AddUserToGroup(ctx context.Context, roomID, userID string) error {
	query := `INSERT INTO room_users (room_id, user_id) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, roomID, userID)
	return err
}

func (r *RoomRepository) RemoveUserFromGroup(ctx context.Context, roomID, userID string) error {
	query := `DELETE FROM room_users WHERE room_id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, query, roomID, userID)
	return err
}
