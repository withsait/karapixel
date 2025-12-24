package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/withsait/karapixel/karapanel/internal/models"
)

type PlayerHandler struct {
	repo *models.PlayerRepository
}

func NewPlayerHandler(repo *models.PlayerRepository) *PlayerHandler {
	return &PlayerHandler{repo: repo}
}

// GET /api/players
func (h *PlayerHandler) ListPlayers(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	search := c.Query("search", "")
	onlineOnly := c.Query("online", "false") == "true"

	if limit > 100 {
		limit = 100
	}

	players, total, err := h.repo.GetAll(limit, offset, search, onlineOnly)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch players", "details": err.Error()})
	}

	return c.JSON(fiber.Map{
		"players": players,
		"total":   total,
		"limit":   limit,
		"offset":  offset,
	})
}

// GET /api/players/:uuid
func (h *PlayerHandler) GetPlayer(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	player, err := h.repo.GetByUUID(uuid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Player not found"})
	}

	return c.JSON(player)
}

// GET /api/players/search/:username
func (h *PlayerHandler) SearchPlayer(c *fiber.Ctx) error {
	username := c.Params("username")

	player, err := h.repo.GetByUsername(username)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Player not found"})
	}

	return c.JSON(player)
}

// POST /api/players
func (h *PlayerHandler) CreateOrUpdatePlayer(c *fiber.Ctx) error {
	var player models.Player
	if err := c.BodyParser(&player); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.Create(&player); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create/update player", "details": err.Error()})
	}

	return c.Status(201).JSON(player)
}

// PATCH /api/players/:uuid/online
func (h *PlayerHandler) UpdateOnlineStatus(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	var body struct {
		IsOnline bool `json:"isOnline"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.UpdateOnlineStatus(uuid, body.IsOnline); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update online status"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// POST /api/players/:uuid/stats
func (h *PlayerHandler) UpdateStats(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	var stats models.PlayerStats
	if err := c.BodyParser(&stats); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	stats.PlayerUUID = uuid
	if err := h.repo.UpdateStats(&stats); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update stats", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

// POST /api/players/:uuid/ip
func (h *PlayerHandler) AddIPRecord(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	var body struct {
		IP string `json:"ip"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.AddIPRecord(uuid, body.IP); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to add IP record"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// GET /api/players/stats
func (h *PlayerHandler) GetPlayerStats(c *fiber.Ctx) error {
	online, _ := h.repo.GetOnlineCount()
	total, _ := h.repo.GetTotalCount()
	newToday, _ := h.repo.GetNewPlayersToday()

	return c.JSON(fiber.Map{
		"online":   online,
		"total":    total,
		"newToday": newToday,
	})
}
