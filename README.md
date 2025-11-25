# 💬 Chat em Tempo Real com Golang

Um chat em tempo real construído com Go e WebSockets.

## 🏗️ Arquitetura

```
┌─────────────────────────────────────────────────────────────┐
│                     Cliente Web                              │
│                    (HTML + JS)                               │
└─────────────────────┬───────────────────────────────────────┘
                      │ WebSocket
┌─────────────────────▼───────────────────────────────────────┐
│                   Chat Server                                │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐          │
│  │     Hub     │  │  Handlers   │  │   Client    │          │
│  │ (Gerencia   │  │ (HTTP/WS)   │  │ (Conexão    │          │
│  │  conexões)  │  │             │  │  WebSocket) │          │
│  └─────────────┘  └─────────────┘  └─────────────┘          │
└─────────────────────────────────────────────────────────────┘
```

## 📁 Estrutura do Projeto

```
chat/
├── cmd/
│   └── server/
│       └── main.go           # Ponto de entrada
├── internal/
│   ├── hub/
│   │   └── hub.go            # Gerenciador de conexões
│   ├── client/
│   │   └── client.go         # Conexão do cliente
│   ├── handlers/
│   │   └── websocket.go      # Handler WebSocket
│   └── models/
│       └── message.go        # Modelo de mensagem
├── web/
│   └── index.html            # Cliente web
├── go.mod
├── go.sum
└── README.md
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.23 ou superior

### Instalação

```bash
# Clone o repositório
git clone https://github.com/lucaspanzera1/chat.git
cd chat

# Instale as dependências
go mod tidy

# Execute o servidor
go run cmd/server/main.go
```

### Uso

1. Acesse `http://localhost:8080` no navegador
2. Digite seu nome de usuário
3. Clique em "Entrar"
4. Comece a conversar!

> 💡 Abra em múltiplas abas para testar a comunicação entre usuários.

## 🛠️ Tecnologias

| Tecnologia | Uso |
|------------|-----|
| [Go](https://golang.org/) | Linguagem principal |
| [Gorilla WebSocket](https://github.com/gorilla/websocket) | Comunicação em tempo real |
| [UUID](https://github.com/google/uuid) | Geração de IDs únicos |

## 📦 Dependências

```go
require (
    github.com/google/uuid v1.6.0
    github.com/gorilla/websocket v1.5.3
)
```

## 🔧 Componentes

### Hub
Gerenciador central que:
- Mantém registro de clientes conectados
- Distribui mensagens (broadcast)
- Gerencia entrada/saída de usuários

### Client
Representa cada conexão:
- `ReadPump`: Lê mensagens do WebSocket
- `WritePump`: Envia mensagens para o WebSocket
- Mantém heartbeat com ping/pong

### Message
Estrutura de dados:
```go
type Message struct {
    ID        string    // UUID único
    Username  string    // Nome do usuário
    Content   string    // Conteúdo da mensagem
    Timestamp time.Time // Data/hora
    Type      string    // "message", "join", "leave"
}
```

