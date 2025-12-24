package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/withsait/karapixel/karapanel/internal/server"
)

type ServerHandler struct {
	manager *server.Manager
}

func NewServerHandler(manager *server.Manager) *ServerHandler {
	return &ServerHandler{manager: manager}
}

// ListServers returns all configured servers with their status
func (h *ServerHandler) ListServers(c *fiber.Ctx) error {
	servers := h.manager.GetAllServers()
	return c.JSON(fiber.Map{
		"servers": servers,
	})
}

// GetServer returns a specific server's status
func (h *ServerHandler) GetServer(c *fiber.Ctx) error {
	id := c.Params("id")

	info, err := h.manager.GetServer(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(info)
}

// StartServer starts a specific server
func (h *ServerHandler) StartServer(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.manager.StartServer(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Server starting",
		"id":      id,
	})
}

// StopServer stops a specific server
func (h *ServerHandler) StopServer(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.manager.StopServer(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Server stopping",
		"id":      id,
	})
}

// RestartServer restarts a specific server
func (h *ServerHandler) RestartServer(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.manager.RestartServer(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Server restarting",
		"id":      id,
	})
}

// KillServer force kills a specific server
func (h *ServerHandler) KillServer(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.manager.KillServer(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Server killed",
		"id":      id,
	})
}
