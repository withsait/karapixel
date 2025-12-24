"use client";

import { useEffect, useRef, useState, useCallback } from "react";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Send, Trash2, Download, Filter } from "lucide-react";

interface ConsoleProps {
  serverId: string;
}

type LogLevel = "all" | "info" | "warn" | "error";

export function Console({ serverId }: ConsoleProps) {
  const terminalRef = useRef<HTMLDivElement>(null);
  const terminalInstance = useRef<any>(null);
  const wsRef = useRef<WebSocket | null>(null);
  const fitAddonRef = useRef<any>(null);
  const [connected, setConnected] = useState(false);
  const [command, setCommand] = useState("");
  const [commandHistory, setCommandHistory] = useState<string[]>([]);
  const [historyIndex, setHistoryIndex] = useState(-1);
  const [logFilter, setLogFilter] = useState<LogLevel>("all");
  const inputRef = useRef<HTMLInputElement>(null);

  const sendCommand = useCallback(() => {
    if (!command.trim() || !wsRef.current || wsRef.current.readyState !== WebSocket.OPEN) return;

    wsRef.current.send(JSON.stringify({ type: "command", data: command }));

    // Add to history
    setCommandHistory(prev => {
      const newHistory = [command, ...prev.filter(c => c !== command)].slice(0, 50);
      return newHistory;
    });
    setCommand("");
    setHistoryIndex(-1);
  }, [command]);

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      sendCommand();
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      if (commandHistory.length > 0) {
        const newIndex = Math.min(historyIndex + 1, commandHistory.length - 1);
        setHistoryIndex(newIndex);
        setCommand(commandHistory[newIndex]);
      }
    } else if (e.key === "ArrowDown") {
      e.preventDefault();
      if (historyIndex > 0) {
        const newIndex = historyIndex - 1;
        setHistoryIndex(newIndex);
        setCommand(commandHistory[newIndex]);
      } else {
        setHistoryIndex(-1);
        setCommand("");
      }
    }
  };

  const clearConsole = () => {
    if (terminalInstance.current) {
      terminalInstance.current.clear();
      terminalInstance.current.writeln("\x1b[35m[Console cleared]\x1b[0m");
    }
  };

  const downloadLogs = async () => {
    try {
      const response = await api.getLogs(serverId, 1000);
      const blob = new Blob([response.logs.join("\n")], { type: "text/plain" });
      const url = URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = `${serverId}-logs-${new Date().toISOString().split("T")[0]}.txt`;
      a.click();
      URL.revokeObjectURL(url);
    } catch (error) {
      console.error("Failed to download logs:", error);
    }
  };

  useEffect(() => {
    let term: any;
    let fitAddon: any;

    const initTerminal = async () => {
      const { Terminal } = await import("xterm");
      const { FitAddon } = await import("xterm-addon-fit");
      const { WebLinksAddon } = await import("xterm-addon-web-links");
      // @ts-ignore
      await import("xterm/css/xterm.css");

      if (!terminalRef.current) return;

      // KaraPixel Obsidian/Indigo Theme
      term = new Terminal({
        theme: {
          background: "#0f0a15",
          foreground: "#e8e4f0",
          cursor: "#a855f7",
          cursorAccent: "#0f0a15",
          selectionBackground: "rgba(168, 85, 247, 0.3)",
          black: "#1a1225",
          red: "#ef4444",
          green: "#22c55e",
          yellow: "#eab308",
          blue: "#6366f1",
          magenta: "#a855f7",
          cyan: "#06b6d4",
          white: "#e8e4f0",
          brightBlack: "#3d2e52",
          brightRed: "#f87171",
          brightGreen: "#4ade80",
          brightYellow: "#facc15",
          brightBlue: "#818cf8",
          brightMagenta: "#c084fc",
          brightCyan: "#22d3ee",
          brightWhite: "#faf5ff",
        },
        cursorBlink: true,
        fontSize: 14,
        fontFamily: "'JetBrains Mono', 'Fira Code', 'Consolas', monospace",
        scrollback: 10000,
        convertEol: true,
      });

      fitAddon = new FitAddon();
      fitAddonRef.current = fitAddon;
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
        term.writeln("\x1b[35m╔════════════════════════════════════════╗\x1b[0m");
        term.writeln("\x1b[35m║   \x1b[1;37mKaraPanel Console Connected\x1b[0m\x1b[35m         ║\x1b[0m");
        term.writeln("\x1b[35m╚════════════════════════════════════════╝\x1b[0m\n");
      };

      ws.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          if (data.type === "log") {
            // Apply filter
            const line = data.data;
            if (logFilter === "all") {
              term.writeln(line);
            } else if (logFilter === "error" && (line.includes("ERROR") || line.includes("SEVERE"))) {
              term.writeln(line);
            } else if (logFilter === "warn" && (line.includes("WARN") || line.includes("ERROR") || line.includes("SEVERE"))) {
              term.writeln(line);
            } else if (logFilter === "info") {
              term.writeln(line);
            }
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
        term.writeln("\x1b[31m[WebSocket connection error]\x1b[0m");
      };

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
  }, [serverId, logFilter]);

  return (
    <div className="flex flex-col h-full rounded-lg overflow-hidden border border-border">
      {/* Toolbar */}
      <div className="flex items-center justify-between px-4 py-2 bg-secondary/50 border-b">
        <div className="flex items-center gap-4">
          <div className="flex items-center gap-2">
            <div className={`w-2 h-2 rounded-full ${connected ? "bg-green-500 animate-pulse" : "bg-red-500"}`} />
            <span className="text-sm font-medium">{connected ? "Connected" : "Disconnected"}</span>
          </div>
          <span className="text-xs text-muted-foreground">Server: {serverId}</span>
        </div>

        <div className="flex items-center gap-2">
          {/* Filter dropdown */}
          <select
            value={logFilter}
            onChange={(e) => setLogFilter(e.target.value as LogLevel)}
            className="text-xs bg-background border rounded px-2 py-1"
          >
            <option value="all">All Logs</option>
            <option value="info">Info+</option>
            <option value="warn">Warn+</option>
            <option value="error">Errors Only</option>
          </select>

          <Button variant="ghost" size="sm" onClick={clearConsole} title="Clear Console">
            <Trash2 className="h-4 w-4" />
          </Button>
          <Button variant="ghost" size="sm" onClick={downloadLogs} title="Download Logs">
            <Download className="h-4 w-4" />
          </Button>
        </div>
      </div>

      {/* Terminal */}
      <div ref={terminalRef} className="flex-1 bg-[#0f0a15]" />

      {/* Command input */}
      <div className="flex items-center gap-2 px-4 py-3 bg-secondary/50 border-t">
        <span className="text-primary font-mono text-sm">{">"}</span>
        <input
          ref={inputRef}
          type="text"
          value={command}
          onChange={(e) => setCommand(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder={connected ? "Enter command..." : "Not connected"}
          disabled={!connected}
          className="flex-1 bg-transparent border-none outline-none text-sm font-mono placeholder:text-muted-foreground disabled:opacity-50"
        />
        <Button
          size="sm"
          onClick={sendCommand}
          disabled={!connected || !command.trim()}
          className="bg-primary hover:bg-primary/80"
        >
          <Send className="h-4 w-4" />
        </Button>
      </div>
    </div>
  );
}
