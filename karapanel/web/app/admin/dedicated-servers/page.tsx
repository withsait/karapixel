"use client";

import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { api, DedicatedServer, NodeWithStats, Allocation } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { formatBytes } from "@/lib/utils";
import Link from "next/link";
import {
  Server,
  Plus,
  Settings,
  Trash2,
  RefreshCw,
  HardDrive,
  Cpu,
  MemoryStick,
  Play,
  Square,
  RotateCw,
  AlertTriangle,
  Search,
  Terminal,
  FolderOpen,
  ExternalLink,
  Power,
  Zap,
  Globe,
  ChevronLeft,
  ChevronRight,
} from "lucide-react";

export default function DedicatedServersPage() {
  const queryClient = useQueryClient();
  const [search, setSearch] = useState("");
  const [selectedNode, setSelectedNode] = useState<number | null>(null);
  const [page, setPage] = useState(1);
  const [showCreateModal, setShowCreateModal] = useState(false);

  const { data: serversData, isLoading, refetch } = useQuery({
    queryKey: ["dedicated-servers", { search, nodeId: selectedNode, page }],
    queryFn: () => api.getDedicatedServers({ search, nodeId: selectedNode || undefined, page, perPage: 12 }),
  });

  const { data: statsData } = useQuery({
    queryKey: ["dedicated-servers-stats"],
    queryFn: () => api.getDedicatedServerStats(),
  });

  const { data: nodesData } = useQuery({
    queryKey: ["nodes"],
    queryFn: () => api.getNodes(true),
  });

  const servers = serversData?.servers || [];
  const pagination = serversData?.pagination;
  const stats = statsData;
  const nodes = nodesData?.nodes || [];

  const getStatusIcon = (status: string) => {
    switch (status) {
      case "running":
        return <Play className="h-4 w-4 text-emerald-500" />;
      case "starting":
      case "restarting":
        return <RotateCw className="h-4 w-4 text-yellow-500 animate-spin" />;
      case "stopping":
        return <Square className="h-4 w-4 text-yellow-500" />;
      case "offline":
        return <Power className="h-4 w-4 text-gray-500" />;
      case "installing":
        return <Zap className="h-4 w-4 text-blue-500 animate-pulse" />;
      case "suspended":
        return <AlertTriangle className="h-4 w-4 text-red-500" />;
      default:
        return <Power className="h-4 w-4 text-gray-500" />;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case "running":
        return "text-emerald-500 bg-emerald-500/20";
      case "starting":
      case "restarting":
      case "stopping":
        return "text-yellow-500 bg-yellow-500/20";
      case "installing":
        return "text-blue-500 bg-blue-500/20";
      case "suspended":
        return "text-red-500 bg-red-500/20";
      default:
        return "text-gray-500 bg-gray-500/20";
    }
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-64">
        <RefreshCw className="h-8 w-8 animate-spin text-primary" />
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Dedicated Servers</h1>
          <p className="text-muted-foreground">Tüm sunucuları yönetin ve izleyin</p>
        </div>
        <div className="flex gap-2">
          <Button variant="outline" size="sm" onClick={() => refetch()}>
            <RefreshCw className="h-4 w-4 mr-2" />
            Yenile
          </Button>
          <Button size="sm" onClick={() => setShowCreateModal(true)}>
            <Plus className="h-4 w-4 mr-2" />
            Sunucu Ekle
          </Button>
        </div>
      </div>

      {/* Stats Overview */}
      <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-primary/20">
                <Server className="h-6 w-6 text-primary" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Toplam</p>
                <p className="text-2xl font-bold">{stats?.total || 0}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-emerald-500/20">
                <Play className="h-6 w-6 text-emerald-500" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Çalışan</p>
                <p className="text-2xl font-bold">{stats?.running || 0}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-red-500/20">
                <AlertTriangle className="h-6 w-6 text-red-500" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Askıda</p>
                <p className="text-2xl font-bold">{stats?.suspended || 0}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-accent/20">
                <MemoryStick className="h-6 w-6 text-accent" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Toplam RAM</p>
                <p className="text-2xl font-bold">{formatBytes((stats?.totalMemory || 0) * 1024 * 1024)}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-blue-500/20">
                <HardDrive className="h-6 w-6 text-blue-500" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Toplam Disk</p>
                <p className="text-2xl font-bold">{formatBytes((stats?.totalDisk || 0) * 1024 * 1024)}</p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Filters */}
      <div className="flex gap-4">
        <div className="relative flex-1 max-w-md">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
          <input
            type="text"
            placeholder="Sunucu ara..."
            value={search}
            onChange={(e) => {
              setSearch(e.target.value);
              setPage(1);
            }}
            className="w-full pl-10 pr-4 py-2 rounded-lg bg-background border focus:border-primary outline-none"
          />
        </div>

        <select
          value={selectedNode || ""}
          onChange={(e) => {
            setSelectedNode(e.target.value ? Number(e.target.value) : null);
            setPage(1);
          }}
          className="px-4 py-2 rounded-lg bg-background border focus:border-primary outline-none"
        >
          <option value="">Tüm Node'lar</option>
          {nodes.map((node) => (
            <option key={node.id} value={node.id}>
              {node.name}
            </option>
          ))}
        </select>
      </div>

      {/* Servers Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {servers.map((server) => (
          <ServerCard key={server.id} server={server} onRefresh={refetch} />
        ))}
      </div>

      {servers.length === 0 && (
        <Card className="border-dashed">
          <CardContent className="flex flex-col items-center justify-center py-12">
            <Server className="h-12 w-12 text-muted-foreground mb-4" />
            <p className="text-lg font-medium">
              {search ? "Sonuç bulunamadı" : "Henüz sunucu eklenmemiş"}
            </p>
            <p className="text-sm text-muted-foreground mb-4">
              {search
                ? "Farklı anahtar kelimeler deneyin"
                : "İlk sunucunuzu oluşturun"}
            </p>
            {!search && (
              <Button onClick={() => setShowCreateModal(true)}>
                <Plus className="h-4 w-4 mr-2" />
                Sunucu Ekle
              </Button>
            )}
          </CardContent>
        </Card>
      )}

      {/* Pagination */}
      {pagination && pagination.totalPages > 1 && (
        <div className="flex items-center justify-center gap-2">
          <Button
            variant="outline"
            size="sm"
            onClick={() => setPage((p) => Math.max(1, p - 1))}
            disabled={page === 1}
          >
            <ChevronLeft className="h-4 w-4" />
          </Button>
          <span className="text-sm text-muted-foreground">
            Sayfa {page} / {pagination.totalPages}
          </span>
          <Button
            variant="outline"
            size="sm"
            onClick={() => setPage((p) => Math.min(pagination.totalPages, p + 1))}
            disabled={page === pagination.totalPages}
          >
            <ChevronRight className="h-4 w-4" />
          </Button>
        </div>
      )}

      {/* Create Server Modal */}
      {showCreateModal && (
        <CreateServerModal
          nodes={nodes}
          onClose={() => setShowCreateModal(false)}
          onSuccess={() => {
            setShowCreateModal(false);
            queryClient.invalidateQueries({ queryKey: ["dedicated-servers"] });
            queryClient.invalidateQueries({ queryKey: ["dedicated-servers-stats"] });
          }}
        />
      )}
    </div>
  );
}

// Server Card Component
function ServerCard({
  server,
  onRefresh,
}: {
  server: DedicatedServer;
  onRefresh: () => void;
}) {
  const queryClient = useQueryClient();

  const powerMutation = useMutation({
    mutationFn: async (action: "start" | "stop" | "restart" | "kill") => {
      switch (action) {
        case "start":
          return api.startDedicatedServer(server.id);
        case "stop":
          return api.stopDedicatedServer(server.id);
        case "restart":
          return api.restartDedicatedServer(server.id);
        case "kill":
          return api.killDedicatedServer(server.id);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["dedicated-servers"] });
    },
  });

  const getStatusColor = (status: string) => {
    switch (status) {
      case "running":
        return "text-emerald-500 bg-emerald-500/20";
      case "starting":
      case "restarting":
      case "stopping":
        return "text-yellow-500 bg-yellow-500/20";
      case "installing":
        return "text-blue-500 bg-blue-500/20";
      case "suspended":
        return "text-red-500 bg-red-500/20";
      default:
        return "text-gray-500 bg-gray-500/20";
    }
  };

  const getStatusText = (status: string) => {
    switch (status) {
      case "running":
        return "Çalışıyor";
      case "starting":
        return "Başlatılıyor";
      case "stopping":
        return "Durduruluyor";
      case "restarting":
        return "Yeniden Başlatılıyor";
      case "installing":
        return "Kuruluyor";
      case "suspended":
        return "Askıda";
      case "offline":
        return "Kapalı";
      default:
        return status;
    }
  };

  return (
    <Card className="hover:border-primary/50 transition-colors">
      <CardHeader className="pb-2">
        <div className="flex items-start justify-between">
          <div className="flex-1 min-w-0">
            <CardTitle className="flex items-center gap-2 truncate">
              <span className="truncate">{server.name}</span>
              {server.suspended && (
                <span className="text-xs px-2 py-0.5 rounded bg-red-500/20 text-red-500 shrink-0">
                  Suspended
                </span>
              )}
            </CardTitle>
            <CardDescription className="flex items-center gap-2 mt-1">
              <Globe className="h-3 w-3 shrink-0" />
              <span className="truncate">
                {server.ip}:{server.port}
              </span>
              <span className="text-xs opacity-50">({server.shortId})</span>
            </CardDescription>
          </div>
          <div className={`px-2 py-1 rounded-full text-xs font-medium ${getStatusColor(server.status)}`}>
            {getStatusText(server.status)}
          </div>
        </div>
      </CardHeader>
      <CardContent className="space-y-4">
        {/* Node & Egg info */}
        <div className="flex items-center justify-between text-sm text-muted-foreground">
          <span className="flex items-center gap-1">
            <Server className="h-3 w-3" />
            {server.nodeName}
          </span>
          {server.eggName && (
            <span className="text-xs bg-muted px-2 py-0.5 rounded">
              {server.eggName}
            </span>
          )}
        </div>

        {/* Resource Usage */}
        <div className="space-y-2">
          {/* Memory */}
          <div className="space-y-1">
            <div className="flex items-center justify-between text-xs">
              <span className="flex items-center gap-1 text-muted-foreground">
                <MemoryStick className="h-3 w-3" />
                RAM
              </span>
              <span>
                {formatBytes(server.memoryUsed * 1024 * 1024)} / {formatBytes(server.memory * 1024 * 1024)}
              </span>
            </div>
            <Progress value={server.memoryPercent} className="h-1.5" />
          </div>

          {/* CPU */}
          <div className="space-y-1">
            <div className="flex items-center justify-between text-xs">
              <span className="flex items-center gap-1 text-muted-foreground">
                <Cpu className="h-3 w-3" />
                CPU
              </span>
              <span>{server.cpuUsed?.toFixed(1) || 0}% / {server.cpu}%</span>
            </div>
            <Progress value={(server.cpuUsed || 0) / (server.cpu / 100)} className="h-1.5" />
          </div>

          {/* Disk */}
          <div className="space-y-1">
            <div className="flex items-center justify-between text-xs">
              <span className="flex items-center gap-1 text-muted-foreground">
                <HardDrive className="h-3 w-3" />
                Disk
              </span>
              <span>
                {formatBytes(server.diskUsed * 1024 * 1024)} / {formatBytes(server.disk * 1024 * 1024)}
              </span>
            </div>
            <Progress value={server.diskPercent} className="h-1.5" />
          </div>
        </div>

        {/* Actions */}
        <div className="flex items-center justify-between pt-2 border-t">
          <div className="flex gap-1">
            {server.status === "offline" ? (
              <Button
                variant="ghost"
                size="icon"
                onClick={() => powerMutation.mutate("start")}
                disabled={powerMutation.isPending || server.suspended}
                title="Başlat"
              >
                <Play className="h-4 w-4 text-emerald-500" />
              </Button>
            ) : (
              <>
                <Button
                  variant="ghost"
                  size="icon"
                  onClick={() => powerMutation.mutate("restart")}
                  disabled={powerMutation.isPending}
                  title="Yeniden Başlat"
                >
                  <RotateCw className="h-4 w-4 text-yellow-500" />
                </Button>
                <Button
                  variant="ghost"
                  size="icon"
                  onClick={() => powerMutation.mutate("stop")}
                  disabled={powerMutation.isPending}
                  title="Durdur"
                >
                  <Square className="h-4 w-4 text-red-500" />
                </Button>
              </>
            )}
          </div>

          <div className="flex gap-1">
            <Link href={`/admin/dedicated-servers/${server.id}/console`}>
              <Button variant="ghost" size="icon" title="Konsol">
                <Terminal className="h-4 w-4" />
              </Button>
            </Link>
            <Link href={`/admin/dedicated-servers/${server.id}/files`}>
              <Button variant="ghost" size="icon" title="Dosyalar">
                <FolderOpen className="h-4 w-4" />
              </Button>
            </Link>
            <Link href={`/admin/dedicated-servers/${server.id}`}>
              <Button variant="ghost" size="icon" title="Detaylar">
                <ExternalLink className="h-4 w-4" />
              </Button>
            </Link>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}

// Create Server Modal
function CreateServerModal({
  nodes,
  onClose,
  onSuccess,
}: {
  nodes: NodeWithStats[];
  onClose: () => void;
  onSuccess: () => void;
}) {
  const [step, setStep] = useState(1);
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    nodeId: 0,
    eggId: 0,
    memory: 1024,
    disk: 10240,
    cpu: 100,
    io: 500,
    swap: 0,
    allocationId: 0,
    image: "",
    startupCommand: "",
    environment: {} as Record<string, string>,
  });
  const [error, setError] = useState("");

  const { data: allocationsData } = useQuery({
    queryKey: ["allocations", formData.nodeId],
    queryFn: () => api.getNodeAllocations(formData.nodeId),
    enabled: formData.nodeId > 0,
  });

  const { data: eggsData } = useQuery({
    queryKey: ["eggs"],
    queryFn: () => api.getEggs(),
  });

  const availableAllocations = allocationsData?.allocations?.filter((a) => !a.assigned) || [];
  const eggs = eggsData?.eggs || [];

  const createMutation = useMutation({
    mutationFn: () => api.createDedicatedServer(formData),
    onSuccess: () => onSuccess(),
    onError: (err: Error) => setError(err.message),
  });

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div className="bg-card border rounded-lg w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <div className="p-6 border-b">
          <h2 className="text-xl font-bold">Yeni Sunucu Oluştur</h2>
          <div className="flex gap-2 mt-2">
            {[1, 2, 3].map((s) => (
              <div
                key={s}
                className={`h-1 flex-1 rounded ${
                  s <= step ? "bg-primary" : "bg-muted"
                }`}
              />
            ))}
          </div>
        </div>

        <div className="p-6 space-y-4">
          {error && (
            <div className="p-3 bg-red-500/20 text-red-500 rounded-lg text-sm">
              {error}
            </div>
          )}

          {step === 1 && (
            <>
              <h3 className="font-medium">Temel Bilgiler</h3>
              <div className="space-y-4">
                <div className="space-y-2">
                  <label className="text-sm font-medium">Sunucu Adı *</label>
                  <input
                    type="text"
                    value={formData.name}
                    onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                    className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    placeholder="My Server"
                  />
                </div>

                <div className="space-y-2">
                  <label className="text-sm font-medium">Açıklama</label>
                  <textarea
                    value={formData.description}
                    onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                    className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    rows={2}
                  />
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">Node *</label>
                    <select
                      value={formData.nodeId}
                      onChange={(e) => setFormData({ ...formData, nodeId: Number(e.target.value), allocationId: 0 })}
                      className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    >
                      <option value={0}>Node Seçin</option>
                      {nodes.filter((n) => n.status === "online" && !n.maintenanceMode).map((node) => (
                        <option key={node.id} value={node.id}>
                          {node.name} ({node.fqdn})
                        </option>
                      ))}
                    </select>
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">Allocation (IP:Port) *</label>
                    <select
                      value={formData.allocationId}
                      onChange={(e) => setFormData({ ...formData, allocationId: Number(e.target.value) })}
                      className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                      disabled={!formData.nodeId}
                    >
                      <option value={0}>Allocation Seçin</option>
                      {availableAllocations.map((alloc) => (
                        <option key={alloc.id} value={alloc.id}>
                          {alloc.ip}:{alloc.port}
                        </option>
                      ))}
                    </select>
                  </div>
                </div>
              </div>
            </>
          )}

          {step === 2 && (
            <>
              <h3 className="font-medium">Kaynaklar</h3>
              <div className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">RAM (MB) *</label>
                    <input
                      type="number"
                      value={formData.memory}
                      onChange={(e) => setFormData({ ...formData, memory: Number(e.target.value) })}
                      className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    />
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">Disk (MB) *</label>
                    <input
                      type="number"
                      value={formData.disk}
                      onChange={(e) => setFormData({ ...formData, disk: Number(e.target.value) })}
                      className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    />
                  </div>
                </div>

                <div className="grid grid-cols-3 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">CPU (%)</label>
                    <input
                      type="number"
                      value={formData.cpu}
                      onChange={(e) => setFormData({ ...formData, cpu: Number(e.target.value) })}
                      className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    />
                    <p className="text-xs text-muted-foreground">100 = 1 çekirdek</p>
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">IO</label>
                    <input
                      type="number"
                      value={formData.io}
                      onChange={(e) => setFormData({ ...formData, io: Number(e.target.value) })}
                      className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    />
                    <p className="text-xs text-muted-foreground">10-1000</p>
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-medium">Swap (MB)</label>
                    <input
                      type="number"
                      value={formData.swap}
                      onChange={(e) => setFormData({ ...formData, swap: Number(e.target.value) })}
                      className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    />
                  </div>
                </div>
              </div>
            </>
          )}

          {step === 3 && (
            <>
              <h3 className="font-medium">Yapılandırma</h3>
              <div className="space-y-4">
                <div className="space-y-2">
                  <label className="text-sm font-medium">Egg (Şablon)</label>
                  <select
                    value={formData.eggId}
                    onChange={(e) => setFormData({ ...formData, eggId: Number(e.target.value) })}
                    className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                  >
                    <option value={0}>Egg Seçin (Opsiyonel)</option>
                    {eggs.map((egg) => (
                      <option key={egg.id} value={egg.id}>
                        {egg.name}
                      </option>
                    ))}
                  </select>
                </div>

                <div className="space-y-2">
                  <label className="text-sm font-medium">Docker Image</label>
                  <input
                    type="text"
                    value={formData.image}
                    onChange={(e) => setFormData({ ...formData, image: e.target.value })}
                    className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    placeholder="ghcr.io/pterodactyl/yolks:java_17"
                  />
                </div>

                <div className="space-y-2">
                  <label className="text-sm font-medium">Başlangıç Komutu</label>
                  <input
                    type="text"
                    value={formData.startupCommand}
                    onChange={(e) => setFormData({ ...formData, startupCommand: e.target.value })}
                    className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                    placeholder="java -Xms128M -Xmx{{SERVER_MEMORY}}M -jar server.jar"
                  />
                </div>
              </div>
            </>
          )}
        </div>

        <div className="p-6 border-t flex justify-between">
          <div>
            {step > 1 && (
              <Button variant="outline" onClick={() => setStep(step - 1)}>
                Geri
              </Button>
            )}
          </div>
          <div className="flex gap-2">
            <Button variant="outline" onClick={onClose}>
              İptal
            </Button>
            {step < 3 ? (
              <Button
                onClick={() => setStep(step + 1)}
                disabled={
                  (step === 1 && (!formData.name || !formData.nodeId || !formData.allocationId))
                }
              >
                İleri
              </Button>
            ) : (
              <Button
                onClick={() => createMutation.mutate()}
                disabled={createMutation.isPending}
              >
                {createMutation.isPending ? "Oluşturuluyor..." : "Oluştur"}
              </Button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
