# 🔌 KaraPixel - Plugin Dokümantasyonu

> Tüm pluginler %100 custom yazılacak, tam Türkçe, mobil uyumlu ve 3D model destekli.

---

## 📋 İçindekiler

1. [Plugin Mimarisi](#plugin-mimarisi)
2. [Core Plugins](#core-plugins)
3. [Auth Plugins](#auth-plugins)
4. [Hub Plugins](#hub-plugins)
5. [Skyblock Plugins](#skyblock-plugins)
6. [Global Plugins](#global-plugins)
7. [Admin Plugins](#admin-plugins)
8. [Dil Desteği Altyapısı](#dil-desteği-altyapısı)
9. [Mobil Uyumluluk](#mobil-uyumluluk)
10. [3D Model Entegrasyonu](#3d-model-entegrasyonu)

---

## Plugin Mimarisi

### Bağımlılık Grafiği

```
┌─────────────────────────────────────────────────────────────────┐
│                    PLUGIN MİMARİSİ                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│                    ┌─────────────────────┐                      │
│                    │   karapixel-core    │                      │
│                    │   (Merkezi Temel)   │                      │
│                    └──────────┬──────────┘                      │
│                               │                                  │
│         ┌─────────────────────┼─────────────────────┐           │
│         │                     │                     │           │
│         ▼                     ▼                     ▼           │
│  ┌─────────────┐      ┌─────────────┐      ┌─────────────┐     │
│  │  karapixel  │      │  karapixel  │      │  karapixel  │     │
│  │  -database  │      │ -messaging  │      │    -ui      │     │
│  │  (HikariCP) │      │ (Redis PS)  │      │ (Menu API)  │     │
│  └──────┬──────┘      └──────┬──────┘      └──────┬──────┘     │
│         │                    │                    │             │
│         └────────────────────┼────────────────────┘             │
│                              │                                  │
│                              ▼                                  │
│         ┌────────────────────────────────────────────┐         │
│         │            FEATURE PLUGINS                  │         │
│         │   (Tüm pluginler core modüllere bağımlı)   │         │
│         └────────────────────────────────────────────┘         │
│                              │                                  │
│    ┌───────┬───────┬────────┼────────┬───────┬───────┐        │
│    ▼       ▼       ▼        ▼        ▼       ▼       ▼        │
│ ┌──────┐┌──────┐┌──────┐┌──────┐┌──────┐┌──────┐┌──────┐     │
│ │ auth ││ hub  ││select││skyblk││econmy││ranks ││ chat │     │
│ └──────┘└──────┘└──────┘└──────┘└──────┘└──────┘└──────┘     │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Plugin Dağılımı (Sunucu Başına)

| Plugin | Limbo | Hub | SB-Spawn | SB-World | Velocity |
|--------|:-----:|:---:|:--------:|:--------:|:--------:|
| karapixel-core | ✅ | ✅ | ✅ | ✅ | - |
| karapixel-ui | ✅ | ✅ | ✅ | ✅ | - |
| karapixel-auth | ✅ | - | - | - | - |
| karapixel-hub | - | ✅ | - | - | - |
| karapixel-selector | - | ✅ | - | - | - |
| karapixel-skyblock | - | - | ✅ | ✅ | - |
| karapixel-generators | - | - | - | ✅ | - |
| karapixel-skills | - | - | - | ✅ | - |
| karapixel-quests | - | - | ✅ | ✅ | - |
| karapixel-economy | - | ✅ | ✅ | ✅ | - |
| karapixel-shop | - | - | ✅ | ✅ | - |
| karapixel-cosmetics | - | ✅ | ✅ | ✅ | - |
| karapixel-pets | - | ✅ | ✅ | ✅ | - |
| karapixel-chat | - | ✅ | ✅ | ✅ | - |
| karapixel-tablist | - | ✅ | ✅ | ✅ | - |
| karapixel-ranks | - | ✅ | ✅ | ✅ | - |
| karapixel-moderation | - | ✅ | ✅ | ✅ | - |
| karapixel-velocity | - | - | - | - | ✅ |

---

## Core Plugins

### 1. karapixel-core

**Ana merkezi kütüphane - Tüm pluginlerin temeli**

```
karapixel-core/
├── src/main/java/net/karapixel/core/
│   ├── KaraCore.java                 # Ana plugin sınıfı
│   │
│   ├── api/                          # Public API
│   │   ├── KaraAPI.java              # Singleton erişim noktası
│   │   ├── player/
│   │   │   ├── KaraPlayer.java       # Platform-agnostic oyuncu wrapper
│   │   │   ├── PlayerManager.java    # Oyuncu yönetimi
│   │   │   └── PlayerPlatform.java   # JAVA, BEDROCK enum
│   │   ├── server/
│   │   │   ├── ServerType.java       # LIMBO, HUB, SKYBLOCK_SPAWN, SKYBLOCK_WORLD
│   │   │   └── ServerInfo.java       # Sunucu bilgisi
│   │   └── event/
│   │       ├── KaraEvent.java        # Base event
│   │       └── EventBus.java         # Lightweight event bus
│   │
│   ├── locale/                       # 🌍 DİL DESTEĞİ SİSTEMİ
│   │   ├── LocaleManager.java        # Dil yönetimi
│   │   ├── Message.java              # Mesaj enum'ları
│   │   ├── MessageProvider.java      # Mesaj sağlayıcı
│   │   └── PlaceholderResolver.java  # {player}, {amount} vb.
│   │
│   ├── platform/                     # 📱 PLATFORM TESPİTİ
│   │   ├── PlatformDetector.java     # Java/Bedrock tespit
│   │   ├── BedrockUtil.java          # Bedrock-specific utilities
│   │   └── GeyserHook.java           # Geyser entegrasyonu
│   │
│   ├── config/                       # Konfigürasyon
│   │   ├── ConfigManager.java
│   │   ├── YamlConfig.java
│   │   └── ConfigReloader.java       # Hot reload desteği
│   │
│   └── util/                         # Yardımcı sınıflar
│       ├── Text.java                 # MiniMessage wrapper
│       ├── TimeUtil.java             # Süre formatlama
│       ├── ItemBuilder.java          # Item oluşturma
│       ├── LocationUtil.java         # Lokasyon işlemleri
│       └── ColorUtil.java            # Renk dönüşümleri
│
└── src/main/resources/
    ├── plugin.yml
    └── locales/                      # 🌍 DİL DOSYALARI
        ├── tr_TR.yml                 # Türkçe (varsayılan)
        ├── en_US.yml                 # İngilizce
        └── de_DE.yml                 # Almanca (opsiyonel)
```

**Örnek Kullanım:**
```java
// Platform tespiti
KaraPlayer player = KaraAPI.getPlayer(bukkitPlayer);
if (player.isBedrock()) {
    // Bedrock-specific işlem
}

// Dil desteği
String message = player.getMessage(Message.WELCOME);
player.sendMessage(message);

// Placeholder ile
String balanceMsg = player.getMessage(Message.BALANCE, 
    Placeholder.of("amount", economy.getBalance(player))
);
```

---

### 2. karapixel-database

**Veritabanı yönetimi ve repository pattern**

```java
// Merkezi database bağlantısı
public class DatabaseManager {
    private HikariDataSource dataSource;
    
    // Connection pool ayarları
    public void initialize() {
        HikariConfig config = new HikariConfig();
        config.setJdbcUrl("jdbc:mysql://localhost:3306/karapixel_db");
        config.setUsername("karapixel");
        config.setPassword(System.getenv("DB_PASSWORD"));
        config.setMaximumPoolSize(20);
        config.setMinimumIdle(5);
        config.setConnectionTimeout(30000);
        config.setIdleTimeout(600000);
        config.setMaxLifetime(1800000);
        
        // Performance optimizations
        config.addDataSourceProperty("cachePrepStmts", "true");
        config.addDataSourceProperty("prepStmtCacheSize", "250");
        config.addDataSourceProperty("prepStmtCacheSqlLimit", "2048");
        
        dataSource = new HikariDataSource(config);
    }
}

// Repository pattern örneği
public interface PlayerRepository {
    Optional<PlayerData> findByUuid(UUID uuid);
    void save(PlayerData player);
    void updateLastLogin(UUID uuid);
}

public class MySQLPlayerRepository implements PlayerRepository {
    @Override
    public Optional<PlayerData> findByUuid(UUID uuid) {
        return database.query(
            "SELECT * FROM players WHERE uuid = ?",
            stmt -> stmt.setString(1, uuid.toString()),
            this::mapToPlayerData
        );
    }
}
```

---

### 3. karapixel-messaging

**Redis Pub/Sub ile cross-server iletişim**

```java
// Mesaj gönderme
public class MessageBroker {
    private final RedisClient redis;
    
    public void publish(String channel, Packet packet) {
        redis.publish("karapixel:" + channel, packet.serialize());
    }
    
    public void subscribe(String channel, PacketHandler handler) {
        redis.subscribe("karapixel:" + channel, data -> {
            Packet packet = Packet.deserialize(data);
            handler.handle(packet);
        });
    }
}

// Packet tipleri
public abstract class Packet {
    public abstract String getType();
    public abstract String serialize();
    public static Packet deserialize(String data);
}

public class TeleportRequestPacket extends Packet {
    private UUID playerId;
    private String targetServer;
    private String targetType;  // "skyblock", "hub", etc.
}

public class PlayerDataSyncPacket extends Packet {
    private UUID playerId;
    private Map<String, Object> data;
}
```

---

### 4. karapixel-ui

**Mobil-first UI framework**

```java
// Platform-aware menu sistemi
public class KaraMenu {
    private final String titleKey;  // Lokalizasyon key
    private final List<MenuItem> items;
    private final int rows;
    
    public void open(KaraPlayer player) {
        String title = player.getMessage(titleKey);
        
        if (player.isBedrock()) {
            // Bedrock: Geyser Forms API
            openBedrockForm(player, title);
        } else {
            // Java: Chest inventory
            openJavaInventory(player, title);
        }
    }
    
    private void openBedrockForm(KaraPlayer player, String title) {
        // Geyser Cumulus API kullanımı
        SimpleForm form = SimpleForm.builder()
            .title(title)
            .content(getDescription(player));
        
        for (MenuItem item : items) {
            form.button(item.getName(player), item.getFormImage());
        }
        
        GeyserHook.sendForm(player, form.build());
    }
    
    private void openJavaInventory(KaraPlayer player, String title) {
        Inventory inv = Bukkit.createInventory(null, rows * 9, title);
        
        for (MenuItem item : items) {
            inv.setItem(item.getSlot(), item.toItemStack(player));
        }
        
        player.getBukkitPlayer().openInventory(inv);
    }
}

// MenuItem tanımı
public class MenuItem {
    private String nameKey;           // Lokalizasyon
    private Material material;
    private String texture;           // Custom skull texture
    private int slot;
    private Consumer<KaraPlayer> onClick;
    private FormImage formImage;      // Bedrock için ikon
    
    public String getName(KaraPlayer player) {
        return player.getMessage(nameKey);
    }
}
```

---

## Auth Plugins

### 5. karapixel-auth

**Login/Register sistemi**

```yaml
# config.yml
auth:
  # Dil mesajları locales/ klasöründen yüklenir
  default-locale: tr_TR
  
  password:
    min-length: 6
    max-length: 32
    hash-algorithm: BCRYPT
    bcrypt-cost: 12
    
  login:
    max-attempts: 5
    lockout-duration: 300
    timeout: 30
    
  captcha:
    enabled: true
    type: AUTO  # Platform'a göre
    
  bedrock:
    auto-login: true  # Floodgate ile otomatik giriş
    prefix: "*"       # Bedrock oyuncu prefix'i
```

```java
// Bedrock otomatik giriş
@EventHandler
public void onJoin(PlayerJoinEvent event) {
    KaraPlayer player = KaraAPI.getPlayer(event.getPlayer());
    
    if (player.isBedrock()) {
        // Floodgate ile giriş yapmış, Xbox auth güvenilir
        if (FloodgateApi.getInstance().isFloodgatePlayer(player.getUuid())) {
            // Otomatik login
            sessionManager.createSession(player);
            player.sendMessage(Message.AUTH_AUTO_LOGIN);
            teleportToHub(player);
            return;
        }
    }
    
    // Java oyuncu - normal auth akışı
    startAuthProcess(player);
}
```

---

## Hub Plugins

### 6. karapixel-hub

**Hub lobby yönetimi**

```java
// Hub özellikleri
public class KaraHub extends JavaPlugin {
    
    @Override
    public void onEnable() {
        // Spawn koruma
        registerListener(new SpawnProtectionListener());
        
        // Double jump
        registerListener(new DoubleJumpListener());
        
        // Hub items (server selector, cosmetics, etc.)
        registerListener(new HubItemsListener());
        
        // Void damage protection
        registerListener(new VoidProtectionListener());
    }
}

// Hub item'ları
public enum HubItem {
    SERVER_SELECTOR(Material.COMPASS, "hub.item.server_selector", 0),
    COSMETICS(Material.CHEST, "hub.item.cosmetics", 4),
    PROFILE(Material.PLAYER_HEAD, "hub.item.profile", 8);
    
    // Bedrock'ta da aynı slot ve fonksiyon
}
```

### 7. karapixel-selector

**Oyun modu seçici (NPC + Portal)**

```java
// NPC sistemi
public class GameSelector {
    
    // 3D model NPC'ler
    private final Map<String, SelectorNPC> npcs = new HashMap<>();
    
    public void createNPC(String id, Location location, String modelId) {
        SelectorNPC npc = new SelectorNPC(id, location);
        npc.setCustomModel(modelId);  // 3D model entegrasyonu
        npc.setClickAction(player -> openGameMenu(player));
        npcs.put(id, npc);
    }
    
    public void openGameMenu(KaraPlayer player) {
        KaraMenu menu = KaraMenu.builder()
            .title("menu.game_selector.title")  // Lokalize
            .rows(3)
            .item(MenuItem.builder()
                .slot(11)
                .material(Material.GRASS_BLOCK)
                .name("menu.game_selector.skyblock.name")
                .lore("menu.game_selector.skyblock.lore")
                .onClick(p -> transferToSkyblock(p))
                .build())
            .item(MenuItem.builder()
                .slot(13)
                .material(Material.IRON_SWORD)
                .name("menu.game_selector.coming_soon.name")
                .build())
            .build();
        
        menu.open(player);
    }
}

// Portal sistemi
public class PortalListener implements Listener {
    
    @EventHandler
    public void onMove(PlayerMoveEvent event) {
        Location to = event.getTo();
        
        for (GamePortal portal : portalManager.getPortals()) {
            if (portal.contains(to)) {
                KaraPlayer player = KaraAPI.getPlayer(event.getPlayer());
                player.sendMessage(Message.PORTAL_TELEPORTING);
                portal.teleport(player);
                break;
            }
        }
    }
}
```

---

## Skyblock Plugins

### 8. karapixel-skyblock

**Ana skyblock sistemi**

```yaml
# config.yml
skyblock:
  island:
    start-size: 100          # Başlangıç ada boyutu
    max-size: 500            # Maksimum ada boyutu
    spacing: 1000            # Adalar arası mesafe
    
  templates:                 # Ada şablonları
    - id: normal
      name: "island.template.normal"  # Lokalize
      schematic: "normal.schem"
      
    - id: desert
      name: "island.template.desert"
      schematic: "desert.schem"
      
    - id: nether
      name: "island.template.nether"
      schematic: "nether.schem"
      unlock-level: 10       # Level 10'da açılır
      
  levels:
    blocks:                  # Block değerleri
      DIAMOND_BLOCK: 100
      EMERALD_BLOCK: 75
      GOLD_BLOCK: 50
      IRON_BLOCK: 25
      
  coop:
    max-members: 5
    roles:
      - OWNER
      - ADMIN
      - MEMBER
      - VISITOR
```

```java
// Ada sistemi
public class IslandManager {
    
    public CompletableFuture<Island> createIsland(KaraPlayer owner, IslandTemplate template) {
        return CompletableFuture.supplyAsync(() -> {
            // En boş world server'ı bul
            String targetServer = findLeastLoadedServer();
            
            // Ada pozisyonu hesapla
            IslandPosition pos = calculateNextPosition(targetServer);
            
            // Database'e kaydet
            Island island = new Island(
                UUID.randomUUID(),
                owner.getUuid(),
                template.getId(),
                targetServer,
                pos,
                Instant.now()
            );
            islandRepository.save(island);
            
            // Schematic yapıştır (async)
            pasteSchematic(template.getSchematic(), pos.toLocation());
            
            // Oyuncuyu transfer et
            messaging.publish("teleport", new TeleportRequestPacket(
                owner.getUuid(),
                targetServer,
                pos.getSpawnLocation()
            ));
            
            return island;
        });
    }
    
    // Ada menüsü (Mobil uyumlu)
    public void openIslandMenu(KaraPlayer player) {
        Island island = getIsland(player);
        
        KaraMenu menu = KaraMenu.builder()
            .title("menu.island.title")
            .rows(6)
            // Ana kontroller
            .item(homeButton(player, island))
            .item(membersButton(player, island))
            .item(settingsButton(player, island))
            .item(upgradesButton(player, island))
            // Bilgi paneli
            .item(levelDisplay(player, island))
            .item(bankDisplay(player, island))
            .build();
        
        menu.open(player);
    }
}
```

### 9. karapixel-generators

**Ore/Crop generator sistemi**

```java
// Generator tipleri
public enum GeneratorType {
    COBBLESTONE("generator.cobble", Material.COBBLESTONE, 20),  // 20 tick = 1sn
    IRON("generator.iron", Material.IRON_ORE, 100),
    GOLD("generator.gold", Material.GOLD_ORE, 200),
    DIAMOND("generator.diamond", Material.DIAMOND_ORE, 400),
    EMERALD("generator.emerald", Material.EMERALD_ORE, 600);
    
    private final String nameKey;  // Lokalizasyon
    private final Material output;
    private final int interval;    // Tick cinsinden
}

// Generator entity (3D model destekli)
public class Generator {
    private UUID id;
    private GeneratorType type;
    private Location location;
    private int tier;              // Upgrade seviyesi
    private String customModelId;  // 3D model ID
    
    public void upgrade() {
        tier++;
        // Yeni 3D model uygula
        updateCustomModel("generator_" + type.name().toLowerCase() + "_tier" + tier);
    }
}
```

### 10. karapixel-skills

**Skill/leveling sistemi**

```java
// Skill tipleri
public enum SkillType {
    MINING("skill.mining", "⛏"),
    FARMING("skill.farming", "🌾"),
    COMBAT("skill.combat", "⚔"),
    FISHING("skill.fishing", "🎣"),
    FORAGING("skill.foraging", "🪓"),
    ENCHANTING("skill.enchanting", "✨");
    
    private final String nameKey;
    private final String icon;  // Bedrock için emoji
}

// XP kazanma
@EventHandler
public void onBlockBreak(BlockBreakEvent event) {
    KaraPlayer player = KaraAPI.getPlayer(event.getPlayer());
    Material block = event.getBlock().getType();
    
    if (isOre(block)) {
        int xp = getXpForBlock(block);
        skillManager.addXp(player, SkillType.MINING, xp);
        
        // Action bar'da göster (her platform)
        player.sendActionBar(player.getMessage(
            Message.SKILL_XP_GAIN,
            Placeholder.of("skill", player.getMessage(SkillType.MINING.getNameKey())),
            Placeholder.of("xp", xp)
        ));
    }
}
```

---

## Global Plugins

### 11. karapixel-economy

**Tek cüzdan ekonomi sistemi**

```java
// Cross-server economy
public class EconomyManager {
    
    // Tüm sunucularda aynı bakiye
    public double getBalance(UUID playerId) {
        return economyRepository.getBalance(playerId);
    }
    
    public CompletableFuture<Boolean> transfer(UUID from, UUID to, double amount) {
        return CompletableFuture.supplyAsync(() -> {
            // Transaction başlat
            try (Transaction tx = database.beginTransaction()) {
                double fromBalance = economyRepository.getBalance(from);
                
                if (fromBalance < amount) {
                    return false;
                }
                
                economyRepository.withdraw(from, amount);
                economyRepository.deposit(to, amount);
                
                // Log transaction
                transactionRepository.log(new EconomyTransaction(
                    UUID.randomUUID(),
                    from, to, amount,
                    TransactionType.TRANSFER,
                    Instant.now()
                ));
                
                tx.commit();
                
                // Cross-server sync
                messaging.publish("economy", new EconomyUpdatePacket(from, to, amount));
                
                return true;
            }
        });
    }
}
```

### 12. karapixel-cosmetics

**Cosmetics sistemi (3D model destekli)**

```java
// Cosmetic tipleri
public enum CosmeticType {
    PARTICLE("cosmetic.type.particle"),
    WING("cosmetic.type.wing"),         // 3D model
    HAT("cosmetic.type.hat"),           // 3D model
    PET("cosmetic.type.pet"),           // 3D model (ayrı plugin)
    TRAIL("cosmetic.type.trail"),
    KILL_EFFECT("cosmetic.type.kill_effect");
}

// 3D model cosmetic
public class WingCosmetic extends Cosmetic {
    private String modelId;      // Resource pack model ID
    private String bedrockModel; // Bedrock geometry ID
    
    @Override
    public void apply(KaraPlayer player) {
        if (player.isBedrock()) {
            // Bedrock custom geometry
            applyBedrockModel(player, bedrockModel);
        } else {
            // Java custom model data
            applyJavaModel(player, modelId);
        }
    }
}

// Cosmetics menüsü
public void openCosmeticsMenu(KaraPlayer player) {
    KaraMenu menu = KaraMenu.builder()
        .title("menu.cosmetics.title")
        .rows(6)
        .item(categoryButton(CosmeticType.PARTICLE, 10))
        .item(categoryButton(CosmeticType.WING, 12))
        .item(categoryButton(CosmeticType.HAT, 14))
        .item(categoryButton(CosmeticType.PET, 16))
        .item(categoryButton(CosmeticType.TRAIL, 28))
        .item(categoryButton(CosmeticType.KILL_EFFECT, 30))
        .build();
    
    menu.open(player);
}
```

### 13. karapixel-pets

**Pet sistemi (3D model)**

```java
// Pet entity
public class Pet {
    private UUID id;
    private UUID ownerId;
    private PetType type;
    private String name;           // Custom isim
    private int level;
    private int xp;
    private String modelId;        // 3D model
    private Map<PetAbility, Integer> abilities;
    
    public void spawn(KaraPlayer owner) {
        // 3D model entity oluştur
        Location loc = owner.getLocation().add(1, 0, 1);
        
        if (owner.isBedrock()) {
            // Bedrock: Custom entity
            spawnBedrockPet(owner, loc);
        } else {
            // Java: ArmorStand + Custom Model
            spawnJavaPet(owner, loc);
        }
        
        // Pet AI başlat
        startFollowTask(owner);
    }
}

// Pet tipleri
public enum PetType {
    DOG("pet.type.dog", "pet_dog"),
    CAT("pet.type.cat", "pet_cat"),
    DRAGON("pet.type.dragon", "pet_dragon"),      // Premium
    PHOENIX("pet.type.phoenix", "pet_phoenix");    // Rare
    
    private final String nameKey;
    private final String modelId;
}
```

---

## Dil Desteği Altyapısı

### Dil Dosyası Yapısı

```yaml
# locales/tr_TR.yml
meta:
  language: "Türkçe"
  code: "tr_TR"
  authors: ["KaraPixel Team"]
  
# Genel mesajlar
general:
  prefix: "<gradient:#FF6B6B:#4ECDC4>[KaraPixel]</gradient> "
  error: "<red>Hata: {message}</red>"
  success: "<green>Başarılı!</green>"
  no_permission: "<red>Bu işlem için yetkiniz yok!</red>"
  
# Auth mesajları
auth:
  welcome: |
    <yellow>═══════════════════════════════════
    <gold>KaraPixel</gold> sunucusuna hoş geldin!
    ═══════════════════════════════════</yellow>
  register_prompt: "<yellow>Kayıt olmak için: /register <şifre> <şifre></yellow>"
  login_prompt: "<yellow>Giriş yapmak için: /login <şifre></yellow>"
  register_success: "<green>Başarıyla kayıt oldun! Hub'a yönlendiriliyorsun...</green>"
  login_success: "<green>Giriş başarılı! Hoş geldin, {player}!</green>"
  wrong_password: "<red>Yanlış şifre! Kalan hak: {remaining}</red>"
  too_many_attempts: "<red>Çok fazla başarısız deneme! {duration} bekleyin.</red>"
  auto_login: "<green>Xbox hesabınız ile otomatik giriş yapıldı!</green>"
  
# Skyblock mesajları
skyblock:
  island:
    created: "<green>Adan oluşturuldu! /is home ile gidebilirsin.</green>"
    teleporting: "<yellow>Adana ışınlanıyorsun...</yellow>"
    level: "<gold>Ada Seviyesi: <yellow>{level}</yellow></gold>"
    
  generator:
    placed: "<green>{generator} yerleştirildi!</green>"
    upgraded: "<green>{generator} seviye {level} oldu!</green>"
    
  coop:
    invited: "<yellow>{player} seni adasına davet etti!</yellow>"
    joined: "<green>{player} adana katıldı!</green>"
    left: "<red>{player} adandan ayrıldı.</red>"

# Menü başlıkları
menu:
  game_selector:
    title: "Oyun Seçici"
    skyblock:
      name: "<green>⛏ Skyblock</green>"
      lore:
        - "<gray>Kendi adanı oluştur ve geliştir!"
        - ""
        - "<yellow>▸ Tıkla ve başla!</yellow>"
        
  island:
    title: "Ada Menüsü"
    home:
      name: "<green>🏠 Eve Git</green>"
    members:
      name: "<aqua>👥 Üyeler</aqua>"
    settings:
      name: "<gold>⚙ Ayarlar</gold>"
    upgrades:
      name: "<light_purple>⬆ Yükseltmeler</light_purple>"
      
  cosmetics:
    title: "Kozmetikler"
    
# Skill isimleri
skill:
  mining: "Madencilik"
  farming: "Çiftçilik"
  combat: "Savaş"
  fishing: "Balıkçılık"
  foraging: "Ormancılık"
  enchanting: "Büyüleme"
```

```yaml
# locales/en_US.yml
meta:
  language: "English"
  code: "en_US"
  authors: ["KaraPixel Team"]

general:
  prefix: "<gradient:#FF6B6B:#4ECDC4>[KaraPixel]</gradient> "
  error: "<red>Error: {message}</red>"
  success: "<green>Success!</green>"
  no_permission: "<red>You don't have permission for this!</red>"

auth:
  welcome: |
    <yellow>═══════════════════════════════════
    Welcome to <gold>KaraPixel</gold>!
    ═══════════════════════════════════</yellow>
  register_prompt: "<yellow>To register: /register <password> <password></yellow>"
  login_prompt: "<yellow>To login: /login <password></yellow>"
  # ... devamı
```

### Dil Yönetimi Kodu

```java
// LocaleManager - Dil yönetimi
public class LocaleManager {
    private final Map<String, YamlConfiguration> locales = new HashMap<>();
    private String defaultLocale = "tr_TR";
    
    public void loadLocales() {
        // Tüm dil dosyalarını yükle
        for (String locale : Arrays.asList("tr_TR", "en_US", "de_DE")) {
            InputStream is = getResource("locales/" + locale + ".yml");
            if (is != null) {
                locales.put(locale, YamlConfiguration.loadConfiguration(
                    new InputStreamReader(is, StandardCharsets.UTF_8)
                ));
            }
        }
    }
    
    public String getMessage(KaraPlayer player, String key, Placeholder... placeholders) {
        String locale = player.getLocale();  // Oyuncunun dil tercihi
        
        // Önce oyuncunun dilini dene
        String message = getRawMessage(locale, key);
        
        // Bulunamazsa varsayılan dil
        if (message == null) {
            message = getRawMessage(defaultLocale, key);
        }
        
        // Hala bulunamazsa key'i döndür
        if (message == null) {
            return key;
        }
        
        // Placeholder'ları uygula
        for (Placeholder ph : placeholders) {
            message = message.replace("{" + ph.getKey() + "}", ph.getValue());
        }
        
        // MiniMessage ile parse et
        return Text.parse(message);
    }
    
    private String getRawMessage(String locale, String key) {
        YamlConfiguration config = locales.get(locale);
        return config != null ? config.getString(key) : null;
    }
}

// Oyuncu dil tercihi
public class KaraPlayer {
    private String locale = "tr_TR";  // Varsayılan Türkçe
    
    // Bedrock oyuncuların dil tespiti
    public void detectLocale() {
        if (isBedrock()) {
            // Geyser'dan dil bilgisi al
            String geyserLocale = GeyserHook.getLocale(uuid);
            if (geyserLocale != null) {
                this.locale = mapGeyserLocale(geyserLocale);
            }
        } else {
            // Java client locale
            String clientLocale = bukkitPlayer.getLocale();
            this.locale = mapClientLocale(clientLocale);
        }
    }
    
    public String getMessage(String key, Placeholder... placeholders) {
        return KaraAPI.getLocaleManager().getMessage(this, key, placeholders);
    }
    
    public void sendMessage(String key, Placeholder... placeholders) {
        bukkitPlayer.sendMessage(getMessage(key, placeholders));
    }
}
```

---

## Mobil Uyumluluk

### Platform Tespiti

```java
// Platform tespit sistemi
public class PlatformDetector {
    
    public static PlayerPlatform detect(Player player) {
        // Floodgate kontrolü
        if (FloodgateApi.getInstance().isFloodgatePlayer(player.getUniqueId())) {
            return PlayerPlatform.BEDROCK;
        }
        
        // Geyser kontrolü (alternatif)
        if (GeyserHook.isBedrockPlayer(player.getUniqueId())) {
            return PlayerPlatform.BEDROCK;
        }
        
        return PlayerPlatform.JAVA;
    }
    
    public static DeviceType getDeviceType(KaraPlayer player) {
        if (!player.isBedrock()) {
            return DeviceType.PC;
        }
        
        // Geyser'dan device bilgisi
        GeyserConnection conn = GeyserHook.getConnection(player.getUuid());
        if (conn != null) {
            return switch (conn.getDeviceOs()) {
                case ANDROID -> DeviceType.MOBILE;
                case IOS -> DeviceType.MOBILE;
                case WIN10 -> DeviceType.PC;
                case XBOX -> DeviceType.CONSOLE;
                case PLAYSTATION -> DeviceType.CONSOLE;
                case SWITCH -> DeviceType.CONSOLE;
                default -> DeviceType.UNKNOWN;
            };
        }
        
        return DeviceType.UNKNOWN;
    }
}

public enum PlayerPlatform {
    JAVA, BEDROCK
}

public enum DeviceType {
    PC, MOBILE, CONSOLE, UNKNOWN
}
```

### Bedrock Forms API Entegrasyonu

```java
// Geyser Cumulus Forms API
public class BedrockForms {
    
    // Simple form (butonlu menü)
    public static void openSimpleForm(KaraPlayer player, String titleKey, 
                                       List<FormButton> buttons) {
        String title = player.getMessage(titleKey);
        
        SimpleForm.Builder builder = SimpleForm.builder()
            .title(title);
        
        for (FormButton button : buttons) {
            builder.button(
                button.getText(player),
                button.getImage()
            );
        }
        
        GeyserHook.sendForm(player, builder.build(), response -> {
            if (response.isClosed()) return;
            
            int clickedIndex = response.clickedButtonId();
            buttons.get(clickedIndex).getAction().accept(player);
        });
    }
    
    // Modal form (onay dialogu)
    public static void openConfirmation(KaraPlayer player, String titleKey,
                                        String contentKey, Consumer<Boolean> callback) {
        ModalForm form = ModalForm.builder()
            .title(player.getMessage(titleKey))
            .content(player.getMessage(contentKey))
            .button1(player.getMessage("general.confirm"))
            .button2(player.getMessage("general.cancel"))
            .build();
        
        GeyserHook.sendForm(player, form, response -> {
            callback.accept(response.clickedButtonId() == 0);
        });
    }
    
    // Custom form (input'lu form)
    public static void openInputForm(KaraPlayer player, String titleKey,
                                     List<FormInput> inputs, Consumer<List<String>> callback) {
        CustomForm.Builder builder = CustomForm.builder()
            .title(player.getMessage(titleKey));
        
        for (FormInput input : inputs) {
            if (input.isDropdown()) {
                builder.dropdown(input.getLabel(player), input.getOptions(player));
            } else {
                builder.input(input.getLabel(player), input.getPlaceholder(player));
            }
        }
        
        GeyserHook.sendForm(player, builder.build(), response -> {
            if (response.isClosed()) return;
            
            List<String> values = new ArrayList<>();
            for (int i = 0; i < inputs.size(); i++) {
                values.add(response.get(i));
            }
            callback.accept(values);
        });
    }
}
```

### UI Tasarım Kuralları

```
┌─────────────────────────────────────────────────────────────────┐
│                 MOBİL-FIRST UI KURALLARI                        │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  1. MENÜ BOYUTU                                                 │
│  ─────────────                                                  │
│  ├── Maksimum 3 satır (27 slot) tercih et                      │
│  ├── 6 satır sadece gerektiğinde (inventory view)              │
│  └── Bedrock'ta scroll yok, her şey görünür olmalı             │
│                                                                 │
│  2. BUTON BOYUTU                                                │
│  ────────────────                                               │
│  ├── Her buton en az 1 slot boşluk ile ayrılmalı               │
│  ├── Touch target: minimum 2x2 slot ideal                      │
│  └── Ana işlevler ortada, büyük ikonlarla                      │
│                                                                 │
│  3. TEXT UZUNLUĞU                                               │
│  ─────────────────                                              │
│  ├── Buton isimleri: max 15 karakter                           │
│  ├── Lore satırı: max 30 karakter                              │
│  └── Form butonları: max 20 karakter                           │
│                                                                 │
│  4. RENK KULLANIMI                                              │
│  ──────────────────                                             │
│  ├── Yüksek kontrast (mobil ekranlarda okunabilirlik)          │
│  ├── Yeşil = pozitif/onay                                      │
│  ├── Kırmızı = negatif/iptal                                   │
│  └── Sarı = dikkat/bilgi                                       │
│                                                                 │
│  5. GERİ/KAPAT BUTONU                                          │
│  ─────────────────────                                          │
│  ├── Her menüde sağ alt köşede                                 │
│  ├── Her zaman aynı yerde (slot 26 veya 53)                    │
│  └── Barrier block veya kırmızı boya                           │
│                                                                 │
│  6. LOADING STATE                                               │
│  ─────────────────                                              │
│  ├── Async işlemlerde loading göster                           │
│  ├── Bedrock: Form'u kapatıp action bar mesajı                 │
│  └── Spam tıklamayı önle (cooldown)                            │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## 3D Model Entegrasyonu

### Resource Pack Yapısı

```
resourcepack/
├── pack.mcmeta
├── pack.png
│
├── assets/
│   └── minecraft/
│       │
│       ├── models/
│       │   ├── item/
│       │   │   ├── paper.json              # Custom item base
│       │   │   ├── generators/
│       │   │   │   ├── cobble_gen_t1.json
│       │   │   │   ├── cobble_gen_t2.json
│       │   │   │   └── diamond_gen_t1.json
│       │   │   ├── pets/
│       │   │   │   ├── dog.json
│       │   │   │   ├── cat.json
│       │   │   │   └── dragon.json
│       │   │   └── cosmetics/
│       │   │       ├── wing_angel.json
│       │   │       ├── wing_demon.json
│       │   │       └── hat_crown.json
│       │   │
│       │   └── entity/                     # Custom entity modelleri
│       │       └── npc/
│       │           ├── selector_npc.json
│       │           └── shop_keeper.json
│       │
│       ├── textures/
│       │   ├── item/
│       │   │   ├── generators/
│       │   │   ├── pets/
│       │   │   └── cosmetics/
│       │   └── entity/
│       │
│       └── sounds/
│           └── custom/
│               ├── level_up.ogg
│               └── generator_upgrade.ogg
```

### Custom Model Data Sistemi

```java
// Java Edition - Custom Model Data
public class ModelManager {
    
    // Model ID → CustomModelData mapping
    private static final Map<String, Integer> MODEL_DATA = new HashMap<>();
    
    static {
        // Generators
        MODEL_DATA.put("generator_cobble_t1", 1001);
        MODEL_DATA.put("generator_cobble_t2", 1002);
        MODEL_DATA.put("generator_iron_t1", 1011);
        MODEL_DATA.put("generator_diamond_t1", 1021);
        
        // Pets
        MODEL_DATA.put("pet_dog", 2001);
        MODEL_DATA.put("pet_cat", 2002);
        MODEL_DATA.put("pet_dragon", 2003);
        
        // Cosmetics
        MODEL_DATA.put("wing_angel", 3001);
        MODEL_DATA.put("wing_demon", 3002);
        MODEL_DATA.put("hat_crown", 3101);
    }
    
    public ItemStack createModelItem(String modelId) {
        Integer cmd = MODEL_DATA.get(modelId);
        if (cmd == null) {
            throw new IllegalArgumentException("Unknown model: " + modelId);
        }
        
        ItemStack item = new ItemStack(Material.PAPER);  // Base item
        ItemMeta meta = item.getItemMeta();
        meta.setCustomModelData(cmd);
        item.setItemMeta(meta);
        
        return item;
    }
}

// Bedrock uyumluluk - Geyser otomatik dönüştürür
// Ancak bazı kompleks modeller için manual mapping gerekebilir
public class BedrockModelMapper {
    
    // Java CustomModelData → Bedrock geometry mapping
    private static final Map<Integer, String> BEDROCK_MAPPING = new HashMap<>();
    
    static {
        BEDROCK_MAPPING.put(2001, "geometry.pet.dog");
        BEDROCK_MAPPING.put(2003, "geometry.pet.dragon");
    }
    
    public String getBedrockGeometry(int customModelData) {
        return BEDROCK_MAPPING.get(customModelData);
    }
}
```

---

## Gradle Multi-Module Setup

```kotlin
// settings.gradle.kts
rootProject.name = "karapixel-plugins"

// Core modules
include("karapixel-core")
include("karapixel-database")
include("karapixel-messaging")
include("karapixel-ui")

// Auth
include("karapixel-auth")

// Hub
include("karapixel-hub")
include("karapixel-selector")

// Skyblock
include("karapixel-skyblock")
include("karapixel-generators")
include("karapixel-skills")
include("karapixel-quests")
include("karapixel-shop")
include("karapixel-upgrades")
include("karapixel-minions")
include("karapixel-enchants")

// Global
include("karapixel-economy")
include("karapixel-cosmetics")
include("karapixel-pets")
include("karapixel-chat")
include("karapixel-tablist")
include("karapixel-ranks")

// Admin
include("karapixel-moderation")
include("karapixel-security")

// Velocity
include("karapixel-velocity")
```

```kotlin
// build.gradle.kts (root)
plugins {
    java
    id("com.github.johnrengelman.shadow") version "8.1.1" apply false
    id("io.papermc.paperweight.userdev") version "1.7.1" apply false
}

subprojects {
    apply(plugin = "java")
    
    group = "net.karapixel"
    version = "1.0.0-SNAPSHOT"
    
    java {
        toolchain.languageVersion.set(JavaLanguageVersion.of(21))
    }
    
    repositories {
        mavenCentral()
        maven("https://repo.papermc.io/repository/maven-public/")
        maven("https://repo.opencollab.dev/main/")  // Geyser
    }
}
```

---

## Build & Deploy

```bash
# Tüm pluginleri build et
./gradlew build

# Tek plugin build et
./gradlew :karapixel-skyblock:build

# Shadow JAR (bağımlılıklar dahil)
./gradlew shadowJar

# Local sunucuya deploy et
./gradlew deployLocal

# Production'a deploy et (CI/CD)
./gradlew deployProd
```

---

*📅 Son güncelleme: 24 Aralık 2024*
