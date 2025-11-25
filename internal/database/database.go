package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL não configurada")
	}

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return err
	}

	config.ConnConfig.RuntimeParams = map[string]string{
		"timezone": "America/Sao_Paulo",
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), "SET timezone = 'America/Sao_Paulo'")
	if err != nil {
		log.Printf("Aviso: não foi possível configurar timezone: %v", err)
	}

	DB = pool
	log.Println("✓ Conectado ao PostgreSQL")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
