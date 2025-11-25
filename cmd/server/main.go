package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lucaspanzera1/chat/internal/database"
	"github.com/lucaspanzera1/chat/internal/handlers"
	"github.com/lucaspanzera1/chat/internal/hub"
	"github.com/lucaspanzera1/chat/internal/repository"
)

func main() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Printf("Aviso: não foi possível carregar timezone, usando UTC: %v", err)
	} else {
		time.Local = loc
		log.Println("✓ Timezone configurado para America/Sao_Paulo")
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado")
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("ERRO: JWT_SECRET não configurado no .env")
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
	roomRepo := repository.NewRoomRepository(database.DB)

	log.Printf("✓ Repositórios inicializados (DB: %v)", database.DB != nil)

	h := hub.NewHub()
	go h.Run()

	authHandler := handlers.NewAuthHandler(userRepo)
	wsHandler := handlers.NewWSHandler(h, userRepo, messageRepo)
	httpHandler := handlers.NewHTTPHandler(messageRepo, roomRepo, userRepo)

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		authHandler.Register(w, r)
	})

	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		authHandler.Login(w, r)
	})

	http.HandleFunc("/ws", wsHandler.ServeWS)

	http.HandleFunc("/api/messages", httpHandler.GetHistory)
	http.HandleFunc("/api/room/messages", httpHandler.GetRoomHistory)
	http.HandleFunc("/api/users", httpHandler.GetUsers)

	http.HandleFunc("/api/room/private", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		httpHandler.CreatePrivateRoom(w, r)
	})

	http.HandleFunc("/api/group/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		httpHandler.CreateGroup(w, r)
	})

	http.HandleFunc("/api/groups", httpHandler.GetUserGroups)
	http.HandleFunc("/api/group/members", httpHandler.GetGroupMembers)

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando em http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
