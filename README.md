# 💬 Chat em Tempo Real com Golang

Um chat em tempo real construído com Go, WebSockets e PostgreSQL.

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
- Docker e Docker Compose (para PostgreSQL)

### Instalação

```bash
# Clone o repositório
git clone https://github.com/lucaspanzera1/chat.git
cd chat

# Copie o arquivo de exemplo .env
cp .env.example .env

# Inicie o PostgreSQL com Docker
docker-compose up -d

# Aguarde o banco estar pronto
docker-compose ps

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

> 💡 Todas as mensagens são persistidas no PostgreSQL

## 🛠️ Tecnologias

| Tecnologia | Uso |
|------------|-----|
| [Go](https://golang.org/) | Linguagem principal |
| [Gorilla WebSocket](https://github.com/gorilla/websocket) | Comunicação em tempo real |
| [PostgreSQL](https://www.postgresql.org/) | Banco de dados |
| [pgx](https://github.com/jackc/pgx) | Driver PostgreSQL |
| [UUID](https://github.com/google/uuid) | Geração de IDs únicos |
| [godotenv](https://github.com/joho/godotenv) | Variáveis de ambiente |

## 📦 Dependências

```go
require (
    github.com/google/uuid v1.6.0
    github.com/gorilla/websocket v1.5.3
    github.com/jackc/pgx/v5 v5.7.2
    github.com/joho/godotenv v1.5.1
)
```

## 🗄️ Banco de Dados

### Tabelas

**users**
- `id` (UUID, PK)
- `username` (VARCHAR(50), UNIQUE)
- `created_at` (TIMESTAMP)

**messages**
- `id` (UUID, PK)
- `user_id` (UUID, FK → users)
- `username` (VARCHAR(50))
- `content` (TEXT)
- `type` (VARCHAR(20))
- `created_at` (TIMESTAMP)

### API Endpoints

- `GET /api/messages?limit=50` - Buscar histórico de mensagens
- `GET /ws?username=nome` - Conectar ao WebSocket

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

## 🗺️ Roadmap

- [x] MVP básico com WebSocket
- [x] Persistência com PostgreSQL
- [ ] Autenticação JWT
- [ ] Salas/Grupos de chat
- [ ] Envio de arquivos
- [ ] Deploy com Docker

