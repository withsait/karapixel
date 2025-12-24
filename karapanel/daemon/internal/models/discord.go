package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type DiscordLink struct {
	ID              int       `json:"id"`
	PlayerUUID      string    `json:"playerUuid"`
	DiscordID       string    `json:"discordId"`
	DiscordUsername string    `json:"discordUsername"`
	IsVerified      bool      `json:"isVerified"`
	LinkedAt        time.Time `json:"linkedAt"`
}

type DiscordSettings struct {
	GuildID         string          `json:"guildId"`
	Prefix          string          `json:"prefix"`
	WelcomeChannel  string          `json:"welcomeChannel,omitempty"`
	WelcomeMessage  string          `json:"welcomeMessage,omitempty"`
	LogChannel      string          `json:"logChannel,omitempty"`
	ModLogChannel   string          `json:"modLogChannel,omitempty"`
	TicketCategory  string          `json:"ticketCategory,omitempty"`
	AutoRole        string          `json:"autoRole,omitempty"`
	Settings        json.RawMessage `json:"settings,omitempty"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}

type DiscordRepository struct {
	db *sql.DB
}

func NewDiscordRepository(db *sql.DB) *DiscordRepository {
	return &DiscordRepository{db: db}
}

// Link operations
func (r *DiscordRepository) GetLink(playerUUID string) (*DiscordLink, error) {
	var link DiscordLink
	var discordUsername sql.NullString

	err := r.db.QueryRow(`
		SELECT id, player_uuid, discord_id, discord_username, is_verified, linked_at
		FROM discord_links WHERE player_uuid = $1
	`, playerUUID).Scan(&link.ID, &link.PlayerUUID, &link.DiscordID, &discordUsername, &link.IsVerified, &link.LinkedAt)

	if err != nil {
		return nil, err
	}

	if discordUsername.Valid {
		link.DiscordUsername = discordUsername.String
	}

	return &link, nil
}

func (r *DiscordRepository) GetLinkByDiscordID(discordID string) (*DiscordLink, error) {
	var link DiscordLink
	var discordUsername sql.NullString

	err := r.db.QueryRow(`
		SELECT id, player_uuid, discord_id, discord_username, is_verified, linked_at
		FROM discord_links WHERE discord_id = $1
	`, discordID).Scan(&link.ID, &link.PlayerUUID, &link.DiscordID, &discordUsername, &link.IsVerified, &link.LinkedAt)

	if err != nil {
		return nil, err
	}

	if discordUsername.Valid {
		link.DiscordUsername = discordUsername.String
	}

	return &link, nil
}

func (r *DiscordRepository) CreateLink(playerUUID, discordID, discordUsername string) (*DiscordLink, error) {
	var link DiscordLink

	err := r.db.QueryRow(`
		INSERT INTO discord_links (player_uuid, discord_id, discord_username)
		VALUES ($1, $2, $3)
		RETURNING id, player_uuid, discord_id, discord_username, is_verified, linked_at
	`, playerUUID, discordID, discordUsername).Scan(&link.ID, &link.PlayerUUID, &link.DiscordID, &link.DiscordUsername, &link.IsVerified, &link.LinkedAt)

	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *DiscordRepository) VerifyLink(playerUUID string) error {
	_, err := r.db.Exec(`UPDATE discord_links SET is_verified = true WHERE player_uuid = $1`, playerUUID)
	return err
}

func (r *DiscordRepository) DeleteLink(playerUUID string) error {
	_, err := r.db.Exec(`DELETE FROM discord_links WHERE player_uuid = $1`, playerUUID)
	return err
}

func (r *DiscordRepository) GetAllLinks(limit, offset int) ([]DiscordLink, int, error) {
	var links []DiscordLink
	var total int

	if err := r.db.QueryRow(`SELECT COUNT(*) FROM discord_links`).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(`
		SELECT id, player_uuid, discord_id, discord_username, is_verified, linked_at
		FROM discord_links ORDER BY linked_at DESC LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var link DiscordLink
		var discordUsername sql.NullString

		err := rows.Scan(&link.ID, &link.PlayerUUID, &link.DiscordID, &discordUsername, &link.IsVerified, &link.LinkedAt)
		if err != nil {
			return nil, 0, err
		}

		if discordUsername.Valid {
			link.DiscordUsername = discordUsername.String
		}
		links = append(links, link)
	}

	return links, total, nil
}

// Settings operations
func (r *DiscordRepository) GetSettings(guildID string) (*DiscordSettings, error) {
	var settings DiscordSettings
	var welcomeChannel, welcomeMessage, logChannel, modLogChannel, ticketCategory, autoRole sql.NullString
	var settingsJSON sql.NullString

	err := r.db.QueryRow(`
		SELECT guild_id, prefix, welcome_channel, welcome_message, log_channel, mod_log_channel,
		       ticket_category, auto_role, settings, updated_at
		FROM discord_settings WHERE guild_id = $1
	`, guildID).Scan(&settings.GuildID, &settings.Prefix, &welcomeChannel, &welcomeMessage,
		&logChannel, &modLogChannel, &ticketCategory, &autoRole, &settingsJSON, &settings.UpdatedAt)

	if err != nil {
		return nil, err
	}

	if welcomeChannel.Valid {
		settings.WelcomeChannel = welcomeChannel.String
	}
	if welcomeMessage.Valid {
		settings.WelcomeMessage = welcomeMessage.String
	}
	if logChannel.Valid {
		settings.LogChannel = logChannel.String
	}
	if modLogChannel.Valid {
		settings.ModLogChannel = modLogChannel.String
	}
	if ticketCategory.Valid {
		settings.TicketCategory = ticketCategory.String
	}
	if autoRole.Valid {
		settings.AutoRole = autoRole.String
	}
	if settingsJSON.Valid {
		settings.Settings = json.RawMessage(settingsJSON.String)
	}

	return &settings, nil
}

func (r *DiscordRepository) SaveSettings(s *DiscordSettings) error {
	_, err := r.db.Exec(`
		INSERT INTO discord_settings (guild_id, prefix, welcome_channel, welcome_message, log_channel,
			mod_log_channel, ticket_category, auto_role, settings)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (guild_id) DO UPDATE SET
			prefix = EXCLUDED.prefix,
			welcome_channel = EXCLUDED.welcome_channel,
			welcome_message = EXCLUDED.welcome_message,
			log_channel = EXCLUDED.log_channel,
			mod_log_channel = EXCLUDED.mod_log_channel,
			ticket_category = EXCLUDED.ticket_category,
			auto_role = EXCLUDED.auto_role,
			settings = EXCLUDED.settings,
			updated_at = NOW()
	`, s.GuildID, s.Prefix, s.WelcomeChannel, s.WelcomeMessage, s.LogChannel,
		s.ModLogChannel, s.TicketCategory, s.AutoRole, s.Settings)

	return err
}

func (r *DiscordRepository) GetAllSettings() ([]DiscordSettings, error) {
	var settingsList []DiscordSettings

	rows, err := r.db.Query(`
		SELECT guild_id, prefix, welcome_channel, welcome_message, log_channel, mod_log_channel,
		       ticket_category, auto_role, settings, updated_at
		FROM discord_settings ORDER BY updated_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var settings DiscordSettings
		var welcomeChannel, welcomeMessage, logChannel, modLogChannel, ticketCategory, autoRole sql.NullString
		var settingsJSON sql.NullString

		err := rows.Scan(&settings.GuildID, &settings.Prefix, &welcomeChannel, &welcomeMessage,
			&logChannel, &modLogChannel, &ticketCategory, &autoRole, &settingsJSON, &settings.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if welcomeChannel.Valid {
			settings.WelcomeChannel = welcomeChannel.String
		}
		if welcomeMessage.Valid {
			settings.WelcomeMessage = welcomeMessage.String
		}
		if logChannel.Valid {
			settings.LogChannel = logChannel.String
		}
		if modLogChannel.Valid {
			settings.ModLogChannel = modLogChannel.String
		}
		if ticketCategory.Valid {
			settings.TicketCategory = ticketCategory.String
		}
		if autoRole.Valid {
			settings.AutoRole = autoRole.String
		}
		if settingsJSON.Valid {
			settings.Settings = json.RawMessage(settingsJSON.String)
		}

		settingsList = append(settingsList, settings)
	}

	return settingsList, nil
}
