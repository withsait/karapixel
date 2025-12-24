# 🗺️ KaraPixel - Geliştirme Yol Haritası

> Proje planı ve milestone'lar.

---

## Genel Bakış

```
┌─────────────────────────────────────────────────────────────────┐
│                    PROJE TIMELINE                               │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  Tahmini Süre: 10-12 Hafta                                     │
│  Başlangıç: Ocak 2025                                          │
│  Hedef Launch: Mart/Nisan 2025                                 │
│                                                                 │
│  ████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ~10%           │
│  Faz 0: Altyapı (Mevcut)                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Fazlar

### 📦 Faz 0: Altyapı Hazırlığı (1 Hafta)

**Durum:** 🟡 Devam Ediyor

| Görev | Durum | Öncelik |
|-------|-------|---------|
| Hetzner 5950X sipariş | ☐ | 🔴 Kritik |
| Dedicated server kurulum | ☐ | 🔴 Kritik |
| MySQL + Redis kurulum | ☐ | 🔴 Kritik |
| Velocity + Geyser kurulum | ☐ | 🔴 Kritik |
| Hyble-core migration | ☐ | 🟡 Yüksek |
| DNS yapılandırması | ☐ | 🟡 Yüksek |
| Firewall ayarları | ☐ | 🟡 Yüksek |
| Backup sistemi kurulum | ☐ | 🟢 Normal |
| Monitoring kurulum | ☐ | 🟢 Normal |

**RAM Dağılımı (Seçenek B - Onaylandı):**
```
├── Infrastructure    : 16GB (MySQL 8GB, Redis 4GB, OS 4GB)
├── Proxy Layer       : 6GB  (Velocity 3GB, Geyser 2GB, Limbo 1GB)
├── Hub Lobby         : 6GB
├── Skyblock Spawn    : 20GB ⭐ (Market, Event, NPC, 3D Model)
├── PvP Arena         : 6GB  ⭐ (Ayrı sunucu)
├── Island World #1   : 24GB
├── Island World #2   : 24GB
├── Nether/End        : 10GB (Paylaşımlı)
└── Reserve           : 16GB (3. World için hazır)
                        ─────
                  TOPLAM: 128GB ✓
```

**Deliverables:**
- [ ] Çalışan dedicated server
- [ ] MySQL database hazır
- [ ] Redis çalışıyor
- [ ] Velocity proxy çalışıyor
- [ ] Geyser Bedrock desteği aktif

---

### 🔧 Faz 1: Core & Auth (1-2 Hafta)

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| karapixel-core plugin | ☐ | 🔴 Kritik |
| KaraPlayer API | ☐ | 🔴 Kritik |
| Platform detection (Java/Bedrock) | ☐ | 🔴 Kritik |
| Lokalizasyon sistemi | ☐ | 🔴 Kritik |
| karapixel-database plugin | ☐ | 🔴 Kritik |
| HikariCP connection pool | ☐ | 🔴 Kritik |
| karapixel-messaging plugin | ☐ | 🔴 Kritik |
| Redis pub/sub | ☐ | 🔴 Kritik |
| karapixel-ui plugin | ☐ | 🔴 Kritik |
| Bedrock Forms entegrasyonu | ☐ | 🔴 Kritik |
| karapixel-auth plugin | ☐ | 🔴 Kritik |
| Captcha sistemi | ☐ | 🟡 Yüksek |
| Session management | ☐ | 🟡 Yüksek |
| Bedrock auto-login | ☐ | 🟡 Yüksek |
| Limbo server kurulum | ☐ | 🟡 Yüksek |
| tr_TR.yml tam çeviri | ☐ | 🟡 Yüksek |

**Deliverables:**
- [ ] Oyuncular login/register olabiliyor
- [ ] Bedrock oyuncular otomatik giriş yapıyor
- [ ] Cross-server messaging çalışıyor
- [ ] Tüm mesajlar Türkçe

---

### 🏠 Faz 2: Hub Lobby (1 Hafta)

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| karapixel-hub plugin | ☐ | 🔴 Kritik |
| Hub map build/import | ☐ | 🔴 Kritik |
| karapixel-selector plugin | ☐ | 🔴 Kritik |
| Oyun seçici NPC | ☐ | 🟡 Yüksek |
| Oyun seçici menü | ☐ | 🟡 Yüksek |
| Portal sistemi | ☐ | 🟡 Yüksek |
| Hub items (hotbar) | ☐ | 🟡 Yüksek |
| Double jump | ☐ | 🟢 Normal |
| Spawn protection | ☐ | 🟢 Normal |
| Resource pack v1 | ☐ | 🟡 Yüksek |

**Deliverables:**
- [ ] Hub lobby çalışıyor
- [ ] Oyuncular Skyblock'a gidebiliyor
- [ ] Resource pack yükleniyor

---

### ⛏️ Faz 3: Skyblock Core (2-3 Hafta)

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| karapixel-skyblock plugin | ☐ | 🔴 Kritik |
| Island creation | ☐ | 🔴 Kritik |
| Island templates | ☐ | 🔴 Kritik |
| Island home/warp | ☐ | 🔴 Kritik |
| Island settings | ☐ | 🟡 Yüksek |
| Island coop system | ☐ | 🟡 Yüksek |
| Island level system | ☐ | 🟡 Yüksek |
| karapixel-generators plugin | ☐ | 🔴 Kritik |
| Generator tiers | ☐ | 🟡 Yüksek |
| Generator upgrades | ☐ | 🟡 Yüksek |
| karapixel-economy plugin | ☐ | 🔴 Kritik |
| Cross-server balance | ☐ | 🔴 Kritik |
| Multi-server world distribution | ☐ | 🟡 Yüksek |
| Skyblock spawn area | ☐ | 🟡 Yüksek |
| Island menu (mobil uyumlu) | ☐ | 🟡 Yüksek |

**Deliverables:**
- [ ] Oyuncular ada oluşturabiliyor
- [ ] Generator sistemi çalışıyor
- [ ] Para sistemi çalışıyor
- [ ] Multi-server load balancing

---

### ⭐ Faz 4: Skyblock Features (2-3 Hafta)

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| karapixel-skills plugin | ☐ | 🟡 Yüksek |
| 6 skill tipi | ☐ | 🟡 Yüksek |
| Skill rewards | ☐ | 🟢 Normal |
| karapixel-quests plugin | ☐ | 🟡 Yüksek |
| Daily/Weekly quests | ☐ | 🟡 Yüksek |
| Quest rewards | ☐ | 🟢 Normal |
| karapixel-shop plugin | ☐ | 🟡 Yüksek |
| Admin shop | ☐ | 🟡 Yüksek |
| Player shop | ☐ | 🟢 Normal |
| karapixel-upgrades plugin | ☐ | 🟡 Yüksek |
| Island upgrades | ☐ | 🟡 Yüksek |
| karapixel-minions plugin | ☐ | 🟢 Normal |
| karapixel-enchants plugin | ☐ | 🟢 Normal |
| Custom enchantlar | ☐ | 🟢 Normal |

**Deliverables:**
- [ ] Skill sistemi çalışıyor
- [ ] Quest sistemi çalışıyor
- [ ] Shop sistemi çalışıyor
- [ ] Ada yükseltmeleri çalışıyor

---

### 🎨 Faz 5: Polish & Cosmetics (1-2 Hafta)

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| karapixel-cosmetics plugin | ☐ | 🟡 Yüksek |
| Particle effects | ☐ | 🟢 Normal |
| Wing models | ☐ | 🟢 Normal |
| Hat models | ☐ | 🟢 Normal |
| karapixel-pets plugin | ☐ | 🟡 Yüksek |
| Pet models | ☐ | 🟡 Yüksek |
| Pet following AI | ☐ | 🟢 Normal |
| karapixel-chat plugin | ☐ | 🟡 Yüksek |
| Chat format | ☐ | 🟡 Yüksek |
| Private messaging | ☐ | 🟢 Normal |
| karapixel-tablist plugin | ☐ | 🟢 Normal |
| Resource pack v2 (3D models) | ☐ | 🟡 Yüksek |
| Bedrock model uyumluluk | ☐ | 🟡 Yüksek |

**Deliverables:**
- [ ] Cosmetic sistemi çalışıyor
- [ ] Pet sistemi çalışıyor
- [ ] 3D modeller görünüyor (Java + Bedrock)

---

### 💰 Faz 6: Monetization (1 Hafta)

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| karapixel-ranks plugin | ☐ | 🟡 Yüksek |
| VIP/MVP ranks | ☐ | 🟡 Yüksek |
| Rank permissions | ☐ | 🟡 Yüksek |
| karapixel-store plugin | ☐ | 🟡 Yüksek |
| In-game store menu | ☐ | 🟡 Yüksek |
| Shopier entegrasyonu | ☐ | 🟡 Yüksek |
| Papara entegrasyonu | ☐ | 🟢 Normal |
| karapixel-crates plugin | ☐ | 🟢 Normal |
| Crate animations | ☐ | 🟢 Normal |
| karapixel-battlepass plugin | ☐ | 🟢 Normal |

**Deliverables:**
- [ ] Oyuncular rank satın alabiliyor
- [ ] Ödeme sistemi çalışıyor
- [ ] Crate sistemi çalışıyor

---

### 🛡️ Faz 7: Launch Prep (1 Hafta)

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| karapixel-security plugin | ☐ | 🔴 Kritik |
| Anti-cheat entegrasyonu | ☐ | 🔴 Kritik |
| karapixel-moderation plugin | ☐ | 🟡 Yüksek |
| Ban/mute/warn sistemi | ☐ | 🟡 Yüksek |
| Staff tools | ☐ | 🟡 Yüksek |
| Stress test (500+ oyuncu) | ☐ | 🔴 Kritik |
| Security audit | ☐ | 🔴 Kritik |
| Performance tuning | ☐ | 🟡 Yüksek |
| Bug fixes | ☐ | 🟡 Yüksek |
| Closed beta (50-100 kişi) | ☐ | 🟡 Yüksek |
| Discord server kurulum | ☐ | 🟢 Normal |
| Website | ☐ | 🟢 Normal |

**Deliverables:**
- [ ] Güvenlik testleri geçildi
- [ ] Performans kabul edilebilir (18+ TPS @ 500 CCU)
- [ ] Beta feedback toplandı
- [ ] Kritik buglar düzeltildi

---

### 🚀 Faz 8: LAUNCH

**Durum:** ⏳ Beklemede

| Görev | Durum | Öncelik |
|-------|-------|---------|
| Son kontroller | ☐ | 🔴 Kritik |
| Vural ile koordinasyon | ☐ | 🔴 Kritik |
| Video #1 yayını | ☐ | 🔴 Kritik |
| 24/7 monitoring aktif | ☐ | 🔴 Kritik |
| Support ekibi hazır | ☐ | 🟡 Yüksek |
| Hotfix deploy pipeline | ☐ | 🟡 Yüksek |

**Deliverables:**
- [ ] 🎉 SUNUCU AÇIK!
- [ ] İlk 24 saat sorunsuz

---

## Timeline Özeti

```
Hafta 1:  [████████████████████] Faz 0: Altyapı
Hafta 2:  [████████████████████] Faz 1: Core
Hafta 3:  [████████████████████] Faz 1: Auth
Hafta 4:  [████████████████████] Faz 2: Hub
Hafta 5:  [████████████████████] Faz 3: Skyblock Core
Hafta 6:  [████████████████████] Faz 3: Skyblock Core
Hafta 7:  [████████████████████] Faz 4: Features
Hafta 8:  [████████████████████] Faz 4: Features
Hafta 9:  [████████████████████] Faz 5: Polish
Hafta 10: [████████████████████] Faz 6: Monetization
Hafta 11: [████████████████████] Faz 7: Launch Prep
Hafta 12: [████████████████████] Faz 8: LAUNCH 🚀
```

---

## Öncelik Matrisi

```
┌─────────────────────────────────────────────────────────────────┐
│                 ÖNCELİK MATRİSİ                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  🔴 KRİTİK (Launch engelleyici)                                │
│  ├── Altyapı kurulumu                                          │
│  ├── Auth sistemi                                              │
│  ├── Temel Skyblock mekaniği                                   │
│  ├── Güvenlik                                                   │
│  └── Performans                                                 │
│                                                                 │
│  🟡 YÜKSEK (Launch için gerekli)                               │
│  ├── Hub lobby                                                  │
│  ├── Generator sistemi                                         │
│  ├── Economy                                                    │
│  ├── Skills & Quests                                           │
│  ├── Monetization                                               │
│  └── 3D modeller                                               │
│                                                                 │
│  🟢 NORMAL (Nice to have)                                      │
│  ├── Advanced cosmetics                                        │
│  ├── Minions                                                    │
│  ├── Battle pass                                               │
│  ├── Player shops                                              │
│  └── Crates                                                     │
│                                                                 │
│  ⚪ DÜŞÜK (Post-launch)                                        │
│  ├── Yeni oyun modları                                         │
│  ├── Events sistemi                                            │
│  └── Leaderboard advanced                                      │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Risk Faktörleri

| Risk | Olasılık | Etki | Mitigation |
|------|----------|------|------------|
| Hetzner server gecikmesi | Orta | Yüksek | Auction'u takip et, alternatif hazırla |
| Geyser uyumluluk sorunları | Orta | Orta | Erken test, fallback UI |
| Performance sorunları | Orta | Yüksek | Sürekli profiling, optimizasyon |
| DDoS saldırısı | Yüksek | Kritik | TCPShield, hazırlık |
| Bug'lar launch'ta | Yüksek | Orta | Beta test, hotfix pipeline |
| Vural video ertelemesi | Düşük | Yüksek | Alternatif tanıtım planı |

---

## Sonraki Adımlar

### Bu Hafta (Faz 0)
1. [ ] Hetzner Auction'dan 5950X sipariş et
2. [ ] Dedicated server kurulumunu tamamla
3. [ ] MySQL + Redis kur
4. [ ] Velocity + Geyser'ı çalıştır
5. [ ] İlk test bağlantısı

### Gelecek Hafta (Faz 1)
1. [ ] Plugin monorepo oluştur
2. [ ] karapixel-core geliştirmeye başla
3. [ ] Database şemasını implement et
4. [ ] Lokalizasyon sistemini kur

---

*📅 Son güncelleme: 24 Aralık 2024*
