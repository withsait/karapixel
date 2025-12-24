# ⚔️ KaraPixel Combat & Savaş Mekanikleri

> **Durum:** 📝 Planlama Aşaması
> **Son Güncelleme:** 2024-12-24

---

## 📋 Cevaplanması Gereken Sorular

### ❤️ Can Barı Sistemi
- [ ] Mob can barı tasarımı? (BossBar vs Hologram vs ActionBar)
- [ ] Hasar göstergesi? (Floating damage numbers)
- [ ] Oyuncu can barı özelleştirmesi?

### 💀 Ölüm Sistemi
- [ ] Özlü söz sistemi detayları?
- [ ] VIP ölüm yeri kalma süresi?
- [ ] Death recap (nasıl öldün özeti)?

### ⚔️ Combat Mekanikleri
- [ ] Combo sistemi?
- [ ] Dodge/parry mekanikleri?
- [ ] Elemental damage types?

---

## ✅ Alınan Kararlar

| Konu | Karar | Tarih |
|------|-------|-------|
| Mob Can Barı | Evet, olacak | 2024-12-24 |
| Ölüm Sözleri | Satın alınabilir özellik | 2024-12-24 |
| VIP Ölüm | Öldüğü yerde kalabilir | 2024-12-24 |

---

## ❤️ Can Barı Sistemi (Planlanacak)

```
┌─────────────────────────────────────────────────────────────────┐
│                    CAN BARI TİPLERİ                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  MOB CAN BARI:                                                  │
│  ├── Normal mob: İsim üstünde [████████--] 80%                 │
│  ├── Elite mob: BossBar (ekran üstü)                           │
│  ├── Boss: BossBar + Phase göstergesi                          │
│  └── Renk: HP'ye göre değişir (yeşil→sarı→kırmızı)            │
│                                                                  │
│  HASAR GÖSTERGESİ:                                              │
│  ├── Floating numbers (vurduğunda)                             │
│  ├── Renk: Normal=beyaz, Crit=kırmızı, Heal=yeşil             │
│  └── Bedrock: ActionBar alternatifi                            │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## 💀 Ölüm Sistemi (Planlanacak)

```
┌─────────────────────────────────────────────────────────────────┐
│                    ÖLÜM MEKANİKLERİ                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  ÖZLÜ SÖZ SİSTEMİ:                                              │
│  ├── Oyuncu öldüğünde seçtiği söz görünür                     │
│  ├── Varsayılan: "[İsim] öldü"                                 │
│  ├── Custom: "[İsim]: Sözüm söz, geri döneceğim!"             │
│  ├── Satın alma: Oyun içi para veya VIP                        │
│  ├── Sınırlama: Max 50 karakter, küfür filtresi               │
│  └── Özel sözler: Eventlerden kazanılır                        │
│                                                                  │
│  VIP ÖLÜM ÖZELLİKLERİ:                                          │
│  ├── Öldüğü yerde mezar taşı kalır (30dk-24saat)              │
│  ├── Mezar taşına tıklayınca özlü söz görünür                 │
│  ├── Itemlar mezar taşında korunur                             │
│  ├── Respawn seçenekleri (ada, spawn, son checkpoint)         │
│  └── Death recap: Kim/ne öldürdü detayı                        │
│                                                                  │
│  NORMAL OYUNCU:                                                 │
│  ├── Standart ölüm mesajı                                      │
│  ├── Item drop (keepInventory ayarına göre)                   │
│  └── Spawn'da respawn                                          │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## ⚔️ Combat Mekanikleri (Planlanacak)

```
[Detaylar eklenecek]
```

---

## 📝 Notlar

- Can barı performans dostu olmalı (çok fazla entity = lag)
- Ölüm sözleri moderasyon gerektirir
- Bedrock için ActionBar tabanlı alternatifler düşünülmeli
