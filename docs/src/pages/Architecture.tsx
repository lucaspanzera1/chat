import React from 'react';
import { Database, Shield, MessageSquare, Globe } from 'lucide-react';

export const Architecture: React.FC = () => {
    return (
        <div className="space-y-12">
            <header className="space-y-4 border-b border-cyber-border pb-8">
                <h1 className="text-4xl md:text-5xl font-bold text-white glitch-text" data-text="ARQUITETURA">
                    ARQUITETURA
                </h1>
                <p className="text-xl text-cyber-dim max-w-2xl">
                    Visão geral técnica dos componentes do sistema e fluxo de dados.
                </p>
            </header>

            <section className="space-y-6">
                <h2 className="text-2xl font-bold text-white flex items-center gap-3">
                    <span className="w-1 h-8 bg-green-500 rounded-full"></span>
                    Diagrama do Sistema
                </h2>
                <div className="bg-black border border-cyber-border rounded p-6 font-mono text-xs md:text-sm overflow-x-auto text-cyber-text leading-relaxed whitespace-pre">
                    {`┌─────────────────────────────────────────────────────────────┐
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
└─────────────────────────────────────────────────────────────┘`}
                </div>
            </section>

            <section className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="bento-card">
                    <div className="flex items-center gap-3 mb-4">
                        <div className="p-2 bg-purple-500/10 rounded border border-purple-500/20 text-purple-400">
                            <Shield size={20} />
                        </div>
                        <h3 className="text-lg font-bold text-white">Auth (JWT)</h3>
                    </div>
                    <p className="text-cyber-dim text-sm mb-4">
                        Sistema de autenticação stateless.
                    </p>
                    <ul className="space-y-2 text-sm text-cyber-dim list-disc list-inside">
                        <li>Gera tokens JWT com expiração de 24h</li>
                        <li>Valida tokens em todas as rotas protegidas</li>
                        <li>Armazena claims: userID, username, email</li>
                    </ul>
                </div>

                <div className="bento-card">
                    <div className="flex items-center gap-3 mb-4">
                        <div className="p-2 bg-blue-500/10 rounded border border-blue-500/20 text-blue-400">
                            <Globe size={20} />
                        </div>
                        <h3 className="text-lg font-bold text-white">Hub</h3>
                    </div>
                    <p className="text-cyber-dim text-sm mb-4">
                        Gerenciador central de salas e conexões WebSocket.
                    </p>
                    <ul className="space-y-2 text-sm text-cyber-dim list-disc list-inside">
                        <li>Mantém mapa de rooms e seus clientes conectados</li>
                        <li>Distribui mensagens apenas para clientes da mesma sala</li>
                        <li>Gerencia contagem de usuários online por sala</li>
                    </ul>
                </div>

                <div className="bento-card">
                    <div className="flex items-center gap-3 mb-4">
                        <div className="p-2 bg-green-500/10 rounded border border-green-500/20 text-green-400">
                            <MessageSquare size={20} />
                        </div>
                        <h3 className="text-lg font-bold text-white">Client</h3>
                    </div>
                    <p className="text-cyber-dim text-sm mb-4">
                        Representação de cada conexão WebSocket ativa.
                    </p>
                    <ul className="space-y-2 text-sm text-cyber-dim list-disc list-inside">
                        <li>ReadPump: Lê mensagens do WebSocket</li>
                        <li>WritePump: Envia mensagens para o WebSocket</li>
                        <li>Mantém heartbeat com ping/pong</li>
                    </ul>
                </div>

                <div className="bento-card">
                    <div className="flex items-center gap-3 mb-4">
                        <div className="p-2 bg-yellow-500/10 rounded border border-yellow-500/20 text-yellow-400">
                            <Database size={20} />
                        </div>
                        <h3 className="text-lg font-bold text-white">Repositories</h3>
                    </div>
                    <p className="text-cyber-dim text-sm mb-4">
                        Camada de abstração para acesso ao banco de dados.
                    </p>
                    <ul className="space-y-2 text-sm text-cyber-dim list-disc list-inside">
                        <li><strong>UserRepository</strong>: Login, registro, busca</li>
                        <li><strong>MessageRepository</strong>: Histórico de mensagens</li>
                        <li><strong>RoomRepository</strong>: Gerenciamento de salas</li>
                    </ul>
                </div>
            </section>

            <section className="space-y-6">
                <h2 className="text-2xl font-bold text-white flex items-center gap-3">
                    <span className="w-1 h-8 bg-green-500 rounded-full"></span>
                    Banco de Dados
                </h2>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div className="border border-cyber-border bg-cyber-bg/50 p-4 rounded">
                        <h4 className="text-white font-bold mb-2 border-b border-cyber-border pb-2">users</h4>
                        <ul className="text-xs font-mono text-cyber-dim space-y-1">
                            <li>id (UUID, PK)</li>
                            <li>username (VARCHAR, UNIQUE)</li>
                            <li>email (VARCHAR, UNIQUE)</li>
                            <li>password_hash (TEXT)</li>
                            <li>created_at (TIMESTAMP)</li>
                        </ul>
                    </div>
                    <div className="border border-cyber-border bg-cyber-bg/50 p-4 rounded">
                        <h4 className="text-white font-bold mb-2 border-b border-cyber-border pb-2">messages</h4>
                        <ul className="text-xs font-mono text-cyber-dim space-y-1">
                            <li>id (UUID, PK)</li>
                            <li>room_id (UUID, FK)</li>
                            <li>user_id (UUID, FK)</li>
                            <li>content (TEXT)</li>
                            <li>type (VARCHAR)</li>
                            <li>created_at (TIMESTAMP)</li>
                        </ul>
                    </div>
                    <div className="border border-cyber-border bg-cyber-bg/50 p-4 rounded">
                        <h4 className="text-white font-bold mb-2 border-b border-cyber-border pb-2">rooms</h4>
                        <ul className="text-xs font-mono text-cyber-dim space-y-1">
                            <li>id (UUID, PK)</li>
                            <li>name (VARCHAR)</li>
                            <li>type (VARCHAR)</li>
                            <li>created_by (UUID, FK)</li>
                            <li>created_at (TIMESTAMP)</li>
                        </ul>
                    </div>
                    <div className="border border-cyber-border bg-cyber-bg/50 p-4 rounded">
                        <h4 className="text-white font-bold mb-2 border-b border-cyber-border pb-2">room_users</h4>
                        <ul className="text-xs font-mono text-cyber-dim space-y-1">
                            <li>room_id (UUID, FK)</li>
                            <li>user_id (UUID, FK)</li>
                            <li>joined_at (TIMESTAMP)</li>
                        </ul>
                    </div>
                </div>
            </section>
        </div>
    );
};
