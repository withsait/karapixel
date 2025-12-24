package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/withsait/karapixel/karapanel/internal/models"
)

type DedicatedServerHandler struct {
	serverRepo *models.ServerRepository
	nodeRepo   *models.NodeRepository
}

func NewDedicatedServerHandler(serverRepo *models.ServerRepository, nodeRepo *models.NodeRepository) *DedicatedServerHandler {
	return &DedicatedServerHandler{
		serverRepo: serverRepo,
		nodeRepo:   nodeRepo,
	}
}

// ListServers returns all dedicated servers
func (h *DedicatedServerHandler) ListServers(c *fiber.Ctx) error {
	nodeID, _ := strconv.ParseInt(c.Query("nodeId", "0"), 10, 64)
	search := c.Query("search", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("perPage", "25"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 25
	}
	offset := (page - 1) * perPage

	servers, total, err := h.serverRepo.GetAll(nodeID, search, perPage, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch servers: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"servers": servers,
		"pagination": fiber.Map{
			"page":       page,
			"perPage":    perPage,
			"total":      total,
			"totalPages": (total + perPage - 1) / perPage,
		},
	})
}

// GetServer returns a single server
func (h *DedicatedServerHandler) GetServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	server, err := h.serverRepo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	return c.JSON(server)
}

// GetServerByUUID returns a server by UUID
func (h *DedicatedServerHandler) GetServerByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	server, err := h.serverRepo.GetByUUID(uuid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	return c.JSON(server)
}

// CreateServer creates a new server
func (h *DedicatedServerHandler) CreateServer(c *fiber.Ctx) error {
	var input struct {
		Name          string            `json:"name"`
		Description   string            `json:"description"`
		NodeID        int64             `json:"nodeId"`
		EggID         int64             `json:"eggId"`
		Memory        int64             `json:"memory"`
		Disk          int64             `json:"disk"`
		CPU           int               `json:"cpu"`
		IO            int               `json:"io"`
		Swap          int64             `json:"swap"`
		Threads       string            `json:"threads"`
		OOMDisabled   bool              `json:"oomDisabled"`
		AllocationID  int64             `json:"allocationId"`
		Image         string            `json:"image"`
		StartupCmd    string            `json:"startupCommand"`
		Environment   map[string]string `json:"environment"`
		BackupLimit   int               `json:"backupLimit"`
		DatabaseLimit int               `json:"databaseLimit"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if input.Name == "" || input.NodeID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name and NodeID are required",
		})
	}

	if input.Memory <= 0 || input.Disk <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Memory and Disk must be positive",
		})
	}

	// Verify node exists
	node, err := h.nodeRepo.GetByID(input.NodeID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Node not found",
		})
	}

	if node.MaintenanceMode {
		return c.Status(400).JSON(fiber.Map{
			"error": "Node is in maintenance mode",
		})
	}

	// Generate short ID
	shortIDBytes := make([]byte, 4)
	rand.Read(shortIDBytes)
	shortID := hex.EncodeToString(shortIDBytes)

	// Set defaults
	if input.CPU == 0 {
		input.CPU = 100
	}
	if input.IO == 0 {
		input.IO = 500
	}

	// Convert environment to JSON
	var envJSON []byte
	if input.Environment != nil {
		envJSON, _ = json.Marshal(input.Environment)
	}

	server := &models.Server{
		UUID:                uuid.New().String(),
		ShortID:             shortID,
		Name:                input.Name,
		Description:         input.Description,
		NodeID:              input.NodeID,
		EggID:               input.EggID,
		Status:              models.ServerStatusOffline,
		Memory:              input.Memory,
		Disk:                input.Disk,
		CPU:                 input.CPU,
		IO:                  input.IO,
		Swap:                input.Swap,
		Threads:             input.Threads,
		OOMDisabled:         input.OOMDisabled,
		DefaultAllocationID: input.AllocationID,
		Image:               input.Image,
		StartupCommand:      input.StartupCmd,
		Environment:         envJSON,
		BackupLimit:         input.BackupLimit,
		DatabaseLimit:       input.DatabaseLimit,
		AllocationLimit:     1,
	}

	if err := h.serverRepo.Create(server); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create server: " + err.Error(),
		})
	}

	// Assign allocation to server if specified
	if input.AllocationID > 0 {
		h.nodeRepo.AssignAllocation(input.AllocationID, server.ID)
	}

	return c.Status(201).JSON(server)
}

// UpdateServer updates a server
func (h *DedicatedServerHandler) UpdateServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	existing, err := h.serverRepo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	var input struct {
		Name          string            `json:"name"`
		Description   string            `json:"description"`
		Memory        int64             `json:"memory"`
		Disk          int64             `json:"disk"`
		CPU           int               `json:"cpu"`
		IO            int               `json:"io"`
		Swap          int64             `json:"swap"`
		Threads       string            `json:"threads"`
		OOMDisabled   bool              `json:"oomDisabled"`
		Image         string            `json:"image"`
		StartupCmd    string            `json:"startupCommand"`
		Environment   map[string]string `json:"environment"`
		BackupLimit   int               `json:"backupLimit"`
		DatabaseLimit int               `json:"databaseLimit"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var envJSON []byte
	if input.Environment != nil {
		envJSON, _ = json.Marshal(input.Environment)
	}

	server := &models.Server{
		ID:                  existing.ID,
		UUID:                existing.UUID,
		ShortID:             existing.ShortID,
		Name:                input.Name,
		Description:         input.Description,
		NodeID:              existing.NodeID,
		Status:              existing.Status,
		Suspended:           existing.Suspended,
		Memory:              input.Memory,
		Disk:                input.Disk,
		CPU:                 input.CPU,
		IO:                  input.IO,
		Swap:                input.Swap,
		Threads:             input.Threads,
		OOMDisabled:         input.OOMDisabled,
		DefaultAllocationID: existing.DefaultAllocationID,
		Image:               input.Image,
		StartupCommand:      input.StartupCmd,
		Environment:         envJSON,
		BackupLimit:         input.BackupLimit,
		DatabaseLimit:       input.DatabaseLimit,
		AllocationLimit:     existing.AllocationLimit,
	}

	if err := h.serverRepo.Update(server); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update server: " + err.Error(),
		})
	}

	return c.JSON(server)
}

// DeleteServer deletes a server
func (h *DedicatedServerHandler) DeleteServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	server, err := h.serverRepo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	// Check if server is running
	if server.Status == models.ServerStatusRunning {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot delete a running server. Please stop it first.",
		})
	}

	// Unassign allocations
	for _, alloc := range server.Allocations {
		h.nodeRepo.UnassignAllocation(alloc.ID)
	}

	if err := h.serverRepo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete server: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Server deleted successfully",
	})
}

// Power actions

// StartServer starts a server
func (h *DedicatedServerHandler) StartServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	server, err := h.serverRepo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	if server.Suspended {
		return c.Status(400).JSON(fiber.Map{
			"error": "Server is suspended",
		})
	}

	if server.Status == models.ServerStatusRunning {
		return c.Status(400).JSON(fiber.Map{
			"error": "Server is already running",
		})
	}

	// Update status to starting
	if err := h.serverRepo.UpdateStatus(id, models.ServerStatusStarting); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update server status",
		})
	}

	// TODO: Implement actual server start logic (Docker/process management)

	return c.JSON(fiber.Map{
		"message": "Server starting",
		"status":  "starting",
	})
}

// StopServer stops a server
func (h *DedicatedServerHandler) StopServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	server, err := h.serverRepo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	if server.Status == models.ServerStatusOffline {
		return c.Status(400).JSON(fiber.Map{
			"error": "Server is already offline",
		})
	}

	if err := h.serverRepo.UpdateStatus(id, models.ServerStatusStopping); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update server status",
		})
	}

	// TODO: Implement actual server stop logic

	return c.JSON(fiber.Map{
		"message": "Server stopping",
		"status":  "stopping",
	})
}

// RestartServer restarts a server
func (h *DedicatedServerHandler) RestartServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	server, err := h.serverRepo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	if server.Suspended {
		return c.Status(400).JSON(fiber.Map{
			"error": "Server is suspended",
		})
	}

	if err := h.serverRepo.UpdateStatus(id, models.ServerStatusRestarting); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update server status",
		})
	}

	// TODO: Implement actual server restart logic

	return c.JSON(fiber.Map{
		"message": "Server restarting",
		"status":  "restarting",
	})
}

// KillServer force kills a server
func (h *DedicatedServerHandler) KillServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	if _, err := h.serverRepo.GetByID(id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	if err := h.serverRepo.UpdateStatus(id, models.ServerStatusOffline); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update server status",
		})
	}

	// TODO: Implement actual server kill logic

	return c.JSON(fiber.Map{
		"message": "Server killed",
		"status":  "offline",
	})
}

// SuspendServer suspends a server
func (h *DedicatedServerHandler) SuspendServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	if err := h.serverRepo.Suspend(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to suspend server: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Server suspended",
	})
}

// UnsuspendServer unsuspends a server
func (h *DedicatedServerHandler) UnsuspendServer(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	if err := h.serverRepo.Unsuspend(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to unsuspend server: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Server unsuspended",
	})
}

// Eggs & Nests

// ListNests returns all nests
func (h *DedicatedServerHandler) ListNests(c *fiber.Ctx) error {
	nests, err := h.serverRepo.GetAllNests()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch nests: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"nests": nests,
	})
}

// ListEggs returns all eggs
func (h *DedicatedServerHandler) ListEggs(c *fiber.Ctx) error {
	eggs, err := h.serverRepo.GetAllEggs()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch eggs: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"eggs": eggs,
	})
}

// GetEgg returns a single egg
func (h *DedicatedServerHandler) GetEgg(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid egg ID",
		})
	}

	egg, err := h.serverRepo.GetEggByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Egg not found",
		})
	}

	return c.JSON(egg)
}

// GetServerStats returns aggregated statistics
func (h *DedicatedServerHandler) GetServerStats(c *fiber.Ctx) error {
	stats, err := h.serverRepo.GetServerStats()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch stats: " + err.Error(),
		})
	}

	return c.JSON(stats)
}

// SendCommand sends a command to a server
func (h *DedicatedServerHandler) SendCommand(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid server ID",
		})
	}

	var input struct {
		Command string `json:"command"`
	}

	if err := c.BodyParser(&input); err != nil || input.Command == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Command is required",
		})
	}

	server, err := h.serverRepo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Server not found",
		})
	}

	if server.Status != models.ServerStatusRunning {
		return c.Status(400).JSON(fiber.Map{
			"error": "Server is not running",
		})
	}

	// TODO: Implement actual command sending logic

	return c.JSON(fiber.Map{
		"message": "Command sent",
		"command": input.Command,
	})
}
