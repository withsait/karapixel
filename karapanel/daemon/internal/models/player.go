package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Player struct {
	UUID          string          `json:"uuid"`
	Username      string          `json:"username"`
	FirstJoin     time.Time       `json:"firstJoin"`
	LastSeen      time.Time       `json:"lastSeen"`
	TotalPlaytime int64           `json:"totalPlaytime"`
	IsOnline      bool            `json:"isOnline"`
	LastIP        string          `json:"lastIp,omitempty"`
	Metadata      json.RawMessage `json:"metadata,omitempty"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
}

type PlayerStats struct {
	ID            int       `json:"id"`
	PlayerUUID    string    `json:"playerUuid"`
	Kills         int       `json:"kills"`
	Deaths        int       `json:"deaths"`
	BlocksBroken  int64     `json:"blocksBroken"`
	BlocksPlaced  int64     `json:"blocksPlaced"`
	DistanceWalked int64    `json:"distanceWalked"`
	MobsKilled    int       `json:"mobsKilled"`
	PlaySessions  int       `json:"playSessions"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type PlayerIP struct {
	ID         int       `json:"id"`
	PlayerUUID string    `json:"playerUuid"`
	IPAddress  string    `json:"ipAddress"`
	FirstUsed  time.Time `json:"firstUsed"`
	LastUsed   time.Time `json:"lastUsed"`
	TimesUsed  int       `json:"timesUsed"`
}

type PlayerWithStats struct {
	Player
	Stats      *PlayerStats `json:"stats,omitempty"`
	IPHistory  []PlayerIP   `json:"ipHistory,omitempty"`
	Punishments []Punishment `json:"punishments,omitempty"`
}

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) GetAll(limit, offset int, search string, onlineOnly bool) ([]Player, int, error) {
	var players []Player
	var total int

	countQuery := "SELECT COUNT(*) FROM players WHERE 1=1"
	query := `SELECT uuid, username, first_join, last_seen, total_playtime, is_online, last_ip, metadata, created_at, updated_at
			  FROM players WHERE 1=1`

	args := []interface{}{}
	argCount := 0

	if search != "" {
		argCount++
		countQuery += " AND (username ILIKE $" + string(rune('0'+argCount)) + ")"
		query += " AND (username ILIKE $" + string(rune('0'+argCount)) + ")"
		args = append(args, "%"+search+"%")
	}

	if onlineOnly {
		countQuery += " AND is_online = true"
		query += " AND is_online = true"
	}

	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	query += " ORDER BY last_seen DESC"
	argCount++
	query += " LIMIT $" + string(rune('0'+argCount))
	args = append(args, limit)
	argCount++
	query += " OFFSET $" + string(rune('0'+argCount))
	args = append(args, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Player
		var lastIP sql.NullString
		var metadata sql.NullString

		err := rows.Scan(&p.UUID, &p.Username, &p.FirstJoin, &p.LastSeen, &p.TotalPlaytime,
			&p.IsOnline, &lastIP, &metadata, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}

		if lastIP.Valid {
			p.LastIP = lastIP.String
		}
		if metadata.Valid {
			p.Metadata = json.RawMessage(metadata.String)
		}

		players = append(players, p)
	}

	return players, total, nil
}

func (r *PlayerRepository) GetByUUID(uuid string) (*PlayerWithStats, error) {
	var p PlayerWithStats
	var lastIP sql.NullString
	var metadata sql.NullString

	err := r.db.QueryRow(`
		SELECT uuid, username, first_join, last_seen, total_playtime, is_online, last_ip, metadata, created_at, updated_at
		FROM players WHERE uuid = $1
	`, uuid).Scan(&p.UUID, &p.Username, &p.FirstJoin, &p.LastSeen, &p.TotalPlaytime,
		&p.IsOnline, &lastIP, &metadata, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return nil, err
	}

	if lastIP.Valid {
		p.LastIP = lastIP.String
	}
	if metadata.Valid {
		p.Metadata = json.RawMessage(metadata.String)
	}

	// Get stats
	var stats PlayerStats
	err = r.db.QueryRow(`
		SELECT id, player_uuid, kills, deaths, blocks_broken, blocks_placed, distance_walked, mobs_killed, play_sessions, updated_at
		FROM player_stats WHERE player_uuid = $1
	`, uuid).Scan(&stats.ID, &stats.PlayerUUID, &stats.Kills, &stats.Deaths, &stats.BlocksBroken,
		&stats.BlocksPlaced, &stats.DistanceWalked, &stats.MobsKilled, &stats.PlaySessions, &stats.UpdatedAt)

	if err == nil {
		p.Stats = &stats
	}

	// Get IP history
	ipRows, err := r.db.Query(`
		SELECT id, player_uuid, ip_address, first_used, last_used, times_used
		FROM player_ips WHERE player_uuid = $1 ORDER BY last_used DESC LIMIT 10
	`, uuid)
	if err == nil {
		defer ipRows.Close()
		for ipRows.Next() {
			var ip PlayerIP
			ipRows.Scan(&ip.ID, &ip.PlayerUUID, &ip.IPAddress, &ip.FirstUsed, &ip.LastUsed, &ip.TimesUsed)
			p.IPHistory = append(p.IPHistory, ip)
		}
	}

	return &p, nil
}

func (r *PlayerRepository) GetByUsername(username string) (*Player, error) {
	var p Player
	var lastIP sql.NullString
	var metadata sql.NullString

	err := r.db.QueryRow(`
		SELECT uuid, username, first_join, last_seen, total_playtime, is_online, last_ip, metadata, created_at, updated_at
		FROM players WHERE username ILIKE $1
	`, username).Scan(&p.UUID, &p.Username, &p.FirstJoin, &p.LastSeen, &p.TotalPlaytime,
		&p.IsOnline, &lastIP, &metadata, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return nil, err
	}

	if lastIP.Valid {
		p.LastIP = lastIP.String
	}
	if metadata.Valid {
		p.Metadata = json.RawMessage(metadata.String)
	}

	return &p, nil
}

func (r *PlayerRepository) Create(p *Player) error {
	_, err := r.db.Exec(`
		INSERT INTO players (uuid, username, first_join, last_seen, total_playtime, is_online, last_ip, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (uuid) DO UPDATE SET
			username = EXCLUDED.username,
			last_seen = EXCLUDED.last_seen,
			total_playtime = players.total_playtime + EXCLUDED.total_playtime,
			is_online = EXCLUDED.is_online,
			last_ip = EXCLUDED.last_ip,
			updated_at = NOW()
	`, p.UUID, p.Username, p.FirstJoin, p.LastSeen, p.TotalPlaytime, p.IsOnline, p.LastIP, p.Metadata)

	return err
}

func (r *PlayerRepository) UpdateOnlineStatus(uuid string, isOnline bool) error {
	_, err := r.db.Exec(`
		UPDATE players SET is_online = $2, last_seen = NOW(), updated_at = NOW() WHERE uuid = $1
	`, uuid, isOnline)
	return err
}

func (r *PlayerRepository) UpdateStats(stats *PlayerStats) error {
	_, err := r.db.Exec(`
		INSERT INTO player_stats (player_uuid, kills, deaths, blocks_broken, blocks_placed, distance_walked, mobs_killed, play_sessions)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (player_uuid) DO UPDATE SET
			kills = EXCLUDED.kills,
			deaths = EXCLUDED.deaths,
			blocks_broken = EXCLUDED.blocks_broken,
			blocks_placed = EXCLUDED.blocks_placed,
			distance_walked = EXCLUDED.distance_walked,
			mobs_killed = EXCLUDED.mobs_killed,
			play_sessions = EXCLUDED.play_sessions,
			updated_at = NOW()
	`, stats.PlayerUUID, stats.Kills, stats.Deaths, stats.BlocksBroken, stats.BlocksPlaced,
		stats.DistanceWalked, stats.MobsKilled, stats.PlaySessions)

	return err
}

func (r *PlayerRepository) AddIPRecord(uuid, ip string) error {
	_, err := r.db.Exec(`
		INSERT INTO player_ips (player_uuid, ip_address)
		VALUES ($1, $2)
		ON CONFLICT (player_uuid, ip_address) DO UPDATE SET
			last_used = NOW(),
			times_used = player_ips.times_used + 1
	`, uuid, ip)
	return err
}

func (r *PlayerRepository) GetOnlineCount() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM players WHERE is_online = true").Scan(&count)
	return count, err
}

func (r *PlayerRepository) GetTotalCount() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM players").Scan(&count)
	return count, err
}

func (r *PlayerRepository) GetNewPlayersToday() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM players WHERE first_join >= CURRENT_DATE").Scan(&count)
	return count, err
}
