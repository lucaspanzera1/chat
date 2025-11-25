package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucaspanzera1/chat/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	query := `INSERT INTO users (username) VALUES ($1) 
			  ON CONFLICT (username) DO UPDATE SET username = EXCLUDED.username
			  RETURNING id, username, created_at`

	err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, created_at FROM users WHERE username = $1`

	err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Register(ctx context.Context, req models.RegisterRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	query := `INSERT INTO users (username, email, password_hash) 
			  VALUES ($1, $2, $3) 
			  RETURNING id, username, email, created_at`

	err = r.db.QueryRow(ctx, query, req.Username, req.Email, string(hashedPassword)).
		Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Login(ctx context.Context, email, password string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE email = $1`

	err := r.db.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("credenciais inválidas")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, created_at FROM users WHERE id = $1`

	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) SetOnline(ctx context.Context, userID string) error {
	query := `UPDATE users SET is_online = TRUE, last_seen = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *UserRepository) SetOffline(ctx context.Context, userID string) error {
	query := `UPDATE users SET is_online = FALSE, last_seen = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *UserRepository) GetAllWithStatus(ctx context.Context, excludeUserID string) ([]models.User, error) {
	query := `SELECT id, username, email, is_online, last_seen 
			  FROM users 
			  WHERE id != $1 
			  ORDER BY is_online DESC, username`

	rows, err := r.db.Query(ctx, query, excludeUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.IsOnline, &user.LastSeen); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByGoogleID(ctx context.Context, googleID string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, google_id, created_at FROM users WHERE google_id = $1`

	err := r.db.QueryRow(ctx, query, googleID).Scan(&user.ID, &user.Username, &user.Email, &user.GoogleID, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash, google_id, created_at FROM users WHERE email = $1`

	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.GoogleID, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateWithGoogle(ctx context.Context, username, email, googleID string) (*models.User, error) {
	user := &models.User{}
	query := `INSERT INTO users (username, email, google_id) 
			  VALUES ($1, $2, $3) 
			  RETURNING id, username, email, google_id, created_at`

	err := r.db.QueryRow(ctx, query, username, email, googleID).
		Scan(&user.ID, &user.Username, &user.Email, &user.GoogleID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) LinkGoogleAccount(ctx context.Context, userID, googleID string) error {
	query := `UPDATE users SET google_id = $1 WHERE id = $2`
	_, err := r.db.Exec(ctx, query, googleID, userID)
	return err
}
