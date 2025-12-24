# 🔧 KaraPaper Fork Dokümantasyonu

> Purpur tabanlı, Skyblock ve Türkçe odaklı custom Minecraft server fork'u.

---

## Genel Bakış

```
┌──────────────────────────────────────────────────────────────────┐
│                     KARAPAPER FORK YAPISI                        │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  Vanilla 1.21.10                                                 │
│       │                                                          │
│       ▼                                                          │
│  CraftBukkit ──► Spigot ──► Paper ──► Pufferfish ──► Purpur     │
│                                                          │       │
│                                                          ▼       │
│                                              ┌───────────────┐   │
│                                              │   KARAPAPER   │   │
│                                              │   v1.21.10    │   │
│                                              └───────────────┘   │
│                                                                  │
│  KATMANLAR:                                                      │
│  ├── 🔴 Core Layer     : Branding, Türkçe, Güvenlik             │
│  ├── 🟡 Optimize Layer : Performance, Memory, Async             │
│  ├── 🟢 Skyblock Layer : Void world, Island, Mob Stacking       │
│  ├── 🔵 Content Layer  : Built-in API'ler                       │
│  └── 🟣 Bedrock Layer  : Geyser optimization                    │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Sürüm Stratejisi

```
┌──────────────────────────────────────────────────────────────────┐
│                      SÜRÜM STRATEJİSİ                            │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  KARAPAPER ANA SÜRÜM: 1.21.10                                   │
│                                                                  │
│  ═══════════════════════════════════════════════════════════    │
│                                                                  │
│  JAVA EDITION DESTEĞİ (ViaVersion + ViaBackwards):              │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                                                         │    │
│  │  1.21.10 ◄── Native (sunucu sürümü)                    │    │
│  │  1.21.x  ◄── ✅ Tam destek (1.21 - 1.21.9)             │    │
│  │  1.20.x  ◄── ✅ Tam destek (1.20 - 1.20.6)             │    │
│  │  1.19.x  ◄── ❌ Desteklenmez                           │    │
│  │                                                         │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                  │
│  BEDROCK EDITION DESTEĞİ (Geyser):                              │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                                                         │    │
│  │  Geyser otomatik olarak en güncel Bedrock'u destekler  │    │
│  │  Bedrock oyuncuları zaten güncel olmak ZORUNDA         │    │
│  │  (Microsoft Store/Xbox otomatik güncelleme)            │    │
│  │                                                         │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                  │
│  GÜNCELLEME POLİTİKASI:                                         │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                                                         │    │
│  │  • Minecraft major update çıktığında:                  │    │
│  │    → Purpur güncellendikten sonra takip et             │    │
│  │    → Test sunucusunda kontrol et                       │    │
│  │    → Sonra production'a al                             │    │
│  │                                                         │    │
│  │  • Minor/hotfix update çıktığında:                     │    │
│  │    → Purpur güncellenir güncellenmez takip et          │    │
│  │                                                         │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Patch Kategorileri

### 🔴 CORE LAYER (Temel)

| # | Patch Adı | Açıklama | Öncelik |
|---|-----------|----------|---------|
| 0001 | Rebrand-to-KaraPaper | Sunucu adı, versiyon, motd | 🔴 |
| 0002 | Turkish-Locale-Default | Varsayılan dil Türkçe | 🔴 |
| 0003 | Turkish-Death-Messages | Türkçe ölüm mesajları | 🔴 |
| 0004 | Turkish-Item-Names | Türkçe item isimleri (chat dahil) | 🔴 |
| 0005 | Turkish-Mob-Names | Türkçe mob isimleri | 🟡 |
| 0006 | Turkish-Enchant-Names | Türkçe büyü isimleri | 🟡 |
| 0007 | Turkish-Potion-Names | Türkçe iksir isimleri | 🟡 |
| 0008 | Turkish-Biome-Names | Türkçe biyom isimleri | 🟢 |
| 0009 | Hide-Server-Brand | /version gizleme, güvenlik | 🔴 |
| 0010 | Exploit-Patches | Bilinen crash exploitleri | 🔴 |
| 0011 | Configurable-MOTD | Dinamik MOTD sistemi | 🟡 |

### 🟡 OPTIMIZE LAYER (Performans)

| # | Patch Adı | Açıklama | Kazanç |
|---|-----------|----------|--------|
| 0020 | Remove-Advancement-System | Başarım sistemi kaldır | +5% |
| 0021 | Remove-Recipe-Book | Tarif kitabı kaldır | +3% |
| 0022 | Remove-Unused-Registries | Kullanılmayan registry temizle | +2% |
| 0023 | Async-Entity-Tracker | Entity tracking async | +10% |
| 0024 | Async-Pathfinding | Mob pathfinding async | +8% |
| 0025 | Optimized-Tick-Loop | Tick loop optimizasyonu | +5% |
| 0026 | Reduced-Packet-Sending | Gereksiz paket azaltma | +5% |
| 0027 | Lazy-Chunk-Loading | Tembel chunk yükleme | +7% |
| 0028 | Entity-Activation-Range | Akıllı entity aktivasyonu | +8% |
| 0029 | Hopper-Optimization | Hopper tick azaltma | +5% |
| 0030 | Redstone-Optimization | Redstone hesaplama optimize | +3% |
| 0031 | Collision-Optimization | Collision check optimize | +4% |
| 0032 | Light-Engine-Optimization | Işık motoru optimize | +5% |
| 0033 | Memory-Pool-Reuse | Object pooling | +3% |
| 0034 | GC-Friendly-Collections | GC-dostu koleksiyonlar | +2% |
| 0035 | Chunk-Cache-Improvement | Chunk cache büyütme | +3% |

### 🟢 SKYBLOCK LAYER (Skyblock Özel)

| # | Patch Adı | Açıklama | Kazanç |
|---|-----------|----------|--------|
| 0040 | Void-World-Generator | Hızlı void world üretimi | +15% |
| 0041 | Island-Chunk-Optimization | Ada chunk'ları optimize | +10% |
| 0042 | Skip-Empty-Chunk-Ticks | Boş chunk tick'leme | +8% |
| 0043 | Island-Border-Caching | Ada sınır cache | +3% |
| 0044 | Optimized-Block-Place | Blok koyma optimize | +2% |
| 0045 | Optimized-Block-Break | Blok kırma optimize | +2% |
| 0046 | Generator-Block-Hook | Tier sistemi plugin'de | - |
| 0047 | Island-Protection-Native | Native ada koruması | +5% |
| 0048 | Coop-Permission-Cache | Coop izin cache | +2% |
| 0049 | Fast-Island-Teleport | Hızlı ada teleport | +1% |
| 0050 | Island-Value-Calculator-Hook | Native değer hesaplama | +3% |
| 0051 | Native-Mob-Stacking | Native mob stacking ⭐ | +10% |
| 0052 | Spawner-Optimization | Spawner optimize | +5% |
| 0053 | Crop-Growth-Batch | Ekin büyüme batch | +3% |


### 🔵 CONTENT LAYER (İçerik & API)

| # | Patch Adı | Açıklama | Detay |
|---|-----------|----------|-------|
| 0060 | KaraPlayer-API | Genişletilmiş Player API | Platform, locale, session |
| 0061 | KaraWorld-API | Genişletilmiş World API | Island support |
| 0062 | KaraEvent-System | Lightweight event sistemi | Async events |
| 0063 | KaraCommand-Framework | Komut framework | Auto-complete, Türkçe |
| 0064 | KaraGUI-Framework | GUI framework | Bedrock uyumlu |
| 0065 | KaraMessaging-API | Cross-server messaging | Redis entegre |
| 0066 | Built-in-Hologram | Native hologram sistemi | Performanslı |
| 0067 | Built-in-Scoreboard | Native scoreboard API | Flicker-free |
| 0068 | Built-in-Bossbar | Native bossbar API | Animasyonlu |
| 0069 | Built-in-Actionbar | Native actionbar API | - |
| 0070 | Built-in-Title | Native title API | Fade control |
| 0071 | Built-in-Tablist | Native tablist API | Sorting, prefix |
| 0072 | Prometheus-Metrics | Built-in metrik endpoint | /metrics |
| 0073 | Server-Profile-System | Runtime profil sistemi | hub/skyblock/auth |
| 0074 | Native-PlaceholderAPI | Built-in placeholder | %player%, %island% |
| 0075 | Native-Economy-Hook | Economy API hook | Vault benzeri |

### 🟣 BEDROCK LAYER (Bedrock Optimizasyon)

| # | Patch Adı | Açıklama | Kazanç |
|---|-----------|----------|--------|
| 0080 | Bedrock-Packet-Batching | Paket birleştirme | +5% |
| 0081 | Bedrock-Chunk-Sending | Chunk gönderim optimize | +5% |
| 0082 | Bedrock-Entity-Metadata | Entity metadata optimize | +3% |
| 0083 | Bedrock-Inventory-Sync | Envanter senkron optimize | +2% |
| 0084 | Geyser-Cumulus-Hook | Floodgate Forms entegre | - |
| 0085 | Bedrock-Resource-Pack | RP gönderim optimize | +2% |
| 0086 | Geyser-Skin-Cache | Skin cache sistemi | +1% |
| 0087 | Floodgate-Deep-Integration | Floodgate derin entegrasyon | - |

---

## Performans Hedefleri

```
┌──────────────────────────────────────────────────────────────────┐
│                    PERFORMANS HEDEFLERİ                          │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  VANILLA PAPER (Baseline)                                        │
│  └── 100 oyuncu @ 20 TPS                                        │
│                                                                  │
│  PURPUR                                                          │
│  └── 120 oyuncu @ 20 TPS (+20%)                                 │
│                                                                  │
│  KARAPAPER (Hedef)                                               │
│  └── 200+ oyuncu @ 19-20 TPS (+100%)                            │
│                                                                  │
│  ═══════════════════════════════════════════════════════════    │
│                                                                  │
│  KATMAN BAZLI KAZANÇ:                                            │
│  ├── Core Layer        :  +0%  (branding, locale)               │
│  ├── Optimize Layer    : +35%  (async, memory, tick)            │
│  ├── Skyblock Layer    : +40%  (void, island, mob stack)        │
│  ├── Content Layer     : +10%  (native API, no external)        │
│  └── Bedrock Layer     : +15%  (packet, chunk optimize)         │
│                          ─────                                   │
│                    TOPLAM: ~100% kapasite artışı                 │
│                                                                  │
│  SKYBLOCK WORLD BAŞINA:                                          │
│  ├── Paper Vanilla     : ~100-120 oyuncu                        │
│  ├── Purpur            : ~140-160 oyuncu                        │
│  └── KaraPaper         : ~250-300 oyuncu ✓                      │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Türkçe Desteği

```
┌──────────────────────────────────────────────────────────────────┐
│                     TÜRKÇE DESTEK SİSTEMİ                        │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  VARSAYILAN DİL: tr_TR                                           │
│  KAPSAM: Chat dahil tüm sistem                                  │
│                                                                  │
│  1. ÖLÜM MESAJLARI                                              │
│  ─────────────────                                               │
│  "X was slain by Y"    → "X, Y tarafından öldürüldü"           │
│  "X fell from high"    → "X yüksekten düşerek öldü"            │
│  "X drowned"           → "X boğularak öldü"                    │
│  "X burned"            → "X yanarak öldü"                      │
│                                                                  │
│  2. ITEM İSİMLERİ (Chat dahil)                                  │
│  ──────────────────────────────                                  │
│  "Diamond Sword"       → "Elmas Kılıç"                         │
│  "Cobblestone"         → "Kaldırım Taşı"                       │
│  "Oak Planks"          → "Meşe Tahta"                          │
│                                                                  │
│  3. MOB İSİMLERİ                                                │
│  ───────────────                                                 │
│  "Zombie"              → "Zombi"                               │
│  "Skeleton"            → "İskelet"                             │
│  "Iron Golem"          → "Demir Golem"                         │
│                                                                  │
│  4. BÜYÜLER                                                     │
│  ──────────                                                      │
│  "Sharpness"           → "Keskinlik"                           │
│  "Protection"          → "Koruma"                              │
│  "Efficiency"          → "Verimlilik"                          │
│  "Fortune"             → "Şans"                                │
│                                                                  │
│  5. İKSİRLER                                                    │
│  ───────────                                                     │
│  "Speed"               → "Hız"                                 │
│  "Strength"            → "Güç"                                 │
│  "Regeneration"        → "Yenilenme"                           │
│                                                                  │
│  6. SİSTEM MESAJLARI                                            │
│  ───────────────────                                             │
│  "Player joined"       → "Oyuncu katıldı"                      │
│  "Unknown command"     → "Bilinmeyen komut"                    │
│                                                                  │
│  AYAR:                                                          │
│  karapaper.yml:                                                 │
│    force-locale: true  # Her zaman tr_TR göster                │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Native Mob Stacking

```
┌──────────────────────────────────────────────────────────────────┐
│                   NATIVE MOB STACKING                            │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  NEDEN FORK'TA (Plugin yerine)?                                 │
│  ──────────────────────────────                                  │
│  ├── Zero event overhead                                        │
│  ├── Native memory layout                                       │
│  ├── No reflection                                              │
│  ├── 7x daha hızlı (benchmark)                                 │
│  └── NPC/Pet bypass native level'da                            │
│                                                                  │
│  STACK LİMİTLERİ:                                               │
│  ─────────────────                                               │
│  │ Mob Türü          │ Max Stack │                             │
│  ├────────────────────┼───────────┤                             │
│  │ Tavuk              │ 128       │                             │
│  │ İnek/Koyun/Domuz   │ 64        │                             │
│  │ Zombi/İskelet      │ 64        │                             │
│  │ Creeper            │ 32        │                             │
│  │ Enderman           │ 16        │                             │
│  │ Blaze              │ 32        │                             │
│  │ Villager           │ 8         │                             │
│  │ Iron Golem         │ 4         │                             │
│  └────────────────────┴───────────┘                             │
│                                                                  │
│  STACK YAPILMAYACAKLAR:                                         │
│  ──────────────────────                                          │
│  ├── NPC'ler (Citizens metadata: "NPC")                        │
│  ├── Pet'ler (Tamed veya "pet" metadata)                       │
│  ├── İsimlendirilmiş moblar (nametag)                          │
│  ├── Boss moblar (Wither, Dragon)                              │
│  ├── ModelEngine mobları                                        │
│  └── Armor stand, Item frame                                   │
│                                                                  │
│  HASAR MEKANİĞİ:                                                │
│  ───────────────                                                 │
│  ├── Tek vuruş → Tüm stack'e hasar                             │
│  ├── Öldürme → 1 mob ölür, stack -1                            │
│  ├── Drop → Her ölen mob için ayrı hesap                       │
│  └── Looting → Her ölüm için ayrı Looting                      │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```


---

## Generator Block Hook

```
┌──────────────────────────────────────────────────────────────────┐
│                   GENERATOR BLOCK HOOK                           │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  YAKLAŞIM: Hook fork'ta, tier sistemi plugin'de                 │
│  ─────────────────────────────────────────────                   │
│                                                                  │
│  NEDEN?                                                          │
│  ├── Tier ayarları sık değişebilir                              │
│  ├── Fork rebuild gerektirmez                                   │
│  ├── Config ile kolay ayarlama                                  │
│  └── Esneklik                                                   │
│                                                                  │
│  FORK'TA:                                                        │
│  ─────────                                                       │
│  // Lav + Su temas event hook                                   │
│  GeneratorBlockEvent event = new GeneratorBlockEvent(loc);      │
│  if (!event.isCancelled()) {                                    │
│      Material ore = event.getResultBlock();                     │
│      setBlock(loc, ore);                                        │
│  }                                                              │
│                                                                  │
│  PLUGIN'DE:                                                      │
│  ───────────                                                     │
│  @EventHandler                                                   │
│  public void onGenerator(GeneratorBlockEvent e) {               │
│      Island island = getIsland(e.getLocation());                │
│      GeneratorTier tier = island.getGeneratorTier();            │
│      e.setResultBlock(tier.getRandomOre());                     │
│  }                                                              │
│                                                                  │
│  TIER ÖRNEĞİ (Config):                                          │
│  ──────────────────────                                          │
│  tiers:                                                         │
│    tier1:                                                       │
│      cobblestone: 85%                                           │
│      coal_ore: 10%                                              │
│      iron_ore: 5%                                               │
│    tier2:                                                       │
│      cobblestone: 70%                                           │
│      coal_ore: 15%                                              │
│      iron_ore: 10%                                              │
│      gold_ore: 5%                                               │
│    # ... devamı                                                 │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Bedrock Forms (Geyser/Cumulus)

```
┌──────────────────────────────────────────────────────────────────┐
│                   BEDROCK FORMS YAKLAŞIMI                        │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  KARAR: Geyser/Cumulus API kullan ✅                            │
│  ─────────────────────────────────────                           │
│                                                                  │
│  SEBEPLER:                                                       │
│  ├── Geyser zaten kullanıyoruz (Bedrock desteği için)          │
│  ├── Cumulus API stabil ve iyi dokümante                       │
│  ├── Floodgate entegrasyonu hazır                               │
│  ├── Daha az kod = daha az hata payı                           │
│  └── Geyser takımı Bedrock uzmanı                              │
│                                                                  │
│  FORK'TA YAPILACAK:                                             │
│  ──────────────────                                              │
│  ├── KaraPaper.isBedrock(player) helper                        │
│  ├── KaraPaper.sendForm(player, form) wrapper                  │
│  └── Cumulus API expose (plugin'ler için)                      │
│                                                                  │
│  KULLANIM ÖRNEĞİ:                                               │
│  ─────────────────                                               │
│  if (KaraPaper.isBedrock(player)) {                            │
│      SimpleForm form = SimpleForm.builder()                    │
│          .title("Ada Menüsü")                                  │
│          .content("Ne yapmak istersin?")                       │
│          .button("Adama Git")                                  │
│          .button("Üyeler")                                     │
│          .button("Ayarlar")                                    │
│          .build();                                             │
│      KaraPaper.sendForm(player, form);                        │
│  } else {                                                      │
│      openJavaGUI(player);                                      │
│  }                                                             │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Built-in API'ler

```java
// KaraPlayer API
KaraPlayer player = KaraPaper.getPlayer(uuid);
player.getPlatform();        // JAVA, BEDROCK
player.getLocale();          // tr_TR, en_US
player.getSession();         // Session bilgisi
player.getIsland();          // Ada bilgisi (nullable)
player.sendForm(form);       // Bedrock form gönder
player.sendGUI(gui);         // GUI aç (platform-aware)

// KaraWorld API
KaraWorld world = KaraPaper.getWorld("skyblock-1");
world.getProfile();          // SKYBLOCK, HUB, AUTH
world.getIslands();          // Ada listesi
world.getIslandAt(loc);      // Lokasyondaki ada

// KaraIsland API
KaraIsland island = player.getIsland();
island.getOwner();           // Ada sahibi
island.getMembers();         // Üyeler
island.getValue();           // Ada değeri
island.getLevel();           // Ada seviyesi

// KaraMessaging API (Redis)
KaraMessaging.publish("channel", data);
KaraMessaging.subscribe("channel", handler);

// KaraGUI API
KaraGUI gui = KaraGUI.builder()
    .title("§6Market")
    .rows(6)
    .item(slot, item, clickHandler)
    .bedrockForm(formBuilder)  // Bedrock alternatif
    .build();
gui.open(player);

// KaraHologram API
KaraHologram holo = KaraHologram.create(location)
    .addLine("§e§lTOP ADALAR")
    .addLine("%top_island_1%")
    .refreshRate(20)
    .build();

// KaraScoreboard API
KaraScoreboard sb = KaraScoreboard.create(player)
    .title("§6§lKARAPİXEL")
    .line(14, "§7Ada: §f%island_name%")
    .line(13, "§7Seviye: §f%island_level%")
    .build();
```

---

## Proje Yapısı

```
karapaper/
├── patches/
│   ├── server/
│   │   ├── 0001-Rebrand-to-KaraPaper.patch
│   │   ├── 0002-Turkish-Locale-Default.patch
│   │   ├── ...
│   │   └── 0087-Floodgate-Integration.patch
│   └── api/
│       ├── 0001-KaraPlayer-API.patch
│       └── ...
│
├── KaraPaper-API/
│   └── src/main/java/
│       └── net/karapixel/paper/
│           ├── KaraPaper.java
│           ├── player/
│           ├── world/
│           ├── island/
│           ├── gui/
│           ├── messaging/
│           ├── display/
│           └── placeholder/
│
├── KaraPaper-Server/
│   └── src/main/java/
│       └── net/karapixel/paper/
│           └── ... (implementation)
│
├── locales/
│   └── tr_TR/
│       ├── items.json
│       ├── mobs.json
│       ├── enchants.json
│       ├── potions.json
│       ├── biomes.json
│       ├── death_messages.json
│       └── system.json
│
├── scripts/
│   ├── build.sh
│   ├── applyPatches.sh
│   ├── rebuildPatches.sh
│   └── updateUpstream.sh
│
├── build.gradle.kts
└── README.md
```

---

## Kararlar Özeti

| Konu | Karar | Sebep |
|------|-------|-------|
| **Base Fork** | Purpur 1.21.10 | En güncel, en optimize |
| **Java Sürüm Desteği** | 1.20.x - 1.21.10 | Son 2 major (ViaVersion) |
| **Bedrock Desteği** | Geyser otomatik | Microsoft zorla günceller |
| **Mob Stacking** | Native (Fork'ta) | 7x daha hızlı, NPC/Pet korumalı |
| **Generator** | Hook fork'ta, tier plugin'de | Esneklik |
| **Türkçe** | Chat dahil tam Türkçe | Türk hedef kitle |
| **Bedrock Forms** | Geyser/Cumulus | Stabil, az hata payı |

---

## Toplam Patch Sayısı

```
🔴 Core Layer     : 11 patch
🟡 Optimize Layer : 16 patch
🟢 Skyblock Layer : 14 patch
🔵 Content Layer  : 16 patch
🟣 Bedrock Layer  :  8 patch
─────────────────────────────
TOPLAM            : 65 patch
```

---

*📅 Son güncelleme: 24 Aralık 2024*
