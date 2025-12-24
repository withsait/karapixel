# 🏰 KaraPixel Dungeon Sistemi

> **Durum:** 📝 Planlama Aşaması
> **Son Güncelleme:** 2024-12-24

---

## ✅ Alınan Kararlar

| Konu | Karar | Tarih |
|------|-------|-------|
| Giriş Sistemi | Model E (Hibrit) - Tier bazlı | 2024-12-24 |
| Party PvP | Kontrollü - Normal kapalı, Chaos event açık | 2024-12-24 |
| Instance | Hibrit - Story özel, Grind shared | 2024-12-24 |
| Mob Modelleri | Tier bazlı - Vanilla/Textured/Custom | 2024-12-24 |
| Instance Yönetimi | Slime World Manager | 2024-12-24 |
| Dungeon Yapısı | Mixed - Tier bazlı farklı yapılar | 2024-12-24 |
| Loot Sistemi | Hibrit - Personal + Roll + Contribution | 2024-12-24 |
| Wipe Sistemi | Casual friendly - 3 can + %10 HP loss | 2024-12-24 |

---

## 🗺️ Dungeon Haritası

```
┌─────────────────────────────────────────────────────────────────┐
│                    GÖLGE REALM YAPISI                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  🌑 GÖLGE SANCTUARY (Hub)                                       │
│  ├── Gölge NPC (rehber, hikaye anlatıcı)                       │
│  ├── Prestige altar                                             │
│  ├── Özel shop (dungeon currency)                              │
│  └── Lore kitaplığı                                            │
│                                                                  │
│  ⚔️ TIER 1-3: BAŞLANGIÇ DUNGEON'LARI                           │
│  ├── Yapı: Lineer (Oda → Oda → Boss)                          │
│  ├── Giriş: Günlük limit (3/gün, VIP: 5)                      │
│  ├── Instance: Özel (party başına)                             │
│  │                                                              │
│  ├── T1: Kayıp Madenleri (Ada Lv.30+)                         │
│  │   └── Boss: Kemik Muhafız                                   │
│  ├── T2: Donmuş Saray (Ada Lv.50+)                            │
│  │   └── Boss: Buz Kraliçesi                                   │
│  └── T3: Lanetli Orman (Ada Lv.75+)                           │
│      └── Boss: Orman Canavarı                                  │
│                                                                  │
│  ⚔️ TIER 4-5: İLERİ DUNGEON'LAR                                │
│  ├── Yapı: Branching (Çoklu yol, farklı ödüller)              │
│  ├── Giriş: Key sistemi (drop + quest reward)                  │
│  ├── Instance: Shared (max 3 party)                            │
│  │                                                              │
│  ├── T4: Cehennem Kapısı (Ada Lv.100+)                        │
│  │   └── Boss: Ateş Lordu                                      │
│  └── T5: Gölge Kalesi (Prestige 1+)                           │
│      └── Boss: Gölge Lordu Kael                                │
│                                                                  │
│  👑 BOSS ARENA                                                  │
│  ├── Giriş: Haftalık (1/hafta)                                │
│  ├── Haftalık Boss (değişen)                                   │
│  ├── Prestige Boss (solo challenge)                            │
│  └── Seasonal Boss (event)                                      │
│                                                                  │
│  🎲 CHAOS DUNGEON (Event)                                       │
│  ├── Yapı: Roguelike (rastgele odalar)                        │
│  ├── PvP: Açık (loot için rekabet)                            │
│  ├── Giriş: Event sırasında özel key                          │
│  └── Instance: Shared + PvP                                    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## 👹 Mob Sistemi

### Mob Tierleri

```
TIER 1: VANILLA MOB (%70 population)
├── Zombie, Skeleton, Spider varyantları
├── Custom isim + level göstergesi
├── Modified stats
└── Bedrock: %100 uyumlu

TIER 2: TEXTURED VANILLA (%25 population)
├── Elite moblar, mini-bosslar
├── Custom texture (resource pack)
├── Özel parçacık efektleri
└── Bedrock: Texture pack ile uyumlu

TIER 3: CUSTOM MODEL (%5 population)
├── Ana bosslar
├── Item Display based
├── Basit animasyonlar
└── Bedrock: Vanilla fallback
```


### Mob Skill Listesi (20 Skill)

```
⚔️ SALDIRI (5):
├── slash      - Yakın kesme (particle trail)
├── shoot      - Projektil fırlatma
├── slam       - Yere vurma (AoE)
├── beam       - Işın saldırısı (line AoE)
└── volley     - Çoklu projektil

🔥 ELEMENT (4):
├── burn       - Ateş DoT
├── freeze     - Yavaşlatma + hasar
├── shock      - Zincir şimşek
└── void       - Karanlık patlama

🏃 HAREKET (4):
├── dash       - Hızlı ileri atılma
├── blink      - Kısa ışınlanma
├── leap       - Havaya zıplayıp inme
└── pull       - Oyuncuyu kendine çekme

🛡️ SAVUNMA (4):
├── shield     - Hasar bloklama (kırılabilir)
├── heal       - HP yenileme
├── summon     - Yardımcı mob çağırma
└── enrage     - Düşük HP buff

⚡ KONTROL (3):
├── stun       - 1-2sn hareket engeli
├── knockback  - Güçlü itme
└── blind      - Görüş kısıtlama
```

### Skill Trigger Sistemi

```
TIMER:         Her X saniyede (basic mobs)
HP_THRESHOLD:  %75, %50, %25'te (elites)
DISTANCE:      Oyuncu uzaktaysa ranged, yakınsa melee
COMBO:         Skill A → Skill B zinciri (bosses)
RANDOM:        Ağırlıklı rastgele seçim
```

---

## 👑 Boss Sistemi

### Boss Tierleri

```
TIER 1 - MINI BOSS (Dungeon ortası):
├── 2 Phase
├── 2-3 skill
├── Süre: 2-3 dakika
└── Wipe: Odanın başından

TIER 2 - DUNGEON BOSS (Dungeon sonu):
├── 3 Phase
├── 4-5 skill + phase özel mekanik
├── Süre: 5-7 dakika
└── Wipe: 3 can hakkı

TIER 3 - RAID BOSS (Özel etkinlik):
├── 5 Phase
├── 8+ skill + arena değişimi
├── Süre: 10-15 dakika
└── Wipe: Checkpoint (her 2 phase)
```

### Wipe Sistemi

```
CASUAL FRIENDLY:
├── Party toplam 3 can hakkı
├── Ölünce: 10sn sonra canlanabilir
├── Can bitince: Boss HP resetlenmez
├── Full wipe: Boss %10 HP kaybeder
└── Felsefe: "Tekrar dene"

HARDCORE MOD (Opsiyonel):
├── 1 can hakkı
├── Wipe = baştan
├── Ekstra ödüller
└── Leaderboard: En hızlı kill
```

---

## 🎁 Loot Sistemi

```
PERSONAL LOOT:
├── Normal drop herkes alır
└── Dungeon currency herkes alır

ROLL SİSTEMİ:
├── Nadir droplar için
├── Need/Greed seçimi
└── Timer: 30sn

CONTRIBUTION BASED:
├── Boss dropları için
├── DPS/Heal/Tank katkısı
└── AFK önleme
```

---

## 🔧 Teknik Gereksinimler

```
INSTANCE YÖNETİMİ:
├── Slime World Manager
├── RAM-based world clone
└── Custom wrapper plugin

CUSTOM PLUGİNLER:
├── karapixel-dungeons (ana sistem)
├── karapixel-mobs (mob yönetimi)
├── karapixel-models (basit model sistemi)
└── karapixel-loot (drop tabloları)
```

---

## 📝 Geliştirme Notları

- [ ] Slime World Manager entegrasyonu
- [ ] Mob skill sistemi geliştirme
- [ ] Boss phase engine
- [ ] Loot distribution sistemi
- [ ] Bedrock uyumluluk testleri
