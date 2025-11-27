import React from 'react';
import { Endpoint } from '../components/Endpoint';

export const ApiReference: React.FC = () => {
    return (
        <div className="space-y-12">
            <header className="space-y-4 border-b border-cyber-border pb-8">
                <h1 className="text-4xl md:text-5xl font-bold text-white glitch-text" data-text="REFERÊNCIA API">
                    REFERÊNCIA API
                </h1>
                <p className="text-xl text-cyber-dim max-w-2xl">
                    Documentação completa para todos os endpoints REST disponíveis.
                </p>
            </header>

            <div className="space-y-12">
                {/* Authentication Section */}
                <section id="auth" className="scroll-mt-20">
                    <h2 className="text-2xl font-bold text-white mb-6 flex items-center gap-3">
                        <span className="text-green-500">//</span> AUTENTICAÇÃO
                    </h2>
                    <div className="space-y-6">
                        <Endpoint
                            method="POST"
                            path="/api/register"
                            description="Registrar uma nova conta de usuário."
                            requestBody={{
                                username: "johndoe",
                                email: "john@example.com",
                                password: "securepassword123"
                            }}
                            responseBody={{
                                token: "eyJhbGciOiJIUzI1Ni...",
                                user: { id: "uuid", username: "johndoe" }
                            }}
                        />
                        <Endpoint
                            method="POST"
                            path="/api/login"
                            description="Autenticar usuário existente e obter token JWT."
                            requestBody={{
                                email: "john@example.com",
                                password: "securepassword123"
                            }}
                            responseBody={{
                                token: "eyJhbGciOiJIUzI1Ni...",
                                user: { id: "uuid", username: "johndoe" }
                            }}
                        />
                        <Endpoint
                            method="GET"
                            path="/api/auth/google"
                            description="Iniciar fluxo de login com Google OAuth 2.0."
                        />
                        <Endpoint
                            method="GET"
                            path="/api/auth/google/callback"
                            description="Callback para o fluxo do Google OAuth."
                        />
                    </div>
                </section>

                {/* Users Section */}
                <section id="users" className="scroll-mt-20">
                    <h2 className="text-2xl font-bold text-white mb-6 flex items-center gap-3">
                        <span className="text-blue-400">//</span> GERENCIAMENTO DE USUÁRIOS
                    </h2>
                    <div className="space-y-6">
                        <Endpoint
                            method="GET"
                            path="/api/users"
                            description="Listar todos os usuários disponíveis."
                            requiresAuth
                            responseBody={[
                                { id: "uuid", username: "alice", isOnline: true },
                                { id: "uuid", username: "bob", isOnline: false }
                            ]}
                        />
                        <Endpoint
                            method="GET"
                            path="/api/user/me"
                            description="Obter perfil do usuário autenticado."
                            requiresAuth
                            responseBody={{
                                id: "uuid",
                                username: "johndoe",
                                email: "john@example.com"
                            }}
                        />
                        <Endpoint
                            method="GET"
                            path="/api/user/profile"
                            description="Obter perfil completo do usuário."
                            requiresAuth
                        />
                        <Endpoint
                            method="POST"
                            path="/api/user/username"
                            description="Atualizar nome de usuário."
                            requiresAuth
                            requestBody={{ username: "new_alias" }}
                            responseBody={{
                                token: "new_token...",
                                user: { username: "new_alias" }
                            }}
                        />
                        <Endpoint
                            method="POST"
                            path="/api/user/password"
                            description="Alterar senha do usuário."
                            requiresAuth
                            requestBody={{ password: "new_secure_password" }}
                        />
                    </div>
                </section>

                {/* Messages Section */}
                <section id="chat" className="scroll-mt-20">
                    <h2 className="text-2xl font-bold text-white mb-6 flex items-center gap-3">
                        <span className="text-purple-400">//</span> MENSAGENS & SALAS
                    </h2>
                    <div className="space-y-6">
                        <Endpoint
                            method="GET"
                            path="/api/messages"
                            description="Obter histórico do chat geral."
                            parameters={[
                                { name: "limit", type: "number", description: "Máximo de mensagens (padrão: 50)" }
                            ]}
                            responseBody={[
                                { id: "msg-1", content: "Olá", senderId: "user-1" }
                            ]}
                        />
                        <Endpoint
                            method="GET"
                            path="/api/room/messages"
                            description="Obter histórico de uma sala específica."
                            parameters={[
                                { name: "roomId", type: "uuid", required: true, description: "ID da sala alvo" },
                                { name: "limit", type: "number", description: "Máximo de mensagens" }
                            ]}
                        />
                        <Endpoint
                            method="POST"
                            path="/api/room/private"
                            description="Criar ou obter sala de chat privada."
                            requiresAuth
                            requestBody={{ otherUserId: "target-uuid" }}
                            responseBody={{ id: "room-uuid", type: "private" }}
                        />
                        <Endpoint
                            method="POST"
                            path="/api/group/create"
                            description="Criar novo grupo de chat."
                            requiresAuth
                            requestBody={{ name: "Squad", userIds: ["u1", "u2"] }}
                            responseBody={{ id: "group-uuid", name: "Squad" }}
                        />
                        <Endpoint
                            method="GET"
                            path="/api/groups"
                            description="Listar grupos do usuário."
                            requiresAuth
                        />
                        <Endpoint
                            method="GET"
                            path="/api/group/members"
                            description="Listar membros de um grupo."
                            requiresAuth
                            parameters={[
                                { name: "roomId", type: "uuid", required: true, description: "ID do grupo" }
                            ]}
                        />
                    </div>
                </section>

                {/* WebSocket Section */}
                <section id="websocket" className="scroll-mt-20">
                    <h2 className="text-2xl font-bold text-white mb-6 flex items-center gap-3">
                        <span className="text-yellow-400">//</span> WEBSOCKET
                    </h2>
                    <div className="space-y-6">
                        <div className="bento-card">
                            <h3 className="text-lg font-bold text-white mb-4">Conexão</h3>
                            <div className="bg-black border border-cyber-border p-4 font-mono text-sm mb-4 text-cyber-text">
                                GET /ws?token=JWT_TOKEN&roomId=ROOM_ID
                            </div>
                            <p className="text-cyber-dim text-sm">
                                Estabeleça uma conexão WebSocket para receber atualizações em tempo real.
                                O token de autenticação deve ser passado como parâmetro de query.
                            </p>
                        </div>
                    </div>
                </section>
            </div>
        </div>
    );
};
