import React from 'react';
import { NavLink } from 'react-router-dom';
import { Terminal, Server, Shield, MessageSquare, Code2, Book, Zap } from 'lucide-react';

export const Sidebar: React.FC = () => {
    const navItems = [
        { path: '/', label: 'INTRODUÇÃO', icon: Terminal },
        { path: '/architecture', label: 'ARQUITETURA', icon: Server },
        { path: '/api', label: 'REFERÊNCIA API', icon: Book },
    ];

    const apiItems = [
        { path: '/api/auth', label: 'AUTENTICAÇÃO', icon: Shield },
        { path: '/api/users', label: 'USUÁRIOS', icon: Code2 },
        { path: '/api/chat', label: 'CHAT & SALAS', icon: MessageSquare },
        { path: '/api/websocket', label: 'WEBSOCKET', icon: Zap },
    ];

    return (
        <aside className="w-64 h-screen fixed left-0 top-0 bg-black border-r border-cyber-border overflow-y-auto hidden md:flex flex-col z-50">
            <div className="p-6 border-b border-cyber-border">
                <h1 className="text-xl font-bold glitch-text text-white" data-text="CHAT // DOCS">
                    CHAT // DOCS
                </h1>
                <p className="text-xs text-cyber-dim mt-2">v1.0.0 // SECURE</p>
            </div>

            <nav className="flex-1 p-4 space-y-8">
                <div>
                    <h3 className="text-xs font-bold text-cyber-dim mb-4 px-2 tracking-widest">COMEÇANDO</h3>
                    <div className="space-y-1">
                        {navItems.map((item) => (
                            <NavLink
                                key={item.path}
                                to={item.path}
                                end={item.path === '/'}
                                className={({ isActive }) =>
                                    `flex items-center gap-3 px-3 py-2 text-sm rounded transition-all duration-200 ${isActive
                                        ? 'bg-green-500/10 text-green-500 border border-green-500/20'
                                        : 'text-cyber-dim hover:text-white hover:bg-white/5'
                                    }`
                                }
                            >
                                <item.icon size={16} />
                                {item.label}
                            </NavLink>
                        ))}
                    </div>
                </div>

                <div>
                    <h3 className="text-xs font-bold text-cyber-dim mb-4 px-2 tracking-widest">ENDPOINTS API</h3>
                    <div className="space-y-1">
                        {apiItems.map((item) => (
                            <NavLink
                                key={item.path}
                                to={item.path}
                                className={({ isActive }) =>
                                    `flex items-center gap-3 px-3 py-2 text-sm rounded transition-all duration-200 ${isActive
                                        ? 'bg-blue-500/10 text-blue-400 border border-blue-500/20'
                                        : 'text-cyber-dim hover:text-white hover:bg-white/5'
                                    }`
                                }
                            >
                                <item.icon size={16} />
                                {item.label}
                            </NavLink>
                        ))}
                    </div>
                </div>
            </nav>

            <div className="p-4 border-t border-cyber-border">
                <div className="bg-cyber-bg border border-cyber-border p-3 rounded text-xs">
                    <div className="flex items-center gap-2 text-green-500 mb-1">
                        <div className="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                        SYSTEM ONLINE
                    </div>
                    <div className="text-cyber-dim">Latency: 12ms</div>
                </div>
            </div>
        </aside>
    );
};
