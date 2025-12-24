"use client";

import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/api";
import { ServerCard } from "@/components/ServerCard";
import { MetricsPanel } from "@/components/MetricsPanel";
import { formatUptime } from "@/lib/utils";
import { Clock, Server, Users } from "lucide-react";

export default function Dashboard() {
  const { data: serversData, refetch: refetchServers } = useQuery({
    queryKey: ["servers"],
    queryFn: () => api.getServers(),
  });

  const { data: metrics } = useQuery({
    queryKey: ["metrics"],
    queryFn: () => api.getMetrics(),
  });

  const servers = serversData?.servers || [];
  const onlineCount = servers.filter((s) => s.status === "online").length;
  const totalPlayers = servers.reduce((acc, s) => acc + s.players, 0);

  return (
    <div className="space-y-6">
      {/* Page Title */}
      <div>
        <h1 className="text-3xl font-bold">Dashboard</h1>
        <p className="text-muted-foreground">KaraPixel Sunucu Kontrol Paneli</p>
      </div>

      {/* Quick Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card">
          <div className="p-3 rounded-full bg-green-500/20">
            <Server className="h-6 w-6 text-green-500" />
          </div>
          <div>
            <p className="text-sm text-muted-foreground">Online Servers</p>
            <p className="text-2xl font-bold">{onlineCount}/{servers.length}</p>
          </div>
        </div>

        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card">
          <div className="p-3 rounded-full bg-blue-500/20">
            <Users className="h-6 w-6 text-blue-500" />
          </div>
          <div>
            <p className="text-sm text-muted-foreground">Total Players</p>
            <p className="text-2xl font-bold">{totalPlayers}</p>
          </div>
        </div>

        <div className="flex items-center gap-4 p-4 rounded-lg border bg-card">
          <div className="p-3 rounded-full bg-purple-500/20">
            <Clock className="h-6 w-6 text-purple-500" />
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

      {/* Servers Grid */}
      <div>
        <h2 className="text-xl font-semibold mb-4">Servers</h2>
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
