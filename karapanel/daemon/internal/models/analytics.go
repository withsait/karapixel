package models

import (
	"database/sql"
	"time"
)

type ServerStat struct {
	ID          int       `json:"id"`
	ServerID    string    `json:"serverId"`
	PlayerCount int       `json:"playerCount"`
	TPS         float64   `json:"tps"`
	MemoryUsed  int64     `json:"memoryUsed"`
	MemoryMax   int64     `json:"memoryMax"`
	RecordedAt  time.Time `json:"recordedAt"`
}

type ActivityLog struct {
	ID         int             `json:"id"`
	ActorType  string          `json:"actorType"` // user, player, system
	ActorID    string          `json:"actorId"`
	ActorName  string          `json:"actorName"`
	Action     string          `json:"action"`
	TargetType string          `json:"targetType,omitempty"`
	TargetID   string          `json:"targetId,omitempty"`
	TargetName string          `json:"targetName,omitempty"`
	Details    string          `json:"details,omitempty"`
	IPAddress  string          `json:"ipAddress,omitempty"`
	CreatedAt  time.Time       `json:"createdAt"`
}

type Notification struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Message    string    `json:"message"`
	Severity   string    `json:"severity"` // info, warning, error, success
	IsRead     bool      `json:"isRead"`
	TargetUser string    `json:"targetUser,omitempty"`
	Metadata   string    `json:"metadata,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Webhook struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Type      string    `json:"type"` // discord, slack, custom
	Events    []string  `json:"events"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

type AnalyticsRepository struct {
	db *sql.DB
}

func NewAnalyticsRepository(db *sql.DB) *AnalyticsRepository {
	return &AnalyticsRepository{db: db}
}

// Server Stats
func (r *AnalyticsRepository) RecordServerStat(stat *ServerStat) error {
	_, err := r.db.Exec(`
		INSERT INTO server_stats (server_id, player_count, tps, memory_used, memory_max)
		VALUES ($1, $2, $3, $4, $5)
	`, stat.ServerID, stat.PlayerCount, stat.TPS, stat.MemoryUsed, stat.MemoryMax)
	return err
}

func (r *AnalyticsRepository) GetServerStats(serverID string, hours int) ([]ServerStat, error) {
	var stats []ServerStat

	rows, err := r.db.Query(`
		SELECT id, server_id, player_count, tps, memory_used, memory_max, recorded_at
		FROM server_stats
		WHERE server_id = $1 AND recorded_at > NOW() - INTERVAL '1 hour' * $2
		ORDER BY recorded_at ASC
	`, serverID, hours)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s ServerStat
		var tps sql.NullFloat64
		var memUsed, memMax sql.NullInt64

		err := rows.Scan(&s.ID, &s.ServerID, &s.PlayerCount, &tps, &memUsed, &memMax, &s.RecordedAt)
		if err != nil {
			return nil, err
		}

		if tps.Valid {
			s.TPS = tps.Float64
		}
		if memUsed.Valid {
			s.MemoryUsed = memUsed.Int64
		}
		if memMax.Valid {
			s.MemoryMax = memMax.Int64
		}
		stats = append(stats, s)
	}

	return stats, nil
}

func (r *AnalyticsRepository) GetPlayerCountHistory(hours int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	rows, err := r.db.Query(`
		SELECT date_trunc('hour', recorded_at) as hour,
		       MAX(player_count) as peak_players,
		       AVG(player_count)::int as avg_players
		FROM server_stats
		WHERE recorded_at > NOW() - INTERVAL '1 hour' * $1
		GROUP BY hour
		ORDER BY hour ASC
	`, hours)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hour time.Time
		var peak, avg int

		if err := rows.Scan(&hour, &peak, &avg); err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"hour":        hour,
			"peakPlayers": peak,
			"avgPlayers":  avg,
		})
	}

	return results, nil
}

// Activity Logs
func (r *AnalyticsRepository) LogActivity(log *ActivityLog) error {
	_, err := r.db.Exec(`
		INSERT INTO activity_logs (actor_type, actor_id, actor_name, action, target_type, target_id, target_name, details, ip_address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, log.ActorType, log.ActorID, log.ActorName, log.Action, log.TargetType, log.TargetID, log.TargetName, log.Details, log.IPAddress)
	return err
}

func (r *AnalyticsRepository) GetActivityLogs(limit, offset int, actorType, action string) ([]ActivityLog, int, error) {
	var logs []ActivityLog
	var total int

	countQuery := `SELECT COUNT(*) FROM activity_logs WHERE 1=1`
	query := `SELECT id, actor_type, actor_id, actor_name, action, target_type, target_id, target_name, details, ip_address, created_at
			  FROM activity_logs WHERE 1=1`

	args := []interface{}{}
	argCount := 0

	if actorType != "" {
		argCount++
		countQuery += " AND actor_type = $" + itoa(argCount)
		query += " AND actor_type = $" + itoa(argCount)
		args = append(args, actorType)
	}

	if action != "" {
		argCount++
		countQuery += " AND action = $" + itoa(argCount)
		query += " AND action = $" + itoa(argCount)
		args = append(args, action)
	}

	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	query += " ORDER BY created_at DESC"
	argCount++
	query += " LIMIT $" + itoa(argCount)
	args = append(args, limit)
	argCount++
	query += " OFFSET $" + itoa(argCount)
	args = append(args, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var log ActivityLog
		var targetType, targetID, targetName, details, ipAddress sql.NullString

		err := rows.Scan(&log.ID, &log.ActorType, &log.ActorID, &log.ActorName, &log.Action,
			&targetType, &targetID, &targetName, &details, &ipAddress, &log.CreatedAt)
		if err != nil {
			return nil, 0, err
		}

		if targetType.Valid {
			log.TargetType = targetType.String
		}
		if targetID.Valid {
			log.TargetID = targetID.String
		}
		if targetName.Valid {
			log.TargetName = targetName.String
		}
		if details.Valid {
			log.Details = details.String
		}
		if ipAddress.Valid {
			log.IPAddress = ipAddress.String
		}
		logs = append(logs, log)
	}

	return logs, total, nil
}

// Notifications
func (r *AnalyticsRepository) CreateNotification(n *Notification) error {
	return r.db.QueryRow(`
		INSERT INTO notifications (type, title, message, severity, target_user, metadata)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`, n.Type, n.Title, n.Message, n.Severity, n.TargetUser, n.Metadata).Scan(&n.ID, &n.CreatedAt)
}

func (r *AnalyticsRepository) GetNotifications(targetUser string, unreadOnly bool, limit int) ([]Notification, error) {
	var notifications []Notification

	query := `SELECT id, type, title, message, severity, is_read, target_user, metadata, created_at
			  FROM notifications WHERE (target_user = $1 OR target_user IS NULL)`

	if unreadOnly {
		query += ` AND is_read = false`
	}
	query += ` ORDER BY created_at DESC LIMIT $2`

	rows, err := r.db.Query(query, targetUser, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n Notification
		var targetUser, metadata sql.NullString

		err := rows.Scan(&n.ID, &n.Type, &n.Title, &n.Message, &n.Severity, &n.IsRead, &targetUser, &metadata, &n.CreatedAt)
		if err != nil {
			return nil, err
		}

		if targetUser.Valid {
			n.TargetUser = targetUser.String
		}
		if metadata.Valid {
			n.Metadata = metadata.String
		}
		notifications = append(notifications, n)
	}

	return notifications, nil
}

func (r *AnalyticsRepository) MarkNotificationRead(id int) error {
	_, err := r.db.Exec(`UPDATE notifications SET is_read = true WHERE id = $1`, id)
	return err
}

func (r *AnalyticsRepository) MarkAllNotificationsRead(targetUser string) error {
	_, err := r.db.Exec(`UPDATE notifications SET is_read = true WHERE target_user = $1 OR target_user IS NULL`, targetUser)
	return err
}

// Webhooks
func (r *AnalyticsRepository) GetWebhooks() ([]Webhook, error) {
	var webhooks []Webhook

	rows, err := r.db.Query(`SELECT id, name, url, type, events, is_active, created_at FROM webhooks ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var w Webhook
		var events []string

		err := rows.Scan(&w.ID, &w.Name, &w.URL, &w.Type, &events, &w.IsActive, &w.CreatedAt)
		if err != nil {
			return nil, err
		}

		w.Events = events
		webhooks = append(webhooks, w)
	}

	return webhooks, nil
}

func (r *AnalyticsRepository) CreateWebhook(w *Webhook) error {
	return r.db.QueryRow(`
		INSERT INTO webhooks (name, url, type, events, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`, w.Name, w.URL, w.Type, w.Events, w.IsActive).Scan(&w.ID, &w.CreatedAt)
}

func (r *AnalyticsRepository) UpdateWebhook(w *Webhook) error {
	_, err := r.db.Exec(`
		UPDATE webhooks SET name = $2, url = $3, type = $4, events = $5, is_active = $6 WHERE id = $1
	`, w.ID, w.Name, w.URL, w.Type, w.Events, w.IsActive)
	return err
}

func (r *AnalyticsRepository) DeleteWebhook(id int) error {
	_, err := r.db.Exec(`DELETE FROM webhooks WHERE id = $1`, id)
	return err
}

func (r *AnalyticsRepository) GetActiveWebhooksForEvent(event string) ([]Webhook, error) {
	var webhooks []Webhook

	rows, err := r.db.Query(`SELECT id, name, url, type, events, is_active, created_at FROM webhooks WHERE is_active = true AND $1 = ANY(events)`, event)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var w Webhook
		var events []string

		err := rows.Scan(&w.ID, &w.Name, &w.URL, &w.Type, &events, &w.IsActive, &w.CreatedAt)
		if err != nil {
			return nil, err
		}

		w.Events = events
		webhooks = append(webhooks, w)
	}

	return webhooks, nil
}

// Dashboard Stats
func (r *AnalyticsRepository) GetDashboardStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Player stats
	var totalPlayers, onlinePlayers, newPlayersToday int
	r.db.QueryRow(`SELECT COUNT(*) FROM players`).Scan(&totalPlayers)
	r.db.QueryRow(`SELECT COUNT(*) FROM players WHERE is_online = true`).Scan(&onlinePlayers)
	r.db.QueryRow(`SELECT COUNT(*) FROM players WHERE first_join >= CURRENT_DATE`).Scan(&newPlayersToday)

	stats["totalPlayers"] = totalPlayers
	stats["onlinePlayers"] = onlinePlayers
	stats["newPlayersToday"] = newPlayersToday

	// Punishment stats
	var activeBans, activeMutes int
	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'ban' AND is_active = true`).Scan(&activeBans)
	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'mute' AND is_active = true`).Scan(&activeMutes)

	stats["activeBans"] = activeBans
	stats["activeMutes"] = activeMutes

	// Discord stats
	var linkedAccounts int
	r.db.QueryRow(`SELECT COUNT(*) FROM discord_links WHERE is_verified = true`).Scan(&linkedAccounts)
	stats["linkedDiscordAccounts"] = linkedAccounts

	return stats, nil
}
