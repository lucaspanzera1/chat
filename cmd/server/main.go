package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucaspanzera1/chat/internal/database"
	"github.com/lucaspanzera1/chat/internal/handlers"
	"github.com/lucaspanzera1/chat/internal/hub"
	"github.com/lucaspanzera1/chat/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado")
	}
	if err := database.Connect(); err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer database.Close()

	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Erro ao executar migrações: %v", err)
	}

	userRepo := repository.NewUserRepository(database.DB)
	messageRepo := repository.NewMessageRepository(database.DB)

	h := hub.NewHub()
	go h.Run()

	wsHandler := handlers.NewWSHandler(h, userRepo, messageRepo)
	httpHandler := handlers.NewHTTPHandler(messageRepo)

	http.HandleFunc("/ws", wsHandler.ServeWS)
	http.HandleFunc("/api/messages", httpHandler.GetHistory)

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando em http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
