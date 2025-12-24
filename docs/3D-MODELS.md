# 🎨 KaraPixel - 3D Model Sistemi

> Custom 3D modeller ile benzersiz görsel deneyim. Pets, wings, NPC'ler ve daha fazlası.

---

## Genel Bakış

```
┌─────────────────────────────────────────────────────────────────┐
│                    3D MODEL SİSTEMİ                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  KULLANIM ALANLARI                                             │
│  ├── 🐾 Pets (Evcil hayvanlar)                                 │
│  ├── 🪽 Wings (Kanatlar)                                       │
│  ├── 👑 Hats (Şapkalar)                                        │
│  ├── 🤖 NPCs (Özel karakterler)                                │
│  ├── ⚙️ Generators (Animasyonlu bloklar)                       │
│  ├── 📦 Crates (Sandıklar)                                     │
│  ├── 🎁 Cosmetics (Dekoratif itemler)                         │
│  └── ⚔️ Custom Items (Özel silahlar/araçlar)                  │
│                                                                 │
│  TEKNOLOJİ STACK                                               │
│  ├── Blockbench - Model oluşturma                             │
│  ├── Custom Model Data - Java Edition                          │
│  ├── Geyser - Bedrock dönüşümü                                │
│  └── ItemsAdder/Oraxen (opsiyonel) - Yönetim                  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Resource Pack Yapısı

```
resourcepack/
├── pack.mcmeta                      # Pack tanımı
├── pack.png                         # Pack ikonu
│
├── assets/
│   └── minecraft/
│       │
│       ├── models/
│       │   │
│       │   ├── item/
│       │   │   ├── paper.json                    # Base model (override'lar)
│       │   │   │
│       │   │   ├── karapixel/
│       │   │   │   │
│       │   │   │   ├── pets/
│       │   │   │   │   ├── dog.json              # Köpek modeli
│       │   │   │   │   ├── cat.json              # Kedi modeli
│       │   │   │   │   ├── dragon.json           # Ejderha modeli
│       │   │   │   │   ├── phoenix.json          # Anka kuşu
│       │   │   │   │   └── robot.json            # Robot pet
│       │   │   │   │
│       │   │   │   ├── wings/
│       │   │   │   │   ├── angel.json            # Melek kanatları
│       │   │   │   │   ├── demon.json            # Şeytan kanatları
│       │   │   │   │   ├── dragon.json           # Ejderha kanatları
│       │   │   │   │   └── butterfly.json        # Kelebek kanatları
│       │   │   │   │
│       │   │   │   ├── hats/
│       │   │   │   │   ├── crown.json            # Taç
│       │   │   │   │   ├── tophat.json           # Silindir şapka
│       │   │   │   │   ├── santa.json            # Noel baba şapkası
│       │   │   │   │   └── helmet_gold.json      # Altın kask
│       │   │   │   │
│       │   │   │   ├── generators/
│       │   │   │   │   ├── cobble_t1.json        # Cobble gen tier 1
│       │   │   │   │   ├── cobble_t2.json        # Cobble gen tier 2
│       │   │   │   │   ├── iron_t1.json          # Iron gen tier 1
│       │   │   │   │   ├── diamond_t1.json       # Diamond gen tier 1
│       │   │   │   │   └── emerald_t1.json       # Emerald gen tier 1
│       │   │   │   │
│       │   │   │   ├── npcs/
│       │   │   │   │   ├── shopkeeper.json       # Mağazacı NPC
│       │   │   │   │   ├── banker.json           # Bankacı NPC
│       │   │   │   │   └── guide.json            # Rehber NPC
│       │   │   │   │
│       │   │   │   └── items/
│       │   │   │       ├── magic_wand.json       # Sihirli değnek
│       │   │   │       ├── island_key.json       # Ada anahtarı
│       │   │   │       └── teleport_rod.json     # Işınlanma asası
│       │   │   │
│       │   │   └── vanilla_overrides/            # Vanilla item değişiklikleri
│       │   │       └── ...
│       │   │
│       │   └── block/
│       │       └── karapixel/
│       │           └── ...                       # Custom bloklar
│       │
│       ├── textures/
│       │   └── karapixel/
│       │       ├── pets/
│       │       │   ├── dog.png
│       │       │   ├── cat.png
│       │       │   └── ...
│       │       ├── wings/
│       │       ├── hats/
│       │       ├── generators/
│       │       ├── npcs/
│       │       └── items/
│       │
│       ├── sounds/
│       │   └── karapixel/
│       │       ├── pet_spawn.ogg
│       │       ├── level_up.ogg
│       │       ├── generator_upgrade.ogg
│       │       └── crate_open.ogg
│       │
│       └── font/
│           └── default.json                      # Custom karakterler (emoji)
│
└── bedrock/                                      # Bedrock-specific (Geyser için)
    └── geometry/
        └── ...                                   # .geo.json dosyaları
```

---

## Custom Model Data Sistemi

### Model ID Mapping

```java
/**
 * Custom Model Data ID'leri
 * Her model benzersiz bir ID'ye sahip
 */
public class ModelRegistry {
    
    // ID Aralıkları:
    // 1000-1999: Generators
    // 2000-2999: Pets
    // 3000-3999: Cosmetics (Wings, Hats, Trails)
    // 4000-4999: NPCs
    // 5000-5999: Custom Items
    // 6000-6999: Blocks
    // 9000-9999: Reserved/Special
    
    // === GENERATORS ===
    public static final int GENERATOR_COBBLE_T1 = 1001;
    public static final int GENERATOR_COBBLE_T2 = 1002;
    public static final int GENERATOR_COBBLE_T3 = 1003;
    public static final int GENERATOR_IRON_T1 = 1011;
    public static final int GENERATOR_IRON_T2 = 1012;
    public static final int GENERATOR_GOLD_T1 = 1021;
    public static final int GENERATOR_DIAMOND_T1 = 1031;
    public static final int GENERATOR_EMERALD_T1 = 1041;
    
    // === PETS ===
    public static final int PET_DOG = 2001;
    public static final int PET_CAT = 2002;
    public static final int PET_DRAGON = 2003;
    public static final int PET_PHOENIX = 2004;
    public static final int PET_ROBOT = 2005;
    public static final int PET_GHOST = 2006;
    public static final int PET_SLIME = 2007;
    
    // === WINGS ===
    public static final int WING_ANGEL = 3001;
    public static final int WING_DEMON = 3002;
    public static final int WING_DRAGON = 3003;
    public static final int WING_BUTTERFLY = 3004;
    public static final int WING_FAIRY = 3005;
    
    // === HATS ===
    public static final int HAT_CROWN = 3101;
    public static final int HAT_TOPHAT = 3102;
    public static final int HAT_SANTA = 3103;
    public static final int HAT_WITCH = 3104;
    public static final int HAT_HELMET_GOLD = 3105;
    
    // === NPCs ===
    public static final int NPC_SHOPKEEPER = 4001;
    public static final int NPC_BANKER = 4002;
    public static final int NPC_GUIDE = 4003;
    public static final int NPC_SELECTOR = 4004;
    
    // === CUSTOM ITEMS ===
    public static final int ITEM_MAGIC_WAND = 5001;
    public static final int ITEM_ISLAND_KEY = 5002;
    public static final int ITEM_TELEPORT_ROD = 5003;
    
    // Registry map
    private static final Map<String, Integer> MODELS = new HashMap<>();
    
    static {
        // Generators
        MODELS.put("generator_cobble_t1", GENERATOR_COBBLE_T1);
        MODELS.put("generator_cobble_t2", GENERATOR_COBBLE_T2);
        MODELS.put("generator_iron_t1", GENERATOR_IRON_T1);
        MODELS.put("generator_diamond_t1", GENERATOR_DIAMOND_T1);
        
        // Pets
        MODELS.put("pet_dog", PET_DOG);
        MODELS.put("pet_cat", PET_CAT);
        MODELS.put("pet_dragon", PET_DRAGON);
        MODELS.put("pet_phoenix", PET_PHOENIX);
        
        // Wings
        MODELS.put("wing_angel", WING_ANGEL);
        MODELS.put("wing_demon", WING_DEMON);
        MODELS.put("wing_dragon", WING_DRAGON);
        
        // Hats
        MODELS.put("hat_crown", HAT_CROWN);
        MODELS.put("hat_tophat", HAT_TOPHAT);
        
        // NPCs
        MODELS.put("npc_shopkeeper", NPC_SHOPKEEPER);
        MODELS.put("npc_guide", NPC_GUIDE);
    }
    
    public static int getModelId(String modelName) {
        return MODELS.getOrDefault(modelName, 0);
    }
    
    public static ItemStack createModelItem(String modelName) {
        int modelId = getModelId(modelName);
        if (modelId == 0) {
            throw new IllegalArgumentException("Unknown model: " + modelName);
        }
        
        ItemStack item = new ItemStack(Material.PAPER);
        ItemMeta meta = item.getItemMeta();
        meta.setCustomModelData(modelId);
        item.setItemMeta(meta);
        
        return item;
    }
}
```

### Base Model Override (paper.json)

```json
{
    "parent": "minecraft:item/generated",
    "textures": {
        "layer0": "minecraft:item/paper"
    },
    "overrides": [
        {"predicate": {"custom_model_data": 1001}, "model": "karapixel/generators/cobble_t1"},
        {"predicate": {"custom_model_data": 1002}, "model": "karapixel/generators/cobble_t2"},
        {"predicate": {"custom_model_data": 1003}, "model": "karapixel/generators/cobble_t3"},
        {"predicate": {"custom_model_data": 1011}, "model": "karapixel/generators/iron_t1"},
        {"predicate": {"custom_model_data": 1021}, "model": "karapixel/generators/gold_t1"},
        {"predicate": {"custom_model_data": 1031}, "model": "karapixel/generators/diamond_t1"},
        
        {"predicate": {"custom_model_data": 2001}, "model": "karapixel/pets/dog"},
        {"predicate": {"custom_model_data": 2002}, "model": "karapixel/pets/cat"},
        {"predicate": {"custom_model_data": 2003}, "model": "karapixel/pets/dragon"},
        {"predicate": {"custom_model_data": 2004}, "model": "karapixel/pets/phoenix"},
        
        {"predicate": {"custom_model_data": 3001}, "model": "karapixel/wings/angel"},
        {"predicate": {"custom_model_data": 3002}, "model": "karapixel/wings/demon"},
        {"predicate": {"custom_model_data": 3003}, "model": "karapixel/wings/dragon"},
        
        {"predicate": {"custom_model_data": 3101}, "model": "karapixel/hats/crown"},
        {"predicate": {"custom_model_data": 3102}, "model": "karapixel/hats/tophat"},
        
        {"predicate": {"custom_model_data": 4001}, "model": "karapixel/npcs/shopkeeper"},
        {"predicate": {"custom_model_data": 4002}, "model": "karapixel/npcs/banker"},
        
        {"predicate": {"custom_model_data": 5001}, "model": "karapixel/items/magic_wand"},
        {"predicate": {"custom_model_data": 5002}, "model": "karapixel/items/island_key"}
    ]
}
```

---

## Model Örnekleri

### Pet Model (dog.json)

```json
{
    "credit": "KaraPixel Team - Blockbench",
    "texture_size": [32, 32],
    "textures": {
        "0": "karapixel/pets/dog"
    },
    "elements": [
        {
            "name": "body",
            "from": [5, 3, 5],
            "to": [11, 9, 14],
            "rotation": {"angle": 0, "axis": "y", "origin": [8, 6, 9.5]},
            "faces": {
                "north": {"uv": [4, 4, 10, 10], "texture": "#0"},
                "east": {"uv": [4, 4, 13, 10], "texture": "#0"},
                "south": {"uv": [4, 4, 10, 10], "texture": "#0"},
                "west": {"uv": [4, 4, 13, 10], "texture": "#0"},
                "up": {"uv": [4, 4, 10, 13], "texture": "#0"},
                "down": {"uv": [4, 4, 10, 13], "texture": "#0"}
            }
        },
        {
            "name": "head",
            "from": [6, 6, 2],
            "to": [10, 11, 6],
            "faces": {
                "north": {"uv": [0, 0, 4, 5], "texture": "#0"},
                "east": {"uv": [0, 0, 4, 5], "texture": "#0"},
                "south": {"uv": [0, 0, 4, 5], "texture": "#0"},
                "west": {"uv": [0, 0, 4, 5], "texture": "#0"},
                "up": {"uv": [0, 0, 4, 4], "texture": "#0"},
                "down": {"uv": [0, 0, 4, 4], "texture": "#0"}
            }
        },
        {
            "name": "tail",
            "from": [7, 6, 14],
            "to": [9, 8, 18],
            "rotation": {"angle": 22.5, "axis": "x", "origin": [8, 7, 14]},
            "faces": {
                "north": {"uv": [14, 0, 16, 2], "texture": "#0"},
                "east": {"uv": [14, 0, 18, 2], "texture": "#0"},
                "south": {"uv": [14, 0, 16, 2], "texture": "#0"},
                "west": {"uv": [14, 0, 18, 2], "texture": "#0"},
                "up": {"uv": [14, 0, 16, 4], "texture": "#0"},
                "down": {"uv": [14, 0, 16, 4], "texture": "#0"}
            }
        },
        {
            "name": "leg_fl",
            "from": [6, 0, 5],
            "to": [8, 3, 7],
            "faces": {
                "north": {"uv": [12, 8, 14, 11], "texture": "#0"},
                "east": {"uv": [12, 8, 14, 11], "texture": "#0"},
                "south": {"uv": [12, 8, 14, 11], "texture": "#0"},
                "west": {"uv": [12, 8, 14, 11], "texture": "#0"},
                "down": {"uv": [12, 8, 14, 10], "texture": "#0"}
            }
        }
        // ... diğer bacaklar
    ],
    "display": {
        "thirdperson_righthand": {
            "rotation": [75, 45, 0],
            "translation": [0, 2.5, 0],
            "scale": [0.375, 0.375, 0.375]
        },
        "firstperson_righthand": {
            "rotation": [0, 45, 0],
            "scale": [0.4, 0.4, 0.4]
        },
        "ground": {
            "translation": [0, 3, 0],
            "scale": [0.5, 0.5, 0.5]
        },
        "head": {
            "translation": [0, 0, -6],
            "scale": [1.5, 1.5, 1.5]
        }
    }
}
```

### Generator Model (cobble_t1.json)

```json
{
    "credit": "KaraPixel Team",
    "texture_size": [32, 32],
    "textures": {
        "0": "karapixel/generators/cobble_t1",
        "particle": "minecraft:block/cobblestone"
    },
    "elements": [
        {
            "name": "base",
            "from": [2, 0, 2],
            "to": [14, 4, 14],
            "faces": {
                "north": {"uv": [0, 12, 12, 16], "texture": "#0"},
                "east": {"uv": [0, 12, 12, 16], "texture": "#0"},
                "south": {"uv": [0, 12, 12, 16], "texture": "#0"},
                "west": {"uv": [0, 12, 12, 16], "texture": "#0"},
                "up": {"uv": [0, 0, 12, 12], "texture": "#0"},
                "down": {"uv": [0, 0, 12, 12], "texture": "#0"}
            }
        },
        {
            "name": "core",
            "from": [4, 4, 4],
            "to": [12, 12, 12],
            "rotation": {"angle": 45, "axis": "y", "origin": [8, 8, 8]},
            "faces": {
                "north": {"uv": [4, 4, 12, 12], "texture": "#0"},
                "east": {"uv": [4, 4, 12, 12], "texture": "#0"},
                "south": {"uv": [4, 4, 12, 12], "texture": "#0"},
                "west": {"uv": [4, 4, 12, 12], "texture": "#0"},
                "up": {"uv": [4, 4, 12, 12], "texture": "#0"},
                "down": {"uv": [4, 4, 12, 12], "texture": "#0"}
            }
        },
        {
            "name": "top_ring",
            "from": [3, 12, 3],
            "to": [13, 14, 13],
            "faces": {
                "north": {"uv": [0, 0, 10, 2], "texture": "#0"},
                "east": {"uv": [0, 0, 10, 2], "texture": "#0"},
                "south": {"uv": [0, 0, 10, 2], "texture": "#0"},
                "west": {"uv": [0, 0, 10, 2], "texture": "#0"},
                "up": {"uv": [0, 0, 10, 10], "texture": "#0"},
                "down": {"uv": [0, 0, 10, 10], "texture": "#0"}
            }
        }
    ],
    "display": {
        "ground": {
            "scale": [0.5, 0.5, 0.5]
        },
        "gui": {
            "rotation": [30, 225, 0],
            "scale": [0.625, 0.625, 0.625]
        },
        "fixed": {
            "scale": [0.5, 0.5, 0.5]
        }
    }
}
```

---

## Pet Sistemi Implementasyonu

### Pet Entity

```java
public class PetEntity {
    private final UUID id;
    private final UUID ownerId;
    private final PetType type;
    private String name;
    private int level;
    private ArmorStand displayEntity;  // Java
    private Location currentLocation;
    private boolean spawned;
    
    /**
     * Pet'i spawn eder
     */
    public void spawn(KaraPlayer owner) {
        if (spawned) return;
        
        Location spawnLoc = calculateSpawnLocation(owner);
        
        if (owner.isBedrock()) {
            spawnBedrockPet(owner, spawnLoc);
        } else {
            spawnJavaPet(owner, spawnLoc);
        }
        
        spawned = true;
        startFollowTask(owner);
        
        // Spawn sesi
        owner.playSound(Sound.ENTITY_PLAYER_LEVELUP, 1.0f, 1.5f);
    }
    
    private void spawnJavaPet(KaraPlayer owner, Location location) {
        // ArmorStand ile 3D model gösterimi
        World world = location.getWorld();
        
        displayEntity = world.spawn(location, ArmorStand.class, stand -> {
            stand.setVisible(false);
            stand.setSmall(true);
            stand.setGravity(false);
            stand.setInvulnerable(true);
            stand.setMarker(true);
            stand.setCustomNameVisible(true);
            stand.customName(Text.parse("<gradient:#FFD700:#FFA500>" + name + "</gradient>"));
            
            // 3D model item
            ItemStack modelItem = ModelRegistry.createModelItem(type.getModelId());
            stand.getEquipment().setHelmet(modelItem);
        });
    }
    
    private void spawnBedrockPet(KaraPlayer owner, Location location) {
        // Bedrock için Geyser otomatik olarak ArmorStand'ı
        // geometry'ye çevirir. Ek yapılandırma gerekebilir.
        spawnJavaPet(owner, location);
    }
    
    /**
     * Sahibini takip et
     */
    private void startFollowTask(KaraPlayer owner) {
        Bukkit.getScheduler().runTaskTimer(plugin, () -> {
            if (!spawned || displayEntity == null || displayEntity.isDead()) {
                return;
            }
            
            Player ownerPlayer = owner.getBukkitPlayer();
            if (ownerPlayer == null || !ownerPlayer.isOnline()) {
                despawn();
                return;
            }
            
            Location ownerLoc = ownerPlayer.getLocation();
            Location petLoc = displayEntity.getLocation();
            
            double distance = ownerLoc.distance(petLoc);
            
            // Çok uzaksa teleport
            if (distance > 15) {
                Location newLoc = calculateSpawnLocation(owner);
                displayEntity.teleport(newLoc);
                return;
            }
            
            // Yakınsa sabit dur
            if (distance < 2) {
                animateIdle();
                return;
            }
            
            // Takip et
            Vector direction = ownerLoc.toVector().subtract(petLoc.toVector()).normalize();
            Location newLoc = petLoc.add(direction.multiply(0.2));
            newLoc.setYaw(getYawToFace(petLoc, ownerLoc));
            
            displayEntity.teleport(newLoc);
            animateWalking();
            
        }, 0L, 1L);  // Her tick
    }
    
    /**
     * Pet'i kaldır
     */
    public void despawn() {
        if (displayEntity != null && !displayEntity.isDead()) {
            displayEntity.remove();
        }
        spawned = false;
    }
    
    private Location calculateSpawnLocation(KaraPlayer owner) {
        Location ownerLoc = owner.getLocation();
        // Sahibinin 1.5 blok arkasında ve yanında
        Vector offset = ownerLoc.getDirection().multiply(-1.5);
        offset.add(new Vector(0.8, 0, 0));  // Biraz yana
        return ownerLoc.add(offset);
    }
    
    private void animateIdle() {
        // Hafif yukarı aşağı hareket
        Location loc = displayEntity.getLocation();
        double y = loc.getY() + Math.sin(System.currentTimeMillis() / 200.0) * 0.05;
        loc.setY(y);
        displayEntity.teleport(loc);
    }
    
    private void animateWalking() {
        // Yürürken hafif zıplama
        Location loc = displayEntity.getLocation();
        double y = loc.getY() + Math.abs(Math.sin(System.currentTimeMillis() / 100.0)) * 0.1;
        loc.setY(y);
        displayEntity.teleport(loc);
    }
}
```

### Pet Tipleri

```java
public enum PetType {
    DOG("pet_dog", "pet.type.dog", Rarity.COMMON),
    CAT("pet_cat", "pet.type.cat", Rarity.COMMON),
    PARROT("pet_parrot", "pet.type.parrot", Rarity.UNCOMMON),
    WOLF("pet_wolf", "pet.type.wolf", Rarity.UNCOMMON),
    FOX("pet_fox", "pet.type.fox", Rarity.RARE),
    DRAGON("pet_dragon", "pet.type.dragon", Rarity.EPIC),
    PHOENIX("pet_phoenix", "pet.type.phoenix", Rarity.LEGENDARY),
    ROBOT("pet_robot", "pet.type.robot", Rarity.LEGENDARY);
    
    private final String modelId;
    private final String nameKey;
    private final Rarity rarity;
    
    PetType(String modelId, String nameKey, Rarity rarity) {
        this.modelId = modelId;
        this.nameKey = nameKey;
        this.rarity = rarity;
    }
    
    public String getModelId() { return modelId; }
    public String getNameKey() { return nameKey; }
    public Rarity getRarity() { return rarity; }
}

public enum Rarity {
    COMMON("<gray>", "rarity.common"),
    UNCOMMON("<green>", "rarity.uncommon"),
    RARE("<blue>", "rarity.rare"),
    EPIC("<dark_purple>", "rarity.epic"),
    LEGENDARY("<gold>", "rarity.legendary");
    
    private final String color;
    private final String nameKey;
}
```

---

## Wings & Cosmetics

### Wing Renderer

```java
public class WingRenderer {
    
    /**
     * Kanat cosmetic'ini oyuncuya uygula
     */
    public static void applyWing(KaraPlayer player, WingType wingType) {
        Player bukkitPlayer = player.getBukkitPlayer();
        
        // Mevcut wing'i kaldır
        removeWing(player);
        
        if (player.isBedrock()) {
            // Bedrock: Elytra texture değişikliği + Geyser mapping
            applyBedrockWing(player, wingType);
        } else {
            // Java: ArmorStand ile backpack tarzı
            applyJavaWing(player, wingType);
        }
    }
    
    private static void applyJavaWing(KaraPlayer player, WingType wingType) {
        Player bukkitPlayer = player.getBukkitPlayer();
        Location spawnLoc = bukkitPlayer.getLocation();
        
        ArmorStand wingStand = spawnLoc.getWorld().spawn(spawnLoc, ArmorStand.class, stand -> {
            stand.setVisible(false);
            stand.setSmall(true);
            stand.setGravity(false);
            stand.setInvulnerable(true);
            stand.setMarker(true);
            
            // Wing model
            ItemStack wingItem = ModelRegistry.createModelItem(wingType.getModelId());
            stand.getEquipment().setHelmet(wingItem);
        });
        
        // Oyuncuya bağla (passenger olarak değil, location sync)
        startWingFollowTask(player, wingStand);
        
        // Metadata'ya kaydet
        bukkitPlayer.setMetadata("wing_stand", new FixedMetadataValue(plugin, wingStand.getUniqueId()));
    }
    
    private static void applyBedrockWing(KaraPlayer player, WingType wingType) {
        // Bedrock için Geyser custom geometry kullanılabilir
        // veya Java ile aynı yöntem (ArmorStand Geyser'da da çalışır)
        applyJavaWing(player, wingType);
    }
    
    private static void startWingFollowTask(KaraPlayer player, ArmorStand wingStand) {
        new BukkitRunnable() {
            @Override
            public void run() {
                Player bukkitPlayer = player.getBukkitPlayer();
                if (bukkitPlayer == null || !bukkitPlayer.isOnline() || wingStand.isDead()) {
                    wingStand.remove();
                    cancel();
                    return;
                }
                
                // Oyuncunun arkasında, biraz yukarıda
                Location playerLoc = bukkitPlayer.getLocation();
                Vector back = playerLoc.getDirection().multiply(-0.3);
                Location wingLoc = playerLoc.add(back).add(0, 1.5, 0);
                wingLoc.setYaw(playerLoc.getYaw());
                wingLoc.setPitch(0);
                
                wingStand.teleport(wingLoc);
            }
        }.runTaskTimer(plugin, 0L, 1L);
    }
    
    public static void removeWing(KaraPlayer player) {
        Player bukkitPlayer = player.getBukkitPlayer();
        if (bukkitPlayer.hasMetadata("wing_stand")) {
            UUID standId = (UUID) bukkitPlayer.getMetadata("wing_stand").get(0).value();
            Entity stand = Bukkit.getEntity(standId);
            if (stand != null) {
                stand.remove();
            }
            bukkitPlayer.removeMetadata("wing_stand", plugin);
        }
    }
}

public enum WingType {
    ANGEL("wing_angel", "wing.type.angel"),
    DEMON("wing_demon", "wing.type.demon"),
    DRAGON("wing_dragon", "wing.type.dragon"),
    BUTTERFLY("wing_butterfly", "wing.type.butterfly"),
    FAIRY("wing_fairy", "wing.type.fairy");
    
    private final String modelId;
    private final String nameKey;
}
```

---

## Bedrock Model Uyumluluğu

### Geyser Custom Model Mapping

```yaml
# Geyser custom mappings config
# plugins/Geyser-Velocity/custom_mappings/karapixel.json

{
    "format_version": "1",
    "items": {
        "minecraft:paper": {
            "custom_model_data": {
                "1001": {
                    "name": "generator_cobble_t1",
                    "icon": "karapixel.generator_cobble_t1"
                },
                "2001": {
                    "name": "pet_dog",
                    "icon": "karapixel.pet_dog"
                },
                "3001": {
                    "name": "wing_angel",
                    "icon": "karapixel.wing_angel"
                }
            }
        }
    }
}
```

### Bedrock Geometry Dosyası

```json
// bedrock/geometry/pet_dog.geo.json
{
    "format_version": "1.16.0",
    "minecraft:geometry": [
        {
            "description": {
                "identifier": "geometry.karapixel.pet_dog",
                "texture_width": 32,
                "texture_height": 32,
                "visible_bounds_width": 2,
                "visible_bounds_height": 2,
                "visible_bounds_offset": [0, 1, 0]
            },
            "bones": [
                {
                    "name": "body",
                    "pivot": [0, 6, 0],
                    "cubes": [
                        {
                            "origin": [-3, 3, -5],
                            "size": [6, 6, 9],
                            "uv": [0, 0]
                        }
                    ]
                },
                {
                    "name": "head",
                    "parent": "body",
                    "pivot": [0, 9, -5],
                    "cubes": [
                        {
                            "origin": [-2, 6, -9],
                            "size": [4, 5, 4],
                            "uv": [0, 16]
                        }
                    ]
                }
            ]
        }
    ]
}
```

---

## Resource Pack Dağıtımı

### Sunucu Taraflı Zorunlu Pack

```yaml
# server.properties
resource-pack=https://karapixel.net/resourcepack/karapixel-pack.zip
resource-pack-sha1=abc123...  # SHA1 hash
require-resource-pack=true
resource-pack-prompt={"text":"KaraPixel oynamak için resource pack gerekli!","color":"gold"}
```

### Pack Versiyonlama

```java
public class ResourcePackManager {
    
    private static final String PACK_URL = "https://karapixel.net/resourcepack/";
    
    /**
     * Versiyon bazlı pack URL döndür
     */
    public static String getPackUrl() {
        String version = getLatestVersion();  // Config'den veya API'den
        return PACK_URL + "karapixel-pack-v" + version + ".zip";
    }
    
    /**
     * Pack hash'ini al
     */
    public static String getPackHash() {
        // SHA1 hash - pack her güncellendiğinde değişir
        return ConfigManager.getString("resourcepack.sha1");
    }
    
    /**
     * Oyuncuya pack gönder
     */
    public static void sendPack(Player player) {
        String url = getPackUrl();
        String hash = getPackHash();
        
        player.setResourcePack(url, hash, true, 
            Text.parse("<gold>KaraPixel için resource pack yükleniyor...</gold>"));
    }
    
    @EventHandler
    public void onPackStatus(PlayerResourcePackStatusEvent event) {
        Player player = event.getPlayer();
        
        switch (event.getStatus()) {
            case ACCEPTED -> {
                // Pack kabul edildi, yükleniyor
            }
            case SUCCESSFULLY_LOADED -> {
                // Pack yüklendi
                player.sendMessage(Text.parse("<green>Resource pack yüklendi!</green>"));
            }
            case DECLINED, FAILED_DOWNLOAD -> {
                // Pack reddedildi veya indirilemedi
                player.kick(Text.parse("<red>Resource pack gerekli!</red>"));
            }
        }
    }
}
```

---

## Blockbench Workflow

### Model Oluşturma Adımları

```
┌─────────────────────────────────────────────────────────────────┐
│                 BLOCKBENCH WORKFLOW                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  1. YENİ PROJE                                                 │
│  ─────────────                                                  │
│  ├── File → New → Java Block/Item                              │
│  ├── Texture size: 32x32 veya 64x64                           │
│  └── Model resolution ayarla                                   │
│                                                                 │
│  2. MODELİ OLUŞTUR                                             │
│  ──────────────────                                             │
│  ├── Cube tool ile şekil oluştur                               │
│  ├── Parent-child ilişkileri kur (animasyon için)             │
│  ├── Pivot noktalarını doğru ayarla                           │
│  └── UV mapping yap                                            │
│                                                                 │
│  3. TEXTURE                                                    │
│  ──────────                                                     │
│  ├── Paint tool ile doğrudan boya                             │
│  ├── Veya external editor'a export et                         │
│  └── 16x16, 32x32 veya 64x64 çözünürlük                       │
│                                                                 │
│  4. ANİMASYON (opsiyonel)                                      │
│  ─────────────────────────                                      │
│  ├── Animate tab'ına geç                                       │
│  ├── Keyframe'ler ekle                                         │
│  └── Bedrock için .animation.json export                       │
│                                                                 │
│  5. DISPLAY SETTINGS                                           │
│  ────────────────────                                           │
│  ├── Display tab'ına geç                                       │
│  ├── Her görünüm için ayarla:                                  │
│  │   ├── thirdperson_righthand                                 │
│  │   ├── firstperson_righthand                                 │
│  │   ├── ground                                                │
│  │   ├── gui                                                   │
│  │   ├── head                                                  │
│  │   └── fixed (item frame)                                    │
│  └── Scale, rotation, translation                              │
│                                                                 │
│  6. EXPORT                                                     │
│  ────────                                                       │
│  ├── Java: File → Export → Export Java Item Model              │
│  ├── Bedrock: File → Export → Export Bedrock Geometry          │
│  └── Texture: Ayrı PNG olarak kaydet                           │
│                                                                 │
│  7. DOSYA YERLEŞTİRME                                          │
│  ─────────────────────                                          │
│  ├── Model: assets/minecraft/models/karapixel/...              │
│  ├── Texture: assets/minecraft/textures/karapixel/...          │
│  └── paper.json'a override ekle                                │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Test & QA

### Model Test Checklist

| Test | Java | Bedrock | Açıklama |
|------|:----:|:-------:|----------|
| Model görünür | ☐ | ☐ | 3D model doğru render ediliyor |
| Texture doğru | ☐ | ☐ | Texture mapping hatasız |
| Scale doğru | ☐ | ☐ | Boyut uygun |
| Hand görünümü | ☐ | ☐ | Elde tutarken doğru |
| Head görünümü | ☐ | ☐ | Kafada (hat) doğru |
| Ground görünümü | ☐ | ☐ | Yerde düşük doğru |
| GUI görünümü | ☐ | ☐ | Envanterde doğru |
| Performans | ☐ | ☐ | FPS düşüşü yok |

---

*📅 Son güncelleme: 24 Aralık 2024*
