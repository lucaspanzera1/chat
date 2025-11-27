import React from 'react';
import { Endpoint } from './components/Endpoint';

export const ApiDocs: React.FC = () => {
    return (
        <div className="min-h-screen dither-bg dither-overlay text-cyber-text font-mono p-4 md:p-8">
            <div className="scanlines"></div>

            <div className="max-w-7xl mx-auto relative z-10">
                <header className="mb-12 text-center">
                    <h1 className="text-4xl md:text-6xl font-bold mb-4 glitch-text" data-text="CHAT API // DOCS">
                        CHAT API // DOCS
                    </h1>
                    <p className="text-cyber-dim text-sm md:text-base max-w-2xl mx-auto border-t border-b border-cyber-border py-2">
                        SECURE COMMUNICATION PROTOCOL REFERENCE // V1.0.0
                    </p>
                </header>

                <div className="bento-grid">
                    {/* Introduction Card */}
                    <div className="bento-card col-span-2 md:col-span-4">
                        <h2 className="text-xl font-bold mb-4 flex items-center gap-2">
                            <span className="w-2 h-2 bg-green-500 rounded-full animate-pulse"></span>
                            SYSTEM STATUS: ONLINE
                        </h2>
                        <p className="text-cyber-dim mb-4">
                            Welcome to the official API documentation. All endpoints are prefixed with <code className="bg-cyber-bg border border-cyber-border px-1 text-green-500">/api</code>.
                            Authentication is handled via JWT tokens in the Authorization header.
                        </p>
                        <div className="grid grid-cols-2 md:grid-cols-4 gap-4 text-xs-custom">
                            <div className="border border-cyber-border p-2 text-center">
                                <div className="text-cyber-dim mb-1">PROTOCOL</div>
                                <div className="text-cyber-text">HTTPS / WSS</div>
                            </div>
                            <div className="border border-cyber-border p-2 text-center">
                                <div className="text-cyber-dim mb-1">AUTH</div>
                                <div className="text-cyber-text">BEARER JWT</div>
                            </div>
                            <div className="border border-cyber-border p-2 text-center">
                                <div className="text-cyber-dim mb-1">FORMAT</div>
                                <div className="text-cyber-text">JSON</div>
                            </div>
                            <div className="border border-cyber-border p-2 text-center">
                                <div className="text-cyber-dim mb-1">VERSION</div>
                                <div className="text-cyber-text">1.0.0</div>
                            </div>
                        </div>
                    </div>

                    {/* Authentication Section */}
                    <div className="bento-card col-span-2 md:col-span-2">
                        <h3 className="text-lg font-bold mb-4 text-green-500 border-b border-cyber-border pb-2">
              // AUTHENTICATION
                        </h3>
                        <div className="space-y-6">
                            <Endpoint
                                method="POST"
                                path="/register"
                                description="Register a new user account."
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
                                path="/login"
                                description="Authenticate existing user."
                                requestBody={{
                                    email: "john@example.com",
                                    password: "securepassword123"
                                }}
                                responseBody={{
                                    token: "eyJhbGciOiJIUzI1Ni...",
                                    user: { id: "uuid", username: "johndoe" }
                                }}
                            />
                        </div>
                    </div>

                    {/* Users Section */}
                    <div className="bento-card col-span-2 md:col-span-2">
                        <h3 className="text-lg font-bold mb-4 text-blue-400 border-b border-cyber-border pb-2">
              // USER MANAGEMENT
                        </h3>
                        <div className="space-y-6">
                            <Endpoint
                                method="GET"
                                path="/users"
                                description="List all users."
                                requiresAuth
                                responseBody={[
                                    { id: "uuid", username: "alice", isOnline: true },
                                    { id: "uuid", username: "bob", isOnline: false }
                                ]}
                            />
                            <Endpoint
                                method="GET"
                                path="/me"
                                description="Get current user profile."
                                requiresAuth
                                responseBody={{
                                    id: "uuid",
                                    username: "johndoe",
                                    email: "john@example.com"
                                }}
                            />
                            <Endpoint
                                method="POST"
                                path="/username"
                                description="Update username."
                                requiresAuth
                                requestBody={{ username: "new_alias" }}
                                responseBody={{
                                    token: "new_token...",
                                    user: { username: "new_alias" }
                                }}
                            />
                        </div>
                    </div>

                    {/* Messages Section */}
                    <div className="bento-card col-span-2 md:col-span-4">
                        <h3 className="text-lg font-bold mb-4 text-purple-400 border-b border-cyber-border pb-2">
              // MESSAGING & ROOMS
                        </h3>
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                            <Endpoint
                                method="GET"
                                path="/history"
                                description="Get global chat history."
                                parameters={[
                                    { name: "limit", type: "number", description: "Max messages (default: 50)" }
                                ]}
                                responseBody={[
                                    { id: "msg-1", content: "Hello", senderId: "user-1" }
                                ]}
                            />
                            <Endpoint
                                method="GET"
                                path="/room-history"
                                description="Get room specific history."
                                parameters={[
                                    { name: "roomId", type: "uuid", required: true, description: "Target room ID" },
                                    { name: "limit", type: "number", description: "Max messages" }
                                ]}
                            />
                            <Endpoint
                                method="POST"
                                path="/private-room"
                                description="Init private chat."
                                requiresAuth
                                requestBody={{ otherUserId: "target-uuid" }}
                                responseBody={{ id: "room-uuid", type: "private" }}
                            />
                            <Endpoint
                                method="POST"
                                path="/groups"
                                description="Create group chat."
                                requiresAuth
                                requestBody={{ name: "Squad", userIds: ["u1", "u2"] }}
                                responseBody={{ id: "group-uuid", name: "Squad" }}
                            />
                        </div>
                    </div>

                    {/* WebSocket Section */}
                    <div className="bento-card col-span-2 md:col-span-4">
                        <h3 className="text-lg font-bold mb-4 text-yellow-400 border-b border-cyber-border pb-2">
              // REAL-TIME PROTOCOL
                        </h3>
                        <div className="flex flex-col md:flex-row gap-8">
                            <div className="flex-1">
                                <h4 className="text-sm font-bold text-cyber-dim mb-2">CONNECTION</h4>
                                <div className="bg-cyber-bg border border-cyber-border p-4 font-mono text-sm mb-4">
                                    ws://api.host/ws?token=JWT_TOKEN&roomId=ROOM_ID
                                </div>
                                <p className="text-cyber-dim text-sm">
                                    Establish a WebSocket connection to receive real-time updates.
                                    Authentication token must be passed as a query parameter.
                                </p>
                            </div>
                            <div className="flex-1">
                                <h4 className="text-sm font-bold text-cyber-dim mb-2">EVENTS</h4>
                                <ul className="space-y-2 text-sm">
                                    <li className="flex items-start gap-2">
                                        <span className="text-green-500">→</span>
                                        <span><strong className="text-white">message</strong>: New chat message received.</span>
                                    </li>
                                    <li className="flex items-start gap-2">
                                        <span className="text-green-500">→</span>
                                        <span><strong className="text-white">join/leave</strong>: User presence updates.</span>
                                    </li>
                                    <li className="flex items-start gap-2">
                                        <span className="text-green-500">→</span>
                                        <span><strong className="text-white">count</strong>: Online user count updates.</span>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>

                    {/* Footer */}
                    <div className="bento-card col-span-2 md:col-span-4 flex flex-col md:flex-row justify-between items-center text-xs-custom">
                        <div>
                            © 2025 CHAT SYSTEM // ALL RIGHTS RESERVED
                        </div>
                        <div className="mt-2 md:mt-0">
                            POWERED BY <a href="#" className="text-white hover:underline decoration-dotted">GO + REACT</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
