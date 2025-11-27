import React from 'react';
import { Terminal, Shield, Zap, Database, Layout } from 'lucide-react';

export const Introduction: React.FC = () => {
    return (
        <div className="space-y-12">
            <header className="space-y-4 border-b border-cyber-border pb-8">
                <h1 className="text-4xl md:text-5xl font-bold text-white glitch-text" data-text="INTRODUÇÃO">
                    INTRODUÇÃO
                </h1>
                <p className="text-xl text-cyber-dim max-w-2xl">
                    Infraestrutura de chat em tempo real, segura e de alta performance construída com Go e React.
                </p>
            </header>

            <section className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="bento-card group">
                    <div className="mb-4 text-green-500 p-3 bg-green-500/10 w-fit rounded border border-green-500/20 group-hover:bg-green-500/20 transition-colors">
                        <Terminal size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-white mb-2">Backend em Go</h3>
                    <p className="text-cyber-dim text-sm leading-relaxed">
                        Construído sobre a robusta biblioteca padrão do Go. Possui roteador HTTP personalizado, arquitetura limpa e tratamento concorrente de WebSockets para máximo throughput.
                    </p>
                </div>

                <div className="bento-card group">
                    <div className="mb-4 text-blue-400 p-3 bg-blue-400/10 w-fit rounded border border-blue-400/20 group-hover:bg-blue-400/20 transition-colors">
                        <Zap size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-white mb-2">WebSocket em Tempo Real</h3>
                    <p className="text-cyber-dim text-sm leading-relaxed">
                        Canal de comunicação full-duplex para entrega instantânea de mensagens, atualizações de presença e indicadores de digitação.
                    </p>
                </div>

                <div className="bento-card group">
                    <div className="mb-4 text-purple-400 p-3 bg-purple-400/10 w-fit rounded border border-purple-400/20 group-hover:bg-purple-400/20 transition-colors">
                        <Shield size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-white mb-2">Segurança por Padrão</h3>
                    <p className="text-cyber-dim text-sm leading-relaxed">
                        Autenticação stateless baseada em JWT, hash de senha com bcrypt e políticas rigorosas de CORS garantem a integridade e segurança dos dados.
                    </p>
                </div>

                <div className="bento-card group">
                    <div className="mb-4 text-yellow-400 p-3 bg-yellow-400/10 w-fit rounded border border-yellow-400/20 group-hover:bg-yellow-400/20 transition-colors">
                        <Database size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-white mb-2">Persistência de Dados</h3>
                    <p className="text-cyber-dim text-sm leading-relaxed">
                        Esquema de banco de dados PostgreSQL otimizado para histórico de chat, relacionamentos de usuários e metadados de grupos.
                    </p>
                </div>

                <div className="bento-card group col-span-1 md:col-span-2">
                    <div className="mb-4 text-pink-400 p-3 bg-pink-400/10 w-fit rounded border border-pink-400/20 group-hover:bg-pink-400/20 transition-colors">
                        <Layout size={24} />
                    </div>
                    <h3 className="text-lg font-bold text-white mb-2">Interface Moderna</h3>
                    <p className="text-cyber-dim text-sm leading-relaxed">
                        Design cyberpunk com tema escuro, responsivo para mobile e desktop. Inclui badges de notificação, contadores de mensagens não lidas e suporte a timezone (Brasília GMT-3).
                    </p>
                </div>
            </section>

            <section className="space-y-6">
                <h2 className="text-2xl font-bold text-white flex items-center gap-3">
                    <span className="w-1 h-8 bg-green-500 rounded-full"></span>
                    Arquitetura
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

            <section className="space-y-6">
                <h2 className="text-2xl font-bold text-white flex items-center gap-3">
                    <span className="w-1 h-8 bg-green-500 rounded-full"></span>
                    Começo Rápido
                </h2>

                <div className="space-y-4">
                    <p className="text-cyber-dim">Clone o repositório e inicie o servidor de desenvolvimento:</p>
                    <div className="bg-black border border-cyber-border rounded p-4 font-mono text-sm relative group">
                        <div className="absolute top-0 right-0 p-2 opacity-0 group-hover:opacity-100 transition-opacity">
                            <span className="text-xs text-cyber-dim">BASH</span>
                        </div>
                        <div className="flex gap-4">
                            <div className="text-cyber-dim select-none">
                                1<br />2<br />3<br />4<br />5
                            </div>
                            <div className="text-white">
                                <span className="text-purple-400">git</span> clone https://github.com/lucaspanzera1/chat.git<br />
                                <span className="text-purple-400">cd</span> chat<br />
                                <span className="text-purple-400">cp</span> .env.example .env<br />
                                <span className="text-purple-400">docker-compose</span> up -d<br />
                                <span className="text-purple-400">go</span> run cmd/server/main.go
                            </div>
                        </div>
                    </div>
                    <p className="text-xs text-cyber-dim mt-2">
                        O servidor estará rodando em <code className="text-green-500">http://localhost:8080</code>.
                    </p>
                </div>
            </section>
        </div>
    );
};
