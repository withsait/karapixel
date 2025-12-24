package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/withsait/karapixel/karapanel/internal/models"
)

type PunishmentHandler struct {
	repo *models.PunishmentRepository
}

func NewPunishmentHandler(repo *models.PunishmentRepository) *PunishmentHandler {
	return &PunishmentHandler{repo: repo}
}

// GET /api/punishments
func (h *PunishmentHandler) ListPunishments(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	punishmentType := c.Query("type", "")
	status := c.Query("status", "")

	if limit > 100 {
		limit = 100
	}

	punishments, total, err := h.repo.GetAll(limit, offset, punishmentType, status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch punishments", "details": err.Error()})
	}

	return c.JSON(fiber.Map{
		"punishments": punishments,
		"total":       total,
		"limit":       limit,
		"offset":      offset,
	})
}

// GET /api/punishments/player/:uuid
func (h *PunishmentHandler) GetPlayerPunishments(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	punishments, err := h.repo.GetByPlayer(uuid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch punishments"})
	}

	return c.JSON(punishments)
}

// POST /api/punishments
func (h *PunishmentHandler) CreatePunishment(c *fiber.Ctx) error {
	var req models.CreatePunishmentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate type
	validTypes := map[string]bool{"ban": true, "kick": true, "mute": true, "warn": true}
	if !validTypes[req.Type] {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid punishment type. Must be: ban, kick, mute, or warn"})
	}

	punishment, err := h.repo.Create(&req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create punishment", "details": err.Error()})
	}

	return c.Status(201).JSON(punishment)
}

// POST /api/punishments/:id/revoke
func (h *PunishmentHandler) RevokePunishment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid punishment ID"})
	}

	var body struct {
		ModeratorName string `json:"moderatorName"`
	}
	c.BodyParser(&body)

	if err := h.repo.Revoke(id, body.ModeratorName); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to revoke punishment"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Punishment revoked"})
}

// POST /api/punishments/:id/appeal
func (h *PunishmentHandler) AppealPunishment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid punishment ID"})
	}

	var body struct {
		Reason string `json:"reason"`
	}
	if err := c.BodyParser(&body); err != nil || body.Reason == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Appeal reason is required"})
	}

	if err := h.repo.Appeal(id, body.Reason); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to submit appeal"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Appeal submitted"})
}

// POST /api/punishments/:id/appeal/handle
func (h *PunishmentHandler) HandleAppeal(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid punishment ID"})
	}

	var body struct {
		Approved bool `json:"approved"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.HandleAppeal(id, body.Approved); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to handle appeal"})
	}

	status := "rejected"
	if body.Approved {
		status = "approved"
	}
	return c.JSON(fiber.Map{"success": true, "message": "Appeal " + status})
}

// GET /api/punishments/check/:uuid
func (h *PunishmentHandler) CheckBan(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	ban, err := h.repo.GetActiveBan(uuid)
	if err != nil {
		return c.JSON(fiber.Map{"banned": false})
	}

	return c.JSON(fiber.Map{
		"banned":    true,
		"reason":    ban.Reason,
		"expiresAt": ban.ExpiresAt,
		"permanent": ban.Duration == nil || *ban.Duration == 0,
	})
}

// GET /api/punishments/templates
func (h *PunishmentHandler) GetTemplates(c *fiber.Ctx) error {
	templates, err := h.repo.GetTemplates()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch templates"})
	}

	return c.JSON(templates)
}

// POST /api/punishments/templates
func (h *PunishmentHandler) CreateTemplate(c *fiber.Ctx) error {
	var template models.PunishmentTemplate
	if err := c.BodyParser(&template); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.CreateTemplate(&template); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create template"})
	}

	return c.Status(201).JSON(template)
}

// DELETE /api/punishments/templates/:id
func (h *PunishmentHandler) DeleteTemplate(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid template ID"})
	}

	if err := h.repo.DeleteTemplate(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete template"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// GET /api/punishments/stats
func (h *PunishmentHandler) GetStats(c *fiber.Ctx) error {
	stats, err := h.repo.GetStats()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch stats"})
	}

	return c.JSON(stats)
}
