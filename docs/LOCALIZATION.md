# 🌍 KaraPixel - Dil Desteği Sistemi

> Varsayılan %100 Türkçe, çoklu dil desteği altyapısı hazır.

---

## Genel Bakış

```
┌─────────────────────────────────────────────────────────────────┐
│                    DİL DESTEĞİ SİSTEMİ                          │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  DESTEKLENEN DİLLER:                                           │
│  ├── 🇹🇷 Türkçe (tr_TR) ← VARSAYILAN                           │
│  ├── 🇬🇧 İngilizce (en_US)                                     │
│  ├── 🇩🇪 Almanca (de_DE) - opsiyonel                           │
│  └── 🌐 Özel diller eklenebilir                                │
│                                                                 │
│  PRENSİPLER:                                                    │
│  ├── Tüm mesajlar externalized (hardcoded string YOK)          │
│  ├── Oyuncu bazlı dil seçimi                                   │
│  ├── Platform dil tespiti (Java/Bedrock)                       │
│  ├── Hot-reload desteği (restart gerektirmez)                  │
│  └── Placeholder sistemi (dinamik değerler)                    │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Dosya Yapısı

```
locales/
├── tr_TR.yml          # Türkçe (varsayılan, tam)
├── en_US.yml          # İngilizce (tam)
├── de_DE.yml          # Almanca (opsiyonel)
│
├── modules/           # Modül bazlı dil dosyaları (büyük projeler için)
│   ├── auth/
│   │   ├── tr_TR.yml
│   │   └── en_US.yml
│   ├── skyblock/
│   │   ├── tr_TR.yml
│   │   └── en_US.yml
│   └── economy/
│       ├── tr_TR.yml
│       └── en_US.yml
│
└── overrides/         # Sunucu özel override'lar
    └── tr_TR.yml      # Sadece değiştirmek istenen key'ler
```

---

## Dil Dosyası Formatı

```yaml
# locales/tr_TR.yml

# Meta bilgiler
meta:
  language: "Türkçe"
  code: "tr_TR"
  version: "1.0.0"
  authors: 
    - "KaraPixel Team"
  
# ═══════════════════════════════════════════════════════════════
# GENEL MESAJLAR
# ═══════════════════════════════════════════════════════════════
general:
  prefix: "<gradient:#FF6B6B:#4ECDC4>[KaraPixel]</gradient> "
  
  # Durum mesajları
  success: "<green>✓ Başarılı!</green>"
  error: "<red>✗ Hata: {message}</red>"
  warning: "<yellow>⚠ Uyarı: {message}</yellow>"
  info: "<aqua>ℹ {message}</aqua>"
  
  # Yetki mesajları
  no_permission: "<red>Bu işlem için yetkiniz yok!</red>"
  player_only: "<red>Bu komut sadece oyuncular tarafından kullanılabilir!</red>"
  console_only: "<red>Bu komut sadece konsoldan kullanılabilir!</red>"
  
  # Oyuncu mesajları
  player_not_found: "<red>Oyuncu bulunamadı: {player}</red>"
  player_offline: "<red>{player} şu anda çevrimdışı.</red>"
  
  # Onay/İptal
  confirm: "<green>Onayla</green>"
  cancel: "<red>İptal</red>"
  yes: "<green>Evet</green>"
  no: "<red>Hayır</red>"
  back: "<gray>« Geri</gray>"
  close: "<red>✕ Kapat</red>"
  
  # Zaman formatları
  time:
    seconds: "{value} saniye"
    minutes: "{value} dakika"
    hours: "{value} saat"
    days: "{value} gün"
    now: "şimdi"
    ago: "{time} önce"
    remaining: "{time} kaldı"

# ═══════════════════════════════════════════════════════════════
# AUTH SİSTEMİ
# ═══════════════════════════════════════════════════════════════
auth:
  # Karşılama
  welcome: |
    <yellow>═══════════════════════════════════════════════════</yellow>
    <gold><bold>KaraPixel</bold></gold> <gray>sunucusuna hoş geldin!</gray>
    <yellow>═══════════════════════════════════════════════════</yellow>
    
  # Kayıt
  register:
    prompt: "<yellow>Kayıt olmak için:</yellow> <white>/register <şifre> <şifre></white>"
    success: "<green>Başarıyla kayıt oldun!</green>"
    password_mismatch: "<red>Şifreler eşleşmiyor!</red>"
    password_too_short: "<red>Şifre en az {min} karakter olmalı!</red>"
    password_too_long: "<red>Şifre en fazla {max} karakter olabilir!</red>"
    already_registered: "<red>Zaten kayıtlısın! Giriş yapmak için: /login <şifre></red>"
    
  # Giriş
  login:
    prompt: "<yellow>Giriş yapmak için:</yellow> <white>/login <şifre></white>"
    success: "<green>Giriş başarılı! Hoş geldin, <white>{player}</white>!</green>"
    wrong_password: "<red>Yanlış şifre! Kalan hak: <white>{remaining}</white></red>"
    not_registered: "<red>Kayıtlı değilsin! Kayıt olmak için: /register <şifre> <şifre></red>"
    
  # Güvenlik
  too_many_attempts: "<red>Çok fazla başarısız deneme! <white>{duration}</white> beklemelisin.</red>"
  timeout: "<red>Giriş süresi doldu! Lütfen tekrar bağlan.</red>"
  session_expired: "<yellow>Oturumun sona erdi. Lütfen tekrar giriş yap.</yellow>"
  
  # Bedrock
  auto_login: "<green>Xbox hesabın ile otomatik giriş yapıldı!</green>"
  
  # Captcha
  captcha:
    prompt: "<yellow>Lütfen captcha'yı çöz:</yellow>"
    success: "<green>Captcha doğrulandı!</green>"
    failed: "<red>Yanlış cevap! Tekrar dene.</red>"

# ═══════════════════════════════════════════════════════════════
# HUB LOBİSİ
# ═══════════════════════════════════════════════════════════════
hub:
  # Genel
  welcome: "<green>Hub'a hoş geldin!</green>"
  teleporting: "<yellow>Hub'a ışınlanıyorsun...</yellow>"
  
  # Items
  items:
    server_selector: "<green>Oyun Seçici</green>"
    server_selector_lore:
      - "<gray>Bir oyun modu seç ve oyna!"
      - ""
      - "<yellow>▸ Sağ tıkla!</yellow>"
      
    cosmetics: "<light_purple>Kozmetikler</light_purple>"
    cosmetics_lore:
      - "<gray>Görünümünü özelleştir!"
      
    profile: "<aqua>Profil</aqua>"
    profile_lore:
      - "<gray>İstatistiklerini görüntüle"

# ═══════════════════════════════════════════════════════════════
# OYUN SEÇİCİ
# ═══════════════════════════════════════════════════════════════
selector:
  menu:
    title: "Oyun Seçici"
    
  games:
    skyblock:
      name: "<green>⛏ Skyblock</green>"
      lore:
        - "<gray>Kendi adanı oluştur ve geliştir!"
        - "<gray>Arkadaşlarınla birlikte oyna!"
        - ""
        - "<dark_gray>Çevrimiçi: <white>{online}</white></dark_gray>"
        - ""
        - "<yellow>▸ Tıkla ve başla!</yellow>"
        
    coming_soon:
      name: "<gray>🔒 Yakında...</gray>"
      lore:
        - "<gray>Bu mod henüz aktif değil."
        - "<gray>Takipte kal!"

# ═══════════════════════════════════════════════════════════════
# SKYBLOCK
# ═══════════════════════════════════════════════════════════════
skyblock:
  # Ada oluşturma
  island:
    creating: "<yellow>Adan oluşturuluyor...</yellow>"
    created: "<green>Adan başarıyla oluşturuldu!</green>"
    teleporting: "<yellow>Adana ışınlanıyorsun...</yellow>"
    no_island: "<red>Henüz bir adan yok! <white>/is create</white> ile oluştur.</red>"
    already_has: "<red>Zaten bir adan var!</red>"
    
  # Ada bilgileri
  info:
    level: "<gold>Ada Seviyesi: <yellow>{level}</yellow></gold>"
    members: "<aqua>Üyeler: <white>{count}/{max}</white></aqua>"
    created: "<gray>Oluşturulma: {date}</gray>"
    
  # Şablonlar
  templates:
    title: "Ada Şablonu Seç"
    normal:
      name: "<green>🌳 Normal Ada</green>"
      lore:
        - "<gray>Klasik skyblock deneyimi"
        - "<gray>Bir ağaç ve temel malzemeler"
    desert:
      name: "<yellow>🏜️ Çöl Adası</yellow>"
      lore:
        - "<gray>Kumdan bir ada"
        - "<gray>Kaktüs ve hurma ağacı"
    nether:
      name: "<red>🔥 Nether Adası</red>"
      lore:
        - "<gray>Cehennem temalı ada"
        - "<gray>Seviye {level} gerektirir"
        
  # Coop
  coop:
    invited: "<yellow>{player} seni adasına davet etti!</yellow>"
    invite_sent: "<green>{player} adresine davet gönderildi.</green>"
    joined: "<green>{player} adana katıldı!</green>"
    left: "<red>{player} adandan ayrıldı.</red>"
    kicked: "<red>{player} adadan atıldı.</red>"
    promoted: "<green>{player} {role} rolüne yükseltildi.</green>"
    demoted: "<yellow>{player} {role} rolüne düşürüldü.</yellow>"
    
  # Roller
  roles:
    owner: "<gold>Ada Sahibi</gold>"
    admin: "<red>Yönetici</red>"
    member: "<green>Üye</green>"
    visitor: "<gray>Ziyaretçi</gray>"
    
  # Generator
  generator:
    placed: "<green>{generator} yerleştirildi!</green>"
    removed: "<yellow>{generator} kaldırıldı.</yellow>"
    upgraded: "<green>{generator} seviye <white>{level}</white> oldu!</green>"
    max_level: "<yellow>Bu generator maksimum seviyede!</yellow>"
    limit_reached: "<red>Generator limitine ulaştın! ({count}/{max})</red>"
    
  # Yükseltmeler
  upgrades:
    purchased: "<green>{upgrade} satın alındı!</green>"
    not_enough_money: "<red>Yeterli paran yok! Gereken: <white>{amount}</white></red>"
    already_max: "<yellow>Bu yükseltme zaten maksimum seviyede!</yellow>"

# ═══════════════════════════════════════════════════════════════
# EKONOMİ
# ═══════════════════════════════════════════════════════════════
economy:
  currency:
    name: "Kara Coin"
    symbol: "⛃"
    format: "{symbol}{amount}"
    
  balance: "<gold>Bakiyen: <white>{balance}</white></gold>"
  balance_other: "<gold>{player} bakiyesi: <white>{balance}</white></gold>"
  
  # Transfer
  transfer:
    success: "<green>{amount} {player} adresine gönderildi.</green>"
    received: "<green>{player} sana {amount} gönderdi!</green>"
    not_enough: "<red>Yeterli bakiyen yok!</red>"
    self: "<red>Kendine para gönderemezsin!</red>"
    minimum: "<red>Minimum transfer miktarı: {amount}</red>"

# ═══════════════════════════════════════════════════════════════
# SKILL SİSTEMİ
# ═══════════════════════════════════════════════════════════════
skills:
  # Skill isimleri
  names:
    mining: "Madencilik"
    farming: "Çiftçilik"
    combat: "Savaş"
    fishing: "Balıkçılık"
    foraging: "Ormancılık"
    enchanting: "Büyüleme"
    
  # Seviye mesajları
  level_up: |
    <gold>═══════════════════════════════════════</gold>
    <yellow>  ⬆ SEVİYE ATLADIN!</yellow>
    <gold>═══════════════════════════════════════</gold>
    <gray>  {skill}: Seviye <white>{level}</white></gray>
    
  xp_gain: "<gray>+{xp} {skill} XP</gray>"
  
  # Menü
  menu:
    title: "Yetenekler"
    progress: "<gray>İlerleme: <white>{current}/{next}</white></gray>"

# ═══════════════════════════════════════════════════════════════
# MENÜ BAŞLIKLARI
# ═══════════════════════════════════════════════════════════════
menu:
  island:
    title: "Ada Menüsü"
    home: "<green>🏠 Eve Git</green>"
    members: "<aqua>👥 Üyeler</aqua>"
    settings: "<gold>⚙ Ayarlar</gold>"
    upgrades: "<light_purple>⬆ Yükseltmeler</light_purple>"
    bank: "<yellow>💰 Banka</yellow>"
    missions: "<blue>📋 Görevler</blue>"
    
  cosmetics:
    title: "Kozmetikler"
    particles: "<light_purple>✨ Partiküller</light_purple>"
    wings: "<aqua>🪽 Kanatlar</aqua>"
    pets: "<green>🐾 Evcil Hayvanlar</green>"
    hats: "<gold>👑 Şapkalar</gold>"
    
  shop:
    title: "Mağaza"
    buy: "<green>Satın Al</green>"
    sell: "<red>Sat</red>"
    price: "<gold>Fiyat: {price}</gold>"
    
# ═══════════════════════════════════════════════════════════════
# RANK SİSTEMİ
# ═══════════════════════════════════════════════════════════════
ranks:
  names:
    default: "<gray>Oyuncu</gray>"
    vip: "<green>VIP</green>"
    vip_plus: "<green>VIP<gold>+</gold></green>"
    mvp: "<aqua>MVP</aqua>"
    mvp_plus: "<aqua>MVP<red>+</red></aqua>"
    admin: "<red>Admin</red>"
    
  # Prefix'ler
  prefix:
    default: "<gray>[Oyuncu]</gray>"
    vip: "<green>[VIP]</green>"
    vip_plus: "<green>[VIP<gold>+</gold>]</green>"
    mvp: "<aqua>[MVP]</aqua>"
    mvp_plus: "<aqua>[MVP<red>+</red>]</aqua>"
    admin: "<red>[Admin]</red>"

# ═══════════════════════════════════════════════════════════════
# CHAT
# ═══════════════════════════════════════════════════════════════
chat:
  format: "{prefix} {player}<gray>:</gray> <white>{message}</white>"
  
  # Chat komutları
  msg:
    send: "<gray>[<white>Sen</white> → <white>{player}</white>]</gray> {message}"
    receive: "<gray>[<white>{player}</white> → <white>Sen</white>]</gray> {message}"
    
  reply:
    no_target: "<red>Yanıtlayacak kimse yok!</red>"
    
  # Global chat
  global:
    prefix: "<red>[G]</red>"
    
  # Party chat
  party:
    prefix: "<blue>[Parti]</blue>"

# ═══════════════════════════════════════════════════════════════
# MODERASYON
# ═══════════════════════════════════════════════════════════════
moderation:
  ban:
    permanent: "<red>Sunucudan yasaklandın!</red>\n<gray>Sebep: {reason}</gray>"
    temporary: "<red>Sunucudan geçici olarak yasaklandın!</red>\n<gray>Sebep: {reason}</gray>\n<gray>Kalan süre: {duration}</gray>"
    broadcast: "<red>{player} sunucudan yasaklandı. Sebep: {reason}</red>"
    
  kick:
    message: "<red>Sunucudan atıldın!</red>\n<gray>Sebep: {reason}</gray>"
    broadcast: "<yellow>{player} sunucudan atıldı.</yellow>"
    
  mute:
    message: "<red>Susturuldun! Kalan süre: {duration}</red>"
    broadcast: "<yellow>{player} susturuldu.</yellow>"
    try_chat: "<red>Susturuldun! Kalan süre: {duration}</red>"
    
  warn:
    message: "<yellow>Uyarıldın! Sebep: {reason}</yellow>"
    broadcast: "<yellow>{player} uyarıldı. ({count}/3)</yellow>"
```

---

## Lokalizasyon API

### LocaleManager Sınıfı

```java
public class LocaleManager {
    private final Plugin plugin;
    private final Map<String, YamlConfiguration> locales = new HashMap<>();
    private String defaultLocale = "tr_TR";
    
    public void initialize() {
        // Varsayılan dilleri yükle
        loadLocale("tr_TR");
        loadLocale("en_US");
        
        // Override'ları uygula
        loadOverrides();
        
        logger.info("Loaded {} locales", locales.size());
    }
    
    private void loadLocale(String code) {
        InputStream is = plugin.getResource("locales/" + code + ".yml");
        if (is == null) {
            logger.warn("Locale not found: {}", code);
            return;
        }
        
        YamlConfiguration config = YamlConfiguration.loadConfiguration(
            new InputStreamReader(is, StandardCharsets.UTF_8)
        );
        locales.put(code, config);
    }
    
    public String get(String locale, String key) {
        YamlConfiguration config = locales.get(locale);
        if (config == null) {
            config = locales.get(defaultLocale);
        }
        
        String value = config.getString(key);
        if (value == null && !locale.equals(defaultLocale)) {
            // Fallback to default
            value = locales.get(defaultLocale).getString(key);
        }
        
        return value != null ? value : key;
    }
    
    public String get(KaraPlayer player, String key, Placeholder... placeholders) {
        String message = get(player.getLocale(), key);
        
        // Placeholder'ları uygula
        for (Placeholder ph : placeholders) {
            message = message.replace("{" + ph.key() + "}", ph.value());
        }
        
        // Prefix ekle (eğer varsa)
        if (!key.startsWith("general.") && !message.startsWith("<")) {
            String prefix = get(player.getLocale(), "general.prefix");
            message = prefix + message;
        }
        
        return message;
    }
    
    // Hot reload
    public void reload() {
        locales.clear();
        initialize();
    }
}
```

### Placeholder Sistemi

```java
public record Placeholder(String key, String value) {
    
    public static Placeholder of(String key, Object value) {
        return new Placeholder(key, String.valueOf(value));
    }
    
    public static Placeholder of(String key, int value) {
        return new Placeholder(key, formatNumber(value));
    }
    
    public static Placeholder of(String key, double value) {
        return new Placeholder(key, formatDecimal(value));
    }
    
    public static Placeholder of(String key, Duration duration) {
        return new Placeholder(key, formatDuration(duration));
    }
    
    private static String formatNumber(int value) {
        return NumberFormat.getIntegerInstance().format(value);
    }
    
    private static String formatDecimal(double value) {
        return String.format("%.2f", value);
    }
    
    private static String formatDuration(Duration duration) {
        // Dinamik süre formatlama
        if (duration.toSeconds() < 60) {
            return duration.toSeconds() + " saniye";
        } else if (duration.toMinutes() < 60) {
            return duration.toMinutes() + " dakika";
        } else if (duration.toHours() < 24) {
            return duration.toHours() + " saat";
        } else {
            return duration.toDays() + " gün";
        }
    }
}
```

### KaraPlayer Entegrasyonu

```java
public class KaraPlayer {
    private String locale;
    
    // Dil tespiti
    public void detectLocale() {
        if (isBedrock()) {
            // Bedrock: Geyser'dan dil bilgisi
            locale = GeyserHook.getPlayerLocale(uuid);
        } else {
            // Java: Client locale
            locale = bukkitPlayer.getLocale();
        }
        
        // Desteklenen dile map'le
        locale = mapToSupportedLocale(locale);
    }
    
    private String mapToSupportedLocale(String clientLocale) {
        // tr, tr_TR, tr_tr -> tr_TR
        if (clientLocale.toLowerCase().startsWith("tr")) {
            return "tr_TR";
        }
        // en, en_US, en_GB -> en_US
        if (clientLocale.toLowerCase().startsWith("en")) {
            return "en_US";
        }
        // Desteklenmeyen dil -> varsayılan
        return "tr_TR";
    }
    
    // Kolay erişim metodları
    public String t(String key, Placeholder... placeholders) {
        return KaraAPI.getLocaleManager().get(this, key, placeholders);
    }
    
    public void sendMessage(String key, Placeholder... placeholders) {
        String message = t(key, placeholders);
        bukkitPlayer.sendMessage(Text.parse(message));
    }
    
    public void sendActionBar(String key, Placeholder... placeholders) {
        String message = t(key, placeholders);
        bukkitPlayer.sendActionBar(Text.parse(message));
    }
    
    public void sendTitle(String titleKey, String subtitleKey, Placeholder... placeholders) {
        String title = t(titleKey, placeholders);
        String subtitle = t(subtitleKey, placeholders);
        bukkitPlayer.showTitle(Title.title(
            Text.parse(title),
            Text.parse(subtitle)
        ));
    }
}
```

---

## Kullanım Örnekleri

```java
// Basit mesaj
player.sendMessage("auth.login.success", 
    Placeholder.of("player", player.getName())
);

// Action bar ile XP gösterimi
player.sendActionBar("skills.xp_gain",
    Placeholder.of("xp", 50),
    Placeholder.of("skill", player.t("skills.names.mining"))
);

// Kompleks mesaj
player.sendMessage("economy.transfer.success",
    Placeholder.of("amount", economy.format(1000)),
    Placeholder.of("player", targetPlayer.getName())
);

// Menü başlığı
String title = player.t("menu.island.title");

// Liste içinde lokalize item
MenuItem.builder()
    .name(player.t("menu.island.home"))
    .lore(player.t("menu.island.home_lore").split("\n"))
    .build();
```

---

## Dil Ekleme Rehberi

### Yeni Dil Ekleme

1. `locales/` klasörüne `{lang_code}.yml` dosyası ekle
2. `tr_TR.yml` dosyasını kopyala ve çevir
3. `LocaleManager`'a dili ekle

```java
// LocaleManager'a yeni dil ekleme
public void initialize() {
    loadLocale("tr_TR");
    loadLocale("en_US");
    loadLocale("de_DE");  // Yeni dil
}
```

### Çeviri Kuralları

1. **Placeholder'ları değiştirme:** `{player}`, `{amount}` gibi placeholder'lar aynen kalmalı
2. **MiniMessage taglerini koru:** `<green>`, `<bold>` gibi formatlar
3. **Satır sonlarını koru:** Multi-line mesajlarda `|` kullan
4. **Emoji/Unicode:** Platform uyumluluğunu kontrol et

---

*📅 Son güncelleme: 24 Aralık 2024*
