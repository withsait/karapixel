package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/withsait/karapixel/karapanel/internal/models"
)

type DiscordHandler struct {
	repo *models.DiscordRepository
}

func NewDiscordHandler(repo *models.DiscordRepository) *DiscordHandler {
	return &DiscordHandler{repo: repo}
}

// GET /api/discord/links
func (h *DiscordHandler) ListLinks(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	if limit > 100 {
		limit = 100
	}

	links, total, err := h.repo.GetAllLinks(limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch links"})
	}

	return c.JSON(fiber.Map{
		"links":  links,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

// GET /api/discord/links/player/:uuid
func (h *DiscordHandler) GetPlayerLink(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	link, err := h.repo.GetLink(uuid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Link not found"})
	}

	return c.JSON(link)
}

// GET /api/discord/links/discord/:discordId
func (h *DiscordHandler) GetLinkByDiscordID(c *fiber.Ctx) error {
	discordID := c.Params("discordId")

	link, err := h.repo.GetLinkByDiscordID(discordID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Link not found"})
	}

	return c.JSON(link)
}

// POST /api/discord/links
func (h *DiscordHandler) CreateLink(c *fiber.Ctx) error {
	var body struct {
		PlayerUUID      string `json:"playerUuid"`
		DiscordID       string `json:"discordId"`
		DiscordUsername string `json:"discordUsername"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	link, err := h.repo.CreateLink(body.PlayerUUID, body.DiscordID, body.DiscordUsername)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create link", "details": err.Error()})
	}

	return c.Status(201).JSON(link)
}

// POST /api/discord/links/:uuid/verify
func (h *DiscordHandler) VerifyLink(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	if err := h.repo.VerifyLink(uuid); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to verify link"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Link verified"})
}

// DELETE /api/discord/links/:uuid
func (h *DiscordHandler) DeleteLink(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	if err := h.repo.DeleteLink(uuid); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete link"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Link deleted"})
}

// GET /api/discord/settings
func (h *DiscordHandler) ListSettings(c *fiber.Ctx) error {
	settings, err := h.repo.GetAllSettings()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch settings"})
	}

	return c.JSON(settings)
}

// GET /api/discord/settings/:guildId
func (h *DiscordHandler) GetSettings(c *fiber.Ctx) error {
	guildID := c.Params("guildId")

	settings, err := h.repo.GetSettings(guildID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Settings not found"})
	}

	return c.JSON(settings)
}

// PUT /api/discord/settings/:guildId
func (h *DiscordHandler) SaveSettings(c *fiber.Ctx) error {
	guildID := c.Params("guildId")

	var settings models.DiscordSettings
	if err := c.BodyParser(&settings); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	settings.GuildID = guildID

	if err := h.repo.SaveSettings(&settings); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save settings", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Settings saved"})
}
