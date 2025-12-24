"use client";

import { useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { formatBytes, formatUptime } from "@/lib/utils";
import Link from "next/link";
import {
  Server,
  ArrowLeft,
  RefreshCw,
  HardDrive,
  Cpu,
  MemoryStick,
  Play,
  Square,
  RotateCw,
  AlertTriangle,
  Terminal,
  FolderOpen,
  Settings,
  Network,
  Trash2,
  Power,
  Clock,
  Globe,
  Zap,
  Ban,
  CheckCircle,
  Activity,
  Save,
} from "lucide-react";

export default function ServerDetailPage() {
  const params = useParams();
  const router = useRouter();
  const queryClient = useQueryClient();
  const serverId = Number(params.id);

  const [activeTab, setActiveTab] = useState<"overview" | "resources" | "startup" | "settings">("overview");
  const [editMode, setEditMode] = useState(false);

  const { data: server, isLoading, refetch } = useQuery({
    queryKey: ["dedicated-server", serverId],
    queryFn: () => api.getDedicatedServer(serverId),
    refetchInterval: 5000,
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

  const suspendMutation = useMutation({
    mutationFn: (suspended: boolean) =>
      suspended ? api.suspendDedicatedServer(serverId) : api.unsuspendDedicatedServer(serverId),
    onSuccess: () => refetch(),
  });

  const deleteMutation = useMutation({
    mutationFn: () => api.deleteDedicatedServer(serverId),
    onSuccess: () => router.push("/admin/dedicated-servers"),
  });

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-64">
        <RefreshCw className="h-8 w-8 animate-spin text-primary" />
      </div>
    );
  }

  if (!server) {
    return (
      <div className="flex flex-col items-center justify-center h-64">
        <AlertTriangle className="h-12 w-12 text-yellow-500 mb-4" />
        <p className="text-lg font-medium">Sunucu bulunamadı</p>
        <Link href="/admin/dedicated-servers">
          <Button variant="outline" className="mt-4">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Geri Dön
          </Button>
        </Link>
      </div>
    );
  }

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
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-start justify-between">
        <div className="flex items-start gap-4">
          <Link href="/admin/dedicated-servers">
            <Button variant="ghost" size="icon">
              <ArrowLeft className="h-5 w-5" />
            </Button>
          </Link>
          <div>
            <div className="flex items-center gap-3">
              <h1 className="text-3xl font-bold">{server.name}</h1>
              <span className={`px-3 py-1 rounded-full text-sm font-medium ${getStatusColor(server.status)}`}>
                {getStatusText(server.status)}
              </span>
              {server.suspended && (
                <span className="px-3 py-1 rounded-full text-sm font-medium bg-red-500/20 text-red-500">
                  Suspended
                </span>
              )}
            </div>
            <div className="flex items-center gap-4 mt-2 text-sm text-muted-foreground">
              <span className="flex items-center gap-1">
                <Globe className="h-4 w-4" />
                {server.ip}:{server.port}
              </span>
              <span className="flex items-center gap-1">
                <Server className="h-4 w-4" />
                {server.nodeName}
              </span>
              {server.eggName && (
                <span className="flex items-center gap-1">
                  <Zap className="h-4 w-4" />
                  {server.eggName}
                </span>
              )}
              <span className="text-xs opacity-50">ID: {server.shortId}</span>
            </div>
          </div>
        </div>

        <div className="flex gap-2">
          {/* Power Controls */}
          {server.status === "offline" ? (
            <Button
              onClick={() => powerMutation.mutate("start")}
              disabled={powerMutation.isPending || server.suspended}
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

      {/* Quick Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-accent/20">
                <MemoryStick className="h-6 w-6 text-accent" />
              </div>
              <div className="flex-1">
                <p className="text-sm text-muted-foreground">RAM</p>
                <p className="text-xl font-bold">
                  {formatBytes(server.memoryUsed * 1024 * 1024)}
                </p>
                <Progress value={server.memoryPercent} className="h-1.5 mt-1" />
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-primary/20">
                <Cpu className="h-6 w-6 text-primary" />
              </div>
              <div className="flex-1">
                <p className="text-sm text-muted-foreground">CPU</p>
                <p className="text-xl font-bold">{server.cpuUsed?.toFixed(1) || 0}%</p>
                <Progress value={(server.cpuUsed || 0)} className="h-1.5 mt-1" />
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
              <div className="flex-1">
                <p className="text-sm text-muted-foreground">Disk</p>
                <p className="text-xl font-bold">
                  {formatBytes(server.diskUsed * 1024 * 1024)}
                </p>
                <Progress value={server.diskPercent} className="h-1.5 mt-1" />
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-emerald-500/20">
                <Clock className="h-6 w-6 text-emerald-500" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Uptime</p>
                <p className="text-xl font-bold">
                  {server.status === "running" ? formatUptime(server.uptime) : "-"}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Quick Actions */}
      <div className="flex gap-2 flex-wrap">
        <Link href={`/admin/dedicated-servers/${serverId}/console`}>
          <Button variant="outline">
            <Terminal className="h-4 w-4 mr-2" />
            Konsol
          </Button>
        </Link>
        <Link href={`/admin/dedicated-servers/${serverId}/files`}>
          <Button variant="outline">
            <FolderOpen className="h-4 w-4 mr-2" />
            Dosyalar
          </Button>
        </Link>
        <Button
          variant="outline"
          onClick={() => suspendMutation.mutate(!server.suspended)}
          disabled={suspendMutation.isPending}
        >
          {server.suspended ? (
            <>
              <CheckCircle className="h-4 w-4 mr-2" />
              Askıyı Kaldır
            </>
          ) : (
            <>
              <Ban className="h-4 w-4 mr-2" />
              Askıya Al
            </>
          )}
        </Button>
        <Button
          variant="destructive"
          onClick={() => {
            if (confirm("Bu sunucuyu silmek istediğinizden emin misiniz? Bu işlem geri alınamaz.")) {
              deleteMutation.mutate();
            }
          }}
          disabled={deleteMutation.isPending || server.status === "running"}
        >
          <Trash2 className="h-4 w-4 mr-2" />
          Sil
        </Button>
      </div>

      {/* Tabs */}
      <div className="border-b">
        <div className="flex gap-4">
          {[
            { id: "overview", label: "Genel Bakış", icon: Activity },
            { id: "resources", label: "Kaynaklar", icon: Cpu },
            { id: "startup", label: "Başlangıç", icon: Zap },
            { id: "settings", label: "Ayarlar", icon: Settings },
          ].map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id as typeof activeTab)}
              className={`flex items-center gap-2 px-4 py-3 border-b-2 transition-colors ${
                activeTab === tab.id
                  ? "border-primary text-primary"
                  : "border-transparent text-muted-foreground hover:text-foreground"
              }`}
            >
              <tab.icon className="h-4 w-4" />
              {tab.label}
            </button>
          ))}
        </div>
      </div>

      {/* Tab Content */}
      {activeTab === "overview" && (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* Server Info */}
          <Card>
            <CardHeader>
              <CardTitle>Sunucu Bilgileri</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">UUID</p>
                  <p className="font-mono text-sm">{server.uuid}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Short ID</p>
                  <p className="font-mono text-sm">{server.shortId}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Node</p>
                  <p>{server.nodeName}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Egg</p>
                  <p>{server.eggName || "Belirtilmemiş"}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Oluşturulma</p>
                  <p>{new Date(server.createdAt).toLocaleString("tr-TR")}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Güncelleme</p>
                  <p>{new Date(server.updatedAt).toLocaleString("tr-TR")}</p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Network */}
          <Card>
            <CardHeader>
              <CardTitle>Ağ</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div>
                <p className="text-sm text-muted-foreground">Birincil Adres</p>
                <p className="font-mono">{server.ip}:{server.port}</p>
              </div>
              {server.allocations && server.allocations.length > 1 && (
                <div>
                  <p className="text-sm text-muted-foreground mb-2">Ek Portlar</p>
                  <div className="flex flex-wrap gap-2">
                    {server.allocations.slice(1).map((alloc) => (
                      <span
                        key={alloc.id}
                        className="px-2 py-1 bg-muted rounded text-sm font-mono"
                      >
                        {alloc.ip}:{alloc.port}
                      </span>
                    ))}
                  </div>
                </div>
              )}
              <div className="grid grid-cols-2 gap-4 pt-4 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Gelen</p>
                  <p className="font-bold">{formatBytes(server.networkRx || 0)}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Giden</p>
                  <p className="font-bold">{formatBytes(server.networkTx || 0)}</p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Description */}
          {server.description && (
            <Card className="lg:col-span-2">
              <CardHeader>
                <CardTitle>Açıklama</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-muted-foreground">{server.description}</p>
              </CardContent>
            </Card>
          )}
        </div>
      )}

      {activeTab === "resources" && (
        <Card>
          <CardHeader>
            <CardTitle>Kaynak Limitleri</CardTitle>
            <CardDescription>
              Sunucunun kullanabileceği maksimum kaynak miktarları
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <span className="flex items-center gap-2">
                    <MemoryStick className="h-4 w-4 text-muted-foreground" />
                    RAM
                  </span>
                  <span className="font-bold">{formatBytes(server.memory * 1024 * 1024)}</span>
                </div>
                <Progress value={server.memoryPercent} />
                <p className="text-xs text-muted-foreground text-right">
                  {formatBytes(server.memoryUsed * 1024 * 1024)} kullanılıyor
                </p>
              </div>

              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <span className="flex items-center gap-2">
                    <HardDrive className="h-4 w-4 text-muted-foreground" />
                    Disk
                  </span>
                  <span className="font-bold">{formatBytes(server.disk * 1024 * 1024)}</span>
                </div>
                <Progress value={server.diskPercent} />
                <p className="text-xs text-muted-foreground text-right">
                  {formatBytes(server.diskUsed * 1024 * 1024)} kullanılıyor
                </p>
              </div>

              <div className="space-y-2">
                <div className="flex items-center justify-between">
                  <span className="flex items-center gap-2">
                    <Cpu className="h-4 w-4 text-muted-foreground" />
                    CPU
                  </span>
                  <span className="font-bold">{server.cpu}%</span>
                </div>
                <Progress value={(server.cpuUsed || 0) / (server.cpu / 100)} />
                <p className="text-xs text-muted-foreground text-right">
                  {server.cpuUsed?.toFixed(1) || 0}% kullanılıyor
                </p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground mb-1">IO Weight</p>
                <p className="font-bold">{server.io}</p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground mb-1">Swap</p>
                <p className="font-bold">{formatBytes(server.swap * 1024 * 1024)}</p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground mb-1">CPU Threads</p>
                <p className="font-bold">{server.threads || "Sınırsız"}</p>
              </div>
            </div>
          </CardContent>
        </Card>
      )}

      {activeTab === "startup" && (
        <Card>
          <CardHeader>
            <CardTitle>Başlangıç Yapılandırması</CardTitle>
            <CardDescription>
              Sunucunun nasıl başlatılacağını belirleyen ayarlar
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="space-y-2">
              <label className="text-sm font-medium">Docker Image</label>
              <input
                type="text"
                value={server.image || ""}
                readOnly={!editMode}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none font-mono text-sm"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Başlangıç Komutu</label>
              <textarea
                value={server.startupCommand || ""}
                readOnly={!editMode}
                rows={3}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none font-mono text-sm"
              />
            </div>

            {server.environment && Object.keys(server.environment).length > 0 && (
              <div className="space-y-2">
                <label className="text-sm font-medium">Ortam Değişkenleri</label>
                <div className="space-y-2">
                  {Object.entries(server.environment).map(([key, value]) => (
                    <div key={key} className="flex gap-2">
                      <input
                        type="text"
                        value={key}
                        readOnly
                        className="flex-1 px-3 py-2 rounded-lg bg-muted border font-mono text-sm"
                      />
                      <input
                        type="text"
                        value={value}
                        readOnly={!editMode}
                        className="flex-1 px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none font-mono text-sm"
                      />
                    </div>
                  ))}
                </div>
              </div>
            )}
          </CardContent>
        </Card>
      )}

      {activeTab === "settings" && (
        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle>Genel Ayarlar</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="space-y-2">
                <label className="text-sm font-medium">Sunucu Adı</label>
                <input
                  type="text"
                  value={server.name}
                  readOnly={!editMode}
                  className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                />
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium">Açıklama</label>
                <textarea
                  value={server.description || ""}
                  readOnly={!editMode}
                  rows={2}
                  className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Limitler</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-3 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Backup Limiti</p>
                  <p className="font-bold">{server.backupLimit}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Database Limiti</p>
                  <p className="font-bold">{server.databaseLimit}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Allocation Limiti</p>
                  <p className="font-bold">{server.allocationLimit}</p>
                </div>
              </div>
            </CardContent>
          </Card>

          <Card className="border-red-500/50">
            <CardHeader>
              <CardTitle className="text-red-500">Tehlikeli Alan</CardTitle>
              <CardDescription>
                Bu işlemler geri alınamaz
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex items-center justify-between p-4 border border-red-500/20 rounded-lg">
                <div>
                  <p className="font-medium">Sunucuyu Sil</p>
                  <p className="text-sm text-muted-foreground">
                    Sunucu ve tüm verileri kalıcı olarak silinecek
                  </p>
                </div>
                <Button
                  variant="destructive"
                  onClick={() => {
                    if (confirm("Bu sunucuyu silmek istediğinizden emin misiniz?")) {
                      deleteMutation.mutate();
                    }
                  }}
                  disabled={server.status === "running"}
                >
                  <Trash2 className="h-4 w-4 mr-2" />
                  Sil
                </Button>
              </div>
            </CardContent>
          </Card>
        </div>
      )}
    </div>
  );
}
