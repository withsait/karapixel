# 👑 KaraPixel Boss Sistemi

> **Durum:** 📝 Planlama Aşaması (Sonra Dönülecek)
> **Son Güncelleme:** 2024-12-24

---

## 📋 Planlanacak Boss'lar

### Tier 1 - Mini Boss'lar
| Boss | Dungeon | Model | Durum |
|------|---------|-------|-------|
| Kemik Muhafız | Kayıp Madenleri | Textured | ⏳ |
| Buz Golemi | Donmuş Saray | Textured | ⏳ |
| Orman Ruhu | Lanetli Orman | Textured | ⏳ |

### Tier 2 - Dungeon Boss'ları
| Boss | Dungeon | Model | Durum |
|------|---------|-------|-------|
| Kemik Kralı | Kayıp Madenleri | Custom 3D | ⏳ |
| Buz Kraliçesi | Donmuş Saray | Custom 3D | ⏳ |
| Orman Canavarı | Lanetli Orman | Custom 3D | ⏳ |
| Ateş Lordu | Cehennem Kapısı | Custom 3D | ⏳ |
| Gölge Lordu Kael | Gölge Kalesi | Custom 3D | ⏳ |

### Tier 3 - Raid Boss'ları
| Boss | Event | Model | Durum |
|------|-------|-------|-------|
| Karanlık İmparator Xar'eth | Sezon Finali | Custom 3D | ⏳ |
| Gölge (Rehber Dark Form) | Prestige Event | Custom 3D | ⏳ |

---

## 🎨 3D Model Gereksinimleri

```
HER BOSS İÇİN:
├── idle.json       - Boşta duruş pozu
├── attack.json     - Saldırı pozu
├── special.json    - Özel skill pozu
├── hurt.json       - Hasar alma pozu
└── death.json      - Ölüm pozu (opsiyonel)

ARAÇLAR:
├── Blockbench (model tasarımı)
├── Resource Pack (texture)
└── Item Display API (render)
```

---

## 📝 Boss Detay Template

```yaml
boss_name:
  display_name: "&c&lBoss İsmi"
  tier: 2  # 1=Mini, 2=Dungeon, 3=Raid
  
  stats:
    base_health: 10000
    base_damage: 50
    armor: 20
    speed: 0.3
  
  phases:
    - name: "Phase 1"
      hp_range: [100, 70]
      skills: [skill1, skill2]
      music: "boss_phase1.ogg"
    - name: "Phase 2"
      hp_range: [70, 35]
      skills: [skill1, skill2, skill3]
      arena_change: true
    - name: "Phase 3"
      hp_range: [35, 0]
      enrage: true
      skills: [skill1, skill2, skill3, ultimate]
  
  model:
    type: CUSTOM_3D
    model_id: "boss_kael"
    scale: 2.0
    hitbox: [2, 3, 2]  # width, height, depth
  
  loot_table: "boss_kael_loot"
  
  mechanics:
    - type: DPS_CHECK
      phase: 2
      time: 60
      fail_action: WIPE
    - type: ARENA_HAZARD
      phase: 3
      type: FLOOR_DAMAGE
```

---

## 🏆 Öncelik Sırası

1. **Gölge Lordu Kael** (Ana hikaye boss'u)
2. **Kemik Kralı** (İlk dungeon boss)
3. **Buz Kraliçesi** (İkinci dungeon boss)
4. Diğerleri...

---

## 📝 Detaylı Planlar (Sonra Doldurulacak)

### Gölge Lordu Kael
```
[Detaylı plan buraya]
```

### Kemik Kralı
```
[Detaylı plan buraya]
```

### Buz Kraliçesi
```
[Detaylı plan buraya]
```

---

## 🔧 Teknik Notlar

- Item Display API kullanılacak (1.19.4+)
- Blockbench model formatı (.json)
- Bedrock fallback: Textured vanilla mob
- Hitbox: Invisible Slime/Armor Stand
