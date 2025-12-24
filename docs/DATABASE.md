# 🗄️ KaraPixel - Veritabanı Şeması

> MySQL 8.0 veritabanı yapısı ve tablo tanımları.

---

## Genel Bakış

```
┌─────────────────────────────────────────────────────────────────┐
│                    DATABASE MİMARİSİ                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  DATABASE: karapixel_db                                        │
│  CHARSET : utf8mb4                                             │
│  COLLATE : utf8mb4_unicode_ci                                  │
│  ENGINE  : InnoDB                                              │
│                                                                 │
│  TABLOLAR:                                                      │
│  ├── players          → Oyuncu ana verileri                    │
│  ├── auth             → Authentication bilgileri               │
│  ├── sessions         → Aktif oturumlar                        │
│  ├── economy          → Para bakiyeleri                        │
│  ├── transactions     → Para transferleri                      │
│  ├── islands          → Ada bilgileri                          │
│  ├── island_members   → Ada üyeleri                            │
│  ├── island_settings  → Ada ayarları                           │
│  ├── generators       → Generator verileri                     │
│  ├── skills           → Skill ilerlemeleri                     │
│  ├── quests           → Quest ilerlemeleri                     │
│  ├── cosmetics        → Sahip olunan cosmetics                │
│  ├── pets             → Pet verileri                           │
│  ├── ranks            → Oyuncu rank'ları                       │
│  ├── punishments      → Ban/mute kayıtları                     │
│  ├── statistics       → İstatistikler                          │
│  └── security_logs    → Güvenlik logları                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Tablo Şemaları

### players

```sql
CREATE TABLE players (
    uuid VARCHAR(36) PRIMARY KEY,
    username VARCHAR(16) NOT NULL,
    username_lower VARCHAR(16) NOT NULL,
    
    -- Platform bilgisi
    platform ENUM('JAVA', 'BEDROCK') NOT NULL DEFAULT 'JAVA',
    first_join TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_join TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_server VARCHAR(32),
    
    -- Oyuncu ayarları
    locale VARCHAR(10) NOT NULL DEFAULT 'tr_TR',
    
    -- İstatistik
    play_time INT UNSIGNED NOT NULL DEFAULT 0,  -- Dakika cinsinden
    
    -- Metadata
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    -- Indexler
    INDEX idx_username (username_lower),
    INDEX idx_last_join (last_join)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### auth

```sql
CREATE TABLE auth (
    uuid VARCHAR(36) PRIMARY KEY,
    
    -- Şifre (bcrypt hash)
    password_hash VARCHAR(60),
    
    -- Kayıt bilgileri
    registered_at TIMESTAMP,
    registered_ip VARCHAR(45),
    
    -- Son giriş
    last_login_at TIMESTAMP,
    last_login_ip VARCHAR(45),
    
    -- Güvenlik
    failed_attempts INT UNSIGNED NOT NULL DEFAULT 0,
    locked_until TIMESTAMP NULL,
    
    -- 2FA (opsiyonel)
    totp_secret VARCHAR(32),
    totp_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Bedrock
    is_bedrock BOOLEAN NOT NULL DEFAULT FALSE,
    xbox_xuid VARCHAR(32),
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### sessions

```sql
CREATE TABLE sessions (
    id VARCHAR(64) PRIMARY KEY,
    uuid VARCHAR(36) NOT NULL,
    
    -- Session bilgileri
    ip VARCHAR(45) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    
    -- Status
    is_valid BOOLEAN NOT NULL DEFAULT TRUE,
    
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE,
    INDEX idx_uuid (uuid),
    INDEX idx_expires (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### economy

```sql
CREATE TABLE economy (
    uuid VARCHAR(36) PRIMARY KEY,
    balance DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    lifetime_earned DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    lifetime_spent DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### transactions

```sql
CREATE TABLE transactions (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    
    -- İşlem tarafları
    from_uuid VARCHAR(36),  -- NULL = sistem
    to_uuid VARCHAR(36),    -- NULL = sistem
    
    -- İşlem detayları
    amount DECIMAL(20, 2) NOT NULL,
    type ENUM('TRANSFER', 'SHOP_BUY', 'SHOP_SELL', 'REWARD', 'ADMIN', 'ISLAND_BANK', 'OTHER') NOT NULL,
    description VARCHAR(255),
    
    -- Metadata
    server VARCHAR(32),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (from_uuid) REFERENCES players(uuid) ON DELETE SET NULL,
    FOREIGN KEY (to_uuid) REFERENCES players(uuid) ON DELETE SET NULL,
    INDEX idx_from (from_uuid),
    INDEX idx_to (to_uuid),
    INDEX idx_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### islands

```sql
CREATE TABLE islands (
    id VARCHAR(36) PRIMARY KEY,
    
    -- Sahip
    owner_uuid VARCHAR(36) NOT NULL,
    
    -- Temel bilgiler
    name VARCHAR(32),
    template VARCHAR(32) NOT NULL DEFAULT 'normal',
    
    -- Lokasyon
    world_server VARCHAR(32) NOT NULL,  -- skyblock-1, skyblock-2, etc.
    center_x INT NOT NULL,
    center_z INT NOT NULL,
    
    -- Seviye ve ilerleme
    level INT UNSIGNED NOT NULL DEFAULT 1,
    experience BIGINT UNSIGNED NOT NULL DEFAULT 0,
    
    -- Ada bankası
    bank_balance DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    
    -- Boyut
    size INT UNSIGNED NOT NULL DEFAULT 100,  -- Çap
    
    -- Status
    is_public BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (owner_uuid) REFERENCES players(uuid),
    INDEX idx_owner (owner_uuid),
    INDEX idx_server (world_server)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### island_members

```sql
CREATE TABLE island_members (
    island_id VARCHAR(36) NOT NULL,
    uuid VARCHAR(36) NOT NULL,
    
    -- Rol
    role ENUM('OWNER', 'ADMIN', 'MEMBER', 'VISITOR') NOT NULL DEFAULT 'MEMBER',
    
    -- Katılım
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    invited_by VARCHAR(36),
    
    PRIMARY KEY (island_id, uuid),
    FOREIGN KEY (island_id) REFERENCES islands(id) ON DELETE CASCADE,
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE,
    INDEX idx_uuid (uuid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### island_settings

```sql
CREATE TABLE island_settings (
    island_id VARCHAR(36) PRIMARY KEY,
    
    -- PvP
    pvp_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Ziyaretçi izinleri
    visitor_enter BOOLEAN NOT NULL DEFAULT TRUE,
    visitor_interact BOOLEAN NOT NULL DEFAULT FALSE,
    visitor_pickup BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Diğer
    mob_spawning BOOLEAN NOT NULL DEFAULT TRUE,
    animal_spawning BOOLEAN NOT NULL DEFAULT TRUE,
    
    -- Warp
    warp_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    warp_x DOUBLE,
    warp_y DOUBLE,
    warp_z DOUBLE,
    warp_yaw FLOAT,
    warp_pitch FLOAT,
    
    FOREIGN KEY (island_id) REFERENCES islands(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### generators

```sql
CREATE TABLE generators (
    id VARCHAR(36) PRIMARY KEY,
    island_id VARCHAR(36) NOT NULL,
    
    -- Generator tipi ve seviyesi
    type VARCHAR(32) NOT NULL,  -- COBBLESTONE, IRON, GOLD, DIAMOND, EMERALD
    tier INT UNSIGNED NOT NULL DEFAULT 1,
    
    -- Lokasyon
    world VARCHAR(64) NOT NULL,
    x INT NOT NULL,
    y INT NOT NULL,
    z INT NOT NULL,
    
    -- İstatistik
    total_generated BIGINT UNSIGNED NOT NULL DEFAULT 0,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (island_id) REFERENCES islands(id) ON DELETE CASCADE,
    INDEX idx_island (island_id),
    UNIQUE KEY uk_location (world, x, y, z)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### skills

```sql
CREATE TABLE skills (
    uuid VARCHAR(36) NOT NULL,
    skill_type VARCHAR(32) NOT NULL,  -- MINING, FARMING, COMBAT, FISHING, FORAGING, ENCHANTING
    
    level INT UNSIGNED NOT NULL DEFAULT 1,
    experience BIGINT UNSIGNED NOT NULL DEFAULT 0,
    
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (uuid, skill_type),
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### quests

```sql
CREATE TABLE quests (
    uuid VARCHAR(36) NOT NULL,
    quest_id VARCHAR(64) NOT NULL,
    
    -- İlerleme
    progress INT UNSIGNED NOT NULL DEFAULT 0,
    target INT UNSIGNED NOT NULL,
    
    -- Status
    status ENUM('ACTIVE', 'COMPLETED', 'CLAIMED') NOT NULL DEFAULT 'ACTIVE',
    
    -- Timestamps
    started_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    
    PRIMARY KEY (uuid, quest_id),
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE,
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### cosmetics

```sql
CREATE TABLE cosmetics (
    uuid VARCHAR(36) NOT NULL,
    cosmetic_id VARCHAR(64) NOT NULL,
    
    -- Tip
    type ENUM('PARTICLE', 'WING', 'HAT', 'TRAIL', 'KILL_EFFECT') NOT NULL,
    
    -- Status
    owned BOOLEAN NOT NULL DEFAULT TRUE,
    equipped BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Satın alma
    purchased_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (uuid, cosmetic_id),
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE,
    INDEX idx_type (type),
    INDEX idx_equipped (uuid, equipped)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### pets

```sql
CREATE TABLE pets (
    id VARCHAR(36) PRIMARY KEY,
    uuid VARCHAR(36) NOT NULL,
    
    -- Pet bilgileri
    type VARCHAR(32) NOT NULL,
    name VARCHAR(32),
    
    -- Seviye
    level INT UNSIGNED NOT NULL DEFAULT 1,
    experience BIGINT UNSIGNED NOT NULL DEFAULT 0,
    
    -- Status
    active BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Timestamps
    obtained_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE,
    INDEX idx_uuid (uuid),
    INDEX idx_active (uuid, active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### ranks

```sql
CREATE TABLE ranks (
    uuid VARCHAR(36) NOT NULL,
    rank_id VARCHAR(32) NOT NULL,  -- VIP, VIP_PLUS, MVP, MVP_PLUS
    
    -- Süre
    obtained_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,  -- NULL = kalıcı
    
    -- Kaynak
    source ENUM('PURCHASE', 'GIFT', 'ADMIN', 'EVENT') NOT NULL,
    transaction_id VARCHAR(64),  -- Ödeme referansı
    
    PRIMARY KEY (uuid, rank_id),
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE,
    INDEX idx_expires (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### punishments

```sql
CREATE TABLE punishments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    uuid VARCHAR(36) NOT NULL,
    
    -- Ceza tipi
    type ENUM('BAN', 'TEMP_BAN', 'MUTE', 'TEMP_MUTE', 'KICK', 'WARN') NOT NULL,
    
    -- Detaylar
    reason VARCHAR(255) NOT NULL,
    
    -- Süre
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,  -- NULL = kalıcı
    
    -- Yetkili
    staff_uuid VARCHAR(36),
    staff_name VARCHAR(16),
    
    -- Status
    active BOOLEAN NOT NULL DEFAULT TRUE,
    pardoned BOOLEAN NOT NULL DEFAULT FALSE,
    pardoned_by VARCHAR(36),
    pardoned_at TIMESTAMP,
    pardon_reason VARCHAR(255),
    
    -- IP (ban için)
    ip VARCHAR(45),
    
    FOREIGN KEY (uuid) REFERENCES players(uuid),
    INDEX idx_uuid (uuid),
    INDEX idx_active (uuid, active),
    INDEX idx_type (type),
    INDEX idx_ip (ip)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### statistics

```sql
CREATE TABLE statistics (
    uuid VARCHAR(36) NOT NULL,
    stat_key VARCHAR(64) NOT NULL,
    
    value BIGINT NOT NULL DEFAULT 0,
    
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (uuid, stat_key),
    FOREIGN KEY (uuid) REFERENCES players(uuid) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Örnek stat_key'ler:
-- blocks_mined, blocks_placed, mobs_killed, deaths
-- islands_created, quests_completed, skills_leveled
-- chat_messages, commands_used
```

### security_logs

```sql
CREATE TABLE security_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    
    -- Oyuncu
    uuid VARCHAR(36),
    username VARCHAR(16),
    ip VARCHAR(45),
    
    -- Olay
    event_type VARCHAR(32) NOT NULL,  -- LOGIN_SUCCESS, LOGIN_FAIL, EXPLOIT_ATTEMPT, etc.
    severity ENUM('INFO', 'WARNING', 'CRITICAL') NOT NULL DEFAULT 'INFO',
    
    -- Detay
    description VARCHAR(500),
    server VARCHAR(32),
    
    -- Timestamp
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_uuid (uuid),
    INDEX idx_event (event_type),
    INDEX idx_severity (severity),
    INDEX idx_created (created_at),
    INDEX idx_ip (ip)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

---

## Redis Yapısı

```
┌─────────────────────────────────────────────────────────────────┐
│                    REDIS KEY YAPISI                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  SESSIONS                                                       │
│  session:{uuid}                                                │
│  └── Hash: {id, ip, created_at, expires_at}                    │
│  TTL: 7 gün                                                    │
│                                                                 │
│  ONLINE STATUS                                                  │
│  online:{uuid}                                                  │
│  └── String: server_name                                       │
│  TTL: Yok (oyuncu çıkınca silinir)                            │
│                                                                 │
│  ISLAND CACHE                                                   │
│  island:{island_id}                                            │
│  └── Hash: {id, owner, level, server, ...}                     │
│  TTL: 30 dakika                                                │
│                                                                 │
│  PLAYER CACHE                                                   │
│  player:{uuid}                                                  │
│  └── Hash: {username, locale, rank, balance, ...}              │
│  TTL: 10 dakika                                                │
│                                                                 │
│  RATE LIMITING                                                  │
│  ratelimit:{type}:{identifier}                                 │
│  └── String: count                                             │
│  TTL: Değişken (1dk - 1saat)                                  │
│                                                                 │
│  PUB/SUB CHANNELS                                              │
│  karapixel:player                                              │
│  karapixel:server                                              │
│  karapixel:teleport                                            │
│  karapixel:chat                                                │
│  karapixel:economy                                             │
│  karapixel:admin                                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Query Örnekleri

### Oyuncu Bilgisi Getir

```sql
SELECT 
    p.*,
    e.balance,
    i.id as island_id,
    i.level as island_level
FROM players p
LEFT JOIN economy e ON p.uuid = e.uuid
LEFT JOIN islands i ON p.uuid = i.owner_uuid
WHERE p.uuid = ?;
```

### Ada Üyelerini Getir

```sql
SELECT 
    p.uuid,
    p.username,
    im.role,
    im.joined_at
FROM island_members im
JOIN players p ON im.uuid = p.uuid
WHERE im.island_id = ?
ORDER BY 
    FIELD(im.role, 'OWNER', 'ADMIN', 'MEMBER', 'VISITOR'),
    im.joined_at;
```

### Leaderboard (Ada Seviyesi)

```sql
SELECT 
    i.id,
    i.name,
    p.username as owner_name,
    i.level,
    i.experience
FROM islands i
JOIN players p ON i.owner_uuid = p.uuid
ORDER BY i.level DESC, i.experience DESC
LIMIT 100;
```

### Aktif Cezalar

```sql
SELECT * FROM punishments
WHERE uuid = ?
AND active = TRUE
AND (expires_at IS NULL OR expires_at > NOW())
ORDER BY created_at DESC;
```

### Para Transfer

```sql
START TRANSACTION;

UPDATE economy SET balance = balance - ? WHERE uuid = ? AND balance >= ?;
UPDATE economy SET balance = balance + ? WHERE uuid = ?;

INSERT INTO transactions (from_uuid, to_uuid, amount, type, description)
VALUES (?, ?, ?, 'TRANSFER', ?);

COMMIT;
```

---

## Index Stratejisi

```sql
-- Sık kullanılan sorgular için ek indexler

-- Oyuncu arama
CREATE INDEX idx_players_search ON players(username_lower);

-- Aktif oturumlar
CREATE INDEX idx_sessions_valid ON sessions(uuid, is_valid, expires_at);

-- Ada araması
CREATE INDEX idx_islands_search ON islands(name);

-- Leaderboard
CREATE INDEX idx_islands_leaderboard ON islands(level DESC, experience DESC);
CREATE INDEX idx_skills_leaderboard ON skills(skill_type, level DESC, experience DESC);

-- Transaction history
CREATE INDEX idx_transactions_history ON transactions(from_uuid, created_at DESC);
CREATE INDEX idx_transactions_history2 ON transactions(to_uuid, created_at DESC);
```

---

## Bakım

### Günlük

```sql
-- Süresi dolmuş session'ları temizle
DELETE FROM sessions WHERE expires_at < NOW();

-- Süresi dolmuş cezaları güncelle
UPDATE punishments 
SET active = FALSE 
WHERE expires_at IS NOT NULL 
AND expires_at < NOW() 
AND active = TRUE;
```

### Haftalık

```sql
-- Tablo istatistiklerini güncelle
ANALYZE TABLE players, economy, islands, transactions;

-- Eski transaction log'larını arşivle (90 gün+)
INSERT INTO transactions_archive 
SELECT * FROM transactions WHERE created_at < DATE_SUB(NOW(), INTERVAL 90 DAY);

DELETE FROM transactions WHERE created_at < DATE_SUB(NOW(), INTERVAL 90 DAY);
```

### Aylık

```sql
-- Eski security log'larını temizle
DELETE FROM security_logs WHERE created_at < DATE_SUB(NOW(), INTERVAL 6 MONTH);

-- Table fragmentation kontrolü
SELECT 
    table_name,
    data_free / 1024 / 1024 AS free_mb
FROM information_schema.tables
WHERE table_schema = 'karapixel_db'
AND data_free > 100 * 1024 * 1024;

-- Gerekirse OPTIMIZE TABLE
```

---

*📅 Son güncelleme: 24 Aralık 2024*
