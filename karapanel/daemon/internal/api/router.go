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
	"github.com/withsait/karapixel/karapanel/internal/database"
	"github.com/withsait/karapixel/karapanel/internal/metrics"
	"github.com/withsait/karapixel/karapanel/internal/models"
	"github.com/withsait/karapixel/karapanel/internal/server"
)

type Router struct {
	app                    *fiber.App
	serverHandler          *handlers.ServerHandler
	consoleHandler         *handlers.ConsoleHandler
	metricsHandler         *handlers.MetricsHandler
	filesHandler           *handlers.FilesHandler
	playerHandler          *handlers.PlayerHandler
	punishmentHandler      *handlers.PunishmentHandler
	discordHandler         *handlers.DiscordHandler
	analyticsHandler       *handlers.AnalyticsHandler
	nodeHandler            *handlers.NodeHandler
	dedicatedServerHandler *handlers.DedicatedServerHandler
	authMiddleware         *middleware.AuthMiddleware
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

	// Initialize database-dependent handlers if database is available
	db := database.GetDB()
	if db != nil {
		playerRepo := models.NewPlayerRepository(db)
		punishmentRepo := models.NewPunishmentRepository(db)
		discordRepo := models.NewDiscordRepository(db)
		analyticsRepo := models.NewAnalyticsRepository(db)
		nodeRepo := models.NewNodeRepository(db)
		serverRepo := models.NewServerRepository(db)

		router.playerHandler = handlers.NewPlayerHandler(playerRepo)
		router.punishmentHandler = handlers.NewPunishmentHandler(punishmentRepo)
		router.discordHandler = handlers.NewDiscordHandler(discordRepo)
		router.analyticsHandler = handlers.NewAnalyticsHandler(analyticsRepo)
		router.nodeHandler = handlers.NewNodeHandler(nodeRepo)
		router.dedicatedServerHandler = handlers.NewDedicatedServerHandler(serverRepo, nodeRepo)
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
	metricsGroup := api.Group("/metrics")
	metricsGroup.Get("/", r.metricsHandler.GetSystemMetrics)
	metricsGroup.Get("/cpu", r.metricsHandler.GetCPUMetrics)
	metricsGroup.Get("/memory", r.metricsHandler.GetMemoryMetrics)
	metricsGroup.Get("/disk", r.metricsHandler.GetDiskMetrics)
	metricsGroup.Get("/network", r.metricsHandler.GetNetworkMetrics)

	// Database-dependent routes
	if r.playerHandler != nil {
		// Player routes
		players := api.Group("/players")
		players.Get("/", r.playerHandler.ListPlayers)
		players.Get("/stats", r.playerHandler.GetPlayerStats)
		players.Get("/search/:username", r.playerHandler.SearchPlayer)
		players.Get("/:uuid", r.playerHandler.GetPlayer)
		players.Post("/", r.playerHandler.CreateOrUpdatePlayer)
		players.Patch("/:uuid/online", r.playerHandler.UpdateOnlineStatus)
		players.Post("/:uuid/stats", r.playerHandler.UpdateStats)
		players.Post("/:uuid/ip", r.playerHandler.AddIPRecord)
	}

	if r.punishmentHandler != nil {
		// Punishment routes
		punishments := api.Group("/punishments")
		punishments.Get("/", r.punishmentHandler.ListPunishments)
		punishments.Get("/stats", r.punishmentHandler.GetStats)
		punishments.Get("/templates", r.punishmentHandler.GetTemplates)
		punishments.Post("/templates", r.punishmentHandler.CreateTemplate)
		punishments.Delete("/templates/:id", r.punishmentHandler.DeleteTemplate)
		punishments.Get("/player/:uuid", r.punishmentHandler.GetPlayerPunishments)
		punishments.Get("/check/:uuid", r.punishmentHandler.CheckBan)
		punishments.Post("/", r.punishmentHandler.CreatePunishment)
		punishments.Post("/:id/revoke", r.punishmentHandler.RevokePunishment)
		punishments.Post("/:id/appeal", r.punishmentHandler.AppealPunishment)
		punishments.Post("/:id/appeal/handle", r.punishmentHandler.HandleAppeal)
	}

	if r.discordHandler != nil {
		// Discord routes
		discord := api.Group("/discord")
		discord.Get("/links", r.discordHandler.ListLinks)
		discord.Get("/links/player/:uuid", r.discordHandler.GetPlayerLink)
		discord.Get("/links/discord/:discordId", r.discordHandler.GetLinkByDiscordID)
		discord.Post("/links", r.discordHandler.CreateLink)
		discord.Post("/links/:uuid/verify", r.discordHandler.VerifyLink)
		discord.Delete("/links/:uuid", r.discordHandler.DeleteLink)
		discord.Get("/settings", r.discordHandler.ListSettings)
		discord.Get("/settings/:guildId", r.discordHandler.GetSettings)
		discord.Put("/settings/:guildId", r.discordHandler.SaveSettings)
	}

	if r.analyticsHandler != nil {
		// Analytics routes
		analytics := api.Group("/analytics")
		analytics.Get("/dashboard", r.analyticsHandler.GetDashboardStats)
		analytics.Get("/players", r.analyticsHandler.GetPlayerHistory)
		analytics.Get("/server/:serverId", r.analyticsHandler.GetServerStats)
		analytics.Post("/server", r.analyticsHandler.RecordServerStat)
		analytics.Get("/logs", r.analyticsHandler.GetActivityLogs)
		analytics.Post("/logs", r.analyticsHandler.CreateActivityLog)

		// Notifications
		api.Get("/notifications", r.analyticsHandler.GetNotifications)
		api.Post("/notifications", r.analyticsHandler.CreateNotification)
		api.Post("/notifications/:id/read", r.analyticsHandler.MarkNotificationRead)
		api.Post("/notifications/read-all", r.analyticsHandler.MarkAllNotificationsRead)

		// Webhooks
		api.Get("/webhooks", r.analyticsHandler.GetWebhooks)
		api.Post("/webhooks", r.analyticsHandler.CreateWebhook)
		api.Put("/webhooks/:id", r.analyticsHandler.UpdateWebhook)
		api.Delete("/webhooks/:id", r.analyticsHandler.DeleteWebhook)
	}

	// Dedicated Server Management Routes (Pterodactyl-like)
	if r.nodeHandler != nil {
		// Locations
		locations := api.Group("/locations")
		locations.Get("/", r.nodeHandler.ListLocations)
		locations.Post("/", r.authMiddleware.RequireRole("admin"), r.nodeHandler.CreateLocation)
		locations.Delete("/:id", r.authMiddleware.RequireRole("admin"), r.nodeHandler.DeleteLocation)

		// Nodes
		nodes := api.Group("/nodes")
		nodes.Get("/", r.nodeHandler.ListNodes)
		nodes.Get("/:id", r.nodeHandler.GetNode)
		nodes.Post("/", r.authMiddleware.RequireRole("admin"), r.nodeHandler.CreateNode)
		nodes.Put("/:id", r.authMiddleware.RequireRole("admin"), r.nodeHandler.UpdateNode)
		nodes.Delete("/:id", r.authMiddleware.RequireRole("admin"), r.nodeHandler.DeleteNode)
		nodes.Post("/:id/regenerate-token", r.authMiddleware.RequireRole("admin"), r.nodeHandler.RegenerateToken)
		nodes.Get("/:id/allocations", r.nodeHandler.GetNodeAllocations)
		nodes.Post("/:id/allocations", r.authMiddleware.RequireRole("admin"), r.nodeHandler.CreateAllocation)
		nodes.Delete("/:id/allocations/:allocId", r.authMiddleware.RequireRole("admin"), r.nodeHandler.DeleteAllocation)
	}

	if r.dedicatedServerHandler != nil {
		// Dedicated Servers
		dedicatedServers := api.Group("/dedicated-servers")
		dedicatedServers.Get("/", r.dedicatedServerHandler.ListServers)
		dedicatedServers.Get("/stats", r.dedicatedServerHandler.GetServerStats)
		dedicatedServers.Get("/:id", r.dedicatedServerHandler.GetServer)
		dedicatedServers.Get("/uuid/:uuid", r.dedicatedServerHandler.GetServerByUUID)
		dedicatedServers.Post("/", r.authMiddleware.RequireRole("admin"), r.dedicatedServerHandler.CreateServer)
		dedicatedServers.Put("/:id", r.authMiddleware.RequireRole("admin"), r.dedicatedServerHandler.UpdateServer)
		dedicatedServers.Delete("/:id", r.authMiddleware.RequireRole("admin"), r.dedicatedServerHandler.DeleteServer)

		// Power actions
		dedicatedServers.Post("/:id/power/start", r.dedicatedServerHandler.StartServer)
		dedicatedServers.Post("/:id/power/stop", r.dedicatedServerHandler.StopServer)
		dedicatedServers.Post("/:id/power/restart", r.dedicatedServerHandler.RestartServer)
		dedicatedServers.Post("/:id/power/kill", r.dedicatedServerHandler.KillServer)

		// Suspend/Unsuspend
		dedicatedServers.Post("/:id/suspend", r.authMiddleware.RequireRole("admin"), r.dedicatedServerHandler.SuspendServer)
		dedicatedServers.Post("/:id/unsuspend", r.authMiddleware.RequireRole("admin"), r.dedicatedServerHandler.UnsuspendServer)

		// Commands
		dedicatedServers.Post("/:id/command", r.dedicatedServerHandler.SendCommand)

		// Eggs & Nests
		api.Get("/nests", r.dedicatedServerHandler.ListNests)
		api.Get("/eggs", r.dedicatedServerHandler.ListEggs)
		api.Get("/eggs/:id", r.dedicatedServerHandler.GetEgg)
	}
}

func (r *Router) Listen(addr string) error {
	return r.app.Listen(addr)
}

func (r *Router) Shutdown() error {
	return r.app.Shutdown()
}
