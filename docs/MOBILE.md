# 📱 KaraPixel - Mobil & Bedrock Desteği

> Türkiye pazarında %40-60 Bedrock oyuncu oranı göz önünde bulundurularak tam mobil uyumluluk.

---

## Genel Bakış

```
┌─────────────────────────────────────────────────────────────────┐
│                 PLATFORM DESTEĞİ                                │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  JAVA EDITION                                                   │
│  ├── Windows, macOS, Linux                                     │
│  ├── Port: 25565 (TCP)                                         │
│  └── Versiyon: 1.21.x (ViaVersion ile 1.8+)                   │
│                                                                 │
│  BEDROCK EDITION (Geyser + Floodgate)                          │
│  ├── Windows 10/11                                              │
│  ├── iOS (iPhone, iPad)                                        │
│  ├── Android (Telefon, Tablet)                                 │
│  ├── Xbox (One, Series X/S)                                    │
│  ├── PlayStation (4, 5)                                        │
│  ├── Nintendo Switch                                            │
│  └── Port: 19132 (UDP)                                         │
│                                                                 │
│  TAHMİNİ DAĞILIM (Türkiye)                                     │
│  ├── Java PC: %40-50                                           │
│  ├── Bedrock Mobil: %30-40                                     │
│  ├── Bedrock Konsol: %10-15                                    │
│  └── Bedrock PC: %5-10                                         │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Geyser & Floodgate Kurulumu

### Velocity Entegrasyonu

```yaml
# Velocity plugins klasörüne:
# - Geyser-Velocity.jar
# - floodgate-velocity.jar

# plugins/Geyser-Velocity/config.yml
bedrock:
  address: 0.0.0.0
  port: 19132
  clone-remote-port: false
  motd1: "KaraPixel"
  motd2: "Bedrock Sunucusu"
  server-name: "KaraPixel"
  compression-level: 6
  enable-proxy-protocol: false

remote:
  address: 127.0.0.1
  port: 25565
  auth-type: floodgate  # ÖNEMLİ: Floodgate kullan

# Floodgate ayarları
floodgate:
  username-prefix: "*"   # Bedrock oyuncu prefix'i
  replace-spaces: true   # Boşlukları _ ile değiştir

# Performans
pending-authentication-timeout: 120000
max-players: 1000

# Özelleştirme
show-cooldown: title
show-coordinates: true
disable-bedrock-scaffolding: false

# Cache (RAM kullanımını azaltır)
cache-images: 0
allow-custom-skulls: true
```

### Floodgate Konfigürasyonu

```yaml
# plugins/floodgate/config.yml

# Bedrock oyuncu prefix'i
# Bu prefix ile gelen oyuncular otomatik olarak Bedrock kabul edilir
username-prefix: "*"

# Bedrock oyuncuların UUID'si nasıl oluşturulsun
uuid-generation: FLOODGATE  # Tutarlı UUID

# Bedrock skin'leri Java'ya nasıl aktarılsın
send-floodgate-data: true

# Whitelist'te Bedrock oyuncular
whitelist:
  enabled: false  # Sunucu zaten non-premium

# Link sistemi (Java-Bedrock hesap bağlama)
player-link:
  enabled: true
  require-link: false  # Zorunlu değil
  allowed-link-names:
    - "global"

# Debug
debug-mode: false
```

---

## Platform Tespiti

### PlatformDetector Sınıfı

```java
public class PlatformDetector {
    
    /**
     * Oyuncunun platformunu tespit eder
     */
    public static PlayerPlatform detect(Player player) {
        UUID uuid = player.getUniqueId();
        
        // Yöntem 1: Floodgate API
        if (isFloodgateAvailable()) {
            FloodgateApi api = FloodgateApi.getInstance();
            if (api.isFloodgatePlayer(uuid)) {
                return PlayerPlatform.BEDROCK;
            }
        }
        
        // Yöntem 2: UUID prefix kontrolü
        // Floodgate UUID'leri belirli bir pattern izler
        if (uuid.getMostSignificantBits() == 0) {
            return PlayerPlatform.BEDROCK;
        }
        
        // Yöntem 3: İsim prefix kontrolü
        if (player.getName().startsWith("*")) {
            return PlayerPlatform.BEDROCK;
        }
        
        return PlayerPlatform.JAVA;
    }
    
    /**
     * Bedrock oyuncunun cihaz tipini tespit eder
     */
    public static DeviceType getDeviceType(Player player) {
        if (!isFloodgateAvailable()) {
            return DeviceType.UNKNOWN;
        }
        
        FloodgateApi api = FloodgateApi.getInstance();
        if (!api.isFloodgatePlayer(player.getUniqueId())) {
            return DeviceType.PC;  // Java = PC
        }
        
        FloodgatePlayer floodgatePlayer = api.getPlayer(player.getUniqueId());
        if (floodgatePlayer == null) {
            return DeviceType.UNKNOWN;
        }
        
        return switch (floodgatePlayer.getDeviceOs()) {
            case ANDROID -> DeviceType.MOBILE;
            case IOS -> DeviceType.MOBILE;
            case WIN10 -> DeviceType.PC;
            case XBOX -> DeviceType.CONSOLE;
            case PLAYSTATION -> DeviceType.CONSOLE;
            case SWITCH -> DeviceType.CONSOLE;
            case AMAZON -> DeviceType.MOBILE;  // Fire tablet
            default -> DeviceType.UNKNOWN;
        };
    }
    
    /**
     * Oyuncunun input tipini tespit eder (touch, controller, keyboard)
     */
    public static InputType getInputType(Player player) {
        if (!isFloodgateAvailable()) {
            return InputType.KEYBOARD_MOUSE;
        }
        
        FloodgateApi api = FloodgateApi.getInstance();
        FloodgatePlayer fp = api.getPlayer(player.getUniqueId());
        if (fp == null) {
            return InputType.KEYBOARD_MOUSE;
        }
        
        return switch (fp.getInputMode()) {
            case TOUCH -> InputType.TOUCH;
            case CONTROLLER -> InputType.CONTROLLER;
            case KEYBOARD_MOUSE -> InputType.KEYBOARD_MOUSE;
            default -> InputType.UNKNOWN;
        };
    }
    
    private static boolean isFloodgateAvailable() {
        return Bukkit.getPluginManager().getPlugin("floodgate") != null;
    }
}

public enum PlayerPlatform {
    JAVA, BEDROCK
}

public enum DeviceType {
    PC, MOBILE, CONSOLE, UNKNOWN
}

public enum InputType {
    KEYBOARD_MOUSE, TOUCH, CONTROLLER, UNKNOWN
}
```

---

## Bedrock Forms API

### Form Tipleri

```
┌─────────────────────────────────────────────────────────────────┐
│                    BEDROCK FORM TİPLERİ                         │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  1. SIMPLE FORM                                                 │
│  ═══════════════                                                │
│  Butonlu basit menü                                            │
│  ┌─────────────────────────────┐                               │
│  │      Oyun Seçici            │                               │
│  ├─────────────────────────────┤                               │
│  │  [🌳] Skyblock              │                               │
│  │  [⚔️] PvP Arena             │                               │
│  │  [🏠] Hub'a Dön             │                               │
│  └─────────────────────────────┘                               │
│                                                                 │
│  2. MODAL FORM                                                  │
│  ═════════════                                                  │
│  İki butonlu onay dialogu                                      │
│  ┌─────────────────────────────┐                               │
│  │   Adayı silmek istediğine   │                               │
│  │   emin misin?               │                               │
│  ├─────────────────────────────┤                               │
│  │  [Evet]         [Hayır]     │                               │
│  └─────────────────────────────┘                               │
│                                                                 │
│  3. CUSTOM FORM                                                 │
│  ═════════════                                                  │
│  Input alanları, slider, dropdown                              │
│  ┌─────────────────────────────┐                               │
│  │   Ada Ayarları              │                               │
│  ├─────────────────────────────┤                               │
│  │  Ada İsmi: [___________]    │                               │
│  │  PvP: [ON/OFF]              │                               │
│  │  Ziyaretçi: [▼ Dropdown]    │                               │
│  │  [Kaydet]                   │                               │
│  └─────────────────────────────┘                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Geyser Cumulus API Kullanımı

```java
import org.geysermc.cumulus.form.SimpleForm;
import org.geysermc.cumulus.form.ModalForm;
import org.geysermc.cumulus.form.CustomForm;
import org.geysermc.cumulus.util.FormImage;
import org.geysermc.floodgate.api.FloodgateApi;

public class BedrockForms {
    
    /**
     * Simple Form - Butonlu menü
     */
    public static void openGameSelector(KaraPlayer player) {
        SimpleForm form = SimpleForm.builder()
            .title(player.t("selector.menu.title"))
            .content(player.t("selector.menu.content"))
            
            // Skyblock butonu
            .button(
                player.t("selector.games.skyblock.name"),
                FormImage.Type.PATH,
                "textures/items/grass_block"
            )
            
            // PvP butonu (yakında)
            .button(
                player.t("selector.games.coming_soon.name"),
                FormImage.Type.PATH,
                "textures/items/iron_sword"
            )
            
            // Hub butonu
            .button(
                player.t("hub.teleporting"),
                FormImage.Type.PATH,
                "textures/items/bed_red"
            )
            
            .validResultHandler(response -> {
                int buttonId = response.clickedButtonId();
                switch (buttonId) {
                    case 0 -> transferToSkyblock(player);
                    case 1 -> player.sendMessage("selector.games.coming_soon.name");
                    case 2 -> transferToHub(player);
                }
            })
            .build();
        
        sendForm(player, form);
    }
    
    /**
     * Modal Form - Onay dialogu
     */
    public static void confirmIslandDelete(KaraPlayer player, Consumer<Boolean> callback) {
        ModalForm form = ModalForm.builder()
            .title(player.t("island.delete.confirm_title"))
            .content(player.t("island.delete.confirm_content"))
            .button1(player.t("general.yes"))   // Sol buton
            .button2(player.t("general.no"))    // Sağ buton
            .validResultHandler(response -> {
                // clickedButtonId: 0 = button1 (Evet), 1 = button2 (Hayır)
                callback.accept(response.clickedButtonId() == 0);
            })
            .build();
        
        sendForm(player, form);
    }
    
    /**
     * Custom Form - Input'lu form
     */
    public static void openIslandSettings(KaraPlayer player, Island island) {
        CustomForm form = CustomForm.builder()
            .title(player.t("island.settings.title"))
            
            // Text input - Ada ismi
            .input(
                player.t("island.settings.name_label"),
                player.t("island.settings.name_placeholder"),
                island.getName()
            )
            
            // Toggle - PvP
            .toggle(
                player.t("island.settings.pvp_label"),
                island.isPvpEnabled()
            )
            
            // Dropdown - Ziyaretçi izni
            .dropdown(
                player.t("island.settings.visitor_label"),
                List.of(
                    player.t("island.settings.visitor_none"),
                    player.t("island.settings.visitor_friends"),
                    player.t("island.settings.visitor_all")
                ),
                island.getVisitorPermission().ordinal()
            )
            
            // Slider - Spawn koruması mesafesi
            .slider(
                player.t("island.settings.protection_label"),
                0, 50, 1,
                island.getProtectionRadius()
            )
            
            .validResultHandler(response -> {
                String newName = response.asInput(0);
                boolean pvpEnabled = response.asToggle(1);
                int visitorPerm = response.asDropdown(2);
                int protectionRadius = (int) response.asSlider(3);
                
                // Ayarları kaydet
                island.setName(newName);
                island.setPvpEnabled(pvpEnabled);
                island.setVisitorPermission(VisitorPermission.values()[visitorPerm]);
                island.setProtectionRadius(protectionRadius);
                
                player.sendMessage("island.settings.saved");
            })
            .build();
        
        sendForm(player, form);
    }
    
    /**
     * Form gönderme helper
     */
    private static void sendForm(KaraPlayer player, org.geysermc.cumulus.form.Form form) {
        FloodgateApi.getInstance().sendForm(player.getUuid(), form);
    }
}
```

---

## Platform-Aware UI Sistemi

### KaraMenu - Unified Menu API

```java
public class KaraMenu {
    private final String titleKey;
    private final int rows;
    private final List<MenuItem> items;
    private final MenuType type;
    
    public void open(KaraPlayer player) {
        if (player.isBedrock() && type.hasBedrockForm()) {
            // Bedrock: Native form kullan
            openBedrockForm(player);
        } else {
            // Java veya Bedrock chest: Inventory kullan
            openInventory(player);
        }
    }
    
    private void openBedrockForm(KaraPlayer player) {
        SimpleForm.Builder builder = SimpleForm.builder()
            .title(player.t(titleKey));
        
        // Content oluştur
        StringBuilder content = new StringBuilder();
        for (MenuItem item : items) {
            if (item.isInfoOnly()) {
                content.append(item.getDescription(player)).append("\n");
            }
        }
        if (content.length() > 0) {
            builder.content(content.toString());
        }
        
        // Butonları ekle
        for (MenuItem item : items) {
            if (!item.isInfoOnly() && item.isClickable()) {
                builder.button(
                    item.getName(player),
                    item.getFormImage()
                );
            }
        }
        
        builder.validResultHandler(response -> {
            int index = response.clickedButtonId();
            List<MenuItem> clickableItems = items.stream()
                .filter(i -> !i.isInfoOnly() && i.isClickable())
                .toList();
            
            if (index >= 0 && index < clickableItems.size()) {
                clickableItems.get(index).onClick(player);
            }
        });
        
        FloodgateApi.getInstance().sendForm(player.getUuid(), builder.build());
    }
    
    private void openInventory(KaraPlayer player) {
        String title = player.t(titleKey);
        Inventory inv = Bukkit.createInventory(null, rows * 9, title);
        
        for (MenuItem item : items) {
            if (item.getSlot() >= 0) {
                inv.setItem(item.getSlot(), item.toItemStack(player));
            }
        }
        
        player.getBukkitPlayer().openInventory(inv);
    }
    
    // Builder pattern
    public static Builder builder() {
        return new Builder();
    }
    
    public static class Builder {
        private String titleKey;
        private int rows = 3;
        private List<MenuItem> items = new ArrayList<>();
        private MenuType type = MenuType.CHEST;
        
        public Builder title(String key) {
            this.titleKey = key;
            return this;
        }
        
        public Builder rows(int rows) {
            this.rows = rows;
            return this;
        }
        
        public Builder item(MenuItem item) {
            this.items.add(item);
            return this;
        }
        
        public Builder type(MenuType type) {
            this.type = type;
            return this;
        }
        
        public KaraMenu build() {
            return new KaraMenu(titleKey, rows, items, type);
        }
    }
}

public enum MenuType {
    CHEST(false),           // Her zaman inventory
    FORM_PREFERRED(true),   // Bedrock'ta form, Java'da inventory
    FORM_ONLY(true);        // Sadece Bedrock'ta açılır
    
    private final boolean hasBedrockForm;
    
    MenuType(boolean hasBedrockForm) {
        this.hasBedrockForm = hasBedrockForm;
    }
    
    public boolean hasBedrockForm() {
        return hasBedrockForm;
    }
}
```

### MenuItem Sınıfı

```java
public class MenuItem {
    private final String nameKey;
    private final List<String> loreKeys;
    private final Material material;
    private final int slot;
    private final Consumer<KaraPlayer> onClick;
    private final FormImage formImage;
    private final boolean clickable;
    private final boolean infoOnly;
    
    // Bedrock için
    public FormImage getFormImage() {
        if (formImage != null) {
            return formImage;
        }
        // Material'dan otomatik texture path
        return FormImage.of(
            FormImage.Type.PATH,
            "textures/items/" + material.name().toLowerCase()
        );
    }
    
    // Java için
    public ItemStack toItemStack(KaraPlayer player) {
        ItemStack item = new ItemStack(material);
        ItemMeta meta = item.getItemMeta();
        
        // İsim
        meta.displayName(Text.parse(player.t(nameKey)));
        
        // Lore
        List<Component> lore = new ArrayList<>();
        for (String loreKey : loreKeys) {
            String loreLine = player.t(loreKey);
            for (String line : loreLine.split("\n")) {
                lore.add(Text.parse(line));
            }
        }
        meta.lore(lore);
        
        // Tıklanamaz ise enchant glow ekle
        if (!clickable) {
            meta.addEnchant(Enchantment.LUCK, 1, true);
            meta.addItemFlags(ItemFlag.HIDE_ENCHANTS);
        }
        
        item.setItemMeta(meta);
        return item;
    }
    
    public void onClick(KaraPlayer player) {
        if (clickable && onClick != null) {
            onClick.accept(player);
        }
    }
    
    // Builder
    public static Builder builder() {
        return new Builder();
    }
    
    public static class Builder {
        private String nameKey;
        private List<String> loreKeys = new ArrayList<>();
        private Material material = Material.STONE;
        private int slot = -1;
        private Consumer<KaraPlayer> onClick;
        private FormImage formImage;
        private boolean clickable = true;
        private boolean infoOnly = false;
        
        public Builder name(String key) {
            this.nameKey = key;
            return this;
        }
        
        public Builder lore(String... keys) {
            this.loreKeys.addAll(Arrays.asList(keys));
            return this;
        }
        
        public Builder material(Material material) {
            this.material = material;
            return this;
        }
        
        public Builder slot(int slot) {
            this.slot = slot;
            return this;
        }
        
        public Builder onClick(Consumer<KaraPlayer> action) {
            this.onClick = action;
            return this;
        }
        
        public Builder formImage(String path) {
            this.formImage = FormImage.of(FormImage.Type.PATH, path);
            return this;
        }
        
        public Builder formImageUrl(String url) {
            this.formImage = FormImage.of(FormImage.Type.URL, url);
            return this;
        }
        
        public Builder notClickable() {
            this.clickable = false;
            return this;
        }
        
        public Builder infoOnly() {
            this.infoOnly = true;
            this.clickable = false;
            return this;
        }
        
        public MenuItem build() {
            return new MenuItem(nameKey, loreKeys, material, slot, 
                               onClick, formImage, clickable, infoOnly);
        }
    }
}
```

---

## Mobil UI Tasarım Kuralları

### Touch-Friendly Design

```
┌─────────────────────────────────────────────────────────────────┐
│                 MOBİL UI TASARIM KURALLARI                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  📐 BOYUT & SPACING                                            │
│  ──────────────────                                             │
│  ├── Minimum touch target: 44x44 piksel (1 slot)              │
│  ├── İdeal touch target: 88x88 piksel (2x2 slot)              │
│  ├── Butonlar arası boşluk: en az 1 slot                       │
│  └── Ekran kenarlarından uzak tut (scroll alanı)              │
│                                                                 │
│  📝 METİN                                                      │
│  ────────                                                       │
│  ├── Buton isimleri: max 15 karakter                           │
│  ├── Lore satırı: max 30 karakter                              │
│  ├── Form butonları: max 20 karakter                           │
│  ├── Emoji kullan (evrensel, hızlı anlam)                      │
│  └── BÜYÜK HARF sadece başlıklarda                             │
│                                                                 │
│  🎨 RENK                                                       │
│  ──────                                                         │
│  ├── Yüksek kontrast (güneş altında okunabilirlik)            │
│  ├── Yeşil = pozitif/onay/git                                  │
│  ├── Kırmızı = negatif/iptal/tehlike                          │
│  ├── Sarı = dikkat/bilgi/uyarı                                │
│  ├── Mavi = bilgi/link                                         │
│  └── Gri = devre dışı/unavailable                             │
│                                                                 │
│  📱 LAYOUT                                                     │
│  ────────                                                       │
│  ├── Maksimum 3 satır tercih et (27 slot)                     │
│  ├── 6 satır sadece envanter görünümü için                    │
│  ├── Önemli butonlar ortada/üstte                             │
│  ├── Geri/Kapat her zaman sağ alt (slot 26)                   │
│  └── Scroll gerektiren tasarımlardan kaçın                     │
│                                                                 │
│  ⏱️ FEEDBACK                                                   │
│  ──────────                                                     │
│  ├── Her tıklamada ses çal                                     │
│  ├── Loading durumlarında action bar mesajı                   │
│  ├── Spam koruması (cooldown)                                  │
│  └── Başarı/hata durumunda görsel feedback                    │
│                                                                 │
│  🎮 BEDROCK-SPECIFIC                                           │
│  ──────────────────                                             │
│  ├── Form tercih et (native, hızlı)                           │
│  ├── Uzun listeler için dropdown kullan                       │
│  ├── Tek sayfa tasarımı (pagination'dan kaçın)                │
│  └── Controller uyumlu navigation                              │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Örnek Layout'lar

```
3 Satır Ana Menü (Mobil-Optimized):
┌─────────────────────────────────────┐
│     │ 10  │     │ 12  │     │ 14  │ │  ← Ana butonlar
│─────│─────│─────│─────│─────│─────│─│
│     │     │     │ 22  │     │     │ │  ← Bilgi
│─────│─────│─────│─────│─────│─────│─│
│     │     │     │     │     │ 26  │ │  ← Geri/Kapat
└─────────────────────────────────────┘

6 Satır Envanter Görünümü:
┌─────────────────────────────────────┐
│  0  │  1  │  2  │  3  │  4  │  5  │ │
│─────│─────│─────│─────│─────│─────│─│
│  9  │ 10  │ 11  │ 12  │ 13  │ 14  │ │  ← Item grid
│─────│─────│─────│─────│─────│─────│─│
│ 18  │ 19  │ 20  │ 21  │ 22  │ 23  │ │
│─────│─────│─────│─────│─────│─────│─│
│ 27  │ 28  │ 29  │ 30  │ 31  │ 32  │ │
│─────│─────│─────│─────│─────│─────│─│
│ 36  │ 37  │ 38  │ 39  │ 40  │ 41  │ │
│─────│─────│─────│─────│─────│─────│─│
│ 45  │ 46  │ 47  │ 48  │ 49  │ BACK │ │  ← Navigation
└─────────────────────────────────────┘
```

---

## Bedrock Otomatik Login

### Floodgate Auth Bypass

```java
// karapixel-auth içinde
@EventHandler(priority = EventPriority.LOWEST)
public void onJoin(PlayerJoinEvent event) {
    Player player = event.getPlayer();
    KaraPlayer karaPlayer = KaraAPI.getPlayer(player);
    
    // Bedrock oyuncu kontrolü
    if (karaPlayer.isBedrock()) {
        // Floodgate ile giriş yapmış, Xbox auth güvenilir
        if (FloodgateApi.getInstance().isFloodgatePlayer(player.getUniqueId())) {
            // Session oluştur
            sessionManager.createSession(karaPlayer);
            
            // Hoşgeldin mesajı
            karaPlayer.sendMessage("auth.auto_login");
            
            // Direkt Hub'a gönder (Limbo bypass)
            Bukkit.getScheduler().runTaskLater(plugin, () -> {
                teleportToHub(karaPlayer);
            }, 20L);  // 1 saniye delay
            
            return;
        }
    }
    
    // Java oyuncu - normal auth akışı
    startAuthProcess(karaPlayer);
}
```

### Bedrock Oyuncu İsmi

```java
// Bedrock oyuncu ismi temizleme
public class BedrockNameUtil {
    
    /**
     * Bedrock oyuncu ismini temizler
     * "*Steve123" -> "Steve123"
     */
    public static String cleanName(String name) {
        if (name.startsWith("*")) {
            return name.substring(1);
        }
        return name;
    }
    
    /**
     * Bedrock oyuncu ismini gösterim için formatlar
     */
    public static String formatName(KaraPlayer player) {
        String name = player.getName();
        if (player.isBedrock()) {
            // Prefix'i kaldır ve Bedrock badge ekle
            return "📱 " + cleanName(name);
        }
        return name;
    }
    
    /**
     * Oyuncu aramasında Bedrock prefix'i handle et
     */
    public static Player findPlayer(String query) {
        // Önce direkt ara
        Player player = Bukkit.getPlayer(query);
        if (player != null) return player;
        
        // Bedrock prefix ile ara
        player = Bukkit.getPlayer("*" + query);
        if (player != null) return player;
        
        // Partial match
        for (Player online : Bukkit.getOnlinePlayers()) {
            String cleanName = cleanName(online.getName());
            if (cleanName.toLowerCase().startsWith(query.toLowerCase())) {
                return online;
            }
        }
        
        return null;
    }
}
```

---

## Performans Optimizasyonu

### Geyser Resource Kullanımı

```yaml
# Geyser performans ayarları
# plugins/Geyser-Velocity/config.yml

# Chunk cache (RAM vs CPU trade-off)
cache-chunks: true
cache-images: 0  # 0 = cache'leme (RAM tasarrufu)

# Sıkıştırma seviyesi (1-9)
# Düşük = daha az CPU, daha fazla bandwidth
# Yüksek = daha fazla CPU, daha az bandwidth
compression-level: 6  # Dengeli

# Pending connection timeout
pending-authentication-timeout: 120000

# Thread pool
# Varsayılan genellikle yeterli
```

### Bedrock Oyuncu Overhead

```
┌─────────────────────────────────────────────────────────────────┐
│              BEDROCK OVERHEAD ANALİZİ                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  JAVA OYUNCU                                                    │
│  ├── Ortalama RAM: ~150-200 MB                                 │
│  └── Network: Native protocol                                  │
│                                                                 │
│  BEDROCK OYUNCU (Geyser)                                       │
│  ├── Ortalama RAM: ~200-280 MB (+30-50% overhead)             │
│  ├── Protocol translation: ~5-10ms latency                     │
│  └── Chunk conversion: CPU intensive                           │
│                                                                 │
│  OPTİMİZASYON ÖNERİLERİ                                        │
│  ├── view-distance: 6-8 (Bedrock için yeterli)                │
│  ├── simulation-distance: 4-6                                  │
│  ├── Chunk pre-generation (startup overhead azalt)            │
│  ├── Entity limits (mobil cihazlar için önemli)               │
│  └── Particle limits (mobil performans)                        │
│                                                                 │
│  TAHMİNİ TOPLAM OVERHEAD                                       │
│  ├── %40 Bedrock oyuncu oranında: ~15-20% ekstra kaynak       │
│  └── Kapasite etkisi: ~100-150 oyuncu azalma (1000→850-900)   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Test Checklist

### Bedrock Test Matrisi

| Özellik | Android | iOS | Xbox | Win10 |
|---------|:-------:|:---:|:----:|:-----:|
| Bağlantı | ☐ | ☐ | ☐ | ☐ |
| Otomatik login | ☐ | ☐ | ☐ | ☐ |
| Form menüler | ☐ | ☐ | ☐ | ☐ |
| Chest menüler | ☐ | ☐ | ☐ | ☐ |
| 3D modeller | ☐ | ☐ | ☐ | ☐ |
| Resource pack | ☐ | ☐ | ☐ | ☐ |
| Chat/Komutlar | ☐ | ☐ | ☐ | ☐ |
| Teleport | ☐ | ☐ | ☐ | ☐ |
| Skyblock oynanış | ☐ | ☐ | ☐ | ☐ |

### Test Senaryoları

1. **İlk bağlantı:** Yeni Bedrock oyuncu bağlanır, otomatik login olur
2. **Form test:** Tüm form tipleri açılır ve çalışır
3. **Cross-platform:** Java ve Bedrock oyuncu aynı adada
4. **Performans:** 100+ Bedrock oyuncu ile TPS testi
5. **Resource pack:** 3D modeller Bedrock'ta görünür

---

*📅 Son güncelleme: 24 Aralık 2024*
