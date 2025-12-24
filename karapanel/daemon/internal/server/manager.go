package server

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/withsait/karapixel/karapanel/internal/config"
)

type ServerStatus string

const (
	StatusOnline   ServerStatus = "online"
	StatusOffline  ServerStatus = "offline"
	StatusStarting ServerStatus = "starting"
	StatusStopping ServerStatus = "stopping"
	StatusUnknown  ServerStatus = "unknown"
)

type ServerInfo struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Status      ServerStatus `json:"status"`
	Uptime      int64        `json:"uptime"` // seconds
	Players     int          `json:"players"`
	MaxPlayers  int          `json:"maxPlayers"`
	TPS         float64      `json:"tps"`
	MemoryUsed  int64        `json:"memoryUsed"`  // bytes
	MemoryMax   int64        `json:"memoryMax"`   // bytes
	CPUPercent  float64      `json:"cpuPercent"`
}

type Manager struct {
	servers    map[string]*config.MCServer
	statuses   map[string]*ServerInfo
	systemd    *SystemdClient
	mu         sync.RWMutex
}

func NewManager(servers []config.MCServer) (*Manager, error) {
	m := &Manager{
		servers:  make(map[string]*config.MCServer),
		statuses: make(map[string]*ServerInfo),
	}

	// Initialize systemd client
	systemd, err := NewSystemdClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create systemd client: %w", err)
	}
	m.systemd = systemd

	// Register servers
	for i := range servers {
		srv := &servers[i]
		m.servers[srv.ID] = srv
		m.statuses[srv.ID] = &ServerInfo{
			ID:     srv.ID,
			Name:   srv.Name,
			Type:   srv.Type,
			Status: StatusUnknown,
		}
	}

	// Start status updater
	go m.statusUpdater()

	return m, nil
}

func (m *Manager) statusUpdater() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		m.updateAllStatuses()
	}
}

func (m *Manager) updateAllStatuses() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for id, srv := range m.servers {
		status, err := m.systemd.GetUnitStatus(srv.ServiceName)
		if err != nil {
			m.statuses[id].Status = StatusUnknown
			continue
		}

		switch status {
		case "active":
			m.statuses[id].Status = StatusOnline
			// Get process metrics
			if pid, err := m.systemd.GetUnitPID(srv.ServiceName); err == nil && pid > 0 {
				m.statuses[id].CPUPercent, m.statuses[id].MemoryUsed = getProcessMetrics(pid)
			}
		case "activating":
			m.statuses[id].Status = StatusStarting
		case "deactivating":
			m.statuses[id].Status = StatusStopping
		default:
			m.statuses[id].Status = StatusOffline
		}
	}
}

func (m *Manager) GetAllServers() []*ServerInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*ServerInfo, 0, len(m.statuses))
	for _, info := range m.statuses {
		result = append(result, info)
	}
	return result
}

func (m *Manager) GetServer(id string) (*ServerInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	info, ok := m.statuses[id]
	if !ok {
		return nil, fmt.Errorf("server not found: %s", id)
	}
	return info, nil
}

func (m *Manager) GetServerConfig(id string) (*config.MCServer, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	srv, ok := m.servers[id]
	if !ok {
		return nil, fmt.Errorf("server not found: %s", id)
	}
	return srv, nil
}

func (m *Manager) StartServer(ctx context.Context, id string) error {
	srv, err := m.GetServerConfig(id)
	if err != nil {
		return err
	}

	m.mu.Lock()
	m.statuses[id].Status = StatusStarting
	m.mu.Unlock()

	if err := m.systemd.StartUnit(srv.ServiceName); err != nil {
		m.mu.Lock()
		m.statuses[id].Status = StatusOffline
		m.mu.Unlock()
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func (m *Manager) StopServer(ctx context.Context, id string) error {
	srv, err := m.GetServerConfig(id)
	if err != nil {
		return err
	}

	m.mu.Lock()
	m.statuses[id].Status = StatusStopping
	m.mu.Unlock()

	if err := m.systemd.StopUnit(srv.ServiceName); err != nil {
		return fmt.Errorf("failed to stop server: %w", err)
	}

	m.mu.Lock()
	m.statuses[id].Status = StatusOffline
	m.mu.Unlock()

	return nil
}

func (m *Manager) RestartServer(ctx context.Context, id string) error {
	srv, err := m.GetServerConfig(id)
	if err != nil {
		return err
	}

	m.mu.Lock()
	m.statuses[id].Status = StatusStopping
	m.mu.Unlock()

	if err := m.systemd.RestartUnit(srv.ServiceName); err != nil {
		return fmt.Errorf("failed to restart server: %w", err)
	}

	m.mu.Lock()
	m.statuses[id].Status = StatusStarting
	m.mu.Unlock()

	return nil
}

func (m *Manager) KillServer(ctx context.Context, id string) error {
	srv, err := m.GetServerConfig(id)
	if err != nil {
		return err
	}

	if err := m.systemd.KillUnit(srv.ServiceName); err != nil {
		return fmt.Errorf("failed to kill server: %w", err)
	}

	m.mu.Lock()
	m.statuses[id].Status = StatusOffline
	m.mu.Unlock()

	return nil
}

func (m *Manager) Close() error {
	if m.systemd != nil {
		return m.systemd.Close()
	}
	return nil
}
