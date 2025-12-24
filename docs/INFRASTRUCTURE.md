# 🏗️ KaraPixel - Altyapı Dokümantasyonu

> Sunucu kurulumu, yapılandırması ve yönetimi.

---

## Donanım Spesifikasyonları

### Ana Sunucu (Dedicated)

```
┌─────────────────────────────────────────────────────────────────┐
│                   HETZNER DEDICATED SERVER                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  MODEL        : Hetzner Auction (Ryzen 9 5950X)                │
│  CPU          : AMD Ryzen 9 5950X                              │
│                16 Core / 32 Thread                             │
│                3.4 GHz Base / 4.9 GHz Boost                    │
│  RAM          : 128 GB DDR4 ECC                                │
│  STORAGE      : 2x 3.84 TB NVMe SSD (RAID 1)                  │
│  NETWORK      : 1 Gbit/s (Dedicated)                          │
│  DDoS         : Hetzner DDoS Protection (Included)            │
│  LOCATION     : FSN (Falkenstein, Germany)                    │
│  OS           : Ubuntu 24.04 LTS                              │
│  COST         : ~€68.70/ay                                    │
│                                                                 │
│  IP           : (Sipariş sonrası atanacak)                    │
│  HOSTNAME     : kp-game-01.karapixel.net                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Hyble-Core VPS

```
┌─────────────────────────────────────────────────────────────────┐
│                    HETZNER CLOUD VPS                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  MODEL        : CX32                                           │
│  CPU          : 4 vCPU (AMD EPYC)                             │
│  RAM          : 8 GB                                           │
│  STORAGE      : 80 GB NVMe SSD                                │
│  NETWORK      : 20 TB Traffic                                  │
│  LOCATION     : FSN (Falkenstein, Germany)                    │
│  OS           : Ubuntu 24.04.3 LTS                            │
│  COST         : €15.59/ay                                     │
│                                                                 │
│  IP           : 159.69.42.124                                 │
│  HOSTNAME     : hyble-core                                     │
│  PURPOSE      : Hyble services (web, api, billing)            │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Storage Box (Backup)

```
┌─────────────────────────────────────────────────────────────────┐
│                   HETZNER STORAGE BOX                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  MODEL        : BX10                                           │
│  CAPACITY     : 1 TB                                           │
│  PROTOCOL     : SFTP, SCP, rsync, Samba, WebDAV               │
│  SNAPSHOTS    : 10 included                                    │
│  COST         : €3.81/ay                                      │
│                                                                 │
│  HOST         : uXXXXXX.your-storagebox.de                    │
│  PORT         : 23 (SSH/SFTP)                                 │
│  PURPOSE      : Off-site backup storage                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Dizin Yapısı

### Game Server Dizin Yapısı

```
/opt/karapixel/
├── servers/
│   ├── velocity/
│   │   ├── velocity.jar
│   │   ├── velocity.toml
│   │   ├── server-icon.png
│   │   ├── plugins/
│   │   │   ├── Geyser-Velocity.jar
│   │   │   ├── floodgate-velocity.jar
│   │   │   ├── karapixel-velocity.jar
│   │   │   └── LuckPerms-Velocity.jar
│   │   ├── logs/
│   │   └── forwarding.secret
│   │
│   ├── limbo/
│   │   ├── karapaper.jar
│   │   ├── config/
│   │   │   ├── karapaper.yml
│   │   │   ├── bukkit.yml
│   │   │   ├── spigot.yml
│   │   │   └── paper-global.yml
│   │   ├── plugins/
│   │   │   ├── karapixel-core.jar
│   │   │   ├── karapixel-auth.jar
│   │   │   └── karapixel-ui.jar
│   │   ├── world/
│   │   └── logs/
│   │
│   ├── hub/
│   │   ├── karapaper.jar
│   │   ├── config/
│   │   ├── plugins/
│   │   │   ├── karapixel-core.jar
│   │   │   ├── karapixel-hub.jar
│   │   │   ├── karapixel-selector.jar
│   │   │   ├── karapixel-cosmetics.jar
│   │   │   └── ...
│   │   ├── world/
│   │   └── logs/
│   │
│   ├── skyblock-spawn/           # 20GB RAM
│   │   ├── karapaper.jar
│   │   ├── config/
│   │   ├── plugins/
│   │   │   ├── karapixel-core.jar
│   │   │   ├── karapixel-skyblock.jar
│   │   │   ├── karapixel-economy.jar
│   │   │   ├── karapixel-shop.jar
│   │   │   ├── karapixel-crates.jar
│   │   │   ├── karapixel-events.jar      # Balık event vb.
│   │   │   ├── ModelEngine.jar
│   │   │   └── ...
│   │   ├── world/
│   │   └── logs/
│   │
│   ├── pvp-arena/                # 6GB RAM (YENİ)
│   │   ├── karapaper.jar
│   │   ├── config/
│   │   ├── plugins/
│   │   │   ├── karapixel-core.jar
│   │   │   ├── karapixel-pvp.jar
│   │   │   └── ...
│   │   ├── world/
│   │   └── logs/
│   │
│   ├── skyblock-1/               # 24GB RAM
│   │   ├── karapaper.jar
│   │   ├── config/
│   │   ├── plugins/
│   │   ├── world/                # Ada world'ü
│   │   └── logs/
│   │
│   ├── skyblock-2/               # 24GB RAM
│   │   └── ...
│   │
│   ├── skyblock-3/               # 16GB RAM (RESERVE)
│   │   └── ... (gerektiğinde aktif)
│   │
│   └── nether-end/               # 10GB RAM (YENİ - Paylaşımlı)
│       ├── karapaper.jar
│       ├── config/
│       ├── plugins/
│       ├── world_nether/
│       ├── world_the_end/
│       └── logs/
│
├── shared/
│   ├── plugins/           # Tüm sunucularda ortak pluginler
│   ├── configs/           # Ortak konfigürasyonlar
│   └── resourcepack/      # Resource pack
│
├── scripts/
│   ├── start-all.sh
│   ├── stop-all.sh
│   ├── restart-server.sh
│   ├── backup-database.sh
│   ├── backup-worlds.sh
│   ├── update-plugins.sh
│   └── health-check.sh
│
├── backups/
│   ├── database/
│   ├── worlds/
│   └── configs/
│
├── logs/
│   ├── access.log
│   ├── error.log
│   └── security.log
│
└── data/
    ├── mysql/             # MySQL data dir
    └── redis/             # Redis data dir
```

---

## Servis Kurulumu

### Java 21 Kurulumu

```bash
# Java 21 (OpenJDK)
sudo apt update
sudo apt install -y openjdk-21-jdk

# Verify
java -version

# JAVA_HOME ayarla
echo 'export JAVA_HOME=/usr/lib/jvm/java-21-openjdk-amd64' >> ~/.bashrc
source ~/.bashrc
```

### MySQL 8 Kurulumu

```bash
# MySQL 8 kurulumu
sudo apt install -y mysql-server

# Güvenlik ayarları
sudo mysql_secure_installation

# Database ve kullanıcı oluştur
sudo mysql -u root -p << EOF
CREATE DATABASE karapixel_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'karapixel'@'localhost' IDENTIFIED BY 'STRONG_PASSWORD_HERE';
GRANT ALL PRIVILEGES ON karapixel_db.* TO 'karapixel'@'localhost';
FLUSH PRIVILEGES;
EOF

# Performans ayarları
sudo nano /etc/mysql/mysql.conf.d/mysqld.cnf
```

```ini
# /etc/mysql/mysql.conf.d/mysqld.cnf
[mysqld]
# InnoDB ayarları
innodb_buffer_pool_size = 6G
innodb_log_file_size = 512M
innodb_flush_log_at_trx_commit = 2
innodb_flush_method = O_DIRECT

# Bağlantı limitleri
max_connections = 500
wait_timeout = 600
interactive_timeout = 600

# Query cache (MySQL 8'de deprecated)
# Performance schema
performance_schema = ON

# Logging
slow_query_log = 1
slow_query_log_file = /var/log/mysql/slow.log
long_query_time = 2
```

### Redis Kurulumu

```bash
# Redis kurulumu
sudo apt install -y redis-server

# Konfigürasyon
sudo nano /etc/redis/redis.conf
```

```ini
# /etc/redis/redis.conf
bind 127.0.0.1
port 6379

# Şifre
requirepass STRONG_REDIS_PASSWORD

# Memory
maxmemory 4gb
maxmemory-policy allkeys-lru

# Persistence
save 900 1
save 300 10
save 60 10000

# Security
rename-command FLUSHDB ""
rename-command FLUSHALL ""
rename-command DEBUG ""
```

```bash
# Redis başlat
sudo systemctl enable redis-server
sudo systemctl start redis-server
```

---

## Systemd Servisleri

### Velocity Proxy Servisi

```ini
# /etc/systemd/system/velocity.service
[Unit]
Description=KaraPixel Velocity Proxy
After=network.target

[Service]
User=minecraft
Group=minecraft
WorkingDirectory=/opt/karapixel/servers/velocity

ExecStart=/usr/bin/java \
    -Xms3G -Xmx3G \
    -XX:+UseG1GC \
    -XX:G1HeapRegionSize=4M \
    -XX:+UnlockExperimentalVMOptions \
    -XX:+ParallelRefProcEnabled \
    -XX:+AlwaysPreTouch \
    -Dlog4j2.formatMsgNoLookups=true \
    -jar velocity.jar

Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

### Game Server Servisi (Template)

```ini
# /etc/systemd/system/karapixel@.service
[Unit]
Description=KaraPixel %i Server
After=network.target mysql.service redis.service

[Service]
User=minecraft
Group=minecraft
WorkingDirectory=/opt/karapixel/servers/%i

ExecStart=/usr/bin/java \
    -Xms${MEMORY} -Xmx${MEMORY} \
    -XX:+UseG1GC \
    -XX:+ParallelRefProcEnabled \
    -XX:MaxGCPauseMillis=200 \
    -XX:+UnlockExperimentalVMOptions \
    -XX:+DisableExplicitGC \
    -XX:+AlwaysPreTouch \
    -XX:G1NewSizePercent=30 \
    -XX:G1MaxNewSizePercent=40 \
    -XX:G1HeapRegionSize=8M \
    -XX:G1ReservePercent=20 \
    -XX:G1HeapWastePercent=5 \
    -XX:G1MixedGCCountTarget=4 \
    -XX:InitiatingHeapOccupancyPercent=15 \
    -XX:G1MixedGCLiveThresholdPercent=90 \
    -XX:G1RSetUpdatingPauseTimePercent=5 \
    -XX:SurvivorRatio=32 \
    -XX:+PerfDisableSharedMem \
    -XX:MaxTenuringThreshold=1 \
    -Dlog4j2.formatMsgNoLookups=true \
    -Daikars.new.flags=true \
    -jar karapaper.jar --nogui

Restart=always
RestartSec=10

EnvironmentFile=/opt/karapixel/servers/%i/server.env

[Install]
WantedBy=multi-user.target
```

### Environment Files

```bash
# /opt/karapixel/servers/hub/server.env
MEMORY=6G

# /opt/karapixel/servers/skyblock-spawn/server.env
MEMORY=20G

# /opt/karapixel/servers/pvp-arena/server.env
MEMORY=6G

# /opt/karapixel/servers/skyblock-1/server.env
MEMORY=24G

# /opt/karapixel/servers/skyblock-2/server.env
MEMORY=24G

# /opt/karapixel/servers/nether-end/server.env
MEMORY=10G

# NOT: skyblock-3 reserve'de, gerektiğinde aktif edilecek
# /opt/karapixel/servers/skyblock-3/server.env (RESERVE)
# MEMORY=16G (reserve'den tahsis edilecek)
```

### Servis Yönetimi

```bash
# Servisleri etkinleştir
sudo systemctl enable velocity
sudo systemctl enable karapixel@hub
sudo systemctl enable karapixel@limbo
sudo systemctl enable karapixel@skyblock-spawn
sudo systemctl enable karapixel@pvp-arena
sudo systemctl enable karapixel@skyblock-1
sudo systemctl enable karapixel@skyblock-2
sudo systemctl enable karapixel@nether-end
# skyblock-3 reserve'de - gerektiğinde enable edilecek

# Başlat
sudo systemctl start velocity
sudo systemctl start karapixel@hub

# Durum kontrol
sudo systemctl status karapixel@hub

# Log izle
sudo journalctl -u karapixel@hub -f
```

---

## Başlatma Scriptleri

### start-all.sh

```bash
#!/bin/bash
# /opt/karapixel/scripts/start-all.sh

echo "Starting KaraPixel servers..."

# Sıralı başlatma (bağımlılık sırasına göre)
echo "Starting Velocity..."
sudo systemctl start velocity
sleep 5

echo "Starting Limbo..."
sudo systemctl start karapixel@limbo
sleep 5

echo "Starting Hub..."
sudo systemctl start karapixel@hub
sleep 5

echo "Starting Skyblock Spawn..."
sudo systemctl start karapixel@skyblock-spawn
sleep 5

echo "Starting PvP Arena..."
sudo systemctl start karapixel@pvp-arena
sleep 3

echo "Starting Skyblock Worlds..."
sudo systemctl start karapixel@skyblock-1
sudo systemctl start karapixel@skyblock-2
sleep 3

echo "Starting Nether/End..."
sudo systemctl start karapixel@nether-end

# NOT: skyblock-3 reserve'de, manuel başlatılacak
# sudo systemctl start karapixel@skyblock-3

echo "All servers started!"

# Durum kontrolü
sleep 10
./health-check.sh
```

### stop-all.sh

```bash
#!/bin/bash
# /opt/karapixel/scripts/stop-all.sh

echo "Stopping KaraPixel servers..."

# Ters sırada durdur
echo "Stopping Nether/End..."
sudo systemctl stop karapixel@nether-end

echo "Stopping Skyblock Worlds..."
sudo systemctl stop karapixel@skyblock-1
sudo systemctl stop karapixel@skyblock-2
# sudo systemctl stop karapixel@skyblock-3  # Reserve

echo "Stopping PvP Arena..."
sudo systemctl stop karapixel@pvp-arena

echo "Stopping Skyblock Spawn..."
sudo systemctl stop karapixel@skyblock-spawn

echo "Stopping Hub..."
sudo systemctl stop karapixel@hub

echo "Stopping Limbo..."
sudo systemctl stop karapixel@limbo

echo "Stopping Velocity..."
sudo systemctl stop velocity

echo "All servers stopped!"
```

### health-check.sh

```bash
#!/bin/bash
# /opt/karapixel/scripts/health-check.sh

echo "=== KaraPixel Health Check ==="
echo ""

check_service() {
    local name=$1
    local port=$2
    
    if systemctl is-active --quiet $name; then
        echo "✅ $name: Running"
        
        if [ ! -z "$port" ]; then
            if nc -z localhost $port 2>/dev/null; then
                echo "   Port $port: Open"
            else
                echo "   Port $port: ⚠️ Not responding"
            fi
        fi
    else
        echo "❌ $name: Stopped"
    fi
}

check_service "velocity" "25565"
check_service "karapixel@limbo" "25566"
check_service "karapixel@hub" "25567"
check_service "karapixel@skyblock-spawn" "25568"
check_service "karapixel@pvp-arena" "25569"
check_service "karapixel@skyblock-1" "25570"
check_service "karapixel@skyblock-2" "25571"
check_service "karapixel@nether-end" "25572"
# check_service "karapixel@skyblock-3" "25573"  # Reserve

echo ""
echo "=== Resource Usage ==="
echo "CPU: $(top -bn1 | grep "Cpu(s)" | awk '{print $2}')%"
echo "RAM: $(free -h | grep Mem | awk '{print $3 "/" $2}')"
echo "Disk: $(df -h /opt/karapixel | tail -1 | awk '{print $3 "/" $2 " (" $5 ")"}')"

echo ""
echo "=== Database Status ==="
if systemctl is-active --quiet mysql; then
    echo "✅ MySQL: Running"
else
    echo "❌ MySQL: Stopped"
fi

if systemctl is-active --quiet redis-server; then
    echo "✅ Redis: Running"
else
    echo "❌ Redis: Stopped"
fi

echo ""
echo "=== RAM Allocation ==="
echo "Spawn: 20GB | PvP: 6GB | World1: 24GB | World2: 24GB | Nether: 10GB"
echo "Reserve: 16GB (3. World için hazır)"
```

---

## Port Yapılandırması

| Servis | Port | Protokol | RAM | Açıklama |
|--------|------|----------|-----|----------|
| Velocity | 25565 | TCP | 3GB | Java Minecraft |
| Geyser | 19132 | UDP | 2GB | Bedrock Minecraft |
| Limbo | 25566 | TCP | 1GB | Authentication |
| Hub | 25567 | TCP | 6GB | Main Lobby |
| Skyblock Spawn | 25568 | TCP | 20GB | Market, Event, NPC |
| PvP Arena | 25569 | TCP | 6GB | Combat Server |
| Skyblock #1 | 25570 | TCP | 24GB | Island World |
| Skyblock #2 | 25571 | TCP | 24GB | Island World |
| Nether/End | 25572 | TCP | 10GB | Shared Dimension |
| Skyblock #3 | 25573 | TCP | 16GB* | Reserve (gerekirse) |
| MySQL | 3306 | TCP | 8GB | localhost only |
| Redis | 6379 | TCP | 4GB | localhost only |
| SSH | 22 | TCP | - | Admin access |

*Reserve'den tahsis edilecek

### Firewall Kuralları

```bash
# UFW kuralları
sudo ufw default deny incoming
sudo ufw default allow outgoing

# SSH
sudo ufw allow 22/tcp

# Minecraft
sudo ufw allow 25565/tcp   # Java
sudo ufw allow 19132/udp   # Bedrock

# Internal sunuculara dışarıdan erişim YOK
# 25566-25571 portları sadece localhost

sudo ufw enable
```

---

## DNS Yapılandırması

### Cloudflare DNS

| Subdomain | Type | Value | Proxy |
|-----------|------|-------|-------|
| karapixel.net | A | GAME_SERVER_IP | ❌ |
| play.karapixel.net | A | GAME_SERVER_IP | ❌ |
| www.karapixel.net | A | WEB_VPS_IP | ✅ |
| api.karapixel.net | A | WEB_VPS_IP | ✅ |
| panel.karapixel.net | A | GAME_SERVER_IP | ❌ |

> ⚠️ **ÖNEMLİ:** Minecraft portları Cloudflare proxy'den geçmemeli!

### SRV Record (Opsiyonel)

```
_minecraft._tcp.karapixel.net  SRV  0 5 25565 play.karapixel.net
```

---

## Monitoring

### Prometheus + Grafana

```yaml
# /opt/karapixel/monitoring/prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'minecraft'
    static_configs:
      - targets:
        - 'localhost:9100'  # Node exporter
        - 'localhost:9225'  # Minecraft exporter (Hub)
        - 'localhost:9226'  # Minecraft exporter (Skyblock-1)
```

### Alert Rules

```yaml
# /opt/karapixel/monitoring/alerts.yml
groups:
  - name: minecraft
    rules:
      - alert: ServerDown
        expr: up{job="minecraft"} == 0
        for: 1m
        labels:
          severity: critical
        annotations:
          summary: "Minecraft server down"
          
      - alert: HighTPS
        expr: minecraft_tps < 18
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "TPS dropped below 18"
          
      - alert: HighMemory
        expr: (node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes) / node_memory_MemTotal_bytes > 0.9
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "Memory usage above 90%"
```

---

*📅 Son güncelleme: 24 Aralık 2024 - Seçenek B (20GB Spawn, 2 World + Reserve) onaylandı*
