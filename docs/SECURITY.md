# 🛡️ KaraPixel - Güvenlik Stratejisi

> ⚠️ **KRİTİK DOKÜMAN** - Bu dosya sunucu güvenliğinin temelini oluşturur.
> Non-premium sunucu olarak ekstra güvenlik önlemleri ZORUNLUDUR.

---

## 📋 İçindekiler

1. [Güvenlik Katmanları](#güvenlik-katmanları)
2. [Saldırı Tipleri ve Savunma](#saldırı-tipleri-ve-savunma)
3. [DDoS Koruması](#ddos-koruması)
4. [Bot ve Spam Koruması](#bot-ve-spam-koruması)
5. [Exploit Koruması](#exploit-koruması)
6. [Hesap Güvenliği](#hesap-güvenliği)
7. [Veri Güvenliği](#veri-güvenliği)
8. [Bilgi Gizleme](#bilgi-gizleme)
9. [Monitoring ve Alerting](#monitoring-ve-alerting)
10. [Incident Response](#incident-response)

---

## Güvenlik Katmanları

```
┌─────────────────────────────────────────────────────────────────┐
│                    6 KATMANLI GÜVENLİK                          │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  KATMAN 6: MONİTORİNG & ALERTİNG                               │
│  ════════════════════════════════                               │
│  Prometheus, Grafana, Discord Alerts                           │
│  Anomaly detection, Audit logging                              │
│                                                                 │
│  KATMAN 5: VERİTABANI & DATA                                   │
│  ═══════════════════════════                                    │
│  SQL injection koruması, encryption                            │
│  Input validation, rate limiting                               │
│                                                                 │
│  KATMAN 4: GAME SERVERS                                        │
│  ══════════════════════                                         │
│  Anti-cheat, exploit patches                                   │
│  Permission lockdown, backend isolation                        │
│                                                                 │
│  KATMAN 3: LIMBO (Authentication)                              │
│  ════════════════════════════════                               │
│  Captcha, login protection, session management                 │
│  Anti-bot, behavioral analysis                                 │
│                                                                 │
│  KATMAN 2: VELOCITY PROXY                                      │
│  ════════════════════════                                       │
│  Rate limiting, connection filtering                           │
│  Modern forwarding, BungeeGuard                                │
│                                                                 │
│  KATMAN 1: NETWORK (DDoS)                                      │
│  ════════════════════════                                       │
│  TCPShield/Cosmic Guard, Hetzner DDoS Protection              │
│  UFW Firewall, Fail2Ban                                        │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Saldırı Tipleri ve Savunma

### Saldırı Matrisi

| Saldırı Tipi | Risk | Olasılık | Savunma Katmanı |
|--------------|------|----------|-----------------|
| DDoS (L3/L4) | 🔴 Kritik | Yüksek | Katman 1 |
| DDoS (L7) | 🔴 Kritik | Yüksek | Katman 1-2 |
| Bot Attack | 🔴 Kritik | Çok Yüksek | Katman 2-3 |
| Brute Force | 🟠 Yüksek | Yüksek | Katman 3 |
| Crash Exploit | 🟠 Yüksek | Orta | Katman 4 |
| SQL Injection | 🟠 Yüksek | Düşük | Katman 5 |
| Dupe/Item Exploit | 🟡 Orta | Orta | Katman 4 |
| Permission Exploit | 🟡 Orta | Düşük | Katman 4 |
| Name Spoofing | 🟡 Orta | Yüksek | Katman 3 |
| Social Engineering | 🟢 Düşük | Orta | Prosedür |

### Detaylı Saldırı-Savunma Tablosu

```
┌─────────────────────────────────────────────────────────────────┐
│              SALDIRI TİPLERİ VE SAVUNMA DETAYI                  │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  1. DDoS SALDIRISI                                              │
│  ══════════════════                                             │
│  Tip: SYN Flood, UDP Flood, HTTP Flood, Slowloris              │
│  Hedef: Sunucuyu erişilemez yapmak                             │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── TCPShield Premium ($20/ay) veya Cosmic Guard              │
│  │   ├── L3/L4 DDoS mitigation                                 │
│  │   ├── L7 (application layer) filtering                      │
│  │   ├── Anycast network                                        │
│  │   └── Real IP forwarding                                    │
│  ├── Hetzner Built-in DDoS Protection                          │
│  ├── iptables rate limiting:                                   │
│  │   └── iptables -A INPUT -p tcp --syn -m limit               │
│  │       --limit 10/s --limit-burst 20 -j ACCEPT               │
│  └── Fail2Ban custom rules                                     │
│                                                                 │
│  2. BOT ATTACK (Fake Players)                                   │
│  ════════════════════════════                                   │
│  Tip: Binlerce sahte bağlantı, resource exhaustion             │
│  Hedef: Sunucu kaynaklarını tüketmek                           │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── Limbo/FakeLobby katmanı (ZORUNLU)                        │
│  │   └── Bot'lar auth'da takılır, Hub'a ulaşamaz              │
│  ├── Captcha sistemi:                                          │
│  │   ├── Map-based captcha (görsel)                            │
│  │   ├── Math captcha (matematik sorusu)                       │
│  │   └── Click pattern (Bedrock uyumlu)                        │
│  ├── Connection rate limiting:                                 │
│  │   └── IP başına max 3 bağlantı/saniye                       │
│  ├── Behavioral analysis:                                      │
│  │   ├── Bağlantı timing analizi                               │
│  │   ├── Movement pattern kontrolü                             │
│  │   └── Chat pattern kontrolü                                 │
│  └── IP Reputation check:                                      │
│      └── Bilinen bot IP'lerini engelle                         │
│                                                                 │
│  3. BRUTE FORCE (Şifre Deneme)                                 │
│  ══════════════════════════════                                 │
│  Tip: Şifre tahmin saldırısı                                   │
│  Hedef: Hesap ele geçirme                                      │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── Rate limiting:                                            │
│  │   └── Max 5 deneme / 5 dakika (sonra 15 dk ban)            │
│  ├── Güçlü hashing:                                            │
│  │   └── bcrypt (cost factor: 12)                              │
│  ├── Login timeout:                                            │
│  │   └── 30 saniye içinde giriş yapmazsa kick                  │
│  ├── Progressive delay:                                        │
│  │   └── Her başarısız denemede artan bekleme                  │
│  └── IP logging:                                               │
│      └── Şüpheli IP'leri kaydet ve analiz et                   │
│                                                                 │
│  4. CRASH/EXPLOIT SALDIRISI                                    │
│  ═══════════════════════════                                    │
│  Tip: Book ban, chunk ban, NBT overflow, packet exploit        │
│  Hedef: Sunucuyu çökertmek veya oyuncuları banlamak           │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── KaraPaper exploit patches:                                │
│  │   ├── Book page limit (max 100 sayfa)                       │
│  │   ├── Sign text limit (max 100 karakter)                    │
│  │   ├── NBT size limit (max 200KB)                            │
│  │   ├── Chunk data validation                                 │
│  │   └── Invalid packet rejection                              │
│  ├── Packet inspection:                                        │
│  │   └── Anormal paket boyutu/frekansı engelle                │
│  ├── Watchdog:                                                 │
│  │   └── Crash sonrası otomatik restart                        │
│  └── Regular backups:                                          │
│      └── Corruption durumunda geri dönüş                       │
│                                                                 │
│  5. ITEM DUPE / ECONOMY EXPLOIT                                │
│  ══════════════════════════════                                 │
│  Tip: Item çoğaltma, para çoğaltma                            │
│  Hedef: Ekonomiyi bozmak                                       │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── Transaction logging:                                      │
│  │   └── Tüm para/item hareketlerini kaydet                    │
│  ├── Anomaly detection:                                        │
│  │   └── Anormal artışları tespit et                           │
│  ├── Rate limiting:                                            │
│  │   └── Max işlem/dakika limiti                               │
│  ├── Inventory validation:                                     │
│  │   └── Her işlemde envanter doğrulama                        │
│  └── Item hash verification:                                   │
│      └── Item metadata bütünlük kontrolü                       │
│                                                                 │
│  6. HESAP ELE GEÇİRME                                          │
│  ════════════════════                                           │
│  Tip: Session hijacking, name spoofing                         │
│  Hedef: Başka oyuncunun hesabını kullanmak                     │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── Session güvenliği:                                        │
│  │   ├── Redis-based sessions (sunucu tarafı)                  │
│  │   ├── Session timeout (7 gün)                               │
│  │   └── IP-bound sessions (opsiyonel)                         │
│  ├── Name protection:                                          │
│  │   ├── Admin/staff isimlerini koru                           │
│  │   ├── Benzer isim kontrolü (l vs I, 0 vs O)                │
│  │   └── Reserved name listesi                                 │
│  └── 2FA (opsiyonel):                                          │
│      └── Değerli hesaplar için TOTP                            │
│                                                                 │
│  7. YETKİ YÜKSELTMESİ                                          │
│  ════════════════════                                           │
│  Tip: Permission exploit, command injection                    │
│  Hedef: Admin yetkisi almak                                    │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── Zero-trust permissions:                                   │
│  │   └── Default: HİÇBİR yetki                                 │
│  ├── OP sistemi devre dışı:                                    │
│  │   └── ops.json boş, OP komutları kapalı                     │
│  ├── Command whitelist:                                        │
│  │   └── Her rank için izin verilen komut listesi              │
│  ├── Input sanitization:                                       │
│  │   └── Command injection önleme                              │
│  └── Audit logging:                                            │
│      └── Tüm admin işlemlerini kaydet                          │
│                                                                 │
│  8. SQL INJECTION                                               │
│  ════════════════                                               │
│  Tip: Veritabanı manipülasyonu                                 │
│  Hedef: Veri çalma veya değiştirme                             │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── Prepared statements:                                      │
│  │   └── HER SQL sorgusu parameterized                         │
│  ├── Input validation:                                         │
│  │   └── Tüm kullanıcı girdilerini filtrele                    │
│  ├── Least privilege:                                          │
│  │   └── Her servis için ayrı DB kullanıcısı                   │
│  └── Database isolation:                                       │
│      └── Localhost only binding                                │
│                                                                 │
│  9. SOCIAL ENGINEERING                                         │
│  ═════════════════════                                          │
│  Tip: Staff taklidi, phishing, güven istismarı                │
│  Hedef: Hassas bilgi alma                                      │
│                                                                 │
│  SAVUNMA:                                                       │
│  ├── Staff verification:                                       │
│  │   └── Resmi staff listesi, doğrulama sistemi               │
│  ├── Official channels:                                        │
│  │   └── Sadece resmi kanallardan iletişim                     │
│  ├── Player education:                                         │
│  │   └── Güvenlik uyarıları, bilgilendirme                    │
│  └── Report system:                                            │
│      └── Kolay şüpheli davranış bildirimi                      │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## DDoS Koruması

### TCPShield / Cosmic Guard Entegrasyonu

```yaml
# velocity.toml - TCPShield için
[advanced]
tcp-fast-open = true
connection-timeout = 5000
login-ratelimit = 3000

# TCPShield real IP forwarding
[servers]
try = ["limbo"]

# proxy-protocol aktif
proxy-protocol = true
```

### Firewall Kuralları (UFW)

```bash
#!/bin/bash
# /opt/karapixel/scripts/setup-firewall.sh

# Reset rules
ufw --force reset

# Default policies
ufw default deny incoming
ufw default allow outgoing

# SSH (sadece belirli IP'lerden - opsiyonel)
ufw allow 22/tcp

# Minecraft ports (TCPShield kullanılıyorsa kapatılabilir)
ufw allow 25565/tcp  # Java
ufw allow 19132/udp  # Bedrock

# Internal services (localhost only)
# MySQL, Redis sadece localhost'tan erişilebilir

# Rate limiting
ufw limit ssh

# Enable
ufw enable

echo "Firewall configured successfully"
```

### Fail2Ban Minecraft Kuralları

```ini
# /etc/fail2ban/jail.d/minecraft.conf

[minecraft-auth]
enabled = true
port = 25565
filter = minecraft-auth
logpath = /opt/karapixel/servers/limbo/logs/latest.log
maxretry = 5
findtime = 300
bantime = 900

[minecraft-connection]
enabled = true
port = 25565
filter = minecraft-connection
logpath = /var/log/karapixel/velocity.log
maxretry = 20
findtime = 60
bantime = 3600
```

```ini
# /etc/fail2ban/filter.d/minecraft-auth.conf
[Definition]
failregex = ^.*\[KaraAuth\] Failed login attempt from <HOST>.*$
            ^.*\[KaraAuth\] Too many attempts from <HOST>.*$
ignoreregex =
```

---

## Bot ve Spam Koruması

### Captcha Sistemi

```java
// karapixel-auth içinde
public class CaptchaSystem {
    
    public enum CaptchaType {
        MAP,      // Görsel captcha (map item üzerinde)
        MATH,     // Matematik sorusu
        CLICK     // Belirli slot'a tıklama (Bedrock uyumlu)
    }
    
    // Bedrock oyuncular için CLICK tercih edilmeli
    public CaptchaType selectCaptchaType(KaraPlayer player) {
        if (player.isBedrock()) {
            return CaptchaType.CLICK;  // Forms API uyumlu
        }
        return CaptchaType.MAP;  // Java için görsel
    }
}
```

### Anti-Bot Pattern Detection

```java
// Şüpheli davranış tespiti
public class BotDetector {
    
    private static final long MIN_JOIN_INTERVAL = 100;  // ms
    private static final int SUSPICIOUS_PACKET_RATE = 1000;  // packets/sec
    
    public BotScore analyze(Player player) {
        BotScore score = new BotScore();
        
        // 1. Bağlantı timing analizi
        if (getJoinInterval(player) < MIN_JOIN_INTERVAL) {
            score.add(30, "Too fast connection");
        }
        
        // 2. Packet rate analizi
        if (getPacketRate(player) > SUSPICIOUS_PACKET_RATE) {
            score.add(40, "Abnormal packet rate");
        }
        
        // 3. Movement pattern
        if (!hasNaturalMovement(player)) {
            score.add(20, "Unnatural movement");
        }
        
        // 4. Client brand check
        if (isSuspiciousBrand(player.getClientBrand())) {
            score.add(10, "Suspicious client");
        }
        
        return score;  // 70+ = kick, 90+ = temp ban
    }
}
```

---

## Exploit Koruması

### KaraPaper Exploit Patches

```java
// KaraPaper patch: Book exploit koruması
// patches/server/0010-Book-exploit-protection.patch

public class BookValidator {
    
    public static final int MAX_PAGES = 100;
    public static final int MAX_PAGE_LENGTH = 32767;
    public static final int MAX_TOTAL_SIZE = 200000;  // 200KB
    
    public static boolean validate(ItemStack book) {
        if (!(book.getItemMeta() instanceof BookMeta meta)) {
            return true;
        }
        
        // Sayfa sayısı kontrolü
        if (meta.getPageCount() > MAX_PAGES) {
            return false;
        }
        
        // Toplam boyut kontrolü
        int totalSize = 0;
        for (String page : meta.getPages()) {
            if (page.length() > MAX_PAGE_LENGTH) {
                return false;
            }
            totalSize += page.length();
        }
        
        return totalSize <= MAX_TOTAL_SIZE;
    }
}
```

### Packet Inspection

```java
// karapixel-security plugin
public class PacketInspector implements Listener {
    
    @EventHandler(priority = EventPriority.LOWEST)
    public void onBookEdit(PlayerEditBookEvent event) {
        if (!BookValidator.validate(event.getNewBookMeta())) {
            event.setCancelled(true);
            SecurityLogger.log(event.getPlayer(), ThreatType.BOOK_EXPLOIT);
            event.getPlayer().kick(text("§cGeçersiz kitap verisi."));
        }
    }
    
    @EventHandler(priority = EventPriority.LOWEST)
    public void onSignChange(SignChangeEvent event) {
        for (String line : event.getLines()) {
            if (line.length() > 100) {
                event.setCancelled(true);
                SecurityLogger.log(event.getPlayer(), ThreatType.SIGN_EXPLOIT);
                return;
            }
        }
    }
    
    @EventHandler(priority = EventPriority.LOWEST)
    public void onInventoryClick(InventoryClickEvent event) {
        ItemStack item = event.getCurrentItem();
        if (item != null && hasOversizedNBT(item)) {
            event.setCancelled(true);
            item.setAmount(0);
            SecurityLogger.log((Player) event.getWhoClicked(), ThreatType.NBT_EXPLOIT);
        }
    }
    
    private boolean hasOversizedNBT(ItemStack item) {
        try {
            ByteArrayOutputStream baos = new ByteArrayOutputStream();
            // NBT serialization kontrolü
            return baos.size() > 200000;  // 200KB limit
        } catch (Exception e) {
            return true;  // Hata = şüpheli
        }
    }
}
```

---

## Hesap Güvenliği

### Auth Configuration

```yaml
# karapixel-auth/config.yml

auth:
  # Şifre gereksinimleri
  password:
    min-length: 6
    max-length: 32
    require-number: false      # Türk kullanıcılar için basit tut
    require-special: false
    blocked-passwords:         # Yaygın şifreleri engelle
      - "123456"
      - "password"
      - "qwerty"
      - "abc123"
  
  # Hashing
  hashing:
    algorithm: BCRYPT
    bcrypt-cost: 12           # Güvenlik/performans dengesi
  
  # Login koruması
  login:
    max-attempts: 5           # 5 deneme hakkı
    lockout-duration: 300     # 5 dakika kilitleme
    attempt-reset-time: 300   # 5 dakika sonra sıfırla
    timeout: 30               # 30 saniye login süresi
    
  # Session
  session:
    enabled: true
    timeout: 604800           # 7 gün (saniye)
    bind-to-ip: false         # Mobil kullanıcılar için false
    
  # Captcha
  captcha:
    enabled: true
    type: AUTO                # Platform'a göre otomatik seç
    on-first-join: true
    on-suspicious-login: true
    on-ip-change: true
    
  # Anti-bot
  anti-bot:
    enabled: true
    max-registrations-per-ip: 3
    max-joins-per-second: 10
    temp-ban-duration: 3600   # 1 saat
```

### Session Management (Redis)

```java
// karapixel-auth session yapısı
public class SessionManager {
    
    private final RedisClient redis;
    private static final String SESSION_PREFIX = "session:";
    private static final long SESSION_TTL = 604800;  // 7 gün
    
    public void createSession(UUID playerId, String ip) {
        String sessionId = generateSecureToken();
        
        Session session = new Session(
            sessionId,
            playerId,
            ip,
            Instant.now(),
            Instant.now().plusSeconds(SESSION_TTL)
        );
        
        redis.setex(
            SESSION_PREFIX + playerId.toString(),
            SESSION_TTL,
            session.serialize()
        );
    }
    
    public boolean validateSession(UUID playerId, String currentIp) {
        String data = redis.get(SESSION_PREFIX + playerId.toString());
        if (data == null) return false;
        
        Session session = Session.deserialize(data);
        
        // Session süresi dolmuş mu?
        if (session.isExpired()) {
            redis.del(SESSION_PREFIX + playerId.toString());
            return false;
        }
        
        // IP değişmiş mi? (opsiyonel kontrol)
        // if (!session.getIp().equals(currentIp)) return false;
        
        return true;
    }
    
    public void invalidateSession(UUID playerId) {
        redis.del(SESSION_PREFIX + playerId.toString());
    }
}
```

---

## Veri Güvenliği

### SQL Injection Koruması

```java
// karapixel-database - DOĞRU KULLANIM
public class PlayerRepository {
    
    // ✅ DOĞRU: Prepared Statement
    public Optional<PlayerData> findByUuid(UUID uuid) {
        return database.query(
            "SELECT * FROM players WHERE uuid = ?",
            stmt -> stmt.setString(1, uuid.toString()),
            this::mapToPlayerData
        );
    }
    
    // ✅ DOĞRU: Prepared Statement
    public void updateBalance(UUID uuid, double amount) {
        database.execute(
            "UPDATE economy SET balance = ? WHERE uuid = ?",
            stmt -> {
                stmt.setDouble(1, amount);
                stmt.setString(2, uuid.toString());
            }
        );
    }
    
    // ❌ YANLIŞ: String concatenation (ASLA KULLANMA)
    // public void UNSAFE_update(String name, double amount) {
    //     database.execute("UPDATE economy SET balance = " + amount + 
    //                      " WHERE name = '" + name + "'");
    // }
}
```

### Database User Permissions

```sql
-- Her servis için ayrı, minimum yetkili kullanıcı

-- Auth servisi
CREATE USER 'kara_auth'@'localhost' IDENTIFIED BY 'strong_password_1';
GRANT SELECT, INSERT, UPDATE ON karapixel_db.players TO 'kara_auth'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON karapixel_db.sessions TO 'kara_auth'@'localhost';

-- Economy servisi
CREATE USER 'kara_economy'@'localhost' IDENTIFIED BY 'strong_password_2';
GRANT SELECT, UPDATE ON karapixel_db.economy TO 'kara_economy'@'localhost';
GRANT INSERT ON karapixel_db.transactions TO 'kara_economy'@'localhost';

-- Skyblock servisi
CREATE USER 'kara_skyblock'@'localhost' IDENTIFIED BY 'strong_password_3';
GRANT SELECT, INSERT, UPDATE, DELETE ON karapixel_db.islands TO 'kara_skyblock'@'localhost';
GRANT SELECT, INSERT, UPDATE ON karapixel_db.island_members TO 'kara_skyblock'@'localhost';

-- READ-ONLY backup kullanıcısı
CREATE USER 'kara_backup'@'localhost' IDENTIFIED BY 'strong_password_4';
GRANT SELECT, LOCK TABLES ON karapixel_db.* TO 'kara_backup'@'localhost';
```

---

## Bilgi Gizleme

### Plugin ve Versiyon Gizleme

```java
// karapixel-core/security/InfoHider.java
public class InfoHider implements Listener {
    
    // Engellenen komutlar
    private static final Set<String> BLOCKED_COMMANDS = Set.of(
        "plugins", "pl", "bukkit:plugins", "bukkit:pl",
        "version", "ver", "about", "bukkit:version", "bukkit:about",
        "icanhasbukkit", "?", "bukkit:?"
    );
    
    // Tab complete'den gizlenen komutlar
    private static final Set<String> HIDDEN_FROM_TAB = Set.of(
        "op", "deop", "ban", "ban-ip", "banlist", "pardon", "pardon-ip",
        "kick", "reload", "stop", "restart", "timings", "debug",
        "whitelist", "save-all", "save-off", "save-on"
    );
    
    @EventHandler(priority = EventPriority.LOWEST)
    public void onCommand(PlayerCommandPreprocessEvent event) {
        String command = event.getMessage().split(" ")[0].substring(1).toLowerCase();
        
        // Ana komut veya alias kontrolü
        String baseCommand = command.contains(":") 
            ? command.split(":")[1] 
            : command;
        
        if (BLOCKED_COMMANDS.contains(baseCommand) || 
            BLOCKED_COMMANDS.contains(command)) {
            event.setCancelled(true);
            event.getPlayer().sendMessage("§cBilinmeyen komut.");
            
            // Log şüpheli aktivite
            SecurityLogger.log(
                event.getPlayer(), 
                ThreatType.INFO_GATHERING,
                "Blocked command: " + command
            );
        }
    }
    
    @EventHandler
    public void onTabComplete(TabCompleteEvent event) {
        if (!(event.getSender() instanceof Player player)) return;
        
        // Admin değilse hassas komutları gösterme
        if (!player.hasPermission("karapixel.admin")) {
            event.getCompletions().removeIf(completion ->
                HIDDEN_FROM_TAB.contains(completion.toLowerCase()) ||
                completion.contains(":")  // Namespace gizle
            );
        }
    }
    
    @EventHandler
    public void onServerPing(PaperServerListPingEvent event) {
        // Server brand'i gizle
        event.setVersion("KaraPixel");
        
        // Gerçek oyuncu sayısını göster ama versiyon bilgisini değil
        // Protocol version -1 yaparsak "uyumsuz versiyon" gösterir
        // Bu bot scan'leri engelleyebilir (opsiyonel)
        // event.setProtocolVersion(-1);
    }
}
```

### Error Message Sanitization

```java
// Hata mesajlarında hassas bilgi gösterme
public class ErrorHandler {
    
    public static void handleError(Player player, Exception e) {
        // Oyuncuya genel mesaj
        player.sendMessage("§cBir hata oluştu. Lütfen daha sonra tekrar deneyin.");
        
        // Detayları sadece loglara yaz
        Logger.error("Error for player " + player.getName(), e);
    }
    
    // Stack trace'leri ASLA oyuncuya gösterme
    public static String sanitizeError(String message) {
        return message
            .replaceAll("at [\\w.]+\\([\\w.]+:\\d+\\)", "")  // Stack trace satırları
            .replaceAll("(?i)password|secret|key|token", "[REDACTED]")
            .replaceAll("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}", "[IP]");
    }
}
```

---

## Monitoring ve Alerting

### Security Event Logging

```java
// Güvenlik olaylarını kaydet
public class SecurityLogger {
    
    public enum ThreatType {
        BOT_DETECTED,
        BRUTE_FORCE,
        EXPLOIT_ATTEMPT,
        INFO_GATHERING,
        SUSPICIOUS_ACTIVITY,
        DUPE_ATTEMPT,
        PERMISSION_VIOLATION
    }
    
    public static void log(Player player, ThreatType type, String details) {
        SecurityEvent event = new SecurityEvent(
            Instant.now(),
            player.getUniqueId(),
            player.getName(),
            player.getAddress().getAddress().getHostAddress(),
            type,
            details,
            getCurrentServer()
        );
        
        // Dosyaya yaz
        FileLogger.logSecurity(event);
        
        // Veritabanına kaydet
        Database.insertSecurityEvent(event);
        
        // Kritik olaylar için Discord alert
        if (type.isCritical()) {
            DiscordWebhook.sendAlert(event);
        }
    }
}
```

### Discord Alert Webhook

```java
// Kritik güvenlik olayları için Discord bildirimi
public class DiscordWebhook {
    
    private static final String WEBHOOK_URL = System.getenv("DISCORD_SECURITY_WEBHOOK");
    
    public static void sendAlert(SecurityEvent event) {
        if (WEBHOOK_URL == null) return;
        
        JsonObject embed = new JsonObject();
        embed.addProperty("title", "⚠️ Güvenlik Uyarısı");
        embed.addProperty("color", event.getType().isCritical() ? 0xFF0000 : 0xFFA500);
        embed.addProperty("description", formatEvent(event));
        embed.addProperty("timestamp", event.getTimestamp().toString());
        
        // Async gönder
        CompletableFuture.runAsync(() -> {
            try {
                HttpClient.send(WEBHOOK_URL, embed);
            } catch (Exception e) {
                Logger.error("Discord webhook failed", e);
            }
        });
    }
    
    private static String formatEvent(SecurityEvent event) {
        return String.format("""
            **Tip:** %s
            **Oyuncu:** %s (%s)
            **IP:** ||%s||
            **Sunucu:** %s
            **Detay:** %s
            """,
            event.getType(),
            event.getPlayerName(),
            event.getPlayerId(),
            event.getIp(),
            event.getServer(),
            event.getDetails()
        );
    }
}
```

### Prometheus Metrics

```java
// Güvenlik metrikleri
public class SecurityMetrics {
    
    private static final Counter BLOCKED_CONNECTIONS = Counter.build()
        .name("karapixel_blocked_connections_total")
        .help("Total blocked connections")
        .labelNames("reason")
        .register();
    
    private static final Counter SECURITY_EVENTS = Counter.build()
        .name("karapixel_security_events_total")
        .help("Total security events")
        .labelNames("type", "server")
        .register();
    
    private static final Gauge ACTIVE_BANS = Gauge.build()
        .name("karapixel_active_bans")
        .help("Currently active bans")
        .labelNames("type")
        .register();
    
    public static void recordBlockedConnection(String reason) {
        BLOCKED_CONNECTIONS.labels(reason).inc();
    }
    
    public static void recordSecurityEvent(ThreatType type, String server) {
        SECURITY_EVENTS.labels(type.name(), server).inc();
    }
}
```

---

## Incident Response

### Saldırı Anında Yapılacaklar

```
┌─────────────────────────────────────────────────────────────────┐
│                 INCIDENT RESPONSE PLANI                         │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  SEVİYE 1: KÜÇÜK OLAY (Bot spam, tekil exploit denemesi)       │
│  ════════════════════════════════════════════════════════      │
│  1. Otomatik sistemler handle etmeli (Limbo, rate limit)       │
│  2. Log'ları kontrol et                                        │
│  3. Gerekirse IP'yi manuel banla                               │
│  4. Devam et                                                    │
│                                                                 │
│  SEVİYE 2: ORTA OLAY (DDoS başlangıcı, aktif exploit)         │
│  ════════════════════════════════════════════════════════      │
│  1. TCPShield/Cosmic Guard'ın çalıştığını doğrula             │
│  2. Rate limitleri geçici olarak sıkılaştır                   │
│  3. Şüpheli IP aralıklarını engelle                           │
│  4. Discord'da ekibi bilgilendir                              │
│  5. Log'ları analiz et, pattern bul                           │
│                                                                 │
│  SEVİYE 3: BÜYÜK OLAY (Ciddi DDoS, veri sızıntısı şüphesi)    │
│  ════════════════════════════════════════════════════════      │
│  1. Sunucuyu maintenance mode'a al                             │
│  2. TCPShield support'a ulaş                                   │
│  3. Tüm session'ları invalidate et                            │
│  4. Backup'lardan data bütünlüğünü kontrol et                 │
│  5. Post-mortem analizi başlat                                │
│  6. Gerekirse kullanıcıları bilgilendir                       │
│                                                                 │
│  SEVİYE 4: KRİTİK OLAY (Sunucu ele geçirme, veri kaybı)       │
│  ════════════════════════════════════════════════════════      │
│  1. SUNUCUYU KAPAT (fiziksel/network izolasyon)               │
│  2. Hetzner support'a ulaş                                     │
│  3. Forensic için snapshot al                                  │
│  4. Tüm credential'ları değiştir                              │
│  5. Clean install + backup'tan restore                         │
│  6. Güvenlik audit'i yap                                       │
│  7. Kullanıcıları şifre değiştirmeye zorla                    │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Acil Komutlar

```bash
# Sunucuyu maintenance mode'a al
/maintenance on

# Tüm bağlantıları kes (Velocity)
/velocity kick-all "Bakım modu aktif"

# Belirli IP aralığını engelle
sudo iptables -A INPUT -s 192.168.1.0/24 -j DROP

# Tüm session'ları temizle (Redis)
redis-cli KEYS "session:*" | xargs redis-cli DEL

# Emergency backup
/opt/karapixel/scripts/emergency-backup.sh

# Sunucuları durdur
/opt/karapixel/scripts/stop-all.sh
```

---

## Güvenlik Checklist

### Günlük
- [ ] Security log'larını kontrol et
- [ ] Anormal trafik var mı?
- [ ] Failed login spike var mı?

### Haftalık
- [ ] Firewall kurallarını gözden geçir
- [ ] Fail2Ban ban listesini kontrol et
- [ ] Backup'ların başarılı olduğunu doğrula

### Aylık
- [ ] Tüm şifreleri rotate et (DB, Redis, admin)
- [ ] Güvenlik patch'lerini uygula
- [ ] Penetration test yap (opsiyonel)
- [ ] İzin yapısını audit et

### Güncelleme Sonrası
- [ ] Yeni exploitler için kontrol et
- [ ] Plugin uyumluluğunu test et
- [ ] Güvenlik config'lerini doğrula

---

*📅 Son güncelleme: 24 Aralık 2024*
*⚠️ Bu doküman gizli tutulmalı ve sadece yetkili kişilerle paylaşılmalıdır.*
