package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/withsait/karapixel/karapanel/internal/models"
)

type NodeHandler struct {
	repo *models.NodeRepository
}

func NewNodeHandler(repo *models.NodeRepository) *NodeHandler {
	return &NodeHandler{repo: repo}
}

// ListNodes returns all nodes
func (h *NodeHandler) ListNodes(c *fiber.Ctx) error {
	includeOffline := c.Query("includeOffline", "true") == "true"

	nodes, err := h.repo.GetAll(includeOffline)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch nodes: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"nodes": nodes,
		"total": len(nodes),
	})
}

// GetNode returns a single node by ID
func (h *NodeHandler) GetNode(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid node ID",
		})
	}

	node, err := h.repo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Node not found",
		})
	}

	return c.JSON(node)
}

// CreateNode creates a new node
func (h *NodeHandler) CreateNode(c *fiber.Ctx) error {
	var input struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		FQDN            string `json:"fqdn"`
		Scheme          string `json:"scheme"`
		DaemonPort      int    `json:"daemonPort"`
		Memory          int64  `json:"memory"`
		MemoryOveralloc int    `json:"memoryOveralloc"`
		Disk            int64  `json:"disk"`
		DiskOveralloc   int    `json:"diskOveralloc"`
		UploadLimit     int    `json:"uploadLimit"`
		DownloadLimit   int    `json:"downloadLimit"`
		LocationID      int64  `json:"locationId"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if input.Name == "" || input.FQDN == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name and FQDN are required",
		})
	}

	if input.Memory <= 0 || input.Disk <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Memory and Disk must be positive",
		})
	}

	if input.Scheme == "" {
		input.Scheme = "https"
	}
	if input.DaemonPort == 0 {
		input.DaemonPort = 8080
	}

	// Generate daemon token
	tokenBytes := make([]byte, 32)
	rand.Read(tokenBytes)
	daemonToken := hex.EncodeToString(tokenBytes)

	node := &models.Node{
		UUID:            uuid.New().String(),
		Name:            input.Name,
		Description:     input.Description,
		FQDN:            input.FQDN,
		Scheme:          input.Scheme,
		DaemonPort:      input.DaemonPort,
		DaemonToken:     daemonToken,
		Memory:          input.Memory,
		MemoryOveralloc: input.MemoryOveralloc,
		Disk:            input.Disk,
		DiskOveralloc:   input.DiskOveralloc,
		UploadLimit:     input.UploadLimit,
		DownloadLimit:   input.DownloadLimit,
		Status:          models.NodeStatusOffline,
		LocationID:      input.LocationID,
	}

	if err := h.repo.Create(node); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create node: " + err.Error(),
		})
	}

	return c.Status(201).JSON(node)
}

// UpdateNode updates a node
func (h *NodeHandler) UpdateNode(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid node ID",
		})
	}

	existing, err := h.repo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Node not found",
		})
	}

	var input struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		FQDN            string `json:"fqdn"`
		Scheme          string `json:"scheme"`
		DaemonPort      int    `json:"daemonPort"`
		Memory          int64  `json:"memory"`
		MemoryOveralloc int    `json:"memoryOveralloc"`
		Disk            int64  `json:"disk"`
		DiskOveralloc   int    `json:"diskOveralloc"`
		UploadLimit     int    `json:"uploadLimit"`
		DownloadLimit   int    `json:"downloadLimit"`
		LocationID      int64  `json:"locationId"`
		MaintenanceMode bool   `json:"maintenanceMode"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	node := &models.Node{
		ID:              existing.ID,
		UUID:            existing.UUID,
		Name:            input.Name,
		Description:     input.Description,
		FQDN:            input.FQDN,
		Scheme:          input.Scheme,
		DaemonPort:      input.DaemonPort,
		DaemonToken:     existing.DaemonToken,
		Memory:          input.Memory,
		MemoryOveralloc: input.MemoryOveralloc,
		Disk:            input.Disk,
		DiskOveralloc:   input.DiskOveralloc,
		UploadLimit:     input.UploadLimit,
		DownloadLimit:   input.DownloadLimit,
		Status:          existing.Status,
		MaintenanceMode: input.MaintenanceMode,
		LocationID:      input.LocationID,
	}

	if err := h.repo.Update(node); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update node: " + err.Error(),
		})
	}

	return c.JSON(node)
}

// DeleteNode deletes a node
func (h *NodeHandler) DeleteNode(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid node ID",
		})
	}

	// Check if node has servers
	node, err := h.repo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Node not found",
		})
	}

	if node.ServerCount > 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot delete node with active servers",
		})
	}

	if err := h.repo.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete node: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Node deleted successfully",
	})
}

// RegenerateToken regenerates the daemon token for a node
func (h *NodeHandler) RegenerateToken(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid node ID",
		})
	}

	existing, err := h.repo.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Node not found",
		})
	}

	// Generate new token
	tokenBytes := make([]byte, 32)
	rand.Read(tokenBytes)
	newToken := hex.EncodeToString(tokenBytes)

	node := &existing.Node
	node.DaemonToken = newToken

	if err := h.repo.Update(node); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to regenerate token: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": newToken,
	})
}

// GetNodeAllocations returns allocations for a node
func (h *NodeHandler) GetNodeAllocations(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid node ID",
		})
	}

	allocations, err := h.repo.GetAllocations(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch allocations: " + err.Error(),
		})
	}

	// Group by IP
	grouped := make(map[string][]models.Allocation)
	for _, a := range allocations {
		grouped[a.IP] = append(grouped[a.IP], a)
	}

	return c.JSON(fiber.Map{
		"allocations": allocations,
		"grouped":     grouped,
		"total":       len(allocations),
	})
}

// CreateAllocation creates a new allocation
func (h *NodeHandler) CreateAllocation(c *fiber.Ctx) error {
	nodeID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid node ID",
		})
	}

	var input struct {
		IP        string `json:"ip"`
		Alias     string `json:"alias"`
		Port      int    `json:"port"`
		PortRange string `json:"portRange"` // e.g., "25565-25575"
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if input.IP == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "IP is required",
		})
	}

	// If port range is specified, create multiple allocations
	if input.PortRange != "" {
		var startPort, endPort int
		_, err := fmt.Sscanf(input.PortRange, "%d-%d", &startPort, &endPort)
		if err != nil || startPort > endPort || startPort < 1 || endPort > 65535 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid port range",
			})
		}

		if endPort-startPort > 1000 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Port range too large (max 1000 ports)",
			})
		}

		if err := h.repo.CreateAllocationRange(nodeID, input.IP, startPort, endPort); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to create allocations: " + err.Error(),
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Allocations created",
			"count":   endPort - startPort + 1,
		})
	}

	// Single port allocation
	if input.Port < 1 || input.Port > 65535 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid port number",
		})
	}

	alloc := &models.Allocation{
		NodeID: nodeID,
		IP:     input.IP,
		Alias:  input.Alias,
		Port:   input.Port,
	}

	if err := h.repo.CreateAllocation(alloc); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create allocation: " + err.Error(),
		})
	}

	return c.Status(201).JSON(alloc)
}

// DeleteAllocation deletes an allocation
func (h *NodeHandler) DeleteAllocation(c *fiber.Ctx) error {
	allocID, err := strconv.ParseInt(c.Params("allocId"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid allocation ID",
		})
	}

	if err := h.repo.DeleteAllocation(allocID); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete allocation: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Allocation deleted",
	})
}

// Locations

// ListLocations returns all locations
func (h *NodeHandler) ListLocations(c *fiber.Ctx) error {
	locations, err := h.repo.GetAllLocations()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch locations: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"locations": locations,
	})
}

// CreateLocation creates a new location
func (h *NodeHandler) CreateLocation(c *fiber.Ctx) error {
	var input struct {
		Short string `json:"short"`
		Long  string `json:"long"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if input.Short == "" || input.Long == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Short and Long are required",
		})
	}

	location := &models.Location{
		Short: input.Short,
		Long:  input.Long,
	}

	if err := h.repo.CreateLocation(location); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create location: " + err.Error(),
		})
	}

	return c.Status(201).JSON(location)
}

// DeleteLocation deletes a location
func (h *NodeHandler) DeleteLocation(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid location ID",
		})
	}

	if err := h.repo.DeleteLocation(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete location: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Location deleted",
	})
}
