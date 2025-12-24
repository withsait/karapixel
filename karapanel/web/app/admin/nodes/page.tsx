"use client";

import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { api, NodeWithStats, Location, Allocation } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { formatBytes } from "@/lib/utils";
import {
  Server,
  Plus,
  Settings,
  Trash2,
  RefreshCw,
  HardDrive,
  Cpu,
  MemoryStick,
  Network,
  MapPin,
  AlertTriangle,
  Check,
  X,
  Key,
  Globe,
} from "lucide-react";

export default function NodesPage() {
  const queryClient = useQueryClient();
  const [showCreateModal, setShowCreateModal] = useState(false);
  const [showAllocationsModal, setShowAllocationsModal] = useState<number | null>(null);
  const [selectedNode, setSelectedNode] = useState<NodeWithStats | null>(null);

  const { data: nodesData, isLoading, refetch } = useQuery({
    queryKey: ["nodes"],
    queryFn: () => api.getNodes(true),
  });

  const { data: locationsData } = useQuery({
    queryKey: ["locations"],
    queryFn: () => api.getLocations(),
  });

  const deleteNodeMutation = useMutation({
    mutationFn: (id: number) => api.deleteNode(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["nodes"] });
    },
  });

  const nodes = nodesData?.nodes || [];
  const locations = locationsData?.locations || [];

  const getStatusColor = (status: string) => {
    switch (status) {
      case "online":
        return "text-emerald-500";
      case "offline":
        return "text-red-500";
      case "maintenance":
        return "text-yellow-500";
      default:
        return "text-gray-500";
    }
  };

  const getStatusBg = (status: string) => {
    switch (status) {
      case "online":
        return "bg-emerald-500/20";
      case "offline":
        return "bg-red-500/20";
      case "maintenance":
        return "bg-yellow-500/20";
      default:
        return "bg-gray-500/20";
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
          <h1 className="text-3xl font-bold">Nodes</h1>
          <p className="text-muted-foreground">Sunucu makinelerini yönetin</p>
        </div>
        <div className="flex gap-2">
          <Button variant="outline" size="sm" onClick={() => refetch()}>
            <RefreshCw className="h-4 w-4 mr-2" />
            Yenile
          </Button>
          <Button size="sm" onClick={() => setShowCreateModal(true)}>
            <Plus className="h-4 w-4 mr-2" />
            Node Ekle
          </Button>
        </div>
      </div>

      {/* Stats Overview */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-primary/20">
                <Server className="h-6 w-6 text-primary" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Toplam Node</p>
                <p className="text-2xl font-bold">{nodes.length}</p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-emerald-500/20">
                <Check className="h-6 w-6 text-emerald-500" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Online</p>
                <p className="text-2xl font-bold">
                  {nodes.filter((n) => n.status === "online").length}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-accent/20">
                <HardDrive className="h-6 w-6 text-accent" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Toplam Sunucu</p>
                <p className="text-2xl font-bold">
                  {nodes.reduce((acc, n) => acc + n.serverCount, 0)}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center gap-4">
              <div className="p-3 rounded-full bg-yellow-500/20">
                <AlertTriangle className="h-6 w-6 text-yellow-500" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Bakımda</p>
                <p className="text-2xl font-bold">
                  {nodes.filter((n) => n.maintenanceMode).length}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Nodes Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {nodes.map((node) => (
          <Card
            key={node.id}
            className="hover:border-primary/50 transition-colors"
          >
            <CardHeader className="pb-2">
              <div className="flex items-start justify-between">
                <div className="flex items-center gap-3">
                  <div className={`p-2 rounded-lg ${getStatusBg(node.status)}`}>
                    <Server className={`h-5 w-5 ${getStatusColor(node.status)}`} />
                  </div>
                  <div>
                    <CardTitle className="flex items-center gap-2">
                      {node.name}
                      {node.maintenanceMode && (
                        <span className="text-xs px-2 py-0.5 rounded bg-yellow-500/20 text-yellow-500">
                          Bakım
                        </span>
                      )}
                    </CardTitle>
                    <CardDescription className="flex items-center gap-1">
                      <Globe className="h-3 w-3" />
                      {node.fqdn}:{node.daemonPort}
                    </CardDescription>
                  </div>
                </div>
                <div className="flex gap-1">
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => setShowAllocationsModal(node.id)}
                    title="Allocations"
                  >
                    <Network className="h-4 w-4" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => setSelectedNode(node)}
                    title="Ayarlar"
                  >
                    <Settings className="h-4 w-4" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    onClick={() => {
                      if (confirm("Bu node'u silmek istediğinizden emin misiniz?")) {
                        deleteNodeMutation.mutate(node.id);
                      }
                    }}
                    disabled={node.serverCount > 0}
                    title={node.serverCount > 0 ? "Sunucular mevcut" : "Sil"}
                  >
                    <Trash2 className="h-4 w-4 text-red-500" />
                  </Button>
                </div>
              </div>
            </CardHeader>
            <CardContent className="space-y-4">
              {/* Location & Stats */}
              <div className="flex items-center justify-between text-sm">
                <div className="flex items-center gap-2 text-muted-foreground">
                  <MapPin className="h-4 w-4" />
                  {node.locationName || "Konum belirtilmemiş"}
                </div>
                <div className="flex items-center gap-4">
                  <span className="text-muted-foreground">
                    {node.onlineServers}/{node.serverCount} sunucu
                  </span>
                  <span className="text-muted-foreground">
                    {node.allocatedPorts} port
                  </span>
                </div>
              </div>

              {/* Resource Usage */}
              <div className="space-y-3">
                {/* Memory */}
                <div className="space-y-1">
                  <div className="flex items-center justify-between text-sm">
                    <span className="flex items-center gap-2 text-muted-foreground">
                      <MemoryStick className="h-4 w-4" />
                      RAM
                    </span>
                    <span>
                      {formatBytes(node.memoryUsed * 1024 * 1024)} / {formatBytes(node.memory * 1024 * 1024)}
                    </span>
                  </div>
                  <Progress value={node.memoryPercent} className="h-2" />
                </div>

                {/* Disk */}
                <div className="space-y-1">
                  <div className="flex items-center justify-between text-sm">
                    <span className="flex items-center gap-2 text-muted-foreground">
                      <HardDrive className="h-4 w-4" />
                      Disk
                    </span>
                    <span>
                      {formatBytes(node.diskUsed * 1024 * 1024)} / {formatBytes(node.disk * 1024 * 1024)}
                    </span>
                  </div>
                  <Progress value={node.diskPercent} className="h-2" />
                </div>

                {/* CPU */}
                <div className="space-y-1">
                  <div className="flex items-center justify-between text-sm">
                    <span className="flex items-center gap-2 text-muted-foreground">
                      <Cpu className="h-4 w-4" />
                      CPU
                    </span>
                    <span>{node.cpuPercent.toFixed(1)}%</span>
                  </div>
                  <Progress value={node.cpuPercent} className="h-2" />
                </div>
              </div>

              {/* Over-allocation info */}
              {(node.memoryOveralloc > 0 || node.diskOveralloc > 0) && (
                <div className="flex gap-4 text-xs text-muted-foreground pt-2 border-t">
                  {node.memoryOveralloc > 0 && (
                    <span>RAM Overalloc: {node.memoryOveralloc}%</span>
                  )}
                  {node.diskOveralloc > 0 && (
                    <span>Disk Overalloc: {node.diskOveralloc}%</span>
                  )}
                </div>
              )}
            </CardContent>
          </Card>
        ))}
      </div>

      {nodes.length === 0 && (
        <Card className="border-dashed">
          <CardContent className="flex flex-col items-center justify-center py-12">
            <Server className="h-12 w-12 text-muted-foreground mb-4" />
            <p className="text-lg font-medium">Henüz node eklenmemiş</p>
            <p className="text-sm text-muted-foreground mb-4">
              Sunucularınızı yönetmek için bir node ekleyin
            </p>
            <Button onClick={() => setShowCreateModal(true)}>
              <Plus className="h-4 w-4 mr-2" />
              İlk Node'u Ekle
            </Button>
          </CardContent>
        </Card>
      )}

      {/* Create Node Modal */}
      {showCreateModal && (
        <CreateNodeModal
          locations={locations}
          onClose={() => setShowCreateModal(false)}
          onSuccess={() => {
            setShowCreateModal(false);
            queryClient.invalidateQueries({ queryKey: ["nodes"] });
          }}
        />
      )}

      {/* Allocations Modal */}
      {showAllocationsModal && (
        <AllocationsModal
          nodeId={showAllocationsModal}
          onClose={() => setShowAllocationsModal(null)}
        />
      )}

      {/* Edit Node Modal */}
      {selectedNode && (
        <EditNodeModal
          node={selectedNode}
          locations={locations}
          onClose={() => setSelectedNode(null)}
          onSuccess={() => {
            setSelectedNode(null);
            queryClient.invalidateQueries({ queryKey: ["nodes"] });
          }}
        />
      )}
    </div>
  );
}

// Create Node Modal Component
function CreateNodeModal({
  locations,
  onClose,
  onSuccess,
}: {
  locations: Location[];
  onClose: () => void;
  onSuccess: () => void;
}) {
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    fqdn: "",
    scheme: "https",
    daemonPort: 8080,
    memory: 8192,
    memoryOveralloc: 0,
    disk: 102400,
    diskOveralloc: 0,
    uploadLimit: 0,
    downloadLimit: 0,
    locationId: 0,
  });
  const [error, setError] = useState("");

  const createMutation = useMutation({
    mutationFn: () => api.createNode(formData),
    onSuccess: () => onSuccess(),
    onError: (err: Error) => setError(err.message),
  });

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div className="bg-card border rounded-lg w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <div className="p-6 border-b">
          <h2 className="text-xl font-bold">Yeni Node Ekle</h2>
          <p className="text-sm text-muted-foreground">
            Sunucu yönetimi için yeni bir makine ekleyin
          </p>
        </div>

        <div className="p-6 space-y-4">
          {error && (
            <div className="p-3 bg-red-500/20 text-red-500 rounded-lg text-sm">
              {error}
            </div>
          )}

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">İsim *</label>
              <input
                type="text"
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                placeholder="Node-1"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Konum</label>
              <select
                value={formData.locationId}
                onChange={(e) => setFormData({ ...formData, locationId: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              >
                <option value={0}>Konum Seçin</option>
                {locations.map((loc) => (
                  <option key={loc.id} value={loc.id}>
                    {loc.long} ({loc.short})
                  </option>
                ))}
              </select>
            </div>
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium">Açıklama</label>
            <textarea
              value={formData.description}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
              className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              rows={2}
              placeholder="Node hakkında açıklama..."
            />
          </div>

          <div className="grid grid-cols-3 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">FQDN/IP *</label>
              <input
                type="text"
                value={formData.fqdn}
                onChange={(e) => setFormData({ ...formData, fqdn: e.target.value })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
                placeholder="node.example.com"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Protokol</label>
              <select
                value={formData.scheme}
                onChange={(e) => setFormData({ ...formData, scheme: e.target.value })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              >
                <option value="https">HTTPS</option>
                <option value="http">HTTP</option>
              </select>
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Daemon Port</label>
              <input
                type="number"
                value={formData.daemonPort}
                onChange={(e) => setFormData({ ...formData, daemonPort: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>
          </div>

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
              <label className="text-sm font-medium">RAM Overalloc (%)</label>
              <input
                type="number"
                value={formData.memoryOveralloc}
                onChange={(e) => setFormData({ ...formData, memoryOveralloc: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">Disk (MB) *</label>
              <input
                type="number"
                value={formData.disk}
                onChange={(e) => setFormData({ ...formData, disk: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Disk Overalloc (%)</label>
              <input
                type="number"
                value={formData.diskOveralloc}
                onChange={(e) => setFormData({ ...formData, diskOveralloc: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>
          </div>
        </div>

        <div className="p-6 border-t flex justify-end gap-2">
          <Button variant="outline" onClick={onClose}>
            İptal
          </Button>
          <Button
            onClick={() => createMutation.mutate()}
            disabled={createMutation.isPending || !formData.name || !formData.fqdn}
          >
            {createMutation.isPending ? "Oluşturuluyor..." : "Oluştur"}
          </Button>
        </div>
      </div>
    </div>
  );
}

// Edit Node Modal Component
function EditNodeModal({
  node,
  locations,
  onClose,
  onSuccess,
}: {
  node: NodeWithStats;
  locations: Location[];
  onClose: () => void;
  onSuccess: () => void;
}) {
  const [formData, setFormData] = useState({
    name: node.name,
    description: node.description || "",
    fqdn: node.fqdn,
    scheme: node.scheme,
    daemonPort: node.daemonPort,
    memory: node.memory,
    memoryOveralloc: node.memoryOveralloc,
    disk: node.disk,
    diskOveralloc: node.diskOveralloc,
    uploadLimit: node.uploadLimit,
    downloadLimit: node.downloadLimit,
    locationId: node.locationId || 0,
    maintenanceMode: node.maintenanceMode,
  });
  const [error, setError] = useState("");
  const [showToken, setShowToken] = useState(false);
  const [newToken, setNewToken] = useState("");

  const updateMutation = useMutation({
    mutationFn: () => api.updateNode(node.id, formData),
    onSuccess: () => onSuccess(),
    onError: (err: Error) => setError(err.message),
  });

  const regenerateTokenMutation = useMutation({
    mutationFn: () => api.regenerateNodeToken(node.id),
    onSuccess: (data) => {
      setNewToken(data.token);
      setShowToken(true);
    },
    onError: (err: Error) => setError(err.message),
  });

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div className="bg-card border rounded-lg w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        <div className="p-6 border-b">
          <h2 className="text-xl font-bold">Node Düzenle</h2>
          <p className="text-sm text-muted-foreground">{node.name}</p>
        </div>

        <div className="p-6 space-y-4">
          {error && (
            <div className="p-3 bg-red-500/20 text-red-500 rounded-lg text-sm">
              {error}
            </div>
          )}

          {showToken && newToken && (
            <div className="p-3 bg-emerald-500/20 text-emerald-500 rounded-lg text-sm">
              <p className="font-medium mb-1">Yeni Token Oluşturuldu:</p>
              <code className="block bg-black/30 p-2 rounded text-xs break-all">
                {newToken}
              </code>
            </div>
          )}

          {/* Same form fields as create modal */}
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">İsim *</label>
              <input
                type="text"
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Konum</label>
              <select
                value={formData.locationId}
                onChange={(e) => setFormData({ ...formData, locationId: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              >
                <option value={0}>Konum Seçin</option>
                {locations.map((loc) => (
                  <option key={loc.id} value={loc.id}>
                    {loc.long} ({loc.short})
                  </option>
                ))}
              </select>
            </div>
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

          <div className="grid grid-cols-3 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">FQDN/IP *</label>
              <input
                type="text"
                value={formData.fqdn}
                onChange={(e) => setFormData({ ...formData, fqdn: e.target.value })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Protokol</label>
              <select
                value={formData.scheme}
                onChange={(e) => setFormData({ ...formData, scheme: e.target.value })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              >
                <option value="https">HTTPS</option>
                <option value="http">HTTP</option>
              </select>
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Daemon Port</label>
              <input
                type="number"
                value={formData.daemonPort}
                onChange={(e) => setFormData({ ...formData, daemonPort: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">RAM (MB)</label>
              <input
                type="number"
                value={formData.memory}
                onChange={(e) => setFormData({ ...formData, memory: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Disk (MB)</label>
              <input
                type="number"
                value={formData.disk}
                onChange={(e) => setFormData({ ...formData, disk: Number(e.target.value) })}
                className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none"
              />
            </div>
          </div>

          <div className="flex items-center gap-4">
            <label className="flex items-center gap-2">
              <input
                type="checkbox"
                checked={formData.maintenanceMode}
                onChange={(e) => setFormData({ ...formData, maintenanceMode: e.target.checked })}
                className="rounded border-gray-600"
              />
              <span className="text-sm">Bakım Modu</span>
            </label>
          </div>

          <div className="pt-4 border-t">
            <Button
              variant="outline"
              onClick={() => regenerateTokenMutation.mutate()}
              disabled={regenerateTokenMutation.isPending}
            >
              <Key className="h-4 w-4 mr-2" />
              {regenerateTokenMutation.isPending ? "Oluşturuluyor..." : "Token Yenile"}
            </Button>
          </div>
        </div>

        <div className="p-6 border-t flex justify-end gap-2">
          <Button variant="outline" onClick={onClose}>
            İptal
          </Button>
          <Button
            onClick={() => updateMutation.mutate()}
            disabled={updateMutation.isPending}
          >
            {updateMutation.isPending ? "Kaydediliyor..." : "Kaydet"}
          </Button>
        </div>
      </div>
    </div>
  );
}

// Allocations Modal Component
function AllocationsModal({
  nodeId,
  onClose,
}: {
  nodeId: number;
  onClose: () => void;
}) {
  const queryClient = useQueryClient();
  const [newAlloc, setNewAlloc] = useState({ ip: "", port: 0, portRange: "" });
  const [error, setError] = useState("");

  const { data, isLoading, refetch } = useQuery({
    queryKey: ["allocations", nodeId],
    queryFn: () => api.getNodeAllocations(nodeId),
  });

  const createMutation = useMutation({
    mutationFn: () => api.createAllocation(nodeId, newAlloc),
    onSuccess: () => {
      setNewAlloc({ ip: "", port: 0, portRange: "" });
      refetch();
    },
    onError: (err: Error) => setError(err.message),
  });

  const deleteMutation = useMutation({
    mutationFn: (allocId: number) => api.deleteAllocation(nodeId, allocId),
    onSuccess: () => refetch(),
  });

  const grouped = data?.grouped || {};

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div className="bg-card border rounded-lg w-full max-w-3xl max-h-[90vh] overflow-hidden flex flex-col">
        <div className="p-6 border-b">
          <h2 className="text-xl font-bold">Port Allocations</h2>
          <p className="text-sm text-muted-foreground">
            {data?.total || 0} allocation tanımlı
          </p>
        </div>

        <div className="p-6 flex-1 overflow-y-auto space-y-4">
          {error && (
            <div className="p-3 bg-red-500/20 text-red-500 rounded-lg text-sm">
              {error}
            </div>
          )}

          {/* Add new allocation */}
          <div className="p-4 border rounded-lg space-y-3">
            <h3 className="font-medium">Yeni Allocation Ekle</h3>
            <div className="grid grid-cols-3 gap-3">
              <div className="space-y-1">
                <label className="text-xs text-muted-foreground">IP Adresi</label>
                <input
                  type="text"
                  value={newAlloc.ip}
                  onChange={(e) => setNewAlloc({ ...newAlloc, ip: e.target.value })}
                  className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none text-sm"
                  placeholder="0.0.0.0"
                />
              </div>
              <div className="space-y-1">
                <label className="text-xs text-muted-foreground">Tek Port</label>
                <input
                  type="number"
                  value={newAlloc.port || ""}
                  onChange={(e) => setNewAlloc({ ...newAlloc, port: Number(e.target.value), portRange: "" })}
                  className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none text-sm"
                  placeholder="25565"
                />
              </div>
              <div className="space-y-1">
                <label className="text-xs text-muted-foreground">veya Port Aralığı</label>
                <input
                  type="text"
                  value={newAlloc.portRange}
                  onChange={(e) => setNewAlloc({ ...newAlloc, portRange: e.target.value, port: 0 })}
                  className="w-full px-3 py-2 rounded-lg bg-background border focus:border-primary outline-none text-sm"
                  placeholder="25565-25575"
                />
              </div>
            </div>
            <Button
              size="sm"
              onClick={() => createMutation.mutate()}
              disabled={createMutation.isPending || !newAlloc.ip || (!newAlloc.port && !newAlloc.portRange)}
            >
              <Plus className="h-4 w-4 mr-2" />
              Ekle
            </Button>
          </div>

          {/* Allocations by IP */}
          {isLoading ? (
            <div className="flex justify-center py-8">
              <RefreshCw className="h-6 w-6 animate-spin" />
            </div>
          ) : (
            <div className="space-y-4">
              {Object.entries(grouped).map(([ip, allocations]) => (
                <div key={ip} className="border rounded-lg overflow-hidden">
                  <div className="px-4 py-2 bg-muted/50 font-medium text-sm">
                    {ip}
                    <span className="text-muted-foreground ml-2">
                      ({(allocations as Allocation[]).length} port)
                    </span>
                  </div>
                  <div className="p-2 flex flex-wrap gap-2">
                    {(allocations as Allocation[]).map((alloc) => (
                      <div
                        key={alloc.id}
                        className={`inline-flex items-center gap-1 px-2 py-1 rounded text-sm ${
                          alloc.assigned
                            ? "bg-primary/20 text-primary"
                            : "bg-muted"
                        }`}
                      >
                        <span>{alloc.port}</span>
                        {!alloc.assigned && (
                          <button
                            onClick={() => deleteMutation.mutate(alloc.id)}
                            className="hover:text-red-500"
                          >
                            <X className="h-3 w-3" />
                          </button>
                        )}
                      </div>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          )}

          {Object.keys(grouped).length === 0 && !isLoading && (
            <div className="text-center py-8 text-muted-foreground">
              <Network className="h-12 w-12 mx-auto mb-2 opacity-50" />
              <p>Henüz allocation eklenmemiş</p>
            </div>
          )}
        </div>

        <div className="p-6 border-t flex justify-end">
          <Button variant="outline" onClick={onClose}>
            Kapat
          </Button>
        </div>
      </div>
    </div>
  );
}
