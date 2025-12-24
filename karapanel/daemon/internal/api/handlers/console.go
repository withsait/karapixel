package handlers

import (
	"context"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/withsait/karapixel/karapanel/internal/server"
)

type ConsoleHandler struct {
	manager        *server.Manager
	consoleManager *server.ConsoleManager
	systemd        *server.SystemdClient
}

func NewConsoleHandler(manager *server.Manager, consoleManager *server.ConsoleManager) *ConsoleHandler {
	systemd, _ := server.NewSystemdClient()
	return &ConsoleHandler{
		manager:        manager,
		consoleManager: consoleManager,
		systemd:        systemd,
	}
}

// GetLogs returns recent logs for a server
func (h *ConsoleHandler) GetLogs(c *fiber.Ctx) error {
	id := c.Params("id")
	lines := c.QueryInt("lines", 100)

	if lines > 1000 {
		lines = 1000
	}

	logs, err := h.systemd.GetRecentLogs(id, lines)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"logs": logs,
	})
}

// StreamLogs handles WebSocket connections for real-time log streaming
func (h *ConsoleHandler) StreamLogs(c *websocket.Conn) {
	serverID := c.Params("id")
	subscriberID := uuid.New().String()

	// Get server config for workDir
	srv, err := h.manager.GetServerConfig(serverID)
	if err != nil {
		c.WriteJSON(fiber.Map{
			"error": err.Error(),
		})
		c.Close()
		return
	}

	// Subscribe to logs
	logChan, err := h.consoleManager.Subscribe(serverID, subscriberID, srv.WorkDir)
	if err != nil {
		c.WriteJSON(fiber.Map{
			"error": err.Error(),
		})
		c.Close()
		return
	}
	defer h.consoleManager.Unsubscribe(serverID, subscriberID)

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle incoming messages (commands)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, msg, err := c.ReadMessage()
				if err != nil {
					cancel()
					return
				}
				// Handle command input if needed
				_ = msg // TODO: Send to RCON or screen session
			}
		}
	}()

	// Send logs to client
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-logChan:
			if !ok {
				return
			}
			if err := c.WriteJSON(fiber.Map{
				"type": "log",
				"data": log,
			}); err != nil {
				return
			}
		}
	}
}

// WebSocketUpgrade middleware checks if request is a WebSocket upgrade
func WebSocketUpgrade() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}
