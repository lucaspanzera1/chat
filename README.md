# 💬 Chat em Tempo Real com Golang

Um chat em tempo real construído com Go, WebSockets e PostgreSQL com autenticação JWT e mensagens privadas.

![ChatPreview](git/preview.gif)

## 🏗️ Arquitetura

```
┌─────────────────────────────────────────────────────────────┐
│                     Cliente Web                              │
│              (HTML + JS + Tailwind)                          │
└─────────────────────┬───────────────────────────────────────┘
                      │ WebSocket + JWT
┌─────────────────────▼───────────────────────────────────────┐
│                   Chat Server (Go)                           │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐          │
│  │     Hub     │  │   Auth      │  │   Rooms     │          │
│  │ (Gerencia   │  │   (JWT)     │  │ (General +  │          │
│  │  conexões)  │  │             │  │  Private)   │          │
│  └─────────────┘  └─────────────┘  └─────────────┘          │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                   PostgreSQL                                 │
│     ┌──────────┐    ┌──────────┐    ┌──────────┐            │
│     │  users   │    │ messages │    │  rooms   │            │
│     │          │    │          │    │          │            │
│     └──────────┘    └──────────┘    └──────────┘            │
└─────────────────────────────────────────────────────────────┘
```

![Chat](git/print.png)

## ✨ Features

### 🔐 Autenticação
- ✅ Sistema de registro com username, email e senha
- ✅ Login com JWT (token expira em 24h)
- ✅ **Login com Google OAuth 2.0**
- ✅ Vinculação automática de contas (email/senha + Google)
- ✅ Hash de senhas com bcrypt
- ✅ Validação de campos no frontend e backend

### 💬 Chat em Tempo Real
- ✅ **Chat Geral**: Canal público para todos os usuários
- ✅ **Mensagens Privadas (DM)**: Chat 1-a-1 entre usuários
- ✅ **Grupos**: Chat com 3 ou mais usuários
- ✅ Contagem de usuários online por sala
- ✅ Badges de notificação para mensagens não lidas
- ✅ Histórico de mensagens persistido no PostgreSQL

### 👥 Grupos
- ✅ Criar grupos com nome personalizado
- ✅ Adicionar múltiplos membros (mínimo 3 usuários)
- ✅ Lista de grupos na sidebar
- ✅ Histórico de mensagens por grupo
- ✅ Notificações de mensagens não lidas por grupo

### 🎨 Interface
- ✅ Design cyberpunk com tema escuro
- ✅ Sidebar com lista de usuários disponíveis
- ✅ Notificações visuais com badges animados
- ✅ Contador de mensagens não lidas no título da página
- ✅ Indicador de status online em tempo real
- ✅ Responsivo para mobile e desktop

### ⏰ Timezone
- ✅ Horários salvos e exibidos no timezone de Brasília (GMT-3)
- ✅ Timestamps precisos em todas as mensagens

## 📁 Estrutura do Projeto

```
chat/
├── cmd/
│   └── server/
│       └── main.go           # Ponto de entrada
├── internal/
│   ├── auth/
│   │   └── jwt.go            # Geração e validação de tokens
│   ├── hub/
│   │   └── hub.go            # Gerenciador de conexões e rooms
│   ├── client/
│   │   └── client.go         # Conexão WebSocket do cliente
│   ├── handlers/
│   │   ├── auth.go           # Handlers de autenticação
│   │   ├── websocket.go      # Handler WebSocket
│   │   └── http.go           # Handlers HTTP
│   ├── database/
│   │   ├── database.go       # Conexão com PostgreSQL
│   │   └── migrations.go     # Migrações do banco
│   ├── repository/
│   │   ├── user_repository.go     # CRUD de usuários
│   │   ├── message_repository.go  # CRUD de mensagens
│   │   └── room_repository.go     # CRUD de salas
│   └── models/
│       ├── user.go           # Modelo de usuário
│       ├── message.go        # Modelo de mensagem
│       └── room.go           # Modelo de sala
├── web/
│   ├── index.html            # Interface do chat
│   └── auth.html             # Página de login/registro
├── docker-compose.yml        # PostgreSQL containerizado
├── .env.example              # Exemplo de variáveis de ambiente
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

# Edite o .env e configure:

# Inicie o PostgreSQL com Docker
docker-compose up -d

# Aguarde o banco estar pronto (aguarde ~10 segundos)
docker-compose ps

# Instale as dependências
go mod tidy

# Execute o servidor
go run cmd/server/main.go
```

### Uso

1. Acesse `http://localhost:8080/auth.html`
2. Registre-se com username, email e senha
3. Faça login com suas credenciais
4. Você será redirecionado para o chat
5. Use o **# GENERAL** para chat público
6. Clique em qualquer **@ username** para iniciar um chat privado
7. Badges vermelhos indicam mensagens não lidas

## 🛠️ Tecnologias

| Tecnologia | Uso |
|------------|-----|
| [Go](https://golang.org/) | Linguagem principal |
| [Gorilla WebSocket](https://github.com/gorilla/websocket) | Comunicação em tempo real |
| [PostgreSQL](https://www.postgresql.org/) | Banco de dados |
| [pgx](https://github.com/jackc/pgx) | Driver PostgreSQL |
| [JWT](https://github.com/golang-jwt/jwt) | Autenticação com tokens |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hash de senhas |
| [UUID](https://github.com/google/uuid) | Geração de IDs únicos |
| [godotenv](https://github.com/joho/godotenv) | Variáveis de ambiente |
| [Tailwind CSS](https://tailwindcss.com/) | Estilização do frontend |

## 📦 Dependências

```go
require (
    github.com/golang-jwt/jwt/v5 v5.2.1
    github.com/google/uuid v1.6.0
    github.com/gorilla/websocket v1.5.3
    github.com/jackc/pgx/v5 v5.7.2
    github.com/joho/godotenv v1.5.1
    golang.org/x/crypto v0.31.0
)
```

## 🗄️ Banco de Dados

### Tabelas

**users**
- `id` (UUID, PK)
- `username` (VARCHAR(50), UNIQUE)
- `email` (VARCHAR(255), UNIQUE)
- `password_hash` (TEXT)
- `created_at` (TIMESTAMP)

**messages**
- `id` (UUID, PK)
- `room_id` (UUID, FK → rooms)
- `user_id` (UUID, FK → users)
- `username` (VARCHAR(50))
- `content` (TEXT)
- `type` (VARCHAR(20)) - "message", "join", "leave"
- `created_at` (TIMESTAMP)

**rooms**
- `id` (UUID, PK)
- `name` (VARCHAR(100), nullable)
- `type` (VARCHAR(20)) - "general", "private" ou "group"
- `created_by` (UUID, FK → users) - Criador do grupo
- `created_at` (TIMESTAMP)

**room_users**
- `room_id` (UUID, FK → rooms)
- `user_id` (UUID, FK → users)
- `joined_at` (TIMESTAMP)

### API Endpoints

#### Autenticação
- `POST /api/register` - Registrar novo usuário
- `POST /api/login` - Login e obter token JWT
- `GET /api/auth/google` - Iniciar login com Google
- `GET /api/auth/google/callback` - Callback do Google OAuth

#### Chat
- `GET /ws?token=JWT&roomId=UUID` - Conectar ao WebSocket
- `GET /api/messages?limit=50` - Histórico do chat geral
- `GET /api/room/messages?roomId=UUID&limit=50` - Histórico de uma sala

#### Usuários e Salas
- `GET /api/users` - Listar usuários disponíveis (requer token)
- `POST /api/room/private` - Criar/obter sala privada (requer token)

#### Grupos
- `POST /api/group/create` - Criar novo grupo (requer token)
- `GET /api/groups` - Listar grupos do usuário (requer token)
- `GET /api/group/members?roomId=UUID` - Listar membros de um grupo

## 🔧 Componentes

### Auth (JWT)
Sistema de autenticação:
- Gera tokens JWT com expiração de 24h
- Valida tokens em todas as rotas protegidas
- Armazena claims: userID, username, email

### Hub
Gerenciador de salas e conexões:
- Mantém mapa de rooms e seus clientes conectados
- Distribui mensagens apenas para clientes da mesma sala
- Gerencia contagem de usuários online por sala
- Suporta múltiplas salas simultâneas

### Client
Representa cada conexão WebSocket:
- `ReadPump`: Lê mensagens do WebSocket
- `WritePump`: Envia mensagens para o WebSocket
- Mantém heartbeat com ping/pong
- Associado a uma sala específica

### Repositories
Camada de acesso a dados:
- **UserRepository**: Login, registro, buscar usuários
- **MessageRepository**: Salvar e buscar mensagens por sala
- **RoomRepository**: Criar salas privadas, listar salas do usuário

## 🔒 Segurança

- ✅ Senhas com hash bcrypt (cost 10)
- ✅ Tokens JWT com expiração
- ✅ Validação de entrada no frontend e backend
- ✅ Proteção contra SQL injection (prepared statements)
- ✅ CORS configurável
- ⚠️ Em produção: usar HTTPS e secret JWT forte

## 🗺️ Roadmap

- [x] MVP básico com WebSocket
- [x] Persistência com PostgreSQL
- [x] Autenticação JWT
- [x] Salas privadas (DMs)
- [x] Notificações de mensagens não lidas
- [x] Grupos de chat (3+ usuários)
- [x] Status online/offline persistente
- [ ] Envio de arquivos/imagens
- [ ] Emojis e reações
- [ ] Deploy com Docker (compose completo)
- [ ] Testes unitários e de integração

## 📝 Variáveis de Ambiente

```env
DATABASE_URL=postgres://user:pass@localhost:5432/db?sslmode=disable&timezone=America/Sao_Paulo
SERVER_PORT=8080
JWT_SECRET=seu-secret-super-seguro-mude-em-producao
GOOGLE_CLIENT_ID=seu-client-id.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=seu-client-secret
GOOGLE_REDIRECT_URL=http://localhost:8080/api/auth/google/callback
```
