package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/withsait/karapixel/karapanel/internal/api"
	"github.com/withsait/karapixel/karapanel/internal/config"
	"github.com/withsait/karapixel/karapanel/internal/metrics"
	"github.com/withsait/karapixel/karapanel/internal/server"
)

var (
	configPath = flag.String("config", "configs/config.yml", "Path to config file")
	version    = "0.1.0"
)

func main() {
	flag.Parse()

	fmt.Printf(`
╔═══════════════════════════════════════════╗
║           KaraPanel Daemon v%s          ║
║      Minecraft Server Control Panel       ║
╚═══════════════════════════════════════════╝
`, version)

	// Load config
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config loaded from %s", *configPath)

	// Initialize server manager
	manager, err := server.NewManager(cfg.Servers)
	if err != nil {
		log.Fatalf("Failed to create server manager: %v", err)
	}
	defer manager.Close()
	log.Printf("Server manager initialized with %d servers", len(cfg.Servers))

	// Initialize console manager
	consoleManager := server.NewConsoleManager()
	log.Println("Console manager initialized")

	// Initialize metrics collector
	collector := metrics.NewCollector()
	log.Println("Metrics collector initialized")

	// Create API router
	router := api.NewRouter(cfg, manager, consoleManager, collector)

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting HTTP server on %s", addr)

	// Handle graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("Shutting down...")
		if err := router.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}()

	// Start listening
	if err := router.Listen(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
