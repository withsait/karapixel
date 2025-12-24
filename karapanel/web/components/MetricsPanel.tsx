"use client";

import { SystemMetrics } from "@/lib/api";
import { formatBytes, formatUptime } from "@/lib/utils";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { Cpu, HardDrive, Network, Clock, Gauge, MemoryStick } from "lucide-react";

interface MetricsPanelProps {
  metrics: SystemMetrics | null;
}

export function MetricsPanel({ metrics }: MetricsPanelProps) {
  if (!metrics) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        {[...Array(4)].map((_, i) => (
          <Card key={i} className="animate-pulse">
            <CardContent className="pt-6">
              <div className="h-16 bg-muted rounded" />
            </CardContent>
          </Card>
        ))}
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      {/* CPU */}
      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm font-medium flex items-center gap-2">
            <Cpu className="h-4 w-4" />
            CPU Usage
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{metrics.cpu.usagePercent.toFixed(1)}%</div>
          <Progress value={metrics.cpu.usagePercent} className="mt-2" />
          <p className="text-xs text-muted-foreground mt-2">
            {metrics.cpu.coreCount} cores | Load: {metrics.cpu.loadAvg[0].toFixed(2)}
          </p>
        </CardContent>
      </Card>

      {/* Memory */}
      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm font-medium flex items-center gap-2">
            <MemoryStick className="h-4 w-4" />
            Memory
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="text-2xl font-bold">{metrics.memory.percent.toFixed(1)}%</div>
          <Progress value={metrics.memory.percent} className="mt-2" />
          <p className="text-xs text-muted-foreground mt-2">
            {formatBytes(metrics.memory.used)} / {formatBytes(metrics.memory.total)}
          </p>
        </CardContent>
      </Card>

      {/* Disk */}
      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm font-medium flex items-center gap-2">
            <HardDrive className="h-4 w-4" />
            Disk I/O
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex justify-between">
            <div>
              <span className="text-xs text-muted-foreground">Read</span>
              <div className="text-lg font-bold">{formatBytes(metrics.disk.readBytes)}/s</div>
            </div>
            <div className="text-right">
              <span className="text-xs text-muted-foreground">Write</span>
              <div className="text-lg font-bold">{formatBytes(metrics.disk.writeBytes)}/s</div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Network */}
      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-sm font-medium flex items-center gap-2">
            <Network className="h-4 w-4" />
            Network
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex justify-between">
            <div>
              <span className="text-xs text-muted-foreground">In</span>
              <div className="text-lg font-bold">{formatBytes(metrics.network.bytesRecv)}/s</div>
            </div>
            <div className="text-right">
              <span className="text-xs text-muted-foreground">Out</span>
              <div className="text-lg font-bold">{formatBytes(metrics.network.bytesSent)}/s</div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
