package main

import (
	"log"
	"net/http"

	"github.com/lucaspanzera1/chat/internal/handlers"
	"github.com/lucaspanzera1/chat/internal/hub"
)

func main() {
	h := hub.NewHub()
	go h.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWS(h, w, r)
	})

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
