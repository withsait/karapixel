const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

class ApiClient {
  private token: string | null = null;

  setToken(token: string) {
    this.token = token;
    if (typeof window !== 'undefined') {
      localStorage.setItem('karapanel_token', token);
      document.cookie = `karapanel_token=${token}; path=/; max-age=${60 * 60 * 24}; SameSite=Strict`;
    }
  }

  getToken(): string | null {
    if (this.token) return this.token;
    if (typeof window !== 'undefined') {
      this.token = localStorage.getItem('karapanel_token');
    }
    return this.token;
  }

  clearToken() {
    this.token = null;
    if (typeof window !== 'undefined') {
      localStorage.removeItem('karapanel_token');
      document.cookie = 'karapanel_token=; path=/; max-age=0';
    }
  }

  private async request<T>(path: string, options: RequestInit = {}): Promise<T> {
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
    };

    const token = this.getToken();
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }

    const response = await fetch(`${API_URL}${path}`, {
      ...options,
      headers: {
        ...headers,
        ...(options.headers as Record<string, string>),
      },
    });

    if (response.status === 401) {
      this.clearToken();
      if (typeof window !== 'undefined') {
        window.location.href = '/login';
      }
      throw new Error('Unauthorized');
    }

    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: 'Request failed' }));
      throw new Error(error.error || 'Request failed');
    }

    return response.json();
  }

  async login(username: string, password: string) {
    const data = await this.request<{ token: string; username: string; role: string }>('/api/auth/login', {
      method: 'POST',
      body: JSON.stringify({ username, password }),
    });
    this.setToken(data.token);
    return data;
  }

  logout() {
    this.clearToken();
  }

  async getServers() {
    return this.request<{ servers: ServerInfo[] }>('/api/servers');
  }

  async getServer(id: string) {
    return this.request<ServerInfo>(`/api/servers/${id}`);
  }

  async startServer(id: string) {
    return this.request<{ message: string }>(`/api/servers/${id}/start`, { method: 'POST' });
  }

  async stopServer(id: string) {
    return this.request<{ message: string }>(`/api/servers/${id}/stop`, { method: 'POST' });
  }

  async restartServer(id: string) {
    return this.request<{ message: string }>(`/api/servers/${id}/restart`, { method: 'POST' });
  }

  async killServer(id: string) {
    return this.request<{ message: string }>(`/api/servers/${id}/kill`, { method: 'POST' });
  }

  async getLogs(id: string, lines: number = 100) {
    return this.request<{ logs: string[] }>(`/api/servers/${id}/logs?lines=${lines}`);
  }

  async listFiles(id: string, path: string = '/') {
    return this.request<{ path: string; files: FileInfo[] }>(`/api/servers/${id}/files?path=${encodeURIComponent(path)}`);
  }

  async getFile(id: string, path: string) {
    return this.request<{ path: string; content: string; size: number }>(`/api/servers/${id}/files/content?path=${encodeURIComponent(path)}`);
  }

  async saveFile(id: string, path: string, content: string) {
    return this.request<{ message: string }>(`/api/servers/${id}/files/content`, {
      method: 'PUT',
      body: JSON.stringify({ path, content }),
    });
  }

  async deleteFile(id: string, path: string) {
    return this.request<{ message: string }>(`/api/servers/${id}/files?path=${encodeURIComponent(path)}`, {
      method: 'DELETE',
    });
  }

  async getMetrics() {
    return this.request<SystemMetrics>('/api/metrics');
  }

  getConsoleWsUrl(id: string): string {
    const wsUrl = API_URL.replace('http', 'ws');
    return `${wsUrl}/api/servers/${id}/console?token=${this.getToken()}`;
  }

  async getPlayers(params: { limit?: number; offset?: number; search?: string; online?: boolean } = {}) {
    const query = new URLSearchParams();
    if (params.limit) query.set('limit', params.limit.toString());
    if (params.offset) query.set('offset', params.offset.toString());
    if (params.search) query.set('search', params.search);
    if (params.online) query.set('online', 'true');
    return this.request<{ players: Player[]; total: number; limit: number; offset: number }>(`/api/players?${query}`);
  }

  async getPlayer(uuid: string) {
    return this.request<PlayerWithStats>(`/api/players/${uuid}`);
  }

  async searchPlayer(username: string) {
    return this.request<Player>(`/api/players/search/${username}`);
  }

  async getPlayerStats() {
    return this.request<{ online: number; total: number; newToday: number }>('/api/players/stats');
  }

  async getPunishments(params: { limit?: number; offset?: number; type?: string; active?: boolean } = {}) {
    const query = new URLSearchParams();
    if (params.limit) query.set('limit', params.limit.toString());
    if (params.offset) query.set('offset', params.offset.toString());
    if (params.type) query.set('type', params.type);
    if (params.active !== undefined) query.set('active', params.active.toString());
    return this.request<{ punishments: Punishment[]; total: number }>(`/api/punishments?${query}`);
  }

  async createPunishment(data: CreatePunishmentRequest) {
    return this.request<Punishment>('/api/punishments', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async revokePunishment(id: number, reason: string) {
    return this.request<{ success: boolean }>(`/api/punishments/${id}/revoke`, {
      method: 'POST',
      body: JSON.stringify({ reason }),
    });
  }

  async getPunishmentTemplates() {
    return this.request<{ templates: PunishmentTemplate[] }>('/api/punishments/templates');
  }

  async getPunishmentStats() {
    return this.request<PunishmentStats>('/api/punishments/stats');
  }

  async getDiscordLinks(params: { limit?: number; offset?: number } = {}) {
    const query = new URLSearchParams();
    if (params.limit) query.set('limit', params.limit.toString());
    if (params.offset) query.set('offset', params.offset.toString());
    return this.request<{ links: DiscordLink[]; total: number }>(`/api/discord/links?${query}`);
  }

  async getDiscordSettings(guildId: string) {
    return this.request<DiscordSettings>(`/api/discord/settings/${guildId}`);
  }

  async updateDiscordSettings(guildId: string, settings: Partial<DiscordSettings>) {
    return this.request<DiscordSettings>(`/api/discord/settings/${guildId}`, {
      method: 'PUT',
      body: JSON.stringify(settings),
    });
  }

  async getDashboardStats() {
    return this.request<DashboardStats>('/api/analytics/dashboard');
  }

  async getPlayerHistory(days: number = 7) {
    return this.request<{ history: PlayerHistoryPoint[] }>(`/api/analytics/players?days=${days}`);
  }

  async getActivityLogs(params: { limit?: number; offset?: number; type?: string } = {}) {
    const query = new URLSearchParams();
    if (params.limit) query.set('limit', params.limit.toString());
    if (params.offset) query.set('offset', params.offset.toString());
    if (params.type) query.set('type', params.type);
    return this.request<{ logs: ActivityLog[]; total: number }>(`/api/analytics/logs?${query}`);
  }

  async getNotifications(unreadOnly: boolean = false) {
    return this.request<{ notifications: AppNotification[] }>(`/api/notifications?unread=${unreadOnly}`);
  }

  async markNotificationRead(id: number) {
    return this.request<{ success: boolean }>(`/api/notifications/${id}/read`, { method: 'POST' });
  }

  async getWebhooks() {
    return this.request<{ webhooks: Webhook[] }>('/api/webhooks');
  }

  async createWebhook(data: { name: string; url: string; events: string[] }) {
    return this.request<Webhook>('/api/webhooks', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async deleteWebhook(id: number) {
    return this.request<{ success: boolean }>(`/api/webhooks/${id}`, { method: 'DELETE' });
  }

  // ==========================================
  // Dedicated Server Management (Pterodactyl-like)
  // ==========================================

  // Locations
  async getLocations() {
    return this.request<{ locations: Location[] }>('/api/locations');
  }

  async createLocation(data: { short: string; long: string }) {
    return this.request<Location>('/api/locations', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async deleteLocation(id: number) {
    return this.request<{ message: string }>(`/api/locations/${id}`, { method: 'DELETE' });
  }

  // Nodes
  async getNodes(includeOffline: boolean = true) {
    return this.request<{ nodes: NodeWithStats[]; total: number }>(`/api/nodes?includeOffline=${includeOffline}`);
  }

  async getNode(id: number) {
    return this.request<NodeWithStats>(`/api/nodes/${id}`);
  }

  async createNode(data: CreateNodeRequest) {
    return this.request<Node>('/api/nodes', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateNode(id: number, data: UpdateNodeRequest) {
    return this.request<Node>(`/api/nodes/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteNode(id: number) {
    return this.request<{ message: string }>(`/api/nodes/${id}`, { method: 'DELETE' });
  }

  async regenerateNodeToken(id: number) {
    return this.request<{ token: string }>(`/api/nodes/${id}/regenerate-token`, { method: 'POST' });
  }

  async getNodeAllocations(nodeId: number) {
    return this.request<{ allocations: Allocation[]; grouped: Record<string, Allocation[]>; total: number }>(`/api/nodes/${nodeId}/allocations`);
  }

  async createAllocation(nodeId: number, data: { ip: string; port?: number; portRange?: string; alias?: string }) {
    return this.request<Allocation | { message: string; count: number }>(`/api/nodes/${nodeId}/allocations`, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async deleteAllocation(nodeId: number, allocId: number) {
    return this.request<{ message: string }>(`/api/nodes/${nodeId}/allocations/${allocId}`, { method: 'DELETE' });
  }

  // Dedicated Servers
  async getDedicatedServers(params: { nodeId?: number; search?: string; page?: number; perPage?: number } = {}) {
    const query = new URLSearchParams();
    if (params.nodeId) query.set('nodeId', params.nodeId.toString());
    if (params.search) query.set('search', params.search);
    if (params.page) query.set('page', params.page.toString());
    if (params.perPage) query.set('perPage', params.perPage.toString());
    return this.request<{ servers: DedicatedServer[]; pagination: Pagination }>(`/api/dedicated-servers?${query}`);
  }

  async getDedicatedServer(id: number) {
    return this.request<DedicatedServer>(`/api/dedicated-servers/${id}`);
  }

  async getDedicatedServerByUUID(uuid: string) {
    return this.request<DedicatedServer>(`/api/dedicated-servers/uuid/${uuid}`);
  }

  async createDedicatedServer(data: CreateDedicatedServerRequest) {
    return this.request<DedicatedServer>('/api/dedicated-servers', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async updateDedicatedServer(id: number, data: UpdateDedicatedServerRequest) {
    return this.request<DedicatedServer>(`/api/dedicated-servers/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteDedicatedServer(id: number) {
    return this.request<{ message: string }>(`/api/dedicated-servers/${id}`, { method: 'DELETE' });
  }

  async getDedicatedServerStats() {
    return this.request<DedicatedServerStats>('/api/dedicated-servers/stats');
  }

  // Power actions
  async startDedicatedServer(id: number) {
    return this.request<{ message: string; status: string }>(`/api/dedicated-servers/${id}/power/start`, { method: 'POST' });
  }

  async stopDedicatedServer(id: number) {
    return this.request<{ message: string; status: string }>(`/api/dedicated-servers/${id}/power/stop`, { method: 'POST' });
  }

  async restartDedicatedServer(id: number) {
    return this.request<{ message: string; status: string }>(`/api/dedicated-servers/${id}/power/restart`, { method: 'POST' });
  }

  async killDedicatedServer(id: number) {
    return this.request<{ message: string; status: string }>(`/api/dedicated-servers/${id}/power/kill`, { method: 'POST' });
  }

  async suspendDedicatedServer(id: number) {
    return this.request<{ message: string }>(`/api/dedicated-servers/${id}/suspend`, { method: 'POST' });
  }

  async unsuspendDedicatedServer(id: number) {
    return this.request<{ message: string }>(`/api/dedicated-servers/${id}/unsuspend`, { method: 'POST' });
  }

  async sendDedicatedServerCommand(id: number, command: string) {
    return this.request<{ message: string }>(`/api/dedicated-servers/${id}/command`, {
      method: 'POST',
      body: JSON.stringify({ command }),
    });
  }

  // Eggs & Nests
  async getNests() {
    return this.request<{ nests: Nest[] }>('/api/nests');
  }

  async getEggs() {
    return this.request<{ eggs: Egg[] }>('/api/eggs');
  }

  async getEgg(id: number) {
    return this.request<Egg>(`/api/eggs/${id}`);
  }
}

export const api = new ApiClient();

export interface ServerInfo {
  id: string;
  name: string;
  type: string;
  status: 'online' | 'offline' | 'starting' | 'stopping' | 'unknown';
  uptime: number;
  players: number;
  maxPlayers: number;
  tps: number;
  memoryUsed: number;
  memoryMax: number;
  cpuPercent: number;
}

export interface FileInfo {
  name: string;
  path: string;
  isDir: boolean;
  size: number;
  modTime: number;
}

export interface SystemMetrics {
  timestamp: number;
  cpu: {
    usagePercent: number;
    coreCount: number;
    loadAvg: [number, number, number];
  };
  memory: {
    total: number;
    used: number;
    free: number;
    available: number;
    percent: number;
  };
  disk: {
    total: number;
    used: number;
    free: number;
    percent: number;
    readBytes: number;
    writeBytes: number;
  };
  network: {
    bytesRecv: number;
    bytesSent: number;
  };
  uptime: number;
}

export interface Player {
  uuid: string;
  username: string;
  firstJoin: string;
  lastSeen: string;
  totalPlaytime: number;
  isOnline: boolean;
  lastIP: string;
}

export interface PlayerStats {
  playerUUID: string;
  kills: number;
  deaths: number;
  blocksBroken: number;
  blocksPlaced: number;
  mobKills: number;
  distanceWalked: number;
  timePlayed: number;
  lastUpdated: string;
}

export interface PlayerIP {
  ip: string;
  firstUsed: string;
  lastUsed: string;
  usageCount: number;
}

export interface PlayerWithStats extends Player {
  stats?: PlayerStats;
  ips?: PlayerIP[];
  punishments?: Punishment[];
}

export interface Punishment {
  id: number;
  playerUUID: string;
  playerUsername?: string;
  type: 'ban' | 'kick' | 'mute' | 'warn';
  reason: string;
  duration: number;
  expiresAt?: string;
  isActive: boolean;
  isPermanent: boolean;
  isAppealed: boolean;
  appealReason?: string;
  appealStatus?: string;
  issuedBy: string;
  issuedAt: string;
  revokedBy?: string;
  revokedAt?: string;
  revokeReason?: string;
}

export interface PunishmentTemplate {
  id: number;
  name: string;
  type: 'ban' | 'kick' | 'mute' | 'warn';
  reason: string;
  duration: number;
  isPermanent: boolean;
}

export interface PunishmentStats {
  totalBans: number;
  totalKicks: number;
  totalMutes: number;
  totalWarns: number;
  activeBans: number;
  activeMutes: number;
  pendingAppeals: number;
}

export interface CreatePunishmentRequest {
  playerUUID: string;
  type: 'ban' | 'kick' | 'mute' | 'warn';
  reason: string;
  duration?: number;
  isPermanent?: boolean;
}

export interface DiscordLink {
  id: number;
  playerUUID: string;
  playerUsername: string;
  discordId: string;
  discordUsername: string;
  isVerified: boolean;
  linkedAt: string;
}

export interface DiscordSettings {
  guildId: string;
  prefix: string;
  welcomeChannelId: string;
  logChannelId: string;
  syncRoles: boolean;
  linkedRoleId: string;
  staffRoleId: string;
  updatedAt: string;
}

export interface DashboardStats {
  totalPlayers: number;
  onlinePlayers: number;
  newPlayersToday: number;
  activeBans: number;
  activeMutes: number;
  linkedDiscordAccounts: number;
}

export interface PlayerHistoryPoint {
  hour: string;
  peakPlayers: number;
  avgPlayers: number;
}

export interface ActivityLog {
  id: number;
  type: string;
  actorName: string;
  targetName?: string;
  action: string;
  details?: string;
  ipAddress?: string;
  createdAt: string;
}

export interface AppNotification {
  id: number;
  type: string;
  title: string;
  message: string;
  isRead: boolean;
  createdAt: string;
}

export interface Webhook {
  id: number;
  name: string;
  url: string;
  events: string[];
  isActive: boolean;
  createdAt: string;
  lastTriggered?: string;
}

// ==========================================
// Dedicated Server Management Types
// ==========================================

export interface Location {
  id: number;
  short: string;
  long: string;
  createdAt: string;
  updatedAt: string;
}

export interface Node {
  id: number;
  uuid: string;
  name: string;
  description?: string;
  fqdn: string;
  scheme: string;
  daemonPort: number;
  daemonToken?: string;
  memory: number;
  memoryOveralloc: number;
  disk: number;
  diskOveralloc: number;
  uploadLimit: number;
  downloadLimit: number;
  status: 'online' | 'offline' | 'maintenance';
  maintenanceMode: boolean;
  locationId?: number;
  createdAt: string;
  updatedAt: string;
}

export interface NodeWithStats extends Node {
  memoryUsed: number;
  memoryPercent: number;
  diskUsed: number;
  diskPercent: number;
  cpuPercent: number;
  uptime: number;
  serverCount: number;
  onlineServers: number;
  allocatedPorts: number;
  locationName?: string;
}

export interface Allocation {
  id: number;
  nodeId: number;
  ip: string;
  alias?: string;
  port: number;
  serverId?: number;
  notes?: string;
  assigned: boolean;
}

export interface Nest {
  id: number;
  uuid: string;
  name: string;
  description?: string;
  author: string;
  createdAt: string;
  updatedAt: string;
}

export interface Egg {
  id: number;
  uuid: string;
  nestId: number;
  name: string;
  description?: string;
  author: string;
  dockerImages: string[];
  defaultImage: string;
  startup: string;
  stopCommand: string;
  configFiles?: object;
  configLogs?: object;
  variables?: EggVariable[];
  createdAt: string;
  updatedAt: string;
}

export interface EggVariable {
  name: string;
  description?: string;
  envVariable: string;
  defaultValue: string;
  userViewable: boolean;
  userEditable: boolean;
  rules: string;
}

export type DedicatedServerStatus = 'offline' | 'starting' | 'running' | 'stopping' | 'restarting' | 'installing' | 'suspended';

export interface DedicatedServer {
  id: number;
  uuid: string;
  shortId: string;
  name: string;
  description?: string;
  nodeId: number;
  ownerId?: number;
  eggId?: number;
  status: DedicatedServerStatus;
  suspended: boolean;
  memory: number;
  disk: number;
  cpu: number;
  io: number;
  swap: number;
  threads?: string;
  oomDisabled: boolean;
  startupCommand?: string;
  defaultAllocationId?: number;
  image?: string;
  backupLimit: number;
  databaseLimit: number;
  allocationLimit: number;
  installed: boolean;
  environment?: Record<string, string>;
  createdAt: string;
  updatedAt: string;
  // Joined fields
  nodeName: string;
  nodeFqdn: string;
  nodeStatus: string;
  ip: string;
  port: number;
  eggName?: string;
  nestName?: string;
  // Real-time stats
  memoryUsed: number;
  memoryPercent: number;
  cpuUsed: number;
  diskUsed: number;
  diskPercent: number;
  networkRx: number;
  networkTx: number;
  uptime: number;
  allocations?: Allocation[];
}

export interface DedicatedServerStats {
  total: number;
  running: number;
  suspended: number;
  totalMemory: number;
  totalDisk: number;
}

export interface Pagination {
  page: number;
  perPage: number;
  total: number;
  totalPages: number;
}

export interface CreateNodeRequest {
  name: string;
  description?: string;
  fqdn: string;
  scheme?: string;
  daemonPort?: number;
  memory: number;
  memoryOveralloc?: number;
  disk: number;
  diskOveralloc?: number;
  uploadLimit?: number;
  downloadLimit?: number;
  locationId?: number;
}

export interface UpdateNodeRequest extends CreateNodeRequest {
  maintenanceMode?: boolean;
}

export interface CreateDedicatedServerRequest {
  name: string;
  description?: string;
  nodeId: number;
  eggId?: number;
  memory: number;
  disk: number;
  cpu?: number;
  io?: number;
  swap?: number;
  threads?: string;
  oomDisabled?: boolean;
  allocationId?: number;
  image?: string;
  startupCommand?: string;
  environment?: Record<string, string>;
  backupLimit?: number;
  databaseLimit?: number;
}

export interface UpdateDedicatedServerRequest {
  name?: string;
  description?: string;
  memory?: number;
  disk?: number;
  cpu?: number;
  io?: number;
  swap?: number;
  threads?: string;
  oomDisabled?: boolean;
  image?: string;
  startupCommand?: string;
  environment?: Record<string, string>;
  backupLimit?: number;
  databaseLimit?: number;
}
