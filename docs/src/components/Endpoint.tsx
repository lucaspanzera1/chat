import React, { useState } from 'react';

interface Parameter {
    name: string;
    type: string;
    required?: boolean;
    description: string;
}

interface EndpointProps {
    method: 'GET' | 'POST' | 'PUT' | 'DELETE';
    path: string;
    description: string;
    parameters?: Parameter[];
    requestBody?: Record<string, any>;
    responseBody?: Record<string, any>;
    requiresAuth?: boolean;
}

export const Endpoint: React.FC<EndpointProps> = ({
    method,
    path,
    description,
    parameters,
    requestBody,
    responseBody,
    requiresAuth
}) => {
    const [isExpanded, setIsExpanded] = useState(false);

    const methodColors = {
        GET: 'text-blue-400 bg-blue-400/10 border-blue-400/20',
        POST: 'text-green-400 bg-green-400/10 border-green-400/20',
        PUT: 'text-yellow-400 bg-yellow-400/10 border-yellow-400/20',
        DELETE: 'text-red-400 bg-red-400/10 border-red-400/20',
    };

    return (
        <div className="border border-cyber-border bg-cyber-bg/30 hover:border-cyber-dim transition-colors group">
            <div
                className="p-3 flex items-center justify-between cursor-pointer"
                onClick={() => setIsExpanded(!isExpanded)}
            >
                <div className="flex items-center gap-3 overflow-hidden">
                    <span className={`text-xs font-bold px-2 py-1 rounded border ${methodColors[method]} w-16 text-center shrink-0`}>
                        {method}
                    </span>
                    <span className="font-mono text-sm text-cyber-text truncate group-hover:text-white transition-colors">
                        {path}
                    </span>
                    {requiresAuth && (
                        <span title="Authentication Required" className="text-[10px] text-yellow-500 border border-yellow-500/30 px-1.5 py-0.5 rounded bg-yellow-500/5 shrink-0">
                            🔒
                        </span>
                    )}
                </div>
                <div className="text-cyber-dim text-xs shrink-0 ml-2">
                    {isExpanded ? '[-]' : '[+]'}
                </div>
            </div>

            {isExpanded && (
                <div className="p-4 border-t border-cyber-border bg-black/20 text-sm animate-fade-in">
                    <p className="text-cyber-dim mb-4 italic border-l-2 border-cyber-border pl-3">
                        {description}
                    </p>

                    {parameters && parameters.length > 0 && (
                        <div className="mb-4">
                            <h4 className="text-xs font-bold text-cyber-dim mb-2 uppercase tracking-wider">Parâmetros</h4>
                            <div className="overflow-x-auto">
                                <table className="w-full text-left border-collapse">
                                    <thead>
                                        <tr className="border-b border-cyber-border text-xs text-cyber-dim">
                                            <th className="py-1 px-2">NOME</th>
                                            <th className="py-1 px-2">TIPO</th>
                                            <th className="py-1 px-2">DESC</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {parameters.map((param) => (
                                            <tr key={param.name} className="border-b border-cyber-border/50 last:border-0">
                                                <td className="py-2 px-2 font-mono text-cyber-text">
                                                    {param.name}
                                                    {param.required && <span className="text-red-500 ml-1">*</span>}
                                                </td>
                                                <td className="py-2 px-2 text-cyber-dim text-xs">{param.type}</td>
                                                <td className="py-2 px-2 text-cyber-dim">{param.description}</td>
                                            </tr>
                                        ))}
                                    </tbody>
                                </table>
                            </div>
                        </div>
                    )}

                    <div className="grid grid-cols-1 gap-4">
                        {requestBody && (
                            <div>
                                <h4 className="text-xs font-bold text-cyber-dim mb-2 uppercase tracking-wider">Corpo da Requisição</h4>
                                <pre className="bg-black border border-cyber-border p-3 rounded text-xs overflow-x-auto text-cyber-text font-mono">
                                    {JSON.stringify(requestBody, null, 2)}
                                </pre>
                            </div>
                        )}

                        {responseBody && (
                            <div>
                                <h4 className="text-xs font-bold text-cyber-dim mb-2 uppercase tracking-wider">Resposta</h4>
                                <pre className="bg-black border border-cyber-border p-3 rounded text-xs overflow-x-auto text-green-400/80 font-mono">
                                    {JSON.stringify(responseBody, null, 2)}
                                </pre>
                            </div>
                        )}
                    </div>
                </div>
            )}
        </div>
    );
};
