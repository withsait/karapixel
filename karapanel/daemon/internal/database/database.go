package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func Connect(cfg *DBConfig) error {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to PostgreSQL database")

	if err := runMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func runMigrations() error {
	migrations := []string{
		// Players table
		`CREATE TABLE IF NOT EXISTS players (
			uuid VARCHAR(36) PRIMARY KEY,
			username VARCHAR(16) NOT NULL,
			first_join TIMESTAMP NOT NULL DEFAULT NOW(),
			last_seen TIMESTAMP NOT NULL DEFAULT NOW(),
			total_playtime BIGINT DEFAULT 0,
			is_online BOOLEAN DEFAULT FALSE,
			last_ip VARCHAR(45),
			metadata JSONB DEFAULT '{}'::jsonb,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Player stats table
		`CREATE TABLE IF NOT EXISTS player_stats (
			id SERIAL PRIMARY KEY,
			player_uuid VARCHAR(36) REFERENCES players(uuid) ON DELETE CASCADE,
			kills INT DEFAULT 0,
			deaths INT DEFAULT 0,
			blocks_broken BIGINT DEFAULT 0,
			blocks_placed BIGINT DEFAULT 0,
			distance_walked BIGINT DEFAULT 0,
			mobs_killed INT DEFAULT 0,
			play_sessions INT DEFAULT 0,
			updated_at TIMESTAMP DEFAULT NOW(),
			UNIQUE(player_uuid)
		)`,

		// Player IP history
		`CREATE TABLE IF NOT EXISTS player_ips (
			id SERIAL PRIMARY KEY,
			player_uuid VARCHAR(36) REFERENCES players(uuid) ON DELETE CASCADE,
			ip_address VARCHAR(45) NOT NULL,
			first_used TIMESTAMP DEFAULT NOW(),
			last_used TIMESTAMP DEFAULT NOW(),
			times_used INT DEFAULT 1,
			UNIQUE(player_uuid, ip_address)
		)`,

		// Punishments table
		`CREATE TABLE IF NOT EXISTS punishments (
			id SERIAL PRIMARY KEY,
			player_uuid VARCHAR(36) REFERENCES players(uuid) ON DELETE CASCADE,
			type VARCHAR(20) NOT NULL,
			reason TEXT NOT NULL,
			moderator_uuid VARCHAR(36),
			moderator_name VARCHAR(16),
			duration BIGINT,
			expires_at TIMESTAMP,
			server VARCHAR(50),
			is_active BOOLEAN DEFAULT TRUE,
			is_appealed BOOLEAN DEFAULT FALSE,
			appeal_reason TEXT,
			appeal_status VARCHAR(20),
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Punishment templates
		`CREATE TABLE IF NOT EXISTS punishment_templates (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			type VARCHAR(20) NOT NULL,
			reason TEXT NOT NULL,
			duration BIGINT,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Discord links
		`CREATE TABLE IF NOT EXISTS discord_links (
			id SERIAL PRIMARY KEY,
			player_uuid VARCHAR(36) REFERENCES players(uuid) ON DELETE CASCADE,
			discord_id VARCHAR(20) NOT NULL UNIQUE,
			discord_username VARCHAR(100),
			is_verified BOOLEAN DEFAULT FALSE,
			linked_at TIMESTAMP DEFAULT NOW(),
			UNIQUE(player_uuid)
		)`,

		// Discord bot settings
		`CREATE TABLE IF NOT EXISTS discord_settings (
			guild_id VARCHAR(20) PRIMARY KEY,
			prefix VARCHAR(10) DEFAULT '!',
			welcome_channel VARCHAR(20),
			welcome_message TEXT,
			log_channel VARCHAR(20),
			mod_log_channel VARCHAR(20),
			ticket_category VARCHAR(20),
			auto_role VARCHAR(20),
			settings JSONB DEFAULT '{}'::jsonb,
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Activity logs
		`CREATE TABLE IF NOT EXISTS activity_logs (
			id SERIAL PRIMARY KEY,
			actor_type VARCHAR(20) NOT NULL,
			actor_id VARCHAR(36) NOT NULL,
			actor_name VARCHAR(100),
			action VARCHAR(50) NOT NULL,
			target_type VARCHAR(20),
			target_id VARCHAR(36),
			target_name VARCHAR(100),
			details JSONB DEFAULT '{}'::jsonb,
			ip_address VARCHAR(45),
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Server statistics (for analytics)
		`CREATE TABLE IF NOT EXISTS server_stats (
			id SERIAL PRIMARY KEY,
			server_id VARCHAR(50) NOT NULL,
			player_count INT DEFAULT 0,
			tps DECIMAL(5,2),
			memory_used BIGINT,
			memory_max BIGINT,
			recorded_at TIMESTAMP DEFAULT NOW()
		)`,

		// Notifications
		`CREATE TABLE IF NOT EXISTS notifications (
			id SERIAL PRIMARY KEY,
			type VARCHAR(50) NOT NULL,
			title VARCHAR(200) NOT NULL,
			message TEXT,
			severity VARCHAR(20) DEFAULT 'info',
			is_read BOOLEAN DEFAULT FALSE,
			target_user VARCHAR(50),
			metadata JSONB DEFAULT '{}'::jsonb,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Webhook configurations
		`CREATE TABLE IF NOT EXISTS webhooks (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			url TEXT NOT NULL,
			type VARCHAR(50) NOT NULL,
			events TEXT[] DEFAULT '{}',
			is_active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Create indexes
		`CREATE INDEX IF NOT EXISTS idx_players_username ON players(username)`,
		`CREATE INDEX IF NOT EXISTS idx_players_last_seen ON players(last_seen)`,
		`CREATE INDEX IF NOT EXISTS idx_punishments_player ON punishments(player_uuid)`,
		`CREATE INDEX IF NOT EXISTS idx_punishments_active ON punishments(is_active)`,
		`CREATE INDEX IF NOT EXISTS idx_punishments_type ON punishments(type)`,
		`CREATE INDEX IF NOT EXISTS idx_activity_logs_actor ON activity_logs(actor_id)`,
		`CREATE INDEX IF NOT EXISTS idx_activity_logs_created ON activity_logs(created_at)`,
		`CREATE INDEX IF NOT EXISTS idx_server_stats_recorded ON server_stats(recorded_at)`,

		// ==========================================
		// Dedicated Server Management Tables
		// ==========================================

		// Locations (Datacenters)
		`CREATE TABLE IF NOT EXISTS locations (
			id SERIAL PRIMARY KEY,
			short VARCHAR(20) NOT NULL UNIQUE,
			long VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Nodes (Host machines)
		`CREATE TABLE IF NOT EXISTS nodes (
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL UNIQUE,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			fqdn VARCHAR(255) NOT NULL,
			scheme VARCHAR(10) DEFAULT 'https',
			daemon_port INT DEFAULT 8080,
			daemon_token VARCHAR(255),
			memory BIGINT NOT NULL,
			memory_overalloc INT DEFAULT 0,
			disk BIGINT NOT NULL,
			disk_overalloc INT DEFAULT 0,
			upload_limit INT DEFAULT 0,
			download_limit INT DEFAULT 0,
			status VARCHAR(20) DEFAULT 'offline',
			maintenance_mode BOOLEAN DEFAULT FALSE,
			location_id INT REFERENCES locations(id) ON DELETE SET NULL,
			metadata JSONB DEFAULT '{}'::jsonb,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Allocations (IP:Port assignments)
		`CREATE TABLE IF NOT EXISTS allocations (
			id SERIAL PRIMARY KEY,
			node_id INT NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
			ip VARCHAR(45) NOT NULL,
			alias VARCHAR(100),
			port INT NOT NULL,
			server_id INT,
			notes TEXT,
			UNIQUE(node_id, ip, port)
		)`,

		// Nests (Game categories)
		`CREATE TABLE IF NOT EXISTS nests (
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL UNIQUE,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			author VARCHAR(100) DEFAULT 'KaraPanel',
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Eggs (Server templates)
		`CREATE TABLE IF NOT EXISTS eggs (
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL UNIQUE,
			nest_id INT NOT NULL REFERENCES nests(id) ON DELETE CASCADE,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			author VARCHAR(100) DEFAULT 'KaraPanel',
			docker_images JSONB DEFAULT '[]'::jsonb,
			default_image VARCHAR(255),
			startup TEXT NOT NULL,
			stop_command VARCHAR(100) DEFAULT 'stop',
			config_files JSONB DEFAULT '{}'::jsonb,
			config_logs JSONB DEFAULT '{}'::jsonb,
			variables JSONB DEFAULT '[]'::jsonb,
			install_script JSONB DEFAULT '{}'::jsonb,
			metadata JSONB DEFAULT '{}'::jsonb,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Servers (Game server instances)
		`CREATE TABLE IF NOT EXISTS servers (
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL UNIQUE,
			short_id VARCHAR(8) NOT NULL UNIQUE,
			name VARCHAR(100) NOT NULL,
			description TEXT,
			node_id INT NOT NULL REFERENCES nodes(id) ON DELETE RESTRICT,
			owner_id INT,
			egg_id INT REFERENCES eggs(id) ON DELETE SET NULL,
			status VARCHAR(20) DEFAULT 'offline',
			suspended BOOLEAN DEFAULT FALSE,
			memory BIGINT NOT NULL,
			disk BIGINT NOT NULL,
			cpu INT DEFAULT 100,
			io INT DEFAULT 500,
			swap BIGINT DEFAULT 0,
			threads VARCHAR(50),
			oom_disabled BOOLEAN DEFAULT FALSE,
			startup_command TEXT,
			default_allocation_id INT,
			image VARCHAR(255),
			backup_limit INT DEFAULT 0,
			database_limit INT DEFAULT 0,
			allocation_limit INT DEFAULT 1,
			installed BOOLEAN DEFAULT FALSE,
			environment JSONB DEFAULT '{}'::jsonb,
			metadata JSONB DEFAULT '{}'::jsonb,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Add foreign key for allocations -> servers
		`DO $$ BEGIN
			ALTER TABLE allocations ADD CONSTRAINT fk_allocations_server
				FOREIGN KEY (server_id) REFERENCES servers(id) ON DELETE SET NULL;
		EXCEPTION
			WHEN duplicate_object THEN NULL;
		END $$`,

		// Add foreign key for servers -> allocations
		`DO $$ BEGIN
			ALTER TABLE servers ADD CONSTRAINT fk_servers_allocation
				FOREIGN KEY (default_allocation_id) REFERENCES allocations(id) ON DELETE SET NULL;
		EXCEPTION
			WHEN duplicate_object THEN NULL;
		END $$`,

		// Server variables (environment variables)
		`CREATE TABLE IF NOT EXISTS server_variables (
			id SERIAL PRIMARY KEY,
			server_id INT NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
			egg_variable_id INT,
			name VARCHAR(100) NOT NULL,
			value TEXT,
			UNIQUE(server_id, name)
		)`,

		// Server backups
		`CREATE TABLE IF NOT EXISTS server_backups (
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL UNIQUE,
			server_id INT NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
			name VARCHAR(100) NOT NULL,
			ignored_files TEXT,
			disk_usage BIGINT DEFAULT 0,
			checksum VARCHAR(64),
			is_successful BOOLEAN DEFAULT FALSE,
			is_locked BOOLEAN DEFAULT FALSE,
			completed_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT NOW()
		)`,

		// Server schedules (automated tasks)
		`CREATE TABLE IF NOT EXISTS server_schedules (
			id SERIAL PRIMARY KEY,
			server_id INT NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
			name VARCHAR(100) NOT NULL,
			cron_minute VARCHAR(10) NOT NULL,
			cron_hour VARCHAR(10) NOT NULL,
			cron_day_of_month VARCHAR(10) NOT NULL,
			cron_month VARCHAR(10) NOT NULL,
			cron_day_of_week VARCHAR(10) NOT NULL,
			is_active BOOLEAN DEFAULT TRUE,
			is_processing BOOLEAN DEFAULT FALSE,
			only_when_online BOOLEAN DEFAULT FALSE,
			last_run_at TIMESTAMP,
			next_run_at TIMESTAMP,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Schedule tasks
		`CREATE TABLE IF NOT EXISTS schedule_tasks (
			id SERIAL PRIMARY KEY,
			schedule_id INT NOT NULL REFERENCES server_schedules(id) ON DELETE CASCADE,
			sequence_id INT NOT NULL,
			action VARCHAR(20) NOT NULL,
			payload TEXT,
			time_offset INT DEFAULT 0,
			continue_on_failure BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)`,

		// Server resource usage history
		`CREATE TABLE IF NOT EXISTS server_resource_usage (
			id SERIAL PRIMARY KEY,
			server_id INT NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
			cpu_percent DECIMAL(5,2),
			memory_used BIGINT,
			memory_percent DECIMAL(5,2),
			disk_used BIGINT,
			disk_percent DECIMAL(5,2),
			network_rx BIGINT,
			network_tx BIGINT,
			recorded_at TIMESTAMP DEFAULT NOW()
		)`,

		// Dedicated server indexes
		`CREATE INDEX IF NOT EXISTS idx_nodes_status ON nodes(status)`,
		`CREATE INDEX IF NOT EXISTS idx_nodes_location ON nodes(location_id)`,
		`CREATE INDEX IF NOT EXISTS idx_allocations_node ON allocations(node_id)`,
		`CREATE INDEX IF NOT EXISTS idx_allocations_server ON allocations(server_id)`,
		`CREATE INDEX IF NOT EXISTS idx_servers_node ON servers(node_id)`,
		`CREATE INDEX IF NOT EXISTS idx_servers_status ON servers(status)`,
		`CREATE INDEX IF NOT EXISTS idx_servers_egg ON servers(egg_id)`,
		`CREATE INDEX IF NOT EXISTS idx_server_backups_server ON server_backups(server_id)`,
		`CREATE INDEX IF NOT EXISTS idx_server_schedules_server ON server_schedules(server_id)`,
		`CREATE INDEX IF NOT EXISTS idx_server_resource_usage_server ON server_resource_usage(server_id)`,
		`CREATE INDEX IF NOT EXISTS idx_server_resource_usage_recorded ON server_resource_usage(recorded_at)`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("migration failed: %w\nQuery: %s", err, migration[:min(100, len(migration))])
		}
	}

	log.Println("Database migrations completed")
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
