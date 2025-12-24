"use client";

import { useQuery } from "@tanstack/react-query";
import { api, DedicatedServer } from "@/lib/api";
import { ServerCard } from "@/components/ServerCard";
import { MetricsPanel } from "@/components/MetricsPanel";
import { formatUptime, formatBytes } from "@/lib/utils";
import { Clock, Server, Users, RefreshCw, HardDrive, Cpu, MemoryStick, Play, AlertTriangle, ExternalLink } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import Link from "next/link";

export default function Dashboard() {
  const { data: serversData, refetch: refetchServers, isRefetching: isRefetchingServers } = useQuery({
    queryKey: ["servers"],
    queryFn: () => api.getServers(),
    refetchInterval: 10000,
  });

  const { data: metrics, refetch: refetchMetrics, isRefetching: isRefetchingMetrics } = useQuery({
    queryKey: ["metrics"],
    queryFn: () => api.getMetrics(),
    refetchInterval: 5000,
  });

  // Dedicated servers data
  const { data: dedicatedData, refetch: refetchDedicated, isRefetching: isRefetchingDedicated } = useQuery({
    queryKey: ["dedicated-servers-dashboard"],
    queryFn: () => api.getDedicatedServers({ perPage: 6 }),
    refetchInterval: 10000,
  });

  const { data: dedicatedStats } = useQuery({
    queryKey: ["dedicated-servers-stats"],
    queryFn: () => api.getDedicatedServerStats(),
    refetchInterval: 10000,
  });

  const { data: nodesData } = useQuery({
    queryKey: ["nodes-dashboard"],
    queryFn: () => api.getNodes(true),
    refetchInterval: 10000,
  });

  const handleRefresh = () => {
    refetchServers();
    refetchMetrics();
    refetchDedicated();
  };

  const servers = serversData?.servers || [];
  const onlineCount = servers.filter((s) => s.status === "online").length;
  const totalPlayers = servers.reduce((acc, s) => acc + s.players, 0);

  const dedicatedServers = dedicatedData?.servers || [];
  const nodes = nodesData?.nodes || [];
  const onlineNodes = nodes.filter((n) => n.status === "online").length;

  return (
    <div className="space-y-6">
      {/* Page Title */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Dashboard</h1>
          <p className="text-muted-foreground">KaraPixel Sunucu Kontrol Paneli</p>
        </div>
        <Button
          variant="outline"
          size="sm"
          onClick={handleRefresh}
          disabled={isRefetchingServers || isRefetchingMetrics || isRefetchingDedicated}
        >
          <RefreshCw className={`h-4 w-4 mr-2 ${(isRefetchingServers || isRefetchingMetrics || isRefetchingDedicated) ? 'animate-spin' : ''}`} />
          Refresh
        </Button>
      </div>

      {/* Quick Stats */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-4">
        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card hover:border-primary/50 transition-colors">
          <div className="p-3 rounded-full bg-emerald-500/20">
            <Server className="h-6 w-6 text-emerald-500" />
          </div>
          <div>
            <p className="text-sm text-muted-foreground">MC Sunucular</p>
            <p className="text-2xl font-bold">{onlineCount}/{servers.length}</p>
          </div>
        </div>

        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card hover:border-primary/50 transition-colors">
          <div className="p-3 rounded-full bg-primary/20">
            <Users className="h-6 w-6 text-primary" />
          </div>
          <div>
            <p className="text-sm text-muted-foreground">Toplam Oyuncu</p>
            <p className="text-2xl font-bold">{totalPlayers}</p>
          </div>
        </div>

        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card hover:border-primary/50 transition-colors">
          <div className="p-3 rounded-full bg-blue-500/20">
            <HardDrive className="h-6 w-6 text-blue-500" />
          </div>
          <div>
            <p className="text-sm text-muted-foreground">Nodes</p>
            <p className="text-2xl font-bold">{onlineNodes}/{nodes.length}</p>
          </div>
        </div>

        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card hover:border-primary/50 transition-colors">
          <div className="p-3 rounded-full bg-purple-500/20">
            <Play className="h-6 w-6 text-purple-500" />
          </div>
          <div>
            <p className="text-sm text-muted-foreground">Dedicated Servers</p>
            <p className="text-2xl font-bold">{dedicatedStats?.running || 0}/{dedicatedStats?.total || 0}</p>
          </div>
        </div>

        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card hover:border-primary/50 transition-colors">
          <div className="p-3 rounded-full bg-accent/20">
            <Clock className="h-6 w-6 text-accent" />
          </div>
          <div>
            <p className="text-sm text-muted-foreground">System Uptime</p>
            <p className="text-2xl font-bold">{metrics ? formatUptime(metrics.uptime) : "-"}</p>
          </div>
        </div>
      </div>

      {/* System Metrics */}
      <div>
        <h2 className="text-xl font-semibold mb-4">System Metrics</h2>
        <MetricsPanel metrics={metrics || null} />
      </div>

      {/* Dedicated Servers Section */}
      {dedicatedServers.length > 0 && (
        <div>
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold">Dedicated Servers</h2>
            <Link href="/admin/dedicated-servers">
              <Button variant="ghost" size="sm">
                Tümünü Gör
                <ExternalLink className="h-4 w-4 ml-2" />
              </Button>
            </Link>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {dedicatedServers.slice(0, 6).map((server) => (
              <DedicatedServerCard key={server.id} server={server} />
            ))}
          </div>
        </div>
      )}

      {/* Nodes Overview */}
      {nodes.length > 0 && (
        <div>
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold">Nodes Overview</h2>
            <Link href="/admin/nodes">
              <Button variant="ghost" size="sm">
                Tümünü Gör
                <ExternalLink className="h-4 w-4 ml-2" />
              </Button>
            </Link>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {nodes.slice(0, 3).map((node) => (
              <Card key={node.id} className="hover:border-primary/50 transition-colors">
                <CardHeader className="pb-2">
                  <CardTitle className="flex items-center justify-between text-base">
                    <span className="flex items-center gap-2">
                      <HardDrive className="h-4 w-4" />
                      {node.name}
                    </span>
                    <span className={`text-xs px-2 py-0.5 rounded-full ${
                      node.status === "online"
                        ? "bg-emerald-500/20 text-emerald-500"
                        : "bg-red-500/20 text-red-500"
                    }`}>
                      {node.status}
                    </span>
                  </CardTitle>
                </CardHeader>
                <CardContent className="space-y-3">
                  <div className="flex items-center justify-between text-sm text-muted-foreground">
                    <span>{node.serverCount} sunucu</span>
                    <span>{node.fqdn}</span>
                  </div>
                  <div className="space-y-2">
                    <div className="flex items-center justify-between text-xs">
                      <span className="flex items-center gap-1">
                        <MemoryStick className="h-3 w-3" /> RAM
                      </span>
                      <span>{node.memoryPercent.toFixed(0)}%</span>
                    </div>
                    <Progress value={node.memoryPercent} className="h-1.5" />
                  </div>
                  <div className="space-y-2">
                    <div className="flex items-center justify-between text-xs">
                      <span className="flex items-center gap-1">
                        <Cpu className="h-3 w-3" /> CPU
                      </span>
                      <span>{node.cpuPercent.toFixed(0)}%</span>
                    </div>
                    <Progress value={node.cpuPercent} className="h-1.5" />
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      )}

      {/* MC Servers Grid */}
      <div>
        <h2 className="text-xl font-semibold mb-4">Minecraft Servers</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {servers.map((server) => (
            <ServerCard
              key={server.id}
              server={server}
              onAction={() => refetchServers()}
            />
          ))}
        </div>

        {servers.length === 0 && (
          <div className="text-center py-12 text-muted-foreground">
            <Server className="h-12 w-12 mx-auto mb-4 opacity-50" />
            <p>No servers configured</p>
          </div>
        )}
      </div>
    </div>
  );
}

// Mini Dedicated Server Card for Dashboard
function DedicatedServerCard({ server }: { server: DedicatedServer }) {
  const getStatusColor = (status: string) => {
    switch (status) {
      case "running":
        return "bg-emerald-500/20 text-emerald-500";
      case "starting":
      case "restarting":
      case "stopping":
        return "bg-yellow-500/20 text-yellow-500";
      case "suspended":
        return "bg-red-500/20 text-red-500";
      default:
        return "bg-gray-500/20 text-gray-500";
    }
  };

  const getStatusText = (status: string) => {
    switch (status) {
      case "running": return "Çalışıyor";
      case "starting": return "Başlatılıyor";
      case "stopping": return "Durduruluyor";
      case "offline": return "Kapalı";
      case "suspended": return "Askıda";
      default: return status;
    }
  };

  return (
    <Link href={`/admin/dedicated-servers/${server.id}`}>
      <Card className="hover:border-primary/50 transition-colors cursor-pointer">
        <CardHeader className="pb-2">
          <CardTitle className="flex items-center justify-between text-base">
            <span className="truncate">{server.name}</span>
            <span className={`text-xs px-2 py-0.5 rounded-full ${getStatusColor(server.status)}`}>
              {getStatusText(server.status)}
            </span>
          </CardTitle>
        </CardHeader>
        <CardContent className="space-y-2">
          <div className="flex items-center justify-between text-sm text-muted-foreground">
            <span>{server.nodeName}</span>
            <span className="font-mono text-xs">{server.ip}:{server.port}</span>
          </div>
          <div className="grid grid-cols-3 gap-2 text-xs">
            <div className="text-center p-2 bg-muted/50 rounded">
              <MemoryStick className="h-3 w-3 mx-auto mb-1 text-muted-foreground" />
              <p className="font-medium">{formatBytes(server.memory * 1024 * 1024)}</p>
            </div>
            <div className="text-center p-2 bg-muted/50 rounded">
              <Cpu className="h-3 w-3 mx-auto mb-1 text-muted-foreground" />
              <p className="font-medium">{server.cpu}%</p>
            </div>
            <div className="text-center p-2 bg-muted/50 rounded">
              <HardDrive className="h-3 w-3 mx-auto mb-1 text-muted-foreground" />
              <p className="font-medium">{formatBytes(server.disk * 1024 * 1024)}</p>
            </div>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}
