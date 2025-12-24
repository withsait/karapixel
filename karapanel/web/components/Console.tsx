"use client";

import { useEffect, useRef, useState } from "react";
import { api } from "@/lib/api";

interface ConsoleProps {
  serverId: string;
}

export function Console({ serverId }: ConsoleProps) {
  const terminalRef = useRef<HTMLDivElement>(null);
  const terminalInstance = useRef<any>(null);
  const wsRef = useRef<WebSocket | null>(null);
  const [connected, setConnected] = useState(false);

  useEffect(() => {
    let term: any;
    let fitAddon: any;

    const initTerminal = async () => {
      // Dynamic import for xterm (client-side only)
      const { Terminal } = await import("xterm");
      const { FitAddon } = await import("xterm-addon-fit");
      const { WebLinksAddon } = await import("xterm-addon-web-links");
      // @ts-ignore - CSS import
      await import("xterm/css/xterm.css");

      if (!terminalRef.current) return;

      term = new Terminal({
        theme: {
          background: "#0a0a0f",
          foreground: "#e5e5e5",
          cursor: "#3b82f6",
          cursorAccent: "#0a0a0f",
          selectionBackground: "#3b82f6",
          black: "#000000",
          red: "#ef4444",
          green: "#22c55e",
          yellow: "#eab308",
          blue: "#3b82f6",
          magenta: "#a855f7",
          cyan: "#06b6d4",
          white: "#e5e5e5",
        },
        cursorBlink: true,
        fontSize: 14,
        fontFamily: "'JetBrains Mono', 'Fira Code', monospace",
        scrollback: 10000,
      });

      fitAddon = new FitAddon();
      term.loadAddon(fitAddon);
      term.loadAddon(new WebLinksAddon());

      term.open(terminalRef.current);
      fitAddon.fit();

      terminalInstance.current = term;

      // Connect WebSocket
      const wsUrl = api.getConsoleWsUrl(serverId);
      const ws = new WebSocket(wsUrl);
      wsRef.current = ws;

      ws.onopen = () => {
        setConnected(true);
        term.writeln("\x1b[32m[Connected to server console]\x1b[0m\n");
      };

      ws.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          if (data.type === "log") {
            term.writeln(data.data);
          } else if (data.error) {
            term.writeln(`\x1b[31m[Error: ${data.error}]\x1b[0m`);
          }
        } catch {
          term.writeln(event.data);
        }
      };

      ws.onclose = () => {
        setConnected(false);
        term.writeln("\n\x1b[33m[Disconnected from server]\x1b[0m");
      };

      ws.onerror = () => {
        term.writeln("\x1b[31m[WebSocket error]\x1b[0m");
      };

      // Handle window resize
      const handleResize = () => {
        fitAddon.fit();
      };
      window.addEventListener("resize", handleResize);

      return () => {
        window.removeEventListener("resize", handleResize);
      };
    };

    initTerminal();

    return () => {
      wsRef.current?.close();
      terminalInstance.current?.dispose();
    };
  }, [serverId]);

  return (
    <div className="flex flex-col h-full">
      {/* Status bar */}
      <div className="flex items-center justify-between px-4 py-2 bg-secondary/50 border-b">
        <div className="flex items-center gap-2">
          <div className={`w-2 h-2 rounded-full ${connected ? "bg-green-500" : "bg-red-500"}`} />
          <span className="text-sm">{connected ? "Connected" : "Disconnected"}</span>
        </div>
        <span className="text-xs text-muted-foreground">Server: {serverId}</span>
      </div>

      {/* Terminal */}
      <div ref={terminalRef} className="flex-1 bg-[#0a0a0f]" />
    </div>
  );
}
