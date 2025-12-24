package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/withsait/karapixel/karapanel/internal/models"
)

type AnalyticsHandler struct {
	repo *models.AnalyticsRepository
}

func NewAnalyticsHandler(repo *models.AnalyticsRepository) *AnalyticsHandler {
	return &AnalyticsHandler{repo: repo}
}

// GET /api/analytics/dashboard
func (h *AnalyticsHandler) GetDashboardStats(c *fiber.Ctx) error {
	stats, err := h.repo.GetDashboardStats()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch dashboard stats"})
	}

	return c.JSON(stats)
}

// GET /api/analytics/players
func (h *AnalyticsHandler) GetPlayerHistory(c *fiber.Ctx) error {
	hours, _ := strconv.Atoi(c.Query("hours", "24"))

	if hours > 168 { // max 1 week
		hours = 168
	}

	history, err := h.repo.GetPlayerCountHistory(hours)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch player history"})
	}

	return c.JSON(history)
}

// GET /api/analytics/server/:serverId
func (h *AnalyticsHandler) GetServerStats(c *fiber.Ctx) error {
	serverID := c.Params("serverId")
	hours, _ := strconv.Atoi(c.Query("hours", "24"))

	if hours > 168 {
		hours = 168
	}

	stats, err := h.repo.GetServerStats(serverID, hours)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch server stats"})
	}

	return c.JSON(stats)
}

// POST /api/analytics/server
func (h *AnalyticsHandler) RecordServerStat(c *fiber.Ctx) error {
	var stat models.ServerStat
	if err := c.BodyParser(&stat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.RecordServerStat(&stat); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to record stat"})
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}

// GET /api/analytics/logs
func (h *AnalyticsHandler) GetActivityLogs(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	actorType := c.Query("actorType", "")
	action := c.Query("action", "")

	if limit > 200 {
		limit = 200
	}

	logs, total, err := h.repo.GetActivityLogs(limit, offset, actorType, action)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch logs"})
	}

	return c.JSON(fiber.Map{
		"logs":   logs,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

// POST /api/analytics/logs
func (h *AnalyticsHandler) CreateActivityLog(c *fiber.Ctx) error {
	var log models.ActivityLog
	if err := c.BodyParser(&log); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.LogActivity(&log); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create log"})
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}

// GET /api/notifications
func (h *AnalyticsHandler) GetNotifications(c *fiber.Ctx) error {
	targetUser := c.Query("user", "admin")
	unreadOnly := c.Query("unread", "false") == "true"
	limit, _ := strconv.Atoi(c.Query("limit", "20"))

	notifications, err := h.repo.GetNotifications(targetUser, unreadOnly, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch notifications"})
	}

	return c.JSON(notifications)
}

// POST /api/notifications
func (h *AnalyticsHandler) CreateNotification(c *fiber.Ctx) error {
	var notification models.Notification
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.CreateNotification(&notification); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create notification"})
	}

	return c.Status(201).JSON(notification)
}

// POST /api/notifications/:id/read
func (h *AnalyticsHandler) MarkNotificationRead(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid notification ID"})
	}

	if err := h.repo.MarkNotificationRead(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to mark notification as read"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// POST /api/notifications/read-all
func (h *AnalyticsHandler) MarkAllNotificationsRead(c *fiber.Ctx) error {
	targetUser := c.Query("user", "admin")

	if err := h.repo.MarkAllNotificationsRead(targetUser); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to mark notifications as read"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// GET /api/webhooks
func (h *AnalyticsHandler) GetWebhooks(c *fiber.Ctx) error {
	webhooks, err := h.repo.GetWebhooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch webhooks"})
	}

	return c.JSON(webhooks)
}

// POST /api/webhooks
func (h *AnalyticsHandler) CreateWebhook(c *fiber.Ctx) error {
	var webhook models.Webhook
	if err := c.BodyParser(&webhook); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.repo.CreateWebhook(&webhook); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create webhook"})
	}

	return c.Status(201).JSON(webhook)
}

// PUT /api/webhooks/:id
func (h *AnalyticsHandler) UpdateWebhook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid webhook ID"})
	}

	var webhook models.Webhook
	if err := c.BodyParser(&webhook); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	webhook.ID = id
	if err := h.repo.UpdateWebhook(&webhook); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update webhook"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// DELETE /api/webhooks/:id
func (h *AnalyticsHandler) DeleteWebhook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid webhook ID"})
	}

	if err := h.repo.DeleteWebhook(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete webhook"})
	}

	return c.JSON(fiber.Map{"success": true})
}
