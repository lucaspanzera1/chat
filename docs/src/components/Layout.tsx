import React from 'react';
import { Outlet } from 'react-router-dom';
import { Sidebar } from './Sidebar';

export const Layout: React.FC = () => {
    return (
        <div className="min-h-screen bg-black text-cyber-text font-mono selection:bg-green-500 selection:text-black">
            <div className="dither-bg dither-overlay fixed inset-0 pointer-events-none z-0"></div>
            <div className="scanlines fixed inset-0 pointer-events-none z-50"></div>

            <Sidebar />

            <main className="md:ml-64 relative z-10 min-h-screen">
                <div className="max-w-5xl mx-auto p-6 md:p-12 animate-fade-in">
                    <Outlet />
                </div>
            </main>
        </div>
    );
};
