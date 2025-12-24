#!/bin/bash
#===============================================================================
# KaraPixel Game Server - Tam Kurulum Scripti
# Sunucu: 46.4.66.233 (Hetzner Ryzen 9 5950X, 128GB RAM)
# Tarih: 24 Aralık 2024
#===============================================================================

set -e  # Hata durumunda dur

echo "╔══════════════════════════════════════════════════════════════════╗"
echo "║           KARAPIXEL SUNUCU KURULUMU BAŞLIYOR                    ║"
echo "╚══════════════════════════════════════════════════════════════════╝"

#===============================================================================
# 1. SİSTEM GÜNCELLEME
#===============================================================================
echo ""
echo "[1/10] Sistem güncelleniyor..."
apt update && apt upgrade -y

#===============================================================================
# 2. TEMEL PAKETLER
#===============================================================================
echo ""
echo "[2/10] Temel paketler kuruluyor..."
apt install -y \
    curl \
    wget \
    git \
    unzip \
    htop \
    nano \
    screen \
    net-tools \
    ufw \
    fail2ban \
    ncdu \
    iotop \
    tmux \
    jq

#===============================================================================
# 3. MINECRAFT KULLANICISI VE DİZİN YAPISI
#===============================================================================
echo ""
echo "[3/10] Minecraft kullanıcısı ve dizinler oluşturuluyor..."

# Kullanıcı oluştur (varsa atla)
if ! id "minecraft" &>/dev/null; then
    useradd -r -m -U -d /opt/karapixel -s /bin/bash minecraft
    echo "minecraft kullanıcısı oluşturuldu"
else
    echo "minecraft kullanıcısı zaten var"
fi

# Dizin yapısı
mkdir -p /opt/karapixel/{servers/{velocity,limbo,hub,skyblock-spawn,pvp-arena,skyblock-1,skyblock-2,nether-end},shared/{plugins,configs,resourcepack},scripts,backups/{database,worlds,configs},logs,data/{mysql,redis}}

# Sahiplik
chown -R minecraft:minecraft /opt/karapixel
chmod -R 755 /opt/karapixel

echo "Dizin yapısı oluşturuldu: /opt/karapixel"

#===============================================================================
# 4. JAVA 21 KURULUMU
#===============================================================================
echo ""
echo "[4/10] Java 21 kuruluyor..."

apt install -y openjdk-21-jdk

# JAVA_HOME ayarla
echo 'export JAVA_HOME=/usr/lib/jvm/java-21-openjdk-amd64' >> /etc/profile.d/java.sh
echo 'export PATH=$JAVA_HOME/bin:$PATH' >> /etc/profile.d/java.sh
source /etc/profile.d/java.sh

java -version
echo "Java 21 kuruldu"

#===============================================================================
# 5. MYSQL 8 KURULUMU
#===============================================================================
echo ""
echo "[5/10] MySQL 8 kuruluyor..."

apt install -y mysql-server

# MySQL başlat
systemctl enable mysql
systemctl start mysql

# Güvenli şifre oluştur
MYSQL_ROOT_PASS=$(openssl rand -base64 24)
MYSQL_KARA_PASS=$(openssl rand -base64 24)

# Database ve kullanıcı oluştur
mysql -u root << EOF
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '${MYSQL_ROOT_PASS}';
CREATE DATABASE IF NOT EXISTS karapixel_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER IF NOT EXISTS 'karapixel'@'localhost' IDENTIFIED BY '${MYSQL_KARA_PASS}';
GRANT ALL PRIVILEGES ON karapixel_db.* TO 'karapixel'@'localhost';
FLUSH PRIVILEGES;
EOF

# MySQL performans ayarları
cat > /etc/mysql/mysql.conf.d/karapixel.cnf << 'EOF'
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

# Performance
performance_schema = ON

# Logging
slow_query_log = 1
slow_query_log_file = /var/log/mysql/slow.log
long_query_time = 2
EOF

systemctl restart mysql

# Şifreleri kaydet
cat > /opt/karapixel/.mysql-credentials << EOF
MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASS}
MYSQL_KARAPIXEL_USER=karapixel
MYSQL_KARAPIXEL_PASSWORD=${MYSQL_KARA_PASS}
MYSQL_DATABASE=karapixel_db
EOF
chmod 600 /opt/karapixel/.mysql-credentials
chown minecraft:minecraft /opt/karapixel/.mysql-credentials

echo "MySQL 8 kuruldu ve yapılandırıldı"
echo "Şifreler: /opt/karapixel/.mysql-credentials"

#===============================================================================
# 6. REDIS KURULUMU
#===============================================================================
echo ""
echo "[6/10] Redis kuruluyor..."

apt install -y redis-server

# Redis şifresi oluştur
REDIS_PASS=$(openssl rand -base64 24)

# Redis yapılandırması
cat > /etc/redis/redis.conf << EOF
bind 127.0.0.1
port 6379

# Authentication
requirepass ${REDIS_PASS}

# Memory
maxmemory 4gb
maxmemory-policy allkeys-lru

# Persistence
save 900 1
save 300 10
save 60 10000

dir /var/lib/redis
dbfilename dump.rdb

# Security
rename-command FLUSHDB ""
rename-command FLUSHALL ""
rename-command DEBUG ""

# Logging
loglevel notice
logfile /var/log/redis/redis-server.log
EOF

systemctl enable redis-server
systemctl restart redis-server

# Redis şifresini kaydet
echo "REDIS_PASSWORD=${REDIS_PASS}" >> /opt/karapixel/.redis-credentials
chmod 600 /opt/karapixel/.redis-credentials
chown minecraft:minecraft /opt/karapixel/.redis-credentials

echo "Redis kuruldu ve yapılandırıldı"
echo "Şifre: /opt/karapixel/.redis-credentials"

#===============================================================================
# 7. FIREWALL (UFW) YAPILANDIRMASI
#===============================================================================
echo ""
echo "[7/10] Firewall yapılandırılıyor..."

ufw default deny incoming
ufw default allow outgoing

# SSH
ufw allow 22/tcp

# Minecraft Java
ufw allow 25565/tcp

# Minecraft Bedrock (Geyser)
ufw allow 19132/udp

# Enable
ufw --force enable

echo "Firewall yapılandırıldı"
ufw status

#===============================================================================
# 8. FAIL2BAN YAPILANDIRMASI
#===============================================================================
echo ""
echo "[8/10] Fail2ban yapılandırılıyor..."

cat > /etc/fail2ban/jail.local << 'EOF'
[DEFAULT]
bantime = 1h
findtime = 10m
maxretry = 5

[sshd]
enabled = true
port = ssh
filter = sshd
logpath = /var/log/auth.log
maxretry = 3
bantime = 24h
EOF

systemctl enable fail2ban
systemctl restart fail2ban

echo "Fail2ban yapılandırıldı"

#===============================================================================
# 9. VELOCITY PROXY KURULUMU
#===============================================================================
echo ""
echo "[9/10] Velocity proxy kuruluyor..."

cd /opt/karapixel/servers/velocity

# Velocity indir (en son sürüm)
VELOCITY_VERSION="3.4.0-SNAPSHOT"
VELOCITY_BUILD=$(curl -s "https://api.papermc.io/v2/projects/velocity/versions/${VELOCITY_VERSION}/builds" | jq -r '.builds[-1].build')
curl -o velocity.jar "https://api.papermc.io/v2/projects/velocity/versions/${VELOCITY_VERSION}/builds/${VELOCITY_BUILD}/downloads/velocity-${VELOCITY_VERSION}-${VELOCITY_BUILD}.jar"

# İlk çalıştırma (config oluşturması için)
cd /opt/karapixel/servers/velocity
timeout 15 java -Xms1G -Xmx1G -jar velocity.jar || true

# Velocity yapılandırması
cat > /opt/karapixel/servers/velocity/velocity.toml << 'EOF'
# KaraPixel Velocity Configuration

config-version = "2.7"
bind = "0.0.0.0:25565"
motd = "<gradient:purple:blue>★ KaraPixel Network ★</gradient>\n<gray>Skyblock & Daha Fazlası!</gray>"
show-max-players = 500
online-mode = true
force-key-authentication = true
prevent-client-proxy-connections = false
player-info-forwarding-mode = "modern"

[servers]
limbo = "127.0.0.1:25566"
hub = "127.0.0.1:25567"
skyblock-spawn = "127.0.0.1:25568"
pvp-arena = "127.0.0.1:25569"
skyblock-1 = "127.0.0.1:25570"
skyblock-2 = "127.0.0.1:25571"
nether-end = "127.0.0.1:25572"

try = ["limbo"]

[forced-hosts]
"play.karapixel.net" = ["limbo"]

[advanced]
compression-threshold = 256
compression-level = -1
login-ratelimit = 3000
connection-timeout = 5000
read-timeout = 30000
haproxy-protocol = false
tcp-fast-open = false
bungee-plugin-message-channel = true
show-ping-requests = false
failover-on-unexpected-server-disconnect = true
announce-proxy-commands = true
log-command-executions = false
log-player-connections = true

[query]
enabled = false
port = 25565
map = "KaraPixel"
show-plugins = false
EOF

# Forwarding secret oluştur
openssl rand -base64 32 > /opt/karapixel/servers/velocity/forwarding.secret

chown -R minecraft:minecraft /opt/karapixel/servers/velocity

echo "Velocity proxy kuruldu"

#===============================================================================
# 10. GEYSER + FLOODGATE KURULUMU
#===============================================================================
echo ""
echo "[10/10] Geyser ve Floodgate kuruluyor..."

mkdir -p /opt/karapixel/servers/velocity/plugins

cd /opt/karapixel/servers/velocity/plugins

# Geyser-Velocity indir
GEYSER_URL=$(curl -s "https://download.geysermc.org/v2/projects/geyser/versions/latest/builds/latest" | jq -r '.downloads.velocity.url // empty')
if [ -n "$GEYSER_URL" ]; then
    curl -o Geyser-Velocity.jar "https://download.geysermc.org${GEYSER_URL}"
else
    curl -o Geyser-Velocity.jar "https://download.geysermc.org/v2/projects/geyser/versions/latest/builds/latest/downloads/velocity"
fi

# Floodgate-Velocity indir
FLOODGATE_URL=$(curl -s "https://download.geysermc.org/v2/projects/floodgate/versions/latest/builds/latest" | jq -r '.downloads.velocity.url // empty')
if [ -n "$FLOODGATE_URL" ]; then
    curl -o floodgate-velocity.jar "https://download.geysermc.org${FLOODGATE_URL}"
else
    curl -o floodgate-velocity.jar "https://download.geysermc.org/v2/projects/floodgate/versions/latest/builds/latest/downloads/velocity"
fi

chown -R minecraft:minecraft /opt/karapixel/servers/velocity

echo "Geyser ve Floodgate kuruldu"

#===============================================================================
# 11. SYSTEMD SERVİSLERİ
#===============================================================================
echo ""
echo "[11/10] Systemd servisleri oluşturuluyor..."

# Velocity servisi
cat > /etc/systemd/system/velocity.service << 'EOF'
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
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
EOF

# Game server template servisi
cat > /etc/systemd/system/karapixel@.service << 'EOF'
[Unit]
Description=KaraPixel %i Server
After=network.target mysql.service redis.service velocity.service

[Service]
User=minecraft
Group=minecraft
WorkingDirectory=/opt/karapixel/servers/%i

EnvironmentFile=/opt/karapixel/servers/%i/server.env

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
    -jar server.jar --nogui

Restart=always
RestartSec=10
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
EOF

# Environment dosyaları
echo "MEMORY=1G" > /opt/karapixel/servers/limbo/server.env
echo "MEMORY=6G" > /opt/karapixel/servers/hub/server.env
echo "MEMORY=20G" > /opt/karapixel/servers/skyblock-spawn/server.env
echo "MEMORY=6G" > /opt/karapixel/servers/pvp-arena/server.env
echo "MEMORY=24G" > /opt/karapixel/servers/skyblock-1/server.env
echo "MEMORY=24G" > /opt/karapixel/servers/skyblock-2/server.env
echo "MEMORY=10G" > /opt/karapixel/servers/nether-end/server.env

chown -R minecraft:minecraft /opt/karapixel/servers

# Systemd reload
systemctl daemon-reload

echo "Systemd servisleri oluşturuldu"

#===============================================================================
# 12. YÖNETİM SCRİPTLERİ
#===============================================================================
echo ""
echo "[12/10] Yönetim scriptleri oluşturuluyor..."

# start-all.sh
cat > /opt/karapixel/scripts/start-all.sh << 'EOF'
#!/bin/bash
echo "KaraPixel sunucuları başlatılıyor..."
systemctl start velocity
sleep 5
echo "Velocity başlatıldı"
EOF
chmod +x /opt/karapixel/scripts/start-all.sh

# stop-all.sh
cat > /opt/karapixel/scripts/stop-all.sh << 'EOF'
#!/bin/bash
echo "KaraPixel sunucuları durduruluyor..."
systemctl stop velocity
echo "Tüm sunucular durduruldu"
EOF
chmod +x /opt/karapixel/scripts/stop-all.sh

# health-check.sh
cat > /opt/karapixel/scripts/health-check.sh << 'EOF'
#!/bin/bash
echo "=== KaraPixel Durum Kontrolü ==="
echo ""
echo "Servisler:"
systemctl is-active velocity && echo "✅ Velocity: Çalışıyor" || echo "❌ Velocity: Durduruldu"
systemctl is-active mysql && echo "✅ MySQL: Çalışıyor" || echo "❌ MySQL: Durduruldu"
systemctl is-active redis-server && echo "✅ Redis: Çalışıyor" || echo "❌ Redis: Durduruldu"
echo ""
echo "Kaynak Kullanımı:"
echo "CPU: $(top -bn1 | grep 'Cpu(s)' | awk '{print $2}')%"
echo "RAM: $(free -h | grep Mem | awk '{print $3 "/" $2}')"
echo "Disk: $(df -h /opt/karapixel | tail -1 | awk '{print $3 "/" $2 " (" $5 ")"}')"
EOF
chmod +x /opt/karapixel/scripts/health-check.sh

chown -R minecraft:minecraft /opt/karapixel/scripts

echo "Yönetim scriptleri oluşturuldu"

#===============================================================================
# TAMAMLANDI
#===============================================================================
echo ""
echo "╔══════════════════════════════════════════════════════════════════╗"
echo "║              KURULUM BAŞARIYLA TAMAMLANDI!                       ║"
echo "╠══════════════════════════════════════════════════════════════════╣"
echo "║                                                                  ║"
echo "║  Kurulan Servisler:                                             ║"
echo "║  ├── Java 21                                                    ║"
echo "║  ├── MySQL 8 (Port: 3306)                                       ║"
echo "║  ├── Redis (Port: 6379)                                         ║"
echo "║  ├── Velocity Proxy (Port: 25565)                               ║"
echo "║  ├── Geyser (Port: 19132/UDP)                                   ║"
echo "║  └── Floodgate                                                  ║"
echo "║                                                                  ║"
echo "║  Şifreler:                                                      ║"
echo "║  ├── MySQL: /opt/karapixel/.mysql-credentials                   ║"
echo "║  └── Redis: /opt/karapixel/.redis-credentials                   ║"
echo "║                                                                  ║"
echo "║  Komutlar:                                                      ║"
echo "║  ├── systemctl start velocity    (Proxy başlat)                 ║"
echo "║  ├── /opt/karapixel/scripts/health-check.sh                     ║"
echo "║  └── journalctl -u velocity -f   (Log izle)                     ║"
echo "║                                                                  ║"
echo "╚══════════════════════════════════════════════════════════════════╝"
echo ""

# Şifreleri göster
echo "=== KAYITLI ŞİFRELER (SAKLA!) ==="
echo ""
cat /opt/karapixel/.mysql-credentials
echo ""
cat /opt/karapixel/.redis-credentials
echo ""
