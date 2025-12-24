"use client";

import { ServerInfo, api } from "@/lib/api";
import { formatBytes, formatUptime, getStatusColor, getStatusBgColor, cn } from "@/lib/utils";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import {
  Play,
  Square,
  RotateCcw,
  Skull,
  Terminal,
  FolderOpen,
  Server,
  Users,
  Cpu,
  HardDrive
} from "lucide-react";
import Link from "next/link";
import { useState } from "react";

interface ServerCardProps {
  server: ServerInfo;
  onAction?: () => void;
}

export function ServerCard({ server, onAction }: ServerCardProps) {
  const [loading, setLoading] = useState<string | null>(null);

  const handleAction = async (action: 'start' | 'stop' | 'restart' | 'kill') => {
    setLoading(action);
    try {
      switch (action) {
        case 'start':
          await api.startServer(server.id);
          break;
        case 'stop':
          await api.stopServer(server.id);
          break;
        case 'restart':
          await api.restartServer(server.id);
          break;
        case 'kill':
          await api.killServer(server.id);
          break;
      }
      onAction?.();
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(null);
    }
  };

  const isOnline = server.status === 'online';
  const isTransitioning = server.status === 'starting' || server.status === 'stopping';

  return (
    <Card className="relative overflow-hidden">
      {/* Status indicator bar */}
      <div className={cn(
        "absolute top-0 left-0 right-0 h-1",
        server.status === 'online' && "bg-green-500",
        server.status === 'offline' && "bg-red-500",
        isTransitioning && "bg-yellow-500 animate-pulse",
      )} />

      <CardHeader className="pb-2">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Server className="h-5 w-5 text-muted-foreground" />
            <CardTitle className="text-lg">{server.name}</CardTitle>
          </div>
          <span className={cn(
            "text-sm font-medium px-2 py-0.5 rounded",
            getStatusBgColor(server.status),
            getStatusColor(server.status)
          )}>
            {server.status.toUpperCase()}
          </span>
        </div>
        <p className="text-sm text-muted-foreground">{server.type}</p>
      </CardHeader>

      <CardContent className="space-y-4">
        {/* Stats Grid */}
        <div className="grid grid-cols-2 gap-3">
          <div className="flex items-center gap-2">
            <Users className="h-4 w-4 text-muted-foreground" />
            <span className="text-sm">
              {server.players}/{server.maxPlayers || '?'}
            </span>
          </div>
          <div className="flex items-center gap-2">
            <Cpu className="h-4 w-4 text-muted-foreground" />
            <span className="text-sm">{server.cpuPercent.toFixed(1)}%</span>
          </div>
          <div className="flex items-center gap-2">
            <HardDrive className="h-4 w-4 text-muted-foreground" />
            <span className="text-sm">
              {formatBytes(server.memoryUsed)}
            </span>
          </div>
          <div className="text-sm text-muted-foreground">
            {isOnline ? formatUptime(server.uptime) : '-'}
          </div>
        </div>

        {/* Action Buttons */}
        <div className="flex gap-2">
          {!isOnline && (
            <Button
              size="sm"
              variant="success"
              onClick={() => handleAction('start')}
              disabled={isTransitioning || loading !== null}
            >
              <Play className="h-4 w-4 mr-1" />
              Start
            </Button>
          )}
          {isOnline && (
            <>
              <Button
                size="sm"
                variant="destructive"
                onClick={() => handleAction('stop')}
                disabled={isTransitioning || loading !== null}
              >
                <Square className="h-4 w-4 mr-1" />
                Stop
              </Button>
              <Button
                size="sm"
                variant="warning"
                onClick={() => handleAction('restart')}
                disabled={isTransitioning || loading !== null}
              >
                <RotateCcw className="h-4 w-4 mr-1" />
                Restart
              </Button>
            </>
          )}
          <Button
            size="sm"
            variant="ghost"
            onClick={() => handleAction('kill')}
            disabled={loading !== null}
            title="Force Kill"
          >
            <Skull className="h-4 w-4" />
          </Button>
        </div>

        {/* Quick Links */}
        <div className="flex gap-2 pt-2 border-t">
          <Link href={`/servers/${server.id}/console`} className="flex-1">
            <Button size="sm" variant="outline" className="w-full">
              <Terminal className="h-4 w-4 mr-1" />
              Console
            </Button>
          </Link>
          <Link href={`/servers/${server.id}/files`} className="flex-1">
            <Button size="sm" variant="outline" className="w-full">
              <FolderOpen className="h-4 w-4 mr-1" />
              Files
            </Button>
          </Link>
        </div>
      </CardContent>
    </Card>
  );
}
