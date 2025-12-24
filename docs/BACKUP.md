# 💾 KaraPixel - Backup Stratejisi

> 3-2-1 kuralı: 3 kopya, 2 farklı medya, 1 off-site.

---

## Backup Türleri

```
┌─────────────────────────────────────────────────────────────────┐
│                    BACKUP TİPLERİ                               │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  1. DATABASE BACKUP (En Kritik)                                │
│  ═══════════════════════════════                                │
│  Sıklık    : Her 6 saatte                                      │
│  Retention : 7 gün (28 backup)                                 │
│  Method    : mysqldump + gzip                                  │
│  Storage   : Local + Storage Box                               │
│                                                                 │
│  2. WORLD BACKUP                                               │
│  ═══════════════════                                            │
│  Sıklık    : Günde 2 kez (06:00, 18:00)                       │
│  Retention : 14 gün                                            │
│  Method    : rsync + incremental                               │
│  Storage   : Local + Storage Box                               │
│                                                                 │
│  3. CONFIG BACKUP                                              │
│  ══════════════════                                             │
│  Sıklık    : Her değişiklik + günlük                          │
│  Retention : 30 gün                                            │
│  Method    : Git repository                                    │
│  Storage   : GitHub private repo                               │
│                                                                 │
│  4. FULL SERVER BACKUP                                         │
│  ══════════════════════                                         │
│  Sıklık    : Haftalık (Pazar 04:00)                           │
│  Retention : 4 hafta                                           │
│  Method    : Full disk snapshot                                │
│  Storage   : Hetzner Snapshot                                  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Backup Scriptleri

### Database Backup

```bash
#!/bin/bash
# /opt/karapixel/scripts/backup-database.sh

# Konfigürasyon
DB_NAME="karapixel_db"
DB_USER="kara_backup"
DB_PASS="${DB_BACKUP_PASSWORD}"
BACKUP_DIR="/opt/karapixel/backups/database"
REMOTE_DIR="/karapixel/database"
RETENTION_DAYS=7

# Timestamp
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="$BACKUP_DIR/karapixel_$TIMESTAMP.sql.gz"

# Log
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" >> /var/log/karapixel/backup.log
}

log "Starting database backup..."

# Dizin oluştur
mkdir -p $BACKUP_DIR

# MySQL dump
mysqldump -u $DB_USER -p$DB_PASS \
    --single-transaction \
    --routines \
    --triggers \
    --quick \
    $DB_NAME | gzip > $BACKUP_FILE

if [ $? -eq 0 ]; then
    log "Database backup created: $BACKUP_FILE"
    
    # Boyut kontrolü
    SIZE=$(du -h $BACKUP_FILE | cut -f1)
    log "Backup size: $SIZE"
    
    # Remote'a kopyala (Storage Box)
    rsync -avz --progress $BACKUP_FILE \
        -e "ssh -p 23" \
        uXXXXXX@uXXXXXX.your-storagebox.de:$REMOTE_DIR/
    
    if [ $? -eq 0 ]; then
        log "Backup uploaded to Storage Box"
    else
        log "ERROR: Failed to upload backup to Storage Box"
    fi
    
    # Eski backup'ları temizle
    find $BACKUP_DIR -name "*.sql.gz" -mtime +$RETENTION_DAYS -delete
    log "Old backups cleaned (older than $RETENTION_DAYS days)"
    
else
    log "ERROR: Database backup failed!"
    # Discord alert gönder
    /opt/karapixel/scripts/alert.sh "Database backup failed!"
fi

log "Database backup completed"
```

### World Backup

```bash
#!/bin/bash
# /opt/karapixel/scripts/backup-worlds.sh

# Konfigürasyon
SERVERS_DIR="/opt/karapixel/servers"
BACKUP_DIR="/opt/karapixel/backups/worlds"
REMOTE_DIR="/karapixel/worlds"
RETENTION_DAYS=14

TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" >> /var/log/karapixel/backup.log
}

log "Starting world backup..."

# Her skyblock server için
for server in skyblock-1 skyblock-2 skyblock-3; do
    SERVER_DIR="$SERVERS_DIR/$server"
    
    if [ ! -d "$SERVER_DIR" ]; then
        continue
    fi
    
    log "Backing up $server..."
    
    # RCON ile save-off (concurrent yazma önle)
    mcrcon -H localhost -P 256${server: -1}9 -p rcon_password "save-off"
    mcrcon -H localhost -P 256${server: -1}9 -p rcon_password "save-all"
    sleep 5
    
    # World klasörlerini yedekle
    WORLD_BACKUP="$BACKUP_DIR/$server"
    mkdir -p $WORLD_BACKUP
    
    # rsync ile incremental backup
    rsync -avz --delete \
        --exclude='session.lock' \
        --exclude='*.tmp' \
        "$SERVER_DIR/world/" \
        "$WORLD_BACKUP/world_$TIMESTAMP/"
    
    rsync -avz --delete \
        --exclude='session.lock' \
        "$SERVER_DIR/world_nether/" \
        "$WORLD_BACKUP/world_nether_$TIMESTAMP/"
    
    rsync -avz --delete \
        --exclude='session.lock' \
        "$SERVER_DIR/world_the_end/" \
        "$WORLD_BACKUP/world_the_end_$TIMESTAMP/"
    
    # RCON ile save-on
    mcrcon -H localhost -P 256${server: -1}9 -p rcon_password "save-on"
    
    log "$server backup completed"
done

# Hub world backup (tek sefer, değişmez)
log "Backing up hub..."
rsync -avz --delete \
    "$SERVERS_DIR/hub/world/" \
    "$BACKUP_DIR/hub/world_$TIMESTAMP/"

# Remote'a sync
log "Syncing to Storage Box..."
rsync -avz --progress $BACKUP_DIR/ \
    -e "ssh -p 23" \
    uXXXXXX@uXXXXXX.your-storagebox.de:$REMOTE_DIR/

# Eski backup'ları temizle
find $BACKUP_DIR -type d -name "world_*" -mtime +$RETENTION_DAYS -exec rm -rf {} +

log "World backup completed"
```

### Config Backup

```bash
#!/bin/bash
# /opt/karapixel/scripts/backup-configs.sh

CONFIG_REPO="/opt/karapixel/config-backup"
SERVERS_DIR="/opt/karapixel/servers"

log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" >> /var/log/karapixel/backup.log
}

log "Starting config backup..."

cd $CONFIG_REPO

# Tüm server config'lerini kopyala
for server in velocity limbo hub skyblock-spawn skyblock-1 skyblock-2 skyblock-3; do
    mkdir -p $CONFIG_REPO/$server
    
    # Ana config dosyaları
    cp -r $SERVERS_DIR/$server/config/* $CONFIG_REPO/$server/ 2>/dev/null
    cp $SERVERS_DIR/$server/*.yml $CONFIG_REPO/$server/ 2>/dev/null
    cp $SERVERS_DIR/$server/*.toml $CONFIG_REPO/$server/ 2>/dev/null
    cp $SERVERS_DIR/$server/*.properties $CONFIG_REPO/$server/ 2>/dev/null
    
    # Plugin config'leri (hassas veriler hariç)
    mkdir -p $CONFIG_REPO/$server/plugins
    for plugin_dir in $SERVERS_DIR/$server/plugins/*/; do
        plugin_name=$(basename $plugin_dir)
        mkdir -p $CONFIG_REPO/$server/plugins/$plugin_name
        cp $plugin_dir/*.yml $CONFIG_REPO/$server/plugins/$plugin_name/ 2>/dev/null
    done
done

# Hassas verileri kaldır
find $CONFIG_REPO -name "*.yml" -exec sed -i 's/password:.*/password: REDACTED/g' {} \;
find $CONFIG_REPO -name "*.yml" -exec sed -i 's/secret:.*/secret: REDACTED/g' {} \;

# Git commit
git add -A
git commit -m "Auto backup: $(date '+%Y-%m-%d %H:%M:%S')"
git push origin main

log "Config backup completed"
```

### Emergency Backup

```bash
#!/bin/bash
# /opt/karapixel/scripts/emergency-backup.sh
# Acil durumlarda tüm verilerin hızlı yedeği

TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
EMERGENCY_DIR="/opt/karapixel/backups/emergency_$TIMESTAMP"

echo "🚨 EMERGENCY BACKUP STARTED"
echo "Timestamp: $TIMESTAMP"

# Tüm sunucuları durdur
echo "Stopping all servers..."
/opt/karapixel/scripts/stop-all.sh

# Tüm verileri yedekle
mkdir -p $EMERGENCY_DIR

echo "Backing up database..."
mysqldump -u root karapixel_db > $EMERGENCY_DIR/database.sql

echo "Backing up servers..."
tar -czf $EMERGENCY_DIR/servers.tar.gz /opt/karapixel/servers

echo "Backing up configs..."
tar -czf $EMERGENCY_DIR/configs.tar.gz /opt/karapixel/shared

echo "Compressing emergency backup..."
tar -czf /opt/karapixel/backups/emergency_$TIMESTAMP.tar.gz $EMERGENCY_DIR

# Boyut
SIZE=$(du -h /opt/karapixel/backups/emergency_$TIMESTAMP.tar.gz | cut -f1)
echo "Emergency backup size: $SIZE"

# Remote'a upload
echo "Uploading to Storage Box..."
rsync -avz --progress /opt/karapixel/backups/emergency_$TIMESTAMP.tar.gz \
    -e "ssh -p 23" \
    uXXXXXX@uXXXXXX.your-storagebox.de:/karapixel/emergency/

echo "🚨 EMERGENCY BACKUP COMPLETED"
echo "Location: /opt/karapixel/backups/emergency_$TIMESTAMP.tar.gz"

# Alert
/opt/karapixel/scripts/alert.sh "Emergency backup completed: $SIZE"
```

---

## Crontab

```bash
# /etc/crontab veya crontab -e

# Database backup - Her 6 saatte
0 */6 * * * root /opt/karapixel/scripts/backup-database.sh

# World backup - Günde 2 kez
0 6 * * * root /opt/karapixel/scripts/backup-worlds.sh
0 18 * * * root /opt/karapixel/scripts/backup-worlds.sh

# Config backup - Her gece
0 3 * * * root /opt/karapixel/scripts/backup-configs.sh

# Log cleanup - Her hafta
0 4 * * 0 root /opt/karapixel/scripts/cleanup-logs.sh

# Health check - Her 5 dakika
*/5 * * * * root /opt/karapixel/scripts/health-check.sh > /dev/null
```

---

## Disaster Recovery

### Recovery Time Objectives (RTO)

| Senaryo | RTO | Prosedür |
|---------|-----|----------|
| Database corruption | 15-30 dk | Son backup'tan restore |
| World corruption | 30-60 dk | Son backup'tan restore |
| Single server failure | 10-15 dk | Systemd restart |
| Full server failure | 2-4 saat | Full restore |
| Hetzner DC failure | 4-8 saat | Storage Box'tan restore |

### Database Restore

```bash
#!/bin/bash
# /opt/karapixel/scripts/restore-database.sh

# Kullanım: ./restore-database.sh <backup_file>

BACKUP_FILE=$1

if [ -z "$BACKUP_FILE" ]; then
    echo "Usage: $0 <backup_file.sql.gz>"
    echo "Available backups:"
    ls -la /opt/karapixel/backups/database/
    exit 1
fi

echo "⚠️  WARNING: This will overwrite the current database!"
echo "Backup file: $BACKUP_FILE"
read -p "Are you sure? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo "Aborted."
    exit 0
fi

# Sunucuları durdur
echo "Stopping game servers..."
/opt/karapixel/scripts/stop-all.sh

# Restore
echo "Restoring database..."
gunzip -c $BACKUP_FILE | mysql -u root karapixel_db

if [ $? -eq 0 ]; then
    echo "✅ Database restored successfully!"
else
    echo "❌ Database restore failed!"
    exit 1
fi

# Sunucuları başlat
echo "Starting game servers..."
/opt/karapixel/scripts/start-all.sh

echo "Recovery complete!"
```

### World Restore

```bash
#!/bin/bash
# /opt/karapixel/scripts/restore-world.sh

# Kullanım: ./restore-world.sh <server> <backup_timestamp>

SERVER=$1
TIMESTAMP=$2

if [ -z "$SERVER" ] || [ -z "$TIMESTAMP" ]; then
    echo "Usage: $0 <server> <backup_timestamp>"
    echo "Example: $0 skyblock-1 20241224_060000"
    exit 1
fi

BACKUP_DIR="/opt/karapixel/backups/worlds/$SERVER/world_$TIMESTAMP"
SERVER_DIR="/opt/karapixel/servers/$SERVER"

if [ ! -d "$BACKUP_DIR" ]; then
    echo "Backup not found: $BACKUP_DIR"
    echo "Available backups:"
    ls -la /opt/karapixel/backups/worlds/$SERVER/
    exit 1
fi

echo "⚠️  WARNING: This will overwrite the current world!"
echo "Server: $SERVER"
echo "Backup: $TIMESTAMP"
read -p "Are you sure? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo "Aborted."
    exit 0
fi

# Server'ı durdur
echo "Stopping $SERVER..."
sudo systemctl stop karapixel@$SERVER

# Mevcut world'ü yedekle (güvenlik için)
EMERGENCY_TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
mv $SERVER_DIR/world $SERVER_DIR/world_corrupted_$EMERGENCY_TIMESTAMP

# Restore
echo "Restoring world..."
cp -r $BACKUP_DIR $SERVER_DIR/world

# Server'ı başlat
echo "Starting $SERVER..."
sudo systemctl start karapixel@$SERVER

echo "✅ World restored successfully!"
```

---

## Hetzner Storage Box Kurulumu

### SSH Key Ekleme

```bash
# Storage Box'a SSH key ekle
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys

# Key'i Storage Box'a kopyala
echo "ecdsa-sha2-nistp256 AAAA...= backup@karapixel" | \
    ssh -p 23 uXXXXXX@uXXXXXX.your-storagebox.de install-ssh-key
```

### Dizin Yapısı

```
/karapixel/
├── database/
│   ├── karapixel_20241224_060000.sql.gz
│   ├── karapixel_20241224_120000.sql.gz
│   └── ...
├── worlds/
│   ├── skyblock-1/
│   │   ├── world_20241224_060000/
│   │   └── ...
│   ├── skyblock-2/
│   └── skyblock-3/
├── emergency/
│   └── emergency_20241224_153000.tar.gz
└── configs/
    └── (Git repo clone)
```

### Auto-mount (Opsiyonel)

```bash
# /etc/fstab
# SSHFS ile mount (performans için önerilmez, sadece browse için)
# uXXXXXX@uXXXXXX.your-storagebox.de:/karapixel /mnt/backup fuse.sshfs port=23,_netdev,allow_other 0 0
```

---

## Doğrulama ve Test

### Backup Integrity Check

```bash
#!/bin/bash
# /opt/karapixel/scripts/verify-backup.sh

echo "=== Backup Verification ==="

# Son database backup kontrolü
echo ""
echo "Database backups:"
ls -lah /opt/karapixel/backups/database/ | tail -5

LATEST_DB=$(ls -t /opt/karapixel/backups/database/*.sql.gz | head -1)
echo ""
echo "Latest: $LATEST_DB"
echo "Testing integrity..."
gunzip -t $LATEST_DB && echo "✅ OK" || echo "❌ CORRUPTED"

# Son world backup kontrolü
echo ""
echo "World backups:"
for server in skyblock-1 skyblock-2 skyblock-3; do
    LATEST=$(ls -td /opt/karapixel/backups/worlds/$server/world_* 2>/dev/null | head -1)
    if [ -n "$LATEST" ]; then
        echo "$server: $(basename $LATEST)"
    else
        echo "$server: No backup found!"
    fi
done

# Remote backup kontrolü
echo ""
echo "Remote storage (Storage Box):"
ssh -p 23 uXXXXXX@uXXXXXX.your-storagebox.de "du -sh /karapixel/*"
```

### Restore Test (Staging)

```bash
# Aylık restore testi yap
# Test ortamında backup'tan restore et ve doğrula
```

---

*📅 Son güncelleme: 24 Aralık 2024*
