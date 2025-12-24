const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

class ApiClient {
  private token: string | null = null;

  setToken(token: string) {
    this.token = token;
    if (typeof window !== 'undefined') {
      localStorage.setItem('karapanel_token', token);
      // Also set as cookie for middleware auth check
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
      // Also clear cookie
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

  // Auth
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

  // Servers
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

  // Logs
  async getLogs(id: string, lines: number = 100) {
    return this.request<{ logs: string[] }>(`/api/servers/${id}/logs?lines=${lines}`);
  }

  // Files
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

  // Metrics
  async getMetrics() {
    return this.request<SystemMetrics>('/api/metrics');
  }

  // WebSocket URL
  getConsoleWsUrl(id: string): string {
    const wsUrl = API_URL.replace('http', 'ws');
    return `${wsUrl}/api/servers/${id}/console?token=${this.getToken()}`;
  }
}

export const api = new ApiClient();

// Types
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
