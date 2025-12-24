package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/withsait/karapixel/karapanel/internal/api/handlers"
	"github.com/withsait/karapixel/karapanel/internal/api/middleware"
	"github.com/withsait/karapixel/karapanel/internal/config"
	"github.com/withsait/karapixel/karapanel/internal/metrics"
	"github.com/withsait/karapixel/karapanel/internal/server"
)

type Router struct {
	app            *fiber.App
	serverHandler  *handlers.ServerHandler
	consoleHandler *handlers.ConsoleHandler
	metricsHandler *handlers.MetricsHandler
	filesHandler   *handlers.FilesHandler
	authMiddleware *middleware.AuthMiddleware
}

func NewRouter(cfg *config.Config, manager *server.Manager, consoleManager *server.ConsoleManager, collector *metrics.Collector) *Router {
	app := fiber.New(fiber.Config{
		AppName:      "KaraPanel Daemon",
		ServerHeader: "KaraPanel",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH",
	}))

	router := &Router{
		app:            app,
		serverHandler:  handlers.NewServerHandler(manager),
		consoleHandler: handlers.NewConsoleHandler(manager, consoleManager),
		metricsHandler: handlers.NewMetricsHandler(collector),
		filesHandler:   handlers.NewFilesHandler(manager),
		authMiddleware: middleware.NewAuthMiddleware(&cfg.Auth),
	}

	router.setupRoutes()

	return router
}

func (r *Router) setupRoutes() {
	// Health check (no auth)
	r.app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"name":   "KaraPanel Daemon",
		})
	})

	// Auth routes
	r.app.Post("/api/auth/login", r.authMiddleware.Login)

	// Protected API routes
	api := r.app.Group("/api", r.authMiddleware.Protect())

	// Server routes
	servers := api.Group("/servers")
	servers.Get("/", r.serverHandler.ListServers)
	servers.Get("/:id", r.serverHandler.GetServer)
	servers.Post("/:id/start", r.serverHandler.StartServer)
	servers.Post("/:id/stop", r.serverHandler.StopServer)
	servers.Post("/:id/restart", r.serverHandler.RestartServer)
	servers.Post("/:id/kill", r.serverHandler.KillServer)

	// Console routes
	servers.Get("/:id/logs", r.consoleHandler.GetLogs)

	// WebSocket for console (needs special handling)
	r.app.Use("/api/servers/:id/console", handlers.WebSocketUpgrade())
	r.app.Get("/api/servers/:id/console", websocket.New(r.consoleHandler.StreamLogs))

	// Files routes
	servers.Get("/:id/files", r.filesHandler.ListFiles)
	servers.Get("/:id/files/content", r.filesHandler.GetFile)
	servers.Put("/:id/files/content", r.filesHandler.SaveFile)
	servers.Post("/:id/files/upload", r.filesHandler.UploadFile)
	servers.Delete("/:id/files", r.filesHandler.DeleteFile)
	servers.Get("/:id/files/download", r.filesHandler.DownloadFile)

	// Metrics routes
	metrics := api.Group("/metrics")
	metrics.Get("/", r.metricsHandler.GetSystemMetrics)
	metrics.Get("/cpu", r.metricsHandler.GetCPUMetrics)
	metrics.Get("/memory", r.metricsHandler.GetMemoryMetrics)
	metrics.Get("/disk", r.metricsHandler.GetDiskMetrics)
	metrics.Get("/network", r.metricsHandler.GetNetworkMetrics)
}

func (r *Router) Listen(addr string) error {
	return r.app.Listen(addr)
}

func (r *Router) Shutdown() error {
	return r.app.Shutdown()
}
