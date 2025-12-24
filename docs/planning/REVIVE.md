# 💚 KaraPixel Revive (Canlandırma) Sistemi

> **Durum:** ✅ Planlandı
> **Son Güncelleme:** 2024-12-24

---

## 🎯 Konsept

Dungeon ve boss fight'larda "downed state" sistemi. Oyuncu ölmek yerine bayılır, takım arkadaşları canlandırabilir.

---

## 💀 Downed State (Bayılma)

```
┌─────────────────────────────────────────────────────────────────┐
│                    DOWNED STATE                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  TETİKLENME:                                                    │
│  ├── HP 0'a düşünce → Ölme, bayıl                             │
│  ├── Yere düş (crawling pose)                                  │
│  ├── 30 saniye bleed-out timer başlar                         │
│  └── Timer bitince: Gerçek ölüm                               │
│                                                                  │
│  DOWNED SIRASINDA:                                              │
│  ├── Hareket edebilir (çok yavaş, sürünerek)                  │
│  ├── Saldıramaz                                                │
│  ├── Item kullanamaz                                           │
│  ├── Skill kullanamaz                                          │
│  ├── Takım arkadaşını görebilir (marker)                      │
│  └── "Yardım!" butonu (ping gönderir)                         │
│                                                                  │
│  GÖRSEL:                                                        │
│  ├── Ekran kenarları kırmızı                                  │
│  ├── Heartbeat ses efekti                                      │
│  ├── Timer bar (ekranda)                                       │
│  └── Particle: Kan efekti                                      │
│                                                                  │
│  BEDROCK:                                                       │
│  ├── Crouch pose + yavaş hareket                              │
│  ├── ActionBar timer                                           │
│  └── Particle efektler aynı                                    │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## 🤝 Revive Mekaniği

```
┌─────────────────────────────────────────────────────────────────┐
│                    REVİVE NASIL ÇALIŞIR                         │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  TAKIMLA CANLANDIRMA:                                           │
│  ├── 1. Downed oyuncuya yaklaş (3 blok)                       │
│  ├── 2. Shift + Sağ tık (channel başlar)                      │
│  ├── 3. 3 saniye bekle (progress bar)                         │
│  ├── 4. Başarılı: %50 HP ile ayağa kalkar                     │
│  └── 5. Interrupt: Hasar alırsan kesilir                      │
│                                                                  │
│  HIZLI REVİVE (Özel Item):                                      │
│  ├── Phoenix Feather kullan                                    │
│  ├── Anında canlandır (%75 HP)                                │
│  ├── Item harcanır                                             │
│  └── Nadir drop / Event reward                                │
│                                                                  │
│  SELF-REVIVE:                                                   │
│  ├── Phoenix Tear gerekli (daha nadir)                        │
│  ├── Downed iken kullanabilir                                 │
│  ├── %30 HP ile kalk                                          │
│  ├── Dungeon'da 1 kez kullanım                                │
│  └── PvP'de kullanılamaz                                       │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## 📍 Nerede Çalışır

| Bölge | Downed State | Revive |
|-------|--------------|--------|
| Normal Ada | ❌ | ❌ (instant death) |
| Hub/Spawn | ❌ | ❌ (ölüm yok) |
| Dungeon | ✅ | ✅ |
| Boss Arena | ✅ | ✅ |
| Raid | ✅ | ✅ (sınırlı) |
| PvP Arena | ⚠️ Opsiyonel | ⚠️ Mod'a bağlı |

---

## ⚙️ Revive Kuralları

### Dungeon Kuralları

```
├── Party: 3 revive hakkı (toplam)
├── Her revive = 1 hak kullanır
├── Self-revive = Hak kullanmaz (item harcar)
├── Haklar bitince: Gerçek ölüm
└── Tam party wipe: Boss %10 HP kaybeder
```

### Raid Kuralları

```
├── Daha sıkı kurallar
├── Revive süresi: 5 saniye (3 değil)
├── Self-revive: Yok
├── Bleed-out: 20 saniye (30 değil)
└── Koordinasyon çok önemli
```

### PvP Kuralları (Opsiyonel Mod)

```
├── "Execution" mod: Downed'ı öldürebilirsin
├── Revive: Takım arkadaşı yapabilir
├── Self-revive: Kapalı
├── Daha hızlı bleed-out: 15 saniye
└── Competitive tension artırır
```

---

## 🎁 Revive Items

| Item | Etki | Kaynak |
|------|------|--------|
| Phoenix Feather | Hızlı revive (başkasına) | Boss drop, Event |
| Phoenix Tear | Self-revive | Raid drop, Prestige shop |
| Revive Potion | +%25 HP revive | Craft, NPC |
| Guardian Angel | Otomatik self-revive (1x) | Mythic drop |

---

## 📊 Neden Bu Sistem

```
┌─────────────────────────────────────────────────────────────────┐
│                    FAYDALARI                                    │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  OYUNCU DENEYİMİ:                                               │
│  ├── Daha affedici (instant death yok)                        │
│  ├── Takım oyununu teşvik eder                                │
│  ├── Clutch anları yaratır                                    │
│  ├── Tension artırır (bleed-out panic)                        │
│  └── İkinci şans = Daha az frustration                       │
│                                                                  │
│  İÇERİK ÜRETME:                                                 │
│  ├── Dramatik anlar (Vural için)                              │
│  ├── "Son anda kurtardım!" klipleri                          │
│  ├── Party coordination highlight                              │
│  └── Fail compilation da eğlenceli                            │
│                                                                  │
│  OYUN TASARIMI:                                                 │
│  ├── Boss fight'lar daha zor olabilir                        │
│  ├── Raid mechanic'leri daha agresif                         │
│  ├── Healer/Support role değerli                              │
│  └── Item sink (revive items)                                 │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## 🔧 Teknik Notlar

```yaml
Plugin: karapixel-revive

Events:
  - PlayerDeathEvent → DownedState başlat
  - PlayerInteractEntityEvent → Revive channel
  - PlayerMoveEvent → Crawling speed limit

Bedrock:
  - Pose: Swimming/Sneaking hybrid
  - Timer: ActionBar
  - Revive UI: Form yerine ActionBar

Performance:
  - Downed player tracking (HashMap)
  - Timer task (async)
  - Particle throttling
```
