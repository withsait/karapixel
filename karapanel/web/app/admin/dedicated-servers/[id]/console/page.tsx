"use client";

import { useState, useEffect, useRef } from "react";
import { useParams } from "next/navigation";
import { useQuery, useMutation } from "@tanstack/react-query";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { formatBytes, formatUptime } from "@/lib/utils";
import Link from "next/link";
import {
  ArrowLeft,
  RefreshCw,
  Play,
  Square,
  RotateCw,
  Send,
  Terminal,
  Trash2,
  Download,
  Cpu,
  MemoryStick,
  HardDrive,
  Clock,
} from "lucide-react";

export default function ServerConsolePage() {
  const params = useParams();
  const serverId = Number(params.id);
  const consoleRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  const [logs, setLogs] = useState<string[]>([]);
  const [command, setCommand] = useState("");
  const [commandHistory, setCommandHistory] = useState<string[]>([]);
  const [historyIndex, setHistoryIndex] = useState(-1);

  const { data: server, refetch } = useQuery({
    queryKey: ["dedicated-server", serverId],
    queryFn: () => api.getDedicatedServer(serverId),
    refetchInterval: 3000,
  });

  const powerMutation = useMutation({
    mutationFn: async (action: "start" | "stop" | "restart" | "kill") => {
      switch (action) {
        case "start":
          return api.startDedicatedServer(serverId);
        case "stop":
          return api.stopDedicatedServer(serverId);
        case "restart":
          return api.restartDedicatedServer(serverId);
        case "kill":
          return api.killDedicatedServer(serverId);
      }
    },
    onSuccess: () => refetch(),
  });

  const sendCommandMutation = useMutation({
    mutationFn: (cmd: string) => api.sendDedicatedServerCommand(serverId, cmd),
    onSuccess: () => {
      setLogs((prev) => [...prev, `> ${command}`]);
      setCommandHistory((prev) => [command, ...prev].slice(0, 50));
      setCommand("");
      setHistoryIndex(-1);
    },
  });

  // Scroll to bottom on new logs
  useEffect(() => {
    if (consoleRef.current) {
      consoleRef.current.scrollTop = consoleRef.current.scrollHeight;
    }
  }, [logs]);

  // Demo logs for display
  useEffect(() => {
    // Simulate some initial logs
    setLogs([
      "[KaraPanel] Console connected",
      "[KaraPanel] Server status: " + (server?.status || "unknown"),
      "[KaraPanel] Type commands below to interact with the server",
      "",
    ]);
  }, [server?.status]);

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && command.trim()) {
      sendCommandMutation.mutate(command);
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      if (historyIndex < commandHistory.length - 1) {
        const newIndex = historyIndex + 1;
        setHistoryIndex(newIndex);
        setCommand(commandHistory[newIndex]);
      }
    } else if (e.key === "ArrowDown") {
      e.preventDefault();
      if (historyIndex > 0) {
        const newIndex = historyIndex - 1;
        setHistoryIndex(newIndex);
        setCommand(commandHistory[newIndex]);
      } else if (historyIndex === 0) {
        setHistoryIndex(-1);
        setCommand("");
      }
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case "running":
        return "bg-emerald-500";
      case "starting":
      case "restarting":
      case "stopping":
        return "bg-yellow-500";
      case "offline":
        return "bg-gray-500";
      default:
        return "bg-red-500";
    }
  };

  return (
    <div className="h-[calc(100vh-120px)] flex flex-col space-y-4">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Link href={`/admin/dedicated-servers/${serverId}`}>
            <Button variant="ghost" size="icon">
              <ArrowLeft className="h-5 w-5" />
            </Button>
          </Link>
          <div>
            <h1 className="text-2xl font-bold flex items-center gap-2">
              <Terminal className="h-6 w-6" />
              {server?.name || "Server"} - Konsol
            </h1>
            <div className="flex items-center gap-2 text-sm text-muted-foreground">
              <div className={`w-2 h-2 rounded-full ${getStatusColor(server?.status || "offline")}`} />
              {server?.status === "running" ? "Çalışıyor" : "Kapalı"}
              {server?.uptime && server.status === "running" && (
                <span className="ml-2">• Uptime: {formatUptime(server.uptime)}</span>
              )}
            </div>
          </div>
        </div>

        <div className="flex gap-2">
          {server?.status === "offline" ? (
            <Button
              onClick={() => powerMutation.mutate("start")}
              disabled={powerMutation.isPending || server?.suspended}
              className="bg-emerald-600 hover:bg-emerald-700"
            >
              <Play className="h-4 w-4 mr-2" />
              Başlat
            </Button>
          ) : (
            <>
              <Button
                variant="outline"
                onClick={() => powerMutation.mutate("restart")}
                disabled={powerMutation.isPending}
              >
                <RotateCw className="h-4 w-4 mr-2" />
                Yeniden Başlat
              </Button>
              <Button
                variant="destructive"
                onClick={() => powerMutation.mutate("stop")}
                disabled={powerMutation.isPending}
              >
                <Square className="h-4 w-4 mr-2" />
                Durdur
              </Button>
            </>
          )}
        </div>
      </div>

      {/* Resource Stats Bar */}
      <div className="grid grid-cols-4 gap-4">
        <Card className="bg-card/50">
          <CardContent className="py-3 flex items-center gap-3">
            <Cpu className="h-5 w-5 text-primary" />
            <div>
              <p className="text-xs text-muted-foreground">CPU</p>
              <p className="font-bold">{server?.cpuUsed?.toFixed(1) || 0}%</p>
            </div>
          </CardContent>
        </Card>

        <Card className="bg-card/50">
          <CardContent className="py-3 flex items-center gap-3">
            <MemoryStick className="h-5 w-5 text-accent" />
            <div>
              <p className="text-xs text-muted-foreground">RAM</p>
              <p className="font-bold">
                {formatBytes((server?.memoryUsed || 0) * 1024 * 1024)} / {formatBytes((server?.memory || 0) * 1024 * 1024)}
              </p>
            </div>
          </CardContent>
        </Card>

        <Card className="bg-card/50">
          <CardContent className="py-3 flex items-center gap-3">
            <HardDrive className="h-5 w-5 text-blue-500" />
            <div>
              <p className="text-xs text-muted-foreground">Disk</p>
              <p className="font-bold">
                {formatBytes((server?.diskUsed || 0) * 1024 * 1024)} / {formatBytes((server?.disk || 0) * 1024 * 1024)}
              </p>
            </div>
          </CardContent>
        </Card>

        <Card className="bg-card/50">
          <CardContent className="py-3 flex items-center gap-3">
            <Clock className="h-5 w-5 text-emerald-500" />
            <div>
              <p className="text-xs text-muted-foreground">Uptime</p>
              <p className="font-bold">
                {server?.status === "running" ? formatUptime(server?.uptime || 0) : "-"}
              </p>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Console */}
      <Card className="flex-1 flex flex-col overflow-hidden">
        <CardHeader className="py-3 border-b flex flex-row items-center justify-between">
          <CardTitle className="text-sm flex items-center gap-2">
            <Terminal className="h-4 w-4" />
            Console Output
          </CardTitle>
          <div className="flex gap-1">
            <Button
              variant="ghost"
              size="icon"
              onClick={() => setLogs([])}
              title="Clear"
            >
              <Trash2 className="h-4 w-4" />
            </Button>
            <Button
              variant="ghost"
              size="icon"
              title="Download Logs"
            >
              <Download className="h-4 w-4" />
            </Button>
          </div>
        </CardHeader>
        <CardContent className="flex-1 p-0 flex flex-col">
          {/* Log output */}
          <div
            ref={consoleRef}
            className="flex-1 overflow-y-auto p-4 font-mono text-sm bg-black/50"
            style={{ minHeight: 0 }}
          >
            {logs.map((log, i) => (
              <div key={i} className="whitespace-pre-wrap text-gray-300">
                {log}
              </div>
            ))}
          </div>

          {/* Command input */}
          <div className="border-t p-3 flex gap-2">
            <div className="flex-1 flex items-center gap-2 bg-black/50 rounded-lg px-3">
              <span className="text-primary font-mono">{">"}</span>
              <input
                ref={inputRef}
                type="text"
                value={command}
                onChange={(e) => setCommand(e.target.value)}
                onKeyDown={handleKeyDown}
                placeholder={server?.status === "running" ? "Enter command..." : "Server is offline"}
                disabled={server?.status !== "running"}
                className="flex-1 py-2 bg-transparent outline-none font-mono text-sm"
              />
            </div>
            <Button
              onClick={() => command.trim() && sendCommandMutation.mutate(command)}
              disabled={!command.trim() || server?.status !== "running"}
            >
              <Send className="h-4 w-4" />
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
