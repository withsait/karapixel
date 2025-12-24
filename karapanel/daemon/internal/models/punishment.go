package models

import (
	"database/sql"
	"time"
)

type Punishment struct {
	ID            int        `json:"id"`
	PlayerUUID    string     `json:"playerUuid"`
	PlayerName    string     `json:"playerName,omitempty"`
	Type          string     `json:"type"` // ban, kick, mute, warn
	Reason        string     `json:"reason"`
	ModeratorUUID string     `json:"moderatorUuid,omitempty"`
	ModeratorName string     `json:"moderatorName,omitempty"`
	Duration      *int64     `json:"duration,omitempty"` // in seconds, null for permanent
	ExpiresAt     *time.Time `json:"expiresAt,omitempty"`
	Server        string     `json:"server,omitempty"`
	IsActive      bool       `json:"isActive"`
	IsAppealed    bool       `json:"isAppealed"`
	AppealReason  string     `json:"appealReason,omitempty"`
	AppealStatus  string     `json:"appealStatus,omitempty"` // pending, approved, rejected
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}

type PunishmentTemplate struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Duration  *int64 `json:"duration,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreatePunishmentRequest struct {
	PlayerUUID    string `json:"playerUuid"`
	PlayerName    string `json:"playerName"`
	Type          string `json:"type"`
	Reason        string `json:"reason"`
	ModeratorUUID string `json:"moderatorUuid"`
	ModeratorName string `json:"moderatorName"`
	Duration      *int64 `json:"duration,omitempty"`
	Server        string `json:"server,omitempty"`
}

type PunishmentRepository struct {
	db *sql.DB
}

func NewPunishmentRepository(db *sql.DB) *PunishmentRepository {
	return &PunishmentRepository{db: db}
}

func (r *PunishmentRepository) GetAll(limit, offset int, punishmentType, status string) ([]Punishment, int, error) {
	var punishments []Punishment
	var total int

	countQuery := `SELECT COUNT(*) FROM punishments p
				   LEFT JOIN players pl ON p.player_uuid = pl.uuid
				   WHERE 1=1`
	query := `SELECT p.id, p.player_uuid, COALESCE(pl.username, ''), p.type, p.reason,
			  p.moderator_uuid, p.moderator_name, p.duration, p.expires_at, p.server,
			  p.is_active, p.is_appealed, p.appeal_reason, p.appeal_status, p.created_at, p.updated_at
			  FROM punishments p
			  LEFT JOIN players pl ON p.player_uuid = pl.uuid
			  WHERE 1=1`

	args := []interface{}{}
	argCount := 0

	if punishmentType != "" && punishmentType != "all" {
		argCount++
		countQuery += " AND p.type = $" + itoa(argCount)
		query += " AND p.type = $" + itoa(argCount)
		args = append(args, punishmentType)
	}

	if status == "active" {
		countQuery += " AND p.is_active = true"
		query += " AND p.is_active = true"
	} else if status == "expired" {
		countQuery += " AND p.is_active = false"
		query += " AND p.is_active = false"
	} else if status == "appealed" {
		countQuery += " AND p.is_appealed = true"
		query += " AND p.is_appealed = true"
	}

	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	query += " ORDER BY p.created_at DESC"
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
		var p Punishment
		var moderatorUUID, moderatorName, server, appealReason, appealStatus sql.NullString
		var duration sql.NullInt64
		var expiresAt sql.NullTime

		err := rows.Scan(&p.ID, &p.PlayerUUID, &p.PlayerName, &p.Type, &p.Reason,
			&moderatorUUID, &moderatorName, &duration, &expiresAt, &server,
			&p.IsActive, &p.IsAppealed, &appealReason, &appealStatus, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}

		if moderatorUUID.Valid {
			p.ModeratorUUID = moderatorUUID.String
		}
		if moderatorName.Valid {
			p.ModeratorName = moderatorName.String
		}
		if duration.Valid {
			p.Duration = &duration.Int64
		}
		if expiresAt.Valid {
			p.ExpiresAt = &expiresAt.Time
		}
		if server.Valid {
			p.Server = server.String
		}
		if appealReason.Valid {
			p.AppealReason = appealReason.String
		}
		if appealStatus.Valid {
			p.AppealStatus = appealStatus.String
		}

		punishments = append(punishments, p)
	}

	return punishments, total, nil
}

func (r *PunishmentRepository) GetByPlayer(uuid string) ([]Punishment, error) {
	var punishments []Punishment

	rows, err := r.db.Query(`
		SELECT id, player_uuid, type, reason, moderator_uuid, moderator_name,
		       duration, expires_at, server, is_active, is_appealed, appeal_reason, appeal_status, created_at, updated_at
		FROM punishments WHERE player_uuid = $1 ORDER BY created_at DESC
	`, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Punishment
		var moderatorUUID, moderatorName, server, appealReason, appealStatus sql.NullString
		var duration sql.NullInt64
		var expiresAt sql.NullTime

		err := rows.Scan(&p.ID, &p.PlayerUUID, &p.Type, &p.Reason,
			&moderatorUUID, &moderatorName, &duration, &expiresAt, &server,
			&p.IsActive, &p.IsAppealed, &appealReason, &appealStatus, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if moderatorUUID.Valid {
			p.ModeratorUUID = moderatorUUID.String
		}
		if moderatorName.Valid {
			p.ModeratorName = moderatorName.String
		}
		if duration.Valid {
			p.Duration = &duration.Int64
		}
		if expiresAt.Valid {
			p.ExpiresAt = &expiresAt.Time
		}
		if server.Valid {
			p.Server = server.String
		}
		if appealReason.Valid {
			p.AppealReason = appealReason.String
		}
		if appealStatus.Valid {
			p.AppealStatus = appealStatus.String
		}

		punishments = append(punishments, p)
	}

	return punishments, nil
}

func (r *PunishmentRepository) Create(req *CreatePunishmentRequest) (*Punishment, error) {
	var p Punishment
	var expiresAt *time.Time

	if req.Duration != nil && *req.Duration > 0 {
		t := time.Now().Add(time.Duration(*req.Duration) * time.Second)
		expiresAt = &t
	}

	err := r.db.QueryRow(`
		INSERT INTO punishments (player_uuid, type, reason, moderator_uuid, moderator_name, duration, expires_at, server, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, true)
		RETURNING id, player_uuid, type, reason, moderator_uuid, moderator_name, duration, expires_at, server, is_active, is_appealed, created_at, updated_at
	`, req.PlayerUUID, req.Type, req.Reason, req.ModeratorUUID, req.ModeratorName, req.Duration, expiresAt, req.Server).Scan(
		&p.ID, &p.PlayerUUID, &p.Type, &p.Reason, &p.ModeratorUUID, &p.ModeratorName,
		&p.Duration, &p.ExpiresAt, &p.Server, &p.IsActive, &p.IsAppealed, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return nil, err
	}

	p.PlayerName = req.PlayerName
	return &p, nil
}

func (r *PunishmentRepository) Revoke(id int, moderatorName string) error {
	_, err := r.db.Exec(`
		UPDATE punishments SET is_active = false, updated_at = NOW() WHERE id = $1
	`, id)
	return err
}

func (r *PunishmentRepository) Appeal(id int, reason string) error {
	_, err := r.db.Exec(`
		UPDATE punishments SET is_appealed = true, appeal_reason = $2, appeal_status = 'pending', updated_at = NOW() WHERE id = $1
	`, id, reason)
	return err
}

func (r *PunishmentRepository) HandleAppeal(id int, approved bool) error {
	status := "rejected"
	if approved {
		status = "approved"
	}

	query := `UPDATE punishments SET appeal_status = $2, updated_at = NOW()`
	if approved {
		query += `, is_active = false`
	}
	query += ` WHERE id = $1`

	_, err := r.db.Exec(query, id, status)
	return err
}

func (r *PunishmentRepository) GetActiveBan(uuid string) (*Punishment, error) {
	var p Punishment
	var moderatorUUID, moderatorName, server sql.NullString
	var duration sql.NullInt64
	var expiresAt sql.NullTime

	err := r.db.QueryRow(`
		SELECT id, player_uuid, type, reason, moderator_uuid, moderator_name, duration, expires_at, server, is_active, created_at
		FROM punishments
		WHERE player_uuid = $1 AND type = 'ban' AND is_active = true
		AND (expires_at IS NULL OR expires_at > NOW())
		ORDER BY created_at DESC LIMIT 1
	`, uuid).Scan(&p.ID, &p.PlayerUUID, &p.Type, &p.Reason, &moderatorUUID, &moderatorName,
		&duration, &expiresAt, &server, &p.IsActive, &p.CreatedAt)

	if err != nil {
		return nil, err
	}

	if moderatorUUID.Valid {
		p.ModeratorUUID = moderatorUUID.String
	}
	if moderatorName.Valid {
		p.ModeratorName = moderatorName.String
	}
	if duration.Valid {
		p.Duration = &duration.Int64
	}
	if expiresAt.Valid {
		p.ExpiresAt = &expiresAt.Time
	}
	if server.Valid {
		p.Server = server.String
	}

	return &p, nil
}

// Templates
func (r *PunishmentRepository) GetTemplates() ([]PunishmentTemplate, error) {
	var templates []PunishmentTemplate

	rows, err := r.db.Query(`SELECT id, name, type, reason, duration, created_at FROM punishment_templates ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t PunishmentTemplate
		var duration sql.NullInt64

		err := rows.Scan(&t.ID, &t.Name, &t.Type, &t.Reason, &duration, &t.CreatedAt)
		if err != nil {
			return nil, err
		}

		if duration.Valid {
			t.Duration = &duration.Int64
		}
		templates = append(templates, t)
	}

	return templates, nil
}

func (r *PunishmentRepository) CreateTemplate(t *PunishmentTemplate) error {
	return r.db.QueryRow(`
		INSERT INTO punishment_templates (name, type, reason, duration)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`, t.Name, t.Type, t.Reason, t.Duration).Scan(&t.ID, &t.CreatedAt)
}

func (r *PunishmentRepository) DeleteTemplate(id int) error {
	_, err := r.db.Exec(`DELETE FROM punishment_templates WHERE id = $1`, id)
	return err
}

// Stats
func (r *PunishmentRepository) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var totalBans, activeBans, totalMutes, activeMutes, totalWarns, totalKicks int

	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'ban'`).Scan(&totalBans)
	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'ban' AND is_active = true`).Scan(&activeBans)
	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'mute'`).Scan(&totalMutes)
	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'mute' AND is_active = true`).Scan(&activeMutes)
	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'warn'`).Scan(&totalWarns)
	r.db.QueryRow(`SELECT COUNT(*) FROM punishments WHERE type = 'kick'`).Scan(&totalKicks)

	stats["totalBans"] = totalBans
	stats["activeBans"] = activeBans
	stats["totalMutes"] = totalMutes
	stats["activeMutes"] = activeMutes
	stats["totalWarns"] = totalWarns
	stats["totalKicks"] = totalKicks

	return stats, nil
}

func itoa(i int) string {
	return string(rune('0' + i))
}
