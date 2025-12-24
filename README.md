# 🎮 KaraPixel - Minecraft Sunucu Projesi

> **Vural Üzül'ün 3M YouTube kanalı ile Türkiye'nin en profesyonel non-premium Minecraft sunucusu**

---

## 📋 Proje Özeti

| Özellik | Değer |
|---------|-------|
| **Hedef Kapasite** | 700-1000 CCU (Concurrent Users) |
| **Oyun Modu** | Skyblock (Multi-server mimari) |
| **Platform** | Java Edition + Bedrock Edition (Geyser/Floodgate) |
| **Donanım** | Ryzen 9 5950X, 128GB RAM, 2x3.84TB NVMe |
| **Yaklaşım** | %100 Custom (KaraPaper Fork + Custom Plugins) |
| **Varsayılan Dil** | Türkçe (Çoklu dil desteği altyapısı hazır) |
| **3D Modeller** | Evet (Pets, Wings, NPC, Cosmetics) |
| **Mobil Destek** | Tam uyumlu (Touch-friendly UI) |

---

## 🏗️ Mimari Genel Bakış

```
                        play.karapixel.net
                               │
                               ▼
                    ┌─────────────────────┐
                    │   VELOCITY PROXY    │
                    │  ┌───────────────┐  │
                    │  │    Geyser     │  │ ← Bedrock → Java çeviri
                    │  │   Floodgate   │  │ ← Xbox auth
                    │  └───────────────┘  │
                    │  • Rate Limiting    │
                    │  • VPN/Proxy Block  │
                    │  • Connection Filter│
                    └──────────┬──────────┘
                               │
                               ▼
                    ┌─────────────────────┐
                    │       LIMBO         │
                    │    (FakeLobby)      │
                    │  ┌───────────────┐  │
                    │  │   Captcha     │  │ ← Bot koruması
                    │  │ Login/Register│  │ ← Şifre sistemi
                    │  │   Session     │  │ ← Redis session
                    │  └───────────────┘  │
                    └──────────┬──────────┘
                               │ ✓ Authenticated
                               ▼
                    ┌─────────────────────┐
                    │     HUB LOBBY       │
                    │  ┌───────────────┐  │
                    │  │  3D NPC'ler   │  │ ← Oyun modu seçici
                    │  │   Portallar   │  │ ← Delik atlama
                    │  │  Cosmetics    │  │ ← Preview
                    │  └───────────────┘  │
                    └──────────┬──────────┘
                               │
              ┌────────────────┼────────────────┐
              ▼                ▼                ▼
     ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
     │   SKYBLOCK  │  │  (GELECEK)  │  │  (GELECEK)  │
     │    SPAWN    │  │   SURVIVAL  │  │   BEDWARS   │
     └──────┬──────┘  └─────────────┘  └─────────────┘
            │
    ┌───────┼───────┐
    ▼       ▼       ▼
┌───────┬───────┬───────┐
│World 1│World 2│World 3│ ← Ada sunucuları (load balancing)
└───────┴───────┴───────┘
```

---

## 📁 Proje Yapısı

```
KaraPixel/
│
├── README.md                        # Bu dosya
│
├── docs/                            # 📚 Dokümantasyon
│   ├── ARCHITECTURE.md              # Detaylı sistem mimarisi
│   ├── PLUGINS.md                   # Plugin listesi ve API dokümantasyonu
│   ├── SECURITY.md                  # ⚠️ Güvenlik stratejisi (ÇOK ÖNEMLİ)
│   ├── INFRASTRUCTURE.md            # Sunucu altyapısı ve kurulum
│   ├── ROADMAP.md                   # Geliştirme yol haritası
│   ├── LOCALIZATION.md              # 🌍 Dil desteği sistemi
│   ├── MOBILE.md                    # 📱 Bedrock/Geyser desteği
│   ├── 3D-MODELS.md                 # 🎨 3D model sistemi
│   ├── BACKUP.md                    # 💾 Yedekleme stratejisi
│   └── DATABASE.md                  # 🗄️ Veritabanı şeması
│
├── karapaper/                       # 🔧 Custom Server Fork
│   ├── patches/                     # Purpur üzerine patch'ler
│   ├── build.gradle.kts
│   └── README.md
│
├── plugins/                         # 🔌 Plugin Monorepo
│   ├── karapixel-core/              # Merkezi kütüphane
│   ├── karapixel-auth/              # Giriş sistemi
│   ├── karapixel-skyblock/          # Ana oyun modu
│   └── ... (32 plugin)
│
├── karapanel/                       # 🖥️ Yönetim Paneli
│   ├── daemon/                      # Go backend
│   └── web/                         # Next.js frontend
│
├── resourcepack/                    # 🎨 Resource Pack
│   ├── assets/minecraft/models/     # 3D modeller
│   ├── assets/minecraft/textures/   # Texture'lar
│   └── pack.mcmeta
│
├── infrastructure/                  # 🏗️ Altyapı
│   ├── scripts/                     # Kurulum scriptleri
│   ├── configs/                     # Production config'ler
│   └── docker/                      # Docker compose (dev)
│
└── locales/                         # 🌍 Dil Dosyaları
    ├── tr_TR.yml                    # Türkçe (varsayılan)
    ├── en_US.yml                    # İngilizce
    └── ...
```

---

## 🎯 Temel Özellikler

### Platform Desteği
- ✅ Java Edition (1.20.x - 1.21.10) - ViaVersion ile
- ✅ Bedrock Edition (Geyser + Floodgate)
- ✅ Mobil (iOS, Android) - Touch-friendly UI
- ✅ Konsol (Xbox, PlayStation, Switch) - Floodgate auth

### Güvenlik
- ✅ 6 katmanlı güvenlik sistemi
- ✅ DDoS koruması (TCPShield/Cosmic Guard)
- ✅ Bot/Spam koruması (Limbo + Captcha)
- ✅ Exploit koruması (KaraPaper patches)
- ✅ Anti-cheat entegrasyonu

### Dil Desteği
- ✅ %100 Türkçe varsayılan
- ✅ i18n altyapısı (çoklu dil hazır)
- ✅ Oyuncu bazlı dil seçimi
- ✅ Tüm mesajlar externalized

### 3D Modeller
- ✅ Custom pet modelleri
- ✅ Wing/kanat modelleri
- ✅ NPC modelleri
- ✅ Generator animasyonları
- ✅ Bedrock otomatik dönüşüm

---

## 💰 Maliyet Tablosu

| Kalem | Açıklama | Aylık Maliyet |
|-------|----------|---------------|
| Dedicated Server | Ryzen 9 5950X, 128GB RAM | €68.70 |
| Hyble-Core VPS | CX32, 4 vCPU, 8GB RAM | €15.59 |
| Storage Box | BX10, 1TB (Backup) | €3.81 |
| TCPShield | Free tier (başlangıç) | €0.00 |
| **TOPLAM** | | **€88.10/ay** |

---

## 📚 Dokümantasyon İndeksi

| Dosya | Açıklama | Öncelik |
|-------|----------|---------|
| [ARCHITECTURE.md](docs/ARCHITECTURE.md) | Sistem mimarisi, sunucu yapısı | 🔴 Yüksek |
| [FORK.md](docs/FORK.md) | KaraPaper fork detayları | 🔴 Yüksek |
| [AUTH.md](docs/AUTH.md) | Authentication sistemi, güvenlik | 🔴 Yüksek |
| [SECURITY.md](docs/SECURITY.md) | Güvenlik önlemleri, saldırı savunma | 🔴 Yüksek |
| [PLUGINS.md](docs/PLUGINS.md) | Plugin listesi, bağımlılıklar | 🔴 Yüksek |
| [INFRASTRUCTURE.md](docs/INFRASTRUCTURE.md) | Sunucu kurulumu | 🟡 Orta |
| [ROADMAP.md](docs/ROADMAP.md) | Geliştirme planı | 🟡 Orta |
| [LOCALIZATION.md](docs/LOCALIZATION.md) | Dil desteği sistemi | 🟡 Orta |
| [MOBILE.md](docs/MOBILE.md) | Bedrock/Geyser desteği | 🟡 Orta |
| [SPAWN.md](docs/SPAWN.md) | Hub/Spawn, NPC, Kasa, Kozmetik, Discord | 🟡 Orta |
| [3D-MODELS.md](docs/3D-MODELS.md) | 3D model sistemi | 🟢 Normal |
| [BACKUP.md](docs/BACKUP.md) | Yedekleme stratejisi | 🟢 Normal |
| [DATABASE.md](docs/DATABASE.md) | Veritabanı şeması | 🟢 Normal |

---

## 🚀 Hızlı Başlangıç

```bash
# 1. Repoyu klonla
git clone https://github.com/hyble/karapixel.git

# 2. Development ortamını kur
cd karapixel/infrastructure
./setup-dev.sh

# 3. Plugin'leri build et
cd ../plugins
./gradlew build

# 4. Local sunucuyu başlat
./gradlew runServer
```

---

## 👥 Ekip

| Rol | İsim | İletişim |
|-----|------|----------|
| Proje Sahibi & Developer | Sait (Hyble) | - |
| İçerik & Tanıtım | Vural Üzül | 3M YouTube |
| AI Assistant | Claude (Anthropic) | - |

---

## 📞 İletişim & Linkler

| Platform | Link |
|----------|------|
| Website | https://karapixel.net |
| Minecraft (Java) | play.karapixel.net:25565 |
| Minecraft (Bedrock) | play.karapixel.net:19132 |
| Discord | (eklenecek) |
| GitHub | (private repo) |

---

## 📄 Lisans

Bu proje özel mülkiyettir. Tüm hakları Hyble'a aittir.

---

*📅 Son güncelleme: 24 Aralık 2024*
*📌 Versiyon: 0.1.0-SNAPSHOT*
