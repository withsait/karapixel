package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// NodeStatus represents the current status of a node
type NodeStatus string

const (
	NodeStatusOnline      NodeStatus = "online"
	NodeStatusOffline     NodeStatus = "offline"
	NodeStatusMaintenance NodeStatus = "maintenance"
)

// Node represents a dedicated server/host machine
type Node struct {
	ID              int64           `json:"id"`
	UUID            string          `json:"uuid"`
	Name            string          `json:"name"`
	Description     string          `json:"description,omitempty"`
	FQDN            string          `json:"fqdn"`      // Fully Qualified Domain Name or IP
	Scheme          string          `json:"scheme"`    // http or https
	DaemonPort      int             `json:"daemonPort"`
	DaemonToken     string          `json:"daemonToken,omitempty"`
	Memory          int64           `json:"memory"`          // Total memory in MB
	MemoryOveralloc int             `json:"memoryOveralloc"` // Memory over-allocation percentage
	Disk            int64           `json:"disk"`            // Total disk in MB
	DiskOveralloc   int             `json:"diskOveralloc"`   // Disk over-allocation percentage
	UploadLimit     int             `json:"uploadLimit"`     // Upload limit in MB/s
	DownloadLimit   int             `json:"downloadLimit"`   // Download limit in MB/s
	Status          NodeStatus      `json:"status"`
	MaintenanceMode bool            `json:"maintenanceMode"`
	LocationID      int64           `json:"locationId,omitempty"`
	Metadata        json.RawMessage `json:"metadata,omitempty"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}

// NodeWithStats includes real-time statistics
type NodeWithStats struct {
	Node
	// Real-time stats
	MemoryUsed     int64   `json:"memoryUsed"`
	MemoryPercent  float64 `json:"memoryPercent"`
	DiskUsed       int64   `json:"diskUsed"`
	DiskPercent    float64 `json:"diskPercent"`
	CPUPercent     float64 `json:"cpuPercent"`
	Uptime         int64   `json:"uptime"`
	ServerCount    int     `json:"serverCount"`
	OnlineServers  int     `json:"onlineServers"`
	AllocatedPorts int     `json:"allocatedPorts"`
	// Location info
	LocationName string `json:"locationName,omitempty"`
}

// Location represents a physical location (datacenter)
type Location struct {
	ID        int64     `json:"id"`
	Short     string    `json:"short"` // Short code like "us-east-1"
	Long      string    `json:"long"`  // Full name like "US East (Virginia)"
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Allocation represents an IP:Port allocation on a node
type Allocation struct {
	ID       int64  `json:"id"`
	NodeID   int64  `json:"nodeId"`
	IP       string `json:"ip"`
	Alias    string `json:"alias,omitempty"` // Optional display alias
	Port     int    `json:"port"`
	ServerID int64  `json:"serverId,omitempty"` // null if unassigned
	Notes    string `json:"notes,omitempty"`
	Assigned bool   `json:"assigned"`
}

// NodeRepository handles database operations for nodes
type NodeRepository struct {
	db *sql.DB
}

func NewNodeRepository(db *sql.DB) *NodeRepository {
	return &NodeRepository{db: db}
}

// GetAll returns all nodes with optional filters
func (r *NodeRepository) GetAll(includeOffline bool) ([]NodeWithStats, error) {
	query := `
		SELECT n.id, n.uuid, n.name, n.description, n.fqdn, n.scheme, n.daemon_port,
			   n.memory, n.memory_overalloc, n.disk, n.disk_overalloc,
			   n.upload_limit, n.download_limit, n.status, n.maintenance_mode,
			   n.location_id, n.metadata, n.created_at, n.updated_at,
			   COALESCE(l.long, '') as location_name,
			   (SELECT COUNT(*) FROM servers s WHERE s.node_id = n.id) as server_count,
			   (SELECT COUNT(*) FROM servers s WHERE s.node_id = n.id AND s.status = 'running') as online_servers,
			   (SELECT COUNT(*) FROM allocations a WHERE a.node_id = n.id AND a.server_id IS NOT NULL) as allocated_ports
		FROM nodes n
		LEFT JOIN locations l ON n.location_id = l.id
		WHERE 1=1`

	if !includeOffline {
		query += " AND n.status != 'offline'"
	}
	query += " ORDER BY n.name ASC"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []NodeWithStats
	for rows.Next() {
		var n NodeWithStats
		var desc, metadata sql.NullString
		var locationID sql.NullInt64

		err := rows.Scan(
			&n.ID, &n.UUID, &n.Name, &desc, &n.FQDN, &n.Scheme, &n.DaemonPort,
			&n.Memory, &n.MemoryOveralloc, &n.Disk, &n.DiskOveralloc,
			&n.UploadLimit, &n.DownloadLimit, &n.Status, &n.MaintenanceMode,
			&locationID, &metadata, &n.CreatedAt, &n.UpdatedAt,
			&n.LocationName, &n.ServerCount, &n.OnlineServers, &n.AllocatedPorts,
		)
		if err != nil {
			return nil, err
		}

		if desc.Valid {
			n.Description = desc.String
		}
		if metadata.Valid {
			n.Metadata = json.RawMessage(metadata.String)
		}
		if locationID.Valid {
			n.LocationID = locationID.Int64
		}

		nodes = append(nodes, n)
	}

	return nodes, nil
}

// GetByID returns a single node by ID
func (r *NodeRepository) GetByID(id int64) (*NodeWithStats, error) {
	var n NodeWithStats
	var desc, metadata, daemonToken sql.NullString
	var locationID sql.NullInt64

	err := r.db.QueryRow(`
		SELECT n.id, n.uuid, n.name, n.description, n.fqdn, n.scheme, n.daemon_port, n.daemon_token,
			   n.memory, n.memory_overalloc, n.disk, n.disk_overalloc,
			   n.upload_limit, n.download_limit, n.status, n.maintenance_mode,
			   n.location_id, n.metadata, n.created_at, n.updated_at,
			   COALESCE(l.long, '') as location_name,
			   (SELECT COUNT(*) FROM servers s WHERE s.node_id = n.id) as server_count,
			   (SELECT COUNT(*) FROM servers s WHERE s.node_id = n.id AND s.status = 'running') as online_servers,
			   (SELECT COUNT(*) FROM allocations a WHERE a.node_id = n.id AND a.server_id IS NOT NULL) as allocated_ports
		FROM nodes n
		LEFT JOIN locations l ON n.location_id = l.id
		WHERE n.id = $1
	`, id).Scan(
		&n.ID, &n.UUID, &n.Name, &desc, &n.FQDN, &n.Scheme, &n.DaemonPort, &daemonToken,
		&n.Memory, &n.MemoryOveralloc, &n.Disk, &n.DiskOveralloc,
		&n.UploadLimit, &n.DownloadLimit, &n.Status, &n.MaintenanceMode,
		&locationID, &metadata, &n.CreatedAt, &n.UpdatedAt,
		&n.LocationName, &n.ServerCount, &n.OnlineServers, &n.AllocatedPorts,
	)

	if err != nil {
		return nil, err
	}

	if desc.Valid {
		n.Description = desc.String
	}
	if metadata.Valid {
		n.Metadata = json.RawMessage(metadata.String)
	}
	if daemonToken.Valid {
		n.DaemonToken = daemonToken.String
	}
	if locationID.Valid {
		n.LocationID = locationID.Int64
	}

	return &n, nil
}

// Create creates a new node
func (r *NodeRepository) Create(n *Node) error {
	return r.db.QueryRow(`
		INSERT INTO nodes (uuid, name, description, fqdn, scheme, daemon_port, daemon_token,
						   memory, memory_overalloc, disk, disk_overalloc,
						   upload_limit, download_limit, status, maintenance_mode, location_id, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		RETURNING id, created_at, updated_at
	`, n.UUID, n.Name, n.Description, n.FQDN, n.Scheme, n.DaemonPort, n.DaemonToken,
		n.Memory, n.MemoryOveralloc, n.Disk, n.DiskOveralloc,
		n.UploadLimit, n.DownloadLimit, n.Status, n.MaintenanceMode, nullInt64(n.LocationID), n.Metadata,
	).Scan(&n.ID, &n.CreatedAt, &n.UpdatedAt)
}

// Update updates an existing node
func (r *NodeRepository) Update(n *Node) error {
	_, err := r.db.Exec(`
		UPDATE nodes SET
			name = $2, description = $3, fqdn = $4, scheme = $5, daemon_port = $6, daemon_token = $7,
			memory = $8, memory_overalloc = $9, disk = $10, disk_overalloc = $11,
			upload_limit = $12, download_limit = $13, status = $14, maintenance_mode = $15,
			location_id = $16, metadata = $17, updated_at = NOW()
		WHERE id = $1
	`, n.ID, n.Name, n.Description, n.FQDN, n.Scheme, n.DaemonPort, n.DaemonToken,
		n.Memory, n.MemoryOveralloc, n.Disk, n.DiskOveralloc,
		n.UploadLimit, n.DownloadLimit, n.Status, n.MaintenanceMode, nullInt64(n.LocationID), n.Metadata)
	return err
}

// Delete deletes a node by ID
func (r *NodeRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM nodes WHERE id = $1", id)
	return err
}

// UpdateStatus updates node status
func (r *NodeRepository) UpdateStatus(id int64, status NodeStatus) error {
	_, err := r.db.Exec("UPDATE nodes SET status = $2, updated_at = NOW() WHERE id = $1", id, status)
	return err
}

// GetAllocations returns all allocations for a node
func (r *NodeRepository) GetAllocations(nodeID int64) ([]Allocation, error) {
	rows, err := r.db.Query(`
		SELECT id, node_id, ip, alias, port, server_id, notes,
			   (server_id IS NOT NULL) as assigned
		FROM allocations
		WHERE node_id = $1
		ORDER BY ip, port
	`, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allocations []Allocation
	for rows.Next() {
		var a Allocation
		var alias, notes sql.NullString
		var serverID sql.NullInt64

		err := rows.Scan(&a.ID, &a.NodeID, &a.IP, &alias, &a.Port, &serverID, &notes, &a.Assigned)
		if err != nil {
			return nil, err
		}

		if alias.Valid {
			a.Alias = alias.String
		}
		if notes.Valid {
			a.Notes = notes.String
		}
		if serverID.Valid {
			a.ServerID = serverID.Int64
		}

		allocations = append(allocations, a)
	}

	return allocations, nil
}

// CreateAllocation creates a new port allocation
func (r *NodeRepository) CreateAllocation(a *Allocation) error {
	return r.db.QueryRow(`
		INSERT INTO allocations (node_id, ip, alias, port, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, a.NodeID, a.IP, nullString(a.Alias), a.Port, nullString(a.Notes)).Scan(&a.ID)
}

// CreateAllocationRange creates multiple allocations for a port range
func (r *NodeRepository) CreateAllocationRange(nodeID int64, ip string, startPort, endPort int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO allocations (node_id, ip, port)
		VALUES ($1, $2, $3)
		ON CONFLICT (node_id, ip, port) DO NOTHING
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for port := startPort; port <= endPort; port++ {
		if _, err := stmt.Exec(nodeID, ip, port); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// DeleteAllocation deletes an allocation
func (r *NodeRepository) DeleteAllocation(id int64) error {
	_, err := r.db.Exec("DELETE FROM allocations WHERE id = $1 AND server_id IS NULL", id)
	return err
}

// AssignAllocation assigns an allocation to a server
func (r *NodeRepository) AssignAllocation(allocationID, serverID int64) error {
	_, err := r.db.Exec("UPDATE allocations SET server_id = $2 WHERE id = $1", allocationID, serverID)
	return err
}

// UnassignAllocation removes server from allocation
func (r *NodeRepository) UnassignAllocation(allocationID int64) error {
	_, err := r.db.Exec("UPDATE allocations SET server_id = NULL WHERE id = $1", allocationID)
	return err
}

// GetAvailableAllocations returns unassigned allocations for a node
func (r *NodeRepository) GetAvailableAllocations(nodeID int64) ([]Allocation, error) {
	rows, err := r.db.Query(`
		SELECT id, node_id, ip, alias, port, notes
		FROM allocations
		WHERE node_id = $1 AND server_id IS NULL
		ORDER BY ip, port
	`, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allocations []Allocation
	for rows.Next() {
		var a Allocation
		var alias, notes sql.NullString

		err := rows.Scan(&a.ID, &a.NodeID, &a.IP, &alias, &a.Port, &notes)
		if err != nil {
			return nil, err
		}

		if alias.Valid {
			a.Alias = alias.String
		}
		if notes.Valid {
			a.Notes = notes.String
		}

		allocations = append(allocations, a)
	}

	return allocations, nil
}

// Location methods

func (r *NodeRepository) GetAllLocations() ([]Location, error) {
	rows, err := r.db.Query(`
		SELECT id, short, long, created_at, updated_at
		FROM locations
		ORDER BY short ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var l Location
		err := rows.Scan(&l.ID, &l.Short, &l.Long, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			return nil, err
		}
		locations = append(locations, l)
	}

	return locations, nil
}

func (r *NodeRepository) CreateLocation(l *Location) error {
	return r.db.QueryRow(`
		INSERT INTO locations (short, long)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`, l.Short, l.Long).Scan(&l.ID, &l.CreatedAt, &l.UpdatedAt)
}

func (r *NodeRepository) DeleteLocation(id int64) error {
	_, err := r.db.Exec("DELETE FROM locations WHERE id = $1", id)
	return err
}

// Helper functions
func nullInt64(v int64) sql.NullInt64 {
	if v == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: v, Valid: true}
}

func nullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: v, Valid: true}
}
