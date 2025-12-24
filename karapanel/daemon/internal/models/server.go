package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// ServerStatus represents the current status of a server
type ServerStatus string

const (
	ServerStatusOffline    ServerStatus = "offline"
	ServerStatusStarting   ServerStatus = "starting"
	ServerStatusRunning    ServerStatus = "running"
	ServerStatusStopping   ServerStatus = "stopping"
	ServerStatusRestarting ServerStatus = "restarting"
	ServerStatusInstalling ServerStatus = "installing"
	ServerStatusSuspended  ServerStatus = "suspended"
)

// Server represents a game server instance
type Server struct {
	ID                  int64           `json:"id"`
	UUID                string          `json:"uuid"`
	ShortID             string          `json:"shortId"` // 8 char short identifier
	Name                string          `json:"name"`
	Description         string          `json:"description,omitempty"`
	NodeID              int64           `json:"nodeId"`
	OwnerID             int64           `json:"ownerId,omitempty"` // Panel user who owns this
	EggID               int64           `json:"eggId,omitempty"`   // Game type/template
	Status              ServerStatus    `json:"status"`
	Suspended           bool            `json:"suspended"`
	Memory              int64           `json:"memory"`     // Memory limit in MB
	Disk                int64           `json:"disk"`       // Disk limit in MB
	CPU                 int             `json:"cpu"`        // CPU limit (100 = 1 core)
	IO                  int             `json:"io"`         // Block IO weight (10-1000)
	Swap                int64           `json:"swap"`       // Swap limit in MB
	Threads             string          `json:"threads"`    // CPU thread pinning
	OOMDisabled         bool            `json:"oomDisabled"`
	StartupCommand      string          `json:"startupCommand"`
	DefaultAllocationID int64           `json:"defaultAllocationId"`
	Image               string          `json:"image"`     // Docker image
	BackupLimit         int             `json:"backupLimit"`
	DatabaseLimit       int             `json:"databaseLimit"`
	AllocationLimit     int             `json:"allocationLimit"`
	Installed           bool            `json:"installed"`
	Environment         json.RawMessage `json:"environment,omitempty"` // Startup variables
	Metadata            json.RawMessage `json:"metadata,omitempty"`
	CreatedAt           time.Time       `json:"createdAt"`
	UpdatedAt           time.Time       `json:"updatedAt"`
}

// ServerWithDetails includes related data
type ServerWithDetails struct {
	Server
	// Node info
	NodeName   string `json:"nodeName"`
	NodeFQDN   string `json:"nodeFqdn"`
	NodeStatus string `json:"nodeStatus"`
	// Allocation info
	IP   string `json:"ip"`
	Port int    `json:"port"`
	// Egg info
	EggName string `json:"eggName,omitempty"`
	NestName string `json:"nestName,omitempty"`
	// Real-time stats
	MemoryUsed    int64   `json:"memoryUsed"`
	MemoryPercent float64 `json:"memoryPercent"`
	CPUUsed       float64 `json:"cpuUsed"`
	DiskUsed      int64   `json:"diskUsed"`
	DiskPercent   float64 `json:"diskPercent"`
	NetworkRx     int64   `json:"networkRx"`
	NetworkTx     int64   `json:"networkTx"`
	Uptime        int64   `json:"uptime"`
	// Additional allocations
	Allocations []Allocation `json:"allocations,omitempty"`
}

// Egg represents a server template (like Minecraft, Rust, etc.)
type Egg struct {
	ID           int64           `json:"id"`
	UUID         string          `json:"uuid"`
	NestID       int64           `json:"nestId"`
	Name         string          `json:"name"`
	Description  string          `json:"description,omitempty"`
	Author       string          `json:"author"`
	DockerImages json.RawMessage `json:"dockerImages"` // Array of available images
	DefaultImage string          `json:"defaultImage"`
	Startup      string          `json:"startup"` // Startup command template
	StopCommand  string          `json:"stopCommand"`
	ConfigFiles  json.RawMessage `json:"configFiles,omitempty"`  // File parser config
	ConfigLogs   json.RawMessage `json:"configLogs,omitempty"`   // Log configuration
	Variables    json.RawMessage `json:"variables,omitempty"`    // Egg variables
	InstallScript json.RawMessage `json:"installScript,omitempty"`
	Metadata     json.RawMessage `json:"metadata,omitempty"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
}

// Nest represents a category of eggs (like Minecraft, Source Games, etc.)
type Nest struct {
	ID          int64     `json:"id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ServerVariable represents an environment variable for a server
type ServerVariable struct {
	ID           int64  `json:"id"`
	ServerID     int64  `json:"serverId"`
	EggVariableID int64 `json:"eggVariableId,omitempty"`
	Name         string `json:"name"`
	Value        string `json:"value"`
}

// ServerRepository handles database operations for servers
type ServerRepository struct {
	db *sql.DB
}

func NewServerRepository(db *sql.DB) *ServerRepository {
	return &ServerRepository{db: db}
}

// GetAll returns all servers with filters
func (r *ServerRepository) GetAll(nodeID int64, search string, limit, offset int) ([]ServerWithDetails, int, error) {
	var total int
	countQuery := "SELECT COUNT(*) FROM servers WHERE 1=1"
	args := []interface{}{}
	argNum := 0

	if nodeID > 0 {
		argNum++
		countQuery += " AND node_id = $" + string(rune('0'+argNum))
		args = append(args, nodeID)
	}
	if search != "" {
		argNum++
		countQuery += " AND (name ILIKE $" + string(rune('0'+argNum)) + " OR uuid ILIKE $" + string(rune('0'+argNum)) + ")"
		args = append(args, "%"+search+"%")
	}

	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT s.id, s.uuid, s.short_id, s.name, s.description, s.node_id, s.owner_id, s.egg_id,
			   s.status, s.suspended, s.memory, s.disk, s.cpu, s.io, s.swap, s.threads,
			   s.oom_disabled, s.startup_command, s.default_allocation_id, s.image,
			   s.backup_limit, s.database_limit, s.allocation_limit, s.installed,
			   s.environment, s.metadata, s.created_at, s.updated_at,
			   n.name as node_name, n.fqdn as node_fqdn, n.status as node_status,
			   COALESCE(a.ip, '') as ip, COALESCE(a.port, 0) as port,
			   COALESCE(e.name, '') as egg_name,
			   COALESCE(ns.name, '') as nest_name
		FROM servers s
		LEFT JOIN nodes n ON s.node_id = n.id
		LEFT JOIN allocations a ON s.default_allocation_id = a.id
		LEFT JOIN eggs e ON s.egg_id = e.id
		LEFT JOIN nests ns ON e.nest_id = ns.id
		WHERE 1=1`

	if nodeID > 0 {
		query += " AND s.node_id = $1"
	}
	if search != "" {
		if nodeID > 0 {
			query += " AND (s.name ILIKE $2 OR s.uuid ILIKE $2)"
		} else {
			query += " AND (s.name ILIKE $1 OR s.uuid ILIKE $1)"
		}
	}

	query += " ORDER BY s.name ASC"

	limitArg := argNum + 1
	offsetArg := argNum + 2
	query += " LIMIT $" + string(rune('0'+limitArg)) + " OFFSET $" + string(rune('0'+offsetArg))
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var servers []ServerWithDetails
	for rows.Next() {
		var s ServerWithDetails
		var desc, threads, env, metadata sql.NullString
		var ownerID, eggID, defaultAllocID sql.NullInt64

		err := rows.Scan(
			&s.ID, &s.UUID, &s.ShortID, &s.Name, &desc, &s.NodeID, &ownerID, &eggID,
			&s.Status, &s.Suspended, &s.Memory, &s.Disk, &s.CPU, &s.IO, &s.Swap, &threads,
			&s.OOMDisabled, &s.StartupCommand, &defaultAllocID, &s.Image,
			&s.BackupLimit, &s.DatabaseLimit, &s.AllocationLimit, &s.Installed,
			&env, &metadata, &s.CreatedAt, &s.UpdatedAt,
			&s.NodeName, &s.NodeFQDN, &s.NodeStatus,
			&s.IP, &s.Port, &s.EggName, &s.NestName,
		)
		if err != nil {
			return nil, 0, err
		}

		if desc.Valid {
			s.Description = desc.String
		}
		if threads.Valid {
			s.Threads = threads.String
		}
		if env.Valid {
			s.Environment = json.RawMessage(env.String)
		}
		if metadata.Valid {
			s.Metadata = json.RawMessage(metadata.String)
		}
		if ownerID.Valid {
			s.OwnerID = ownerID.Int64
		}
		if eggID.Valid {
			s.EggID = eggID.Int64
		}
		if defaultAllocID.Valid {
			s.DefaultAllocationID = defaultAllocID.Int64
		}

		servers = append(servers, s)
	}

	return servers, total, nil
}

// GetByID returns a server by ID with full details
func (r *ServerRepository) GetByID(id int64) (*ServerWithDetails, error) {
	var s ServerWithDetails
	var desc, threads, env, metadata sql.NullString
	var ownerID, eggID, defaultAllocID sql.NullInt64

	err := r.db.QueryRow(`
		SELECT s.id, s.uuid, s.short_id, s.name, s.description, s.node_id, s.owner_id, s.egg_id,
			   s.status, s.suspended, s.memory, s.disk, s.cpu, s.io, s.swap, s.threads,
			   s.oom_disabled, s.startup_command, s.default_allocation_id, s.image,
			   s.backup_limit, s.database_limit, s.allocation_limit, s.installed,
			   s.environment, s.metadata, s.created_at, s.updated_at,
			   n.name as node_name, n.fqdn as node_fqdn, n.status as node_status,
			   COALESCE(a.ip, '') as ip, COALESCE(a.port, 0) as port,
			   COALESCE(e.name, '') as egg_name,
			   COALESCE(ns.name, '') as nest_name
		FROM servers s
		LEFT JOIN nodes n ON s.node_id = n.id
		LEFT JOIN allocations a ON s.default_allocation_id = a.id
		LEFT JOIN eggs e ON s.egg_id = e.id
		LEFT JOIN nests ns ON e.nest_id = ns.id
		WHERE s.id = $1
	`, id).Scan(
		&s.ID, &s.UUID, &s.ShortID, &s.Name, &desc, &s.NodeID, &ownerID, &eggID,
		&s.Status, &s.Suspended, &s.Memory, &s.Disk, &s.CPU, &s.IO, &s.Swap, &threads,
		&s.OOMDisabled, &s.StartupCommand, &defaultAllocID, &s.Image,
		&s.BackupLimit, &s.DatabaseLimit, &s.AllocationLimit, &s.Installed,
		&env, &metadata, &s.CreatedAt, &s.UpdatedAt,
		&s.NodeName, &s.NodeFQDN, &s.NodeStatus,
		&s.IP, &s.Port, &s.EggName, &s.NestName,
	)

	if err != nil {
		return nil, err
	}

	if desc.Valid {
		s.Description = desc.String
	}
	if threads.Valid {
		s.Threads = threads.String
	}
	if env.Valid {
		s.Environment = json.RawMessage(env.String)
	}
	if metadata.Valid {
		s.Metadata = json.RawMessage(metadata.String)
	}
	if ownerID.Valid {
		s.OwnerID = ownerID.Int64
	}
	if eggID.Valid {
		s.EggID = eggID.Int64
	}
	if defaultAllocID.Valid {
		s.DefaultAllocationID = defaultAllocID.Int64
	}

	// Get all allocations for this server
	allocRows, err := r.db.Query(`
		SELECT id, node_id, ip, alias, port, notes
		FROM allocations
		WHERE server_id = $1
		ORDER BY port
	`, id)
	if err == nil {
		defer allocRows.Close()
		for allocRows.Next() {
			var a Allocation
			var alias, notes sql.NullString
			allocRows.Scan(&a.ID, &a.NodeID, &a.IP, &alias, &a.Port, &notes)
			if alias.Valid {
				a.Alias = alias.String
			}
			if notes.Valid {
				a.Notes = notes.String
			}
			a.ServerID = id
			a.Assigned = true
			s.Allocations = append(s.Allocations, a)
		}
	}

	return &s, nil
}

// GetByUUID returns a server by UUID
func (r *ServerRepository) GetByUUID(uuid string) (*ServerWithDetails, error) {
	var id int64
	err := r.db.QueryRow("SELECT id FROM servers WHERE uuid = $1", uuid).Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

// Create creates a new server
func (r *ServerRepository) Create(s *Server) error {
	return r.db.QueryRow(`
		INSERT INTO servers (uuid, short_id, name, description, node_id, owner_id, egg_id,
							 status, suspended, memory, disk, cpu, io, swap, threads,
							 oom_disabled, startup_command, default_allocation_id, image,
							 backup_limit, database_limit, allocation_limit, installed, environment, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
		RETURNING id, created_at, updated_at
	`, s.UUID, s.ShortID, s.Name, nullString(s.Description), s.NodeID, nullInt64(s.OwnerID), nullInt64(s.EggID),
		s.Status, s.Suspended, s.Memory, s.Disk, s.CPU, s.IO, s.Swap, nullString(s.Threads),
		s.OOMDisabled, s.StartupCommand, nullInt64(s.DefaultAllocationID), s.Image,
		s.BackupLimit, s.DatabaseLimit, s.AllocationLimit, s.Installed, s.Environment, s.Metadata,
	).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}

// Update updates a server
func (r *ServerRepository) Update(s *Server) error {
	_, err := r.db.Exec(`
		UPDATE servers SET
			name = $2, description = $3, status = $4, suspended = $5,
			memory = $6, disk = $7, cpu = $8, io = $9, swap = $10, threads = $11,
			oom_disabled = $12, startup_command = $13, default_allocation_id = $14, image = $15,
			backup_limit = $16, database_limit = $17, allocation_limit = $18,
			environment = $19, metadata = $20, updated_at = NOW()
		WHERE id = $1
	`, s.ID, s.Name, nullString(s.Description), s.Status, s.Suspended,
		s.Memory, s.Disk, s.CPU, s.IO, s.Swap, nullString(s.Threads),
		s.OOMDisabled, s.StartupCommand, nullInt64(s.DefaultAllocationID), s.Image,
		s.BackupLimit, s.DatabaseLimit, s.AllocationLimit, s.Environment, s.Metadata)
	return err
}

// Delete deletes a server
func (r *ServerRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM servers WHERE id = $1", id)
	return err
}

// UpdateStatus updates server status
func (r *ServerRepository) UpdateStatus(id int64, status ServerStatus) error {
	_, err := r.db.Exec("UPDATE servers SET status = $2, updated_at = NOW() WHERE id = $1", id, status)
	return err
}

// Suspend suspends a server
func (r *ServerRepository) Suspend(id int64) error {
	_, err := r.db.Exec("UPDATE servers SET suspended = true, updated_at = NOW() WHERE id = $1", id)
	return err
}

// Unsuspend unsuspends a server
func (r *ServerRepository) Unsuspend(id int64) error {
	_, err := r.db.Exec("UPDATE servers SET suspended = false, updated_at = NOW() WHERE id = $1", id)
	return err
}

// MarkInstalled marks server as installed
func (r *ServerRepository) MarkInstalled(id int64) error {
	_, err := r.db.Exec("UPDATE servers SET installed = true, updated_at = NOW() WHERE id = $1", id)
	return err
}

// GetServersByNode returns servers on a specific node
func (r *ServerRepository) GetServersByNode(nodeID int64) ([]Server, error) {
	rows, err := r.db.Query(`
		SELECT id, uuid, short_id, name, status, memory, disk, cpu
		FROM servers
		WHERE node_id = $1
		ORDER BY name
	`, nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servers []Server
	for rows.Next() {
		var s Server
		err := rows.Scan(&s.ID, &s.UUID, &s.ShortID, &s.Name, &s.Status, &s.Memory, &s.Disk, &s.CPU)
		if err != nil {
			return nil, err
		}
		servers = append(servers, s)
	}

	return servers, nil
}

// Egg methods

func (r *ServerRepository) GetAllEggs() ([]Egg, error) {
	rows, err := r.db.Query(`
		SELECT id, uuid, nest_id, name, description, author, docker_images, default_image,
			   startup, stop_command, config_files, config_logs, variables, install_script,
			   metadata, created_at, updated_at
		FROM eggs
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eggs []Egg
	for rows.Next() {
		var e Egg
		var desc, dockerImages, configFiles, configLogs, variables, installScript, metadata sql.NullString

		err := rows.Scan(
			&e.ID, &e.UUID, &e.NestID, &e.Name, &desc, &e.Author, &dockerImages, &e.DefaultImage,
			&e.Startup, &e.StopCommand, &configFiles, &configLogs, &variables, &installScript,
			&metadata, &e.CreatedAt, &e.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if desc.Valid {
			e.Description = desc.String
		}
		if dockerImages.Valid {
			e.DockerImages = json.RawMessage(dockerImages.String)
		}
		if configFiles.Valid {
			e.ConfigFiles = json.RawMessage(configFiles.String)
		}
		if configLogs.Valid {
			e.ConfigLogs = json.RawMessage(configLogs.String)
		}
		if variables.Valid {
			e.Variables = json.RawMessage(variables.String)
		}
		if installScript.Valid {
			e.InstallScript = json.RawMessage(installScript.String)
		}
		if metadata.Valid {
			e.Metadata = json.RawMessage(metadata.String)
		}

		eggs = append(eggs, e)
	}

	return eggs, nil
}

func (r *ServerRepository) GetEggByID(id int64) (*Egg, error) {
	var e Egg
	var desc, dockerImages, configFiles, configLogs, variables, installScript, metadata sql.NullString

	err := r.db.QueryRow(`
		SELECT id, uuid, nest_id, name, description, author, docker_images, default_image,
			   startup, stop_command, config_files, config_logs, variables, install_script,
			   metadata, created_at, updated_at
		FROM eggs
		WHERE id = $1
	`, id).Scan(
		&e.ID, &e.UUID, &e.NestID, &e.Name, &desc, &e.Author, &dockerImages, &e.DefaultImage,
		&e.Startup, &e.StopCommand, &configFiles, &configLogs, &variables, &installScript,
		&metadata, &e.CreatedAt, &e.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if desc.Valid {
		e.Description = desc.String
	}
	if dockerImages.Valid {
		e.DockerImages = json.RawMessage(dockerImages.String)
	}
	if configFiles.Valid {
		e.ConfigFiles = json.RawMessage(configFiles.String)
	}
	if configLogs.Valid {
		e.ConfigLogs = json.RawMessage(configLogs.String)
	}
	if variables.Valid {
		e.Variables = json.RawMessage(variables.String)
	}
	if installScript.Valid {
		e.InstallScript = json.RawMessage(installScript.String)
	}
	if metadata.Valid {
		e.Metadata = json.RawMessage(metadata.String)
	}

	return &e, nil
}

// Nest methods

func (r *ServerRepository) GetAllNests() ([]Nest, error) {
	rows, err := r.db.Query(`
		SELECT id, uuid, name, description, author, created_at, updated_at
		FROM nests
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nests []Nest
	for rows.Next() {
		var n Nest
		var desc sql.NullString

		err := rows.Scan(&n.ID, &n.UUID, &n.Name, &desc, &n.Author, &n.CreatedAt, &n.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if desc.Valid {
			n.Description = desc.String
		}

		nests = append(nests, n)
	}

	return nests, nil
}

// GetServerStats returns aggregated server statistics
func (r *ServerRepository) GetServerStats() (map[string]interface{}, error) {
	var total, running, suspended int
	var totalMemory, totalDisk int64

	err := r.db.QueryRow(`
		SELECT
			COUNT(*) as total,
			COUNT(*) FILTER (WHERE status = 'running') as running,
			COUNT(*) FILTER (WHERE suspended = true) as suspended,
			COALESCE(SUM(memory), 0) as total_memory,
			COALESCE(SUM(disk), 0) as total_disk
		FROM servers
	`).Scan(&total, &running, &suspended, &totalMemory, &totalDisk)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total":       total,
		"running":     running,
		"suspended":   suspended,
		"totalMemory": totalMemory,
		"totalDisk":   totalDisk,
	}, nil
}
