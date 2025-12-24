# 🤖 Gölge - Discord Bot Sistemi

> KaraPixel Minecraft sunucusu için özel Discord botu.

---

## Genel Bakış

```
┌──────────────────────────────────────────────────────────────────┐
│                      GÖLGE VİZYONU                               │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  AMAÇ:                                                          │
│  ──────                                                          │
│  Gölge, Discord ve Minecraft arasında köprü kurarak             │
│  topluluk deneyimini zenginleştirir.                            │
│                                                                  │
│  TEMEL İLKELER:                                                 │
│  ────────────────                                                │
│  ├── Gerçek zamanlı senkronizasyon (MC ↔ Discord)              │
│  ├── Oyuncu etkileşimini artırma                               │
│  ├── Topluluk yönetimini kolaylaştırma                         │
│  ├── Event ve duyuru otomasyonu                                │
│  └── Staff iş yükünü azaltma                                   │
│                                                                  │
│  KAPSAM:                                                        │
│  ────────                                                        │
│  ├── Hesap bağlama & doğrulama                                 │
│  ├── Çift yönlü chat köprüsü                                   │
│  ├── Sunucu durumu & istatistikler                             │
│  ├── Leaderboard & profil sistemi                              │
│  ├── Event yönetimi & bildirimleri                             │
│  ├── Ticket & destek sistemi                                   │
│  ├── Giveaway & çekiliş sistemi                                │
│  ├── Moderasyon araçları                                       │
│  └── Özel komutlar & eğlence                                   │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Teknoloji Stack

```
┌──────────────────────────────────────────────────────────────────┐
│                    TEKNOLOJİ STACK                               │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  RUNTIME & FRAMEWORK:                                           │
│  ─────────────────────                                           │
│  ├── Runtime: Bun (hızlı, modern)                              │
│  ├── Language: TypeScript                                       │
│  ├── Framework: Discord.js v14                                 │
│  └── Build: esbuild (Bun native)                               │
│                                                                  │
│  DATABASE:                                                      │
│  ──────────                                                      │
│  ├── Primary: PostgreSQL (Hyble DB ile paylaşımlı)            │
│  ├── ORM: Prisma                                               │
│  ├── Cache: Redis                                              │
│  └── Migrations: Prisma Migrate                                │
│                                                                  │
│  API & İLETİŞİM:                                                │
│  ────────────────                                                │
│  ├── MC → Bot: Redis Pub/Sub                                   │
│  ├── Bot → MC: REST API (plugin tarafı)                        │
│  ├── WebSocket: Real-time events                               │
│  └── HTTP Client: Bun native fetch                             │
│                                                                  │
│  DEPLOYMENT:                                                    │
│  ────────────                                                    │
│  ├── Container: Docker                                         │
│  ├── Process Manager: Docker Compose                           │
│  ├── Logging: Winston + Loki                                   │
│  └── Monitoring: Prometheus + Grafana                          │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## Proje Yapısı

```
golge/
├── src/
│   ├── index.ts                 # Entry point
│   ├── client.ts                # Discord client setup
│   ├── config.ts                # Configuration
│   │
│   ├── commands/                # Slash commands
│   │   ├── general/
│   │   │   ├── ping.ts
│   │   │   ├── help.ts
│   │   │   └── info.ts
│   │   ├── minecraft/
│   │   │   ├── link.ts
│   │   │   ├── unlink.ts
│   │   │   ├── profile.ts
│   │   │   └── leaderboard.ts
│   │   ├── moderation/
│   │   │   ├── ban.ts
│   │   │   ├── mute.ts
│   │   │   └── warn.ts
│   │   ├── ticket/
│   │   │   ├── create.ts
│   │   │   └── close.ts
│   │   ├── giveaway/
│   │   │   ├── start.ts
│   │   │   ├── end.ts
│   │   │   └── reroll.ts
│   │   └── admin/
│   │       ├── announce.ts
│   │       ├── status.ts
│   │       └── sync.ts
│   │
│   ├── events/                  # Discord events
│   │   ├── ready.ts
│   │   ├── interactionCreate.ts
│   │   ├── messageCreate.ts
│   │   ├── guildMemberAdd.ts
│   │   └── guildMemberRemove.ts
│   │
│   ├── services/                # Business logic
│   │   ├── linking.service.ts
│   │   ├── chat-bridge.service.ts
│   │   ├── status.service.ts
│   │   ├── leaderboard.service.ts
│   │   ├── ticket.service.ts
│   │   ├── giveaway.service.ts
│   │   └── minecraft-api.service.ts
│   │
│   ├── listeners/               # Redis/WebSocket listeners
│   │   ├── minecraft.listener.ts
│   │   ├── chat.listener.ts
│   │   └── event.listener.ts
│   │
│   ├── handlers/                # Event handlers
│   │   ├── command.handler.ts
│   │   ├── button.handler.ts
│   │   ├── modal.handler.ts
│   │   └── select-menu.handler.ts
│   │
│   ├── utils/                   # Utilities
│   │   ├── embed.util.ts
│   │   ├── logger.util.ts
│   │   ├── cache.util.ts
│   │   └── format.util.ts
│   │
│   ├── types/                   # TypeScript types
│   │   ├── discord.d.ts
│   │   ├── minecraft.d.ts
│   │   └── database.d.ts
│   │
│   └── database/                # Prisma
│       ├── schema.prisma
│       └── client.ts
│
├── prisma/
│   ├── schema.prisma
│   └── migrations/
│
├── Dockerfile
├── docker-compose.yml
├── package.json
├── tsconfig.json
├── .env.example
└── README.md
```

---

## Database Şeması

```prisma
// prisma/schema.prisma

generator client {
  provider = "prisma-client-js"
}

datasource db {
  provider = "postgresql"
  url      = env("DATABASE_URL")
}

// ═══════════════════════════════════════════════════════════
// HESAP BAĞLAMA
// ═══════════════════════════════════════════════════════════

model LinkedAccount {
  id            String   @id @default(cuid())
  minecraftUuid String   @unique @map("minecraft_uuid")
  minecraftName String   @map("minecraft_name")
  discordId     String   @unique @map("discord_id")
  discordTag    String   @map("discord_tag")
  linkedAt      DateTime @default(now()) @map("linked_at")
  isVerified    Boolean  @default(true) @map("is_verified")
  lastSync      DateTime @default(now()) @map("last_sync")

  @@map("discord_linked_accounts")
}

model LinkCode {
  id            String   @id @default(cuid())
  code          String   @unique
  minecraftUuid String   @map("minecraft_uuid")
  minecraftName String   @map("minecraft_name")
  createdAt     DateTime @default(now()) @map("created_at")
  expiresAt     DateTime @map("expires_at")
  used          Boolean  @default(false)

  @@map("discord_link_codes")
}

// ═══════════════════════════════════════════════════════════
// TICKET SİSTEMİ
// ═══════════════════════════════════════════════════════════

model Ticket {
  id          String       @id @default(cuid())
  ticketId    Int          @unique @default(autoincrement()) @map("ticket_id")
  channelId   String       @unique @map("channel_id")
  userId      String       @map("user_id")
  category    TicketCategory
  status      TicketStatus @default(OPEN)
  subject     String?
  priority    TicketPriority @default(NORMAL)
  assignedTo  String?      @map("assigned_to")
  createdAt   DateTime     @default(now()) @map("created_at")
  closedAt    DateTime?    @map("closed_at")
  closedBy    String?      @map("closed_by")
  closeReason String?      @map("close_reason")
  messages    TicketMessage[]

  @@map("discord_tickets")
}

model TicketMessage {
  id        String   @id @default(cuid())
  ticketId  String   @map("ticket_id")
  ticket    Ticket   @relation(fields: [ticketId], references: [id], onDelete: Cascade)
  userId    String   @map("user_id")
  content   String
  isStaff   Boolean  @default(false) @map("is_staff")
  createdAt DateTime @default(now()) @map("created_at")

  @@map("discord_ticket_messages")
}

enum TicketCategory {
  BUG_REPORT
  PAYMENT_ISSUE
  PLAYER_REPORT
  GENERAL_QUESTION
  REWARD_CLAIM
  APPEAL
  OTHER
}

enum TicketStatus {
  OPEN
  IN_PROGRESS
  WAITING_RESPONSE
  RESOLVED
  CLOSED
}

enum TicketPriority {
  LOW
  NORMAL
  HIGH
  URGENT
}

// ═══════════════════════════════════════════════════════════
// GIVEAWAY SİSTEMİ
// ═══════════════════════════════════════════════════════════

model Giveaway {
  id            String    @id @default(cuid())
  messageId     String    @unique @map("message_id")
  channelId     String    @map("channel_id")
  hostId        String    @map("host_id")
  prize         String
  description   String?
  winnerCount   Int       @default(1) @map("winner_count")
  endsAt        DateTime  @map("ends_at")
  endedAt       DateTime? @map("ended_at")
  isEnded       Boolean   @default(false) @map("is_ended")
  createdAt     DateTime  @default(now()) @map("created_at")

  // Şartlar
  requireLinked   Boolean @default(false) @map("require_linked")
  requireRole     String? @map("require_role")
  minPlaytime     Int?    @map("min_playtime") // dakika
  minLevel        Int?    @map("min_level")

  entries  GiveawayEntry[]
  winners  GiveawayWinner[]

  @@map("discord_giveaways")
}

model GiveawayEntry {
  id          String   @id @default(cuid())
  giveawayId  String   @map("giveaway_id")
  giveaway    Giveaway @relation(fields: [giveawayId], references: [id], onDelete: Cascade)
  userId      String   @map("user_id")
  enteredAt   DateTime @default(now()) @map("entered_at")

  @@unique([giveawayId, userId])
  @@map("discord_giveaway_entries")
}

model GiveawayWinner {
  id          String   @id @default(cuid())
  giveawayId  String   @map("giveaway_id")
  giveaway    Giveaway @relation(fields: [giveawayId], references: [id], onDelete: Cascade)
  userId      String   @map("user_id")
  claimed     Boolean  @default(false)
  claimedAt   DateTime? @map("claimed_at")

  @@map("discord_giveaway_winners")
}

// ═══════════════════════════════════════════════════════════
// MODERASYON
// ═══════════════════════════════════════════════════════════

model Warning {
  id        String   @id @default(cuid())
  userId    String   @map("user_id")
  moderator String   @map("moderator_id")
  reason    String
  createdAt DateTime @default(now()) @map("created_at")
  expiresAt DateTime? @map("expires_at")
  isActive  Boolean  @default(true) @map("is_active")

  @@map("discord_warnings")
}

model ModAction {
  id        String       @id @default(cuid())
  userId    String       @map("user_id")
  moderator String       @map("moderator_id")
  action    ModActionType
  reason    String?
  duration  Int?         // dakika
  createdAt DateTime     @default(now()) @map("created_at")
  expiresAt DateTime?    @map("expires_at")

  @@map("discord_mod_actions")
}

enum ModActionType {
  WARN
  MUTE
  KICK
  BAN
  UNMUTE
  UNBAN
}

// ═══════════════════════════════════════════════════════════
// CHAT KÖPRÜSÜ LOG
// ═══════════════════════════════════════════════════════════

model ChatLog {
  id          String      @id @default(cuid())
  direction   ChatDirection
  userId      String      @map("user_id")
  username    String
  content     String
  server      String?     // MC server name
  channelId   String?     @map("channel_id")
  messageId   String?     @map("message_id")
  createdAt   DateTime    @default(now()) @map("created_at")

  @@index([userId, createdAt])
  @@map("discord_chat_logs")
}

enum ChatDirection {
  MC_TO_DISCORD
  DISCORD_TO_MC
}

// ═══════════════════════════════════════════════════════════
// BOT AYARLARI
// ═══════════════════════════════════════════════════════════

model BotConfig {
  key       String   @id
  value     String
  updatedAt DateTime @updatedAt @map("updated_at")

  @@map("discord_bot_config")
}
```

---

## Hesap Bağlama Sistemi

### Bağlama Akışı

```
┌─────────────────────────────────────────────────────────────────┐
│                     BAĞLAMA AKIŞ DİYAGRAMI                      │
│                                                                 │
│  MINECRAFT                           DISCORD                   │
│  ──────────                           ───────                   │
│                                                                 │
│  /discord link                                                  │
│       │                                                         │
│       ▼                                                         │
│  ┌─────────────┐                                               │
│  │ Kod üret:   │                                               │
│  │ "ABC123"    │                                               │
│  │ (5dk geçerli)                                               │
│  └─────────────┘                                               │
│       │                                                         │
│       │ Redis'e kaydet                                         │
│       ▼                                                         │
│  "Discord'da /bağla ABC123 yaz"                                │
│                                                                 │
│                                      /bağla ABC123             │
│                                           │                    │
│                                           ▼                    │
│                                    ┌─────────────┐             │
│                                    │ Kod kontrol │             │
│                                    │ Redis'te    │             │
│                                    │ var mı?     │             │
│                                    └─────────────┘             │
│                                      │    │                    │
│                              Yok ◄───┘    └───► Var            │
│                               │                  │             │
│                               ▼                  ▼             │
│                          "Geçersiz         ┌──────────┐        │
│                           kod!"            │ Hesapları│        │
│                                            │ bağla    │        │
│                                            │ (DB)     │        │
│                                            └──────────┘        │
│                                                 │              │
│                                                 ▼              │
│                                           ┌──────────┐         │
│                                           │Rol ata   │         │
│                                           │Nick sync │         │
│                                           └──────────┘         │
│                                                 │              │
│       ┌─────────────────────────────────────────┘              │
│       │                                                        │
│       ▼                                                        │
│  "Hesabın bağlandı!"                "Hesabın bağlandı!"       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### MC Komutu Mesajı

```
┌─────────────────────────────────────────────────────────────────┐
│  &6&l⚡ &eDiscord Bağlantısı                                    │
│                                                                 │
│  &7Bağlantı kodun: &f&lABC123                                  │
│                                                                 │
│  &7Discord'da şu komutu yaz:                                   │
│  &b/bağla ABC123                                               │
│                                                                 │
│  &8Kod 5 dakika geçerlidir.                                    │
│  &8Discord: &7discord.gg/karapixel                             │
└─────────────────────────────────────────────────────────────────┘
```

### Discord Komutu Kodu

```typescript
// src/commands/minecraft/link.ts

import { SlashCommandBuilder, ChatInputCommandInteraction } from 'discord.js';
import { prisma } from '../../database/client';
import { redis } from '../../utils/cache.util';
import { successEmbed, errorEmbed } from '../../utils/embed.util';
import { config } from '../../config';

export const data = new SlashCommandBuilder()
  .setName('bağla')
  .setDescription('Minecraft hesabını Discord ile bağla')
  .addStringOption(option =>
    option
      .setName('kod')
      .setDescription('Oyun içinde aldığın bağlantı kodu')
      .setRequired(true)
      .setMinLength(6)
      .setMaxLength(6)
  );

export async function execute(interaction: ChatInputCommandInteraction) {
  const code = interaction.options.getString('kod', true).toUpperCase();
  const userId = interaction.user.id;

  // Redis'ten kodu kontrol et
  const linkData = await redis.get(`link:${code}`);

  if (!linkData) {
    return interaction.reply({
      embeds: [errorEmbed('Geçersiz veya süresi dolmuş kod!')],
      ephemeral: true
    });
  }

  const { uuid, name } = JSON.parse(linkData);

  // Zaten bağlı mı kontrol et
  const existing = await prisma.linkedAccount.findFirst({
    where: {
      OR: [
        { discordId: userId },
        { minecraftUuid: uuid }
      ]
    }
  });

  if (existing) {
    return interaction.reply({
      embeds: [errorEmbed('Bu hesaplardan biri zaten bağlı!')],
      ephemeral: true
    });
  }

  // Hesabı bağla
  await prisma.linkedAccount.create({
    data: {
      minecraftUuid: uuid,
      minecraftName: name,
      discordId: userId,
      discordTag: interaction.user.tag
    }
  });

  // Redis'ten kodu sil
  await redis.del(`link:${code}`);

  // Rol ata
  const member = interaction.member as GuildMember;
  const role = interaction.guild!.roles.cache.get(config.roles.linked);
  if (role) {
    await member.roles.add(role);
  }

  // Nickname güncelle
  try {
    await member.setNickname(name);
  } catch (e) {
    // Bot'un yetkisi yetmeyebilir
  }

  // MC'ye bildir (Redis Pub/Sub)
  await redis.publish('discord:linked', JSON.stringify({
    uuid,
    discordId: userId
  }));

  return interaction.reply({
    embeds: [successEmbed(
      `Hesabın başarıyla bağlandı!\n` +
      `Minecraft: **${name}**`
    )]
  });
}
```

### Bağlantı Avantajları

| Avantaj | Açıklama |
|---------|----------|
| 🎮 Oyuncu Rolü | Discord'da özel "Oyuncu" rolü |
| 📝 Nick Sync | MC nick = Discord nick |
| 👑 Rank Sync | VIP → Discord VIP rolü |
| 🎁 Günlük Bonus | +1 Vote Key (Discord bağlıysa) |
| 📢 DM Bildirimi | Event bildirimleri DM'den |
| 🎉 Giveaway | Bazı çekilişlere katılım hakkı |
| 📊 Profil | /profil komutu kullanımı |

---

## Chat Köprüsü

### Çift Yönlü İletişim

```
┌─────────────────────────────────────────────────────────────────┐
│                      CHAT KÖPRÜSÜ                               │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  MC → DISCORD:                                                  │
│  ──────────────                                                 │
│  ├── Global chat → #oyun-sohbet kanalı                        │
│  ├── Staff chat → #staff-chat kanalı (kilitli)                │
│  ├── Ölüm mesajları → #olaylar kanalı                         │
│  ├── Achievement → #olaylar kanalı                            │
│  └── Join/Leave → #olaylar kanalı                             │
│                                                                 │
│  DISCORD → MC:                                                  │
│  ──────────────                                                 │
│  ├── #oyun-sohbet → Global chat                               │
│  ├── Format: [Discord] kullanıcı: mesaj                       │
│  └── Sadece bağlı hesaplar (opsiyonel)                        │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Discord'da Görünüm

```
┌─────────────────────────────────────────────────────────────────┐
│  #oyun-sohbet                                                   │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  🎮 [VIP] PlayerName: Merhaba!                                 │
│  ──────────────────────────────────────────────────────────     │
│  🎉 PlayerName bir başarım kazandı: İlk Elmas!                 │
│  ──────────────────────────────────────────────────────────     │
│  💀 PlayerName düşerek öldü                                    │
│  ──────────────────────────────────────────────────────────     │
│  ➡️ PlayerName sunucuya katıldı (127 online)                  │
│  ──────────────────────────────────────────────────────────     │
│  ⬅️ PlayerName sunucudan ayrıldı (126 online)                 │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Chat Bridge Service

```typescript
// src/services/chat-bridge.service.ts

import { TextChannel, EmbedBuilder } from 'discord.js';
import { redis } from '../utils/cache.util';
import { prisma } from '../database/client';
import { client } from '../client';
import { config } from '../config';

export class ChatBridgeService {
  private chatChannel: TextChannel | null = null;
  private eventsChannel: TextChannel | null = null;

  async initialize() {
    this.chatChannel = await client.channels.fetch(config.channels.chat) as TextChannel;
    this.eventsChannel = await client.channels.fetch(config.channels.events) as TextChannel;

    // Redis subscriber
    const subscriber = redis.duplicate();
    await subscriber.subscribe('mc:chat', 'mc:death', 'mc:achievement', 'mc:join', 'mc:leave');

    subscriber.on('message', async (channel, message) => {
      const data = JSON.parse(message);
      
      switch (channel) {
        case 'mc:chat':
          await this.handleMCChat(data);
          break;
        case 'mc:death':
          await this.handleDeath(data);
          break;
        case 'mc:achievement':
          await this.handleAchievement(data);
          break;
        case 'mc:join':
          await this.handleJoin(data);
          break;
        case 'mc:leave':
          await this.handleLeave(data);
          break;
      }
    });
  }

  async handleMCChat(data: { uuid: string; name: string; rank?: string; message: string }) {
    if (!this.chatChannel) return;

    const rankPrefix = data.rank ? `[${data.rank}] ` : '';
    await this.chatChannel.send(`🎮 ${rankPrefix}**${data.name}**: ${data.message}`);

    // Log
    await prisma.chatLog.create({
      data: {
        direction: 'MC_TO_DISCORD',
        userId: data.uuid,
        username: data.name,
        content: data.message,
        channelId: this.chatChannel.id
      }
    });
  }

  async handleDeath(data: { name: string; message: string }) {
    if (!this.eventsChannel) return;
    await this.eventsChannel.send(`💀 ${data.message}`);
  }

  async handleAchievement(data: { name: string; achievement: string }) {
    if (!this.eventsChannel) return;
    await this.eventsChannel.send(`🎉 **${data.name}** bir başarım kazandı: **${data.achievement}**!`);
  }

  async handleJoin(data: { name: string; online: number }) {
    if (!this.eventsChannel) return;
    await this.eventsChannel.send(`➡️ **${data.name}** sunucuya katıldı (${data.online} online)`);
  }

  async handleLeave(data: { name: string; online: number }) {
    if (!this.eventsChannel) return;
    await this.eventsChannel.send(`⬅️ **${data.name}** sunucudan ayrıldı (${data.online} online)`);
  }

  // Discord → MC
  async sendToMinecraft(discordId: string, username: string, content: string) {
    // Bağlı hesap kontrolü (opsiyonel)
    const linked = await prisma.linkedAccount.findUnique({
      where: { discordId }
    });

    const mcName = linked?.minecraftName || username;

    await redis.publish('discord:chat', JSON.stringify({
      discordId,
      username: mcName,
      content,
      isLinked: !!linked
    }));

    // Log
    await prisma.chatLog.create({
      data: {
        direction: 'DISCORD_TO_MC',
        userId: discordId,
        username,
        content
      }
    });
  }
}

export const chatBridge = new ChatBridgeService();
```

---

## Sunucu Durumu

### Auto-Update Embed

```
┌─────────────────────────────────────────────────────────────────┐
│  🌟 KaraPixel Sunucu Durumu 🌟                                 │
│  ─────────────────────────────────                              │
│                                                                 │
│  📊 Durum: 🟢 Çevrimiçi                                        │
│  👥 Oyuncular: 127/500                                         │
│  ⏰ Uptime: 5 gün 3 saat                                       │
│  🎮 TPS: 19.8                                                  │
│                                                                 │
│  📈 Bugün:                                                     │
│  ├── Benzersiz Giriş: 450                                     │
│  ├── Yeni Kayıt: 23                                           │
│  └── Peak: 203 (14:30)                                        │
│                                                                 │
│  🌐 IP: play.karapixel.net                                    │
│  📅 Son Güncelleme: <timestamp>                                │
└─────────────────────────────────────────────────────────────────┘
```

### Status Service

```typescript
// src/services/status.service.ts

import { TextChannel, EmbedBuilder } from 'discord.js';
import { redis } from '../utils/cache.util';
import { client } from '../client';
import { config } from '../config';

interface ServerStatus {
  online: boolean;
  players: {
    online: number;
    max: number;
    list: string[];
  };
  tps: number;
  uptime: number;
  todayStats: {
    uniqueJoins: number;
    newRegistrations: number;
    peakPlayers: number;
    peakTime: string;
  };
}

export class StatusService {
  private statusMessage: Message | null = null;
  private statusChannel: TextChannel | null = null;

  async initialize() {
    this.statusChannel = await client.channels.fetch(config.channels.status) as TextChannel;
    
    // Mevcut status mesajını bul veya oluştur
    const messages = await this.statusChannel.messages.fetch({ limit: 10 });
    this.statusMessage = messages.find(m => m.author.id === client.user!.id) || null;

    if (!this.statusMessage) {
      this.statusMessage = await this.statusChannel.send({ embeds: [this.createEmbed(null)] });
    }

    // Her 60 saniyede güncelle
    setInterval(() => this.updateStatus(), 60000);
    
    // İlk güncelleme
    await this.updateStatus();
  }

  async updateStatus() {
    const status = await this.fetchStatus();
    const embed = this.createEmbed(status);
    
    if (this.statusMessage) {
      await this.statusMessage.edit({ embeds: [embed] });
    }

    // Bot status güncelle
    client.user?.setActivity(`${status?.players.online || 0} oyuncu | play.karapixel.net`, {
      type: ActivityType.Watching
    });
  }

  async fetchStatus(): Promise<ServerStatus | null> {
    try {
      const data = await redis.get('server:status');
      return data ? JSON.parse(data) : null;
    } catch {
      return null;
    }
  }

  createEmbed(status: ServerStatus | null): EmbedBuilder {
    const embed = new EmbedBuilder()
      .setTitle('🌟 KaraPixel Sunucu Durumu 🌟')
      .setColor(status?.online ? 0x00FF00 : 0xFF0000)
      .setTimestamp();

    if (!status) {
      embed.setDescription('⚠️ Sunucu durumu alınamıyor...');
      return embed;
    }

    const uptimeStr = this.formatUptime(status.uptime);
    const tpsColor = status.tps >= 19 ? '🟢' : status.tps >= 15 ? '🟡' : '🔴';

    embed.addFields(
      { name: '📊 Durum', value: status.online ? '🟢 Çevrimiçi' : '🔴 Çevrimdışı', inline: true },
      { name: '👥 Oyuncular', value: `${status.players.online}/${status.players.max}`, inline: true },
      { name: '⏰ Uptime', value: uptimeStr, inline: true },
      { name: '🎮 TPS', value: `${tpsColor} ${status.tps.toFixed(1)}`, inline: true },
      { name: '\u200B', value: '\u200B', inline: true },
      { name: '\u200B', value: '\u200B', inline: true },
      { name: '📈 Bugün', value: [
        `├── Benzersiz Giriş: **${status.todayStats.uniqueJoins}**`,
        `├── Yeni Kayıt: **${status.todayStats.newRegistrations}**`,
        `└── Peak: **${status.todayStats.peakPlayers}** (${status.todayStats.peakTime})`
      ].join('\n'), inline: false },
      { name: '🌐 IP', value: '`play.karapixel.net`', inline: false }
    );

    return embed;
  }

  formatUptime(seconds: number): string {
    const days = Math.floor(seconds / 86400);
    const hours = Math.floor((seconds % 86400) / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);

    const parts = [];
    if (days > 0) parts.push(`${days} gün`);
    if (hours > 0) parts.push(`${hours} saat`);
    if (minutes > 0) parts.push(`${minutes} dk`);

    return parts.join(' ') || '< 1 dk';
  }
}

export const statusService = new StatusService();
```

---

## Leaderboard Sistemi

### Leaderboard Komutu

```typescript
// src/commands/minecraft/leaderboard.ts

import { SlashCommandBuilder, ChatInputCommandInteraction, EmbedBuilder } from 'discord.js';

const CATEGORIES = {
  ada: { name: 'Top Adalar', emoji: '🏝️', field: 'islandValue' },
  zengin: { name: 'En Zenginler', emoji: '💰', field: 'balance' },
  seviye: { name: 'En Yüksek Seviye', emoji: '⭐', field: 'level' },
  sure: { name: 'En Aktif Oyuncular', emoji: '⏰', field: 'playtime' }
};

export const data = new SlashCommandBuilder()
  .setName('sıralama')
  .setDescription('Sunucu sıralamalarını görüntüle')
  .addStringOption(option =>
    option
      .setName('kategori')
      .setDescription('Sıralama kategorisi')
      .setRequired(true)
      .addChoices(
        { name: '🏝️ Ada Değeri', value: 'ada' },
        { name: '💰 Para', value: 'zengin' },
        { name: '⭐ Seviye', value: 'seviye' },
        { name: '⏰ Oynama Süresi', value: 'sure' }
      )
  );

export async function execute(interaction: ChatInputCommandInteraction) {
  await interaction.deferReply();

  const category = interaction.options.getString('kategori', true) as keyof typeof CATEGORIES;
  const { name, emoji, field } = CATEGORIES[category];

  // API'den veri çek
  const data = await fetch(`${config.api.url}/leaderboard/${field}?limit=10`)
    .then(res => res.json());

  const medals = ['🥇', '🥈', '🥉'];
  const lines = data.map((entry: any, index: number) => {
    const medal = medals[index] || `${index + 1}.`;
    const value = formatValue(entry.value, field);
    const bar = createBar(entry.value, data[0].value);
    return `${medal} **${entry.name}** │ ${value} │ ${bar}`;
  });

  const embed = new EmbedBuilder()
    .setTitle(`${emoji} ${name} ${emoji}`)
    .setDescription(lines.join('\n'))
    .setColor(0x9B59B6)
    .setFooter({ text: 'Güncelleme: Her 5 dakika' })
    .setTimestamp();

  await interaction.editReply({ embeds: [embed] });
}

function formatValue(value: number, field: string): string {
  switch (field) {
    case 'islandValue':
    case 'balance':
      return formatNumber(value);
    case 'playtime':
      return `${Math.floor(value / 60)} saat`;
    default:
      return value.toString();
  }
}

function formatNumber(num: number): string {
  if (num >= 1000000) return `${(num / 1000000).toFixed(1)}M`;
  if (num >= 1000) return `${(num / 1000).toFixed(1)}K`;
  return num.toString();
}

function createBar(value: number, max: number): string {
  const filled = Math.round((value / max) * 10);
  return '█'.repeat(filled) + '░'.repeat(10 - filled);
}
```

### Embed Örneği

```
┌─────────────────────────────────────────────────────────────────┐
│  🏝️ Top Adalar 🏝️                                              │
│  ─────────────────────────────────────────────────────          │
│                                                                 │
│  🥇 **TeamAlpha**     │ 15.2M │ ██████████████████████████████  │
│  🥈 **DragonLords**   │ 12.4M │ ████████████████████████        │
│  🥉 **SkyMasters**    │ 10.2M │ ████████████████████            │
│  4. **CoolKids**      │ 8.7M  │ █████████████████               │
│  5. **ProPlayers**    │ 7.6M  │ ██████████████                  │
│  6. **TeamBeta**      │ 6.5M  │ ████████████                    │
│  7. **Legends**       │ 5.4M  │ ██████████                      │
│  8. **Warriors**      │ 4.3M  │ ████████                        │
│  9. **Knights**       │ 3.2M  │ ██████                          │
│  10. **Heroes**       │ 2.1M  │ ████                            │
│                                                                 │
│  ─────────────────────────────────────────────────────          │
│  Güncelleme: Her 5 dakika                                      │
└─────────────────────────────────────────────────────────────────┘
```

---

## Profil Sistemi

### Profil Komutu

```typescript
// src/commands/minecraft/profile.ts

import { SlashCommandBuilder, ChatInputCommandInteraction, EmbedBuilder } from 'discord.js';
import { prisma } from '../../database/client';

export const data = new SlashCommandBuilder()
  .setName('profil')
  .setDescription('Minecraft profil bilgilerini görüntüle')
  .addUserOption(option =>
    option
      .setName('kullanıcı')
      .setDescription('Profili görüntülenecek kullanıcı')
      .setRequired(false)
  );

export async function execute(interaction: ChatInputCommandInteraction) {
  await interaction.deferReply();

  const targetUser = interaction.options.getUser('kullanıcı') || interaction.user;

  // Bağlı hesap kontrolü
  const linked = await prisma.linkedAccount.findUnique({
    where: { discordId: targetUser.id }
  });

  if (!linked) {
    return interaction.editReply({
      embeds: [errorEmbed(
        targetUser.id === interaction.user.id
          ? 'Hesabın bağlı değil! `/bağla` komutuyla bağlayabilirsin.'
          : 'Bu kullanıcının bağlı Minecraft hesabı yok.'
      )]
    });
  }

  // MC API'den profil verisi çek
  const profile = await fetch(`${config.api.url}/player/${linked.minecraftUuid}`)
    .then(res => res.json());

  const embed = new EmbedBuilder()
    .setTitle(`👤 ${profile.name} Profili`)
    .setThumbnail(`https://mc-heads.net/avatar/${linked.minecraftUuid}`)
    .setColor(getRankColor(profile.rank))
    .addFields(
      { name: '🎖️ Rank', value: profile.rank || 'Oyuncu', inline: true },
      { name: '📊 Seviye', value: profile.level.toString(), inline: true },
      { name: '💰 Para', value: formatNumber(profile.balance) + ' Koin', inline: true },
      { name: '🏝️ Ada', value: profile.island ? `"${profile.island.name}" (Lv.${profile.island.level})` : 'Yok', inline: true },
      { name: '⏰ Toplam Süre', value: formatPlaytime(profile.playtime), inline: true },
      { name: '📅 İlk Giriş', value: formatDate(profile.firstJoin), inline: true },
      { name: '🕐 Son Görülme', value: profile.online ? '🟢 Şu an online' : formatRelative(profile.lastSeen), inline: false }
    );

  // Başarımlar
  if (profile.achievements?.length > 0) {
    const achievementList = profile.achievements
      .slice(0, 5)
      .map((a: any) => `⭐ ${a.name}`)
      .join('\n');
    
    embed.addFields({
      name: `🏆 Başarımlar (${profile.achievementCount}/120)`,
      value: achievementList,
      inline: false
    });
  }

  // Discord bağlantısı
  embed.addFields({
    name: '🔗 Discord',
    value: `<@${targetUser.id}>`,
    inline: false
  });

  embed.setFooter({ text: 'play.karapixel.net' });
  embed.setTimestamp();

  await interaction.editReply({ embeds: [embed] });
}
```

### Embed Örneği

```
┌─────────────────────────────────────────────────────────────────┐
│  👤 PlayerName Profili                                         │
│  ─────────────────────────                                      │
│                                    ┌──────────┐                 │
│  🎖️ Rank: VIP+                    │  [HEAD]  │                 │
│  📊 Seviye: 45                     │          │                 │
│  💰 Para: 1.234.567 Koin          └──────────┘                 │
│  🏝️ Ada: "SkyHaven" (Lv.12)                                    │
│  ⏰ Toplam Süre: 127 saat                                      │
│  📅 İlk Giriş: 15 Ekim 2024                                    │
│  🕐 Son Görülme: 🟢 Şu an online                               │
│                                                                 │
│  🏆 Başarımlar (45/120)                                        │
│  ├── ⭐ İlk Elmas                                              │
│  ├── ⭐ Ada Ustası                                             │
│  ├── ⭐ Zenginlik (Lv.5)                                       │
│  ├── ⭐ Madenci                                                │
│  └── ⭐ İlk Pet                                                │
│                                                                 │
│  🔗 Discord: @username                                         │
│  ─────────────────────────────────────────────────────          │
│  play.karapixel.net                                            │
└─────────────────────────────────────────────────────────────────┘
```

---

## Ticket Sistemi

### Ticket Oluşturma

```
┌─────────────────────────────────────────────────────────────────┐
│  #destek kanalında buton:                                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  🎫 Destek Talebi Oluştur                                      │
│  ─────────────────────────────                                  │
│                                                                 │
│  Aşağıdaki kategorilerden birini seçerek                       │
│  destek talebi oluşturabilirsin.                               │
│                                                                 │
│  ┌─────────────────────────────────────────────────────┐       │
│  │ 🐛 Bug Bildirimi                                    │       │
│  │ 💰 Ödeme Sorunu                                     │       │
│  │ 🚫 Oyuncu Şikayeti                                  │       │
│  │ ❓ Genel Soru                                       │       │
│  │ 🎁 Ödül Talebi                                      │       │
│  │ 📝 İtiraz                                           │       │
│  └─────────────────────────────────────────────────────┘       │
│                                                                 │
│  [ 📩 Ticket Oluştur ]                                         │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Ticket Service

```typescript
// src/services/ticket.service.ts

import { 
  TextChannel, 
  CategoryChannel, 
  PermissionFlagsBits,
  EmbedBuilder,
  ActionRowBuilder,
  ButtonBuilder,
  ButtonStyle,
  StringSelectMenuBuilder
} from 'discord.js';
import { prisma } from '../database/client';
import { config } from '../config';

const CATEGORIES = {
  BUG_REPORT: { name: 'Bug Bildirimi', emoji: '🐛', color: 0xFF6B6B },
  PAYMENT_ISSUE: { name: 'Ödeme Sorunu', emoji: '💰', color: 0xFFD93D },
  PLAYER_REPORT: { name: 'Oyuncu Şikayeti', emoji: '🚫', color: 0xFF6B6B },
  GENERAL_QUESTION: { name: 'Genel Soru', emoji: '❓', color: 0x4ECDC4 },
  REWARD_CLAIM: { name: 'Ödül Talebi', emoji: '🎁', color: 0x95E1D3 },
  APPEAL: { name: 'İtiraz', emoji: '📝', color: 0xDDA0DD }
};

export class TicketService {
  async createTicket(
    guild: Guild,
    userId: string,
    category: TicketCategory,
    subject?: string
  ): Promise<TextChannel> {
    const user = await guild.members.fetch(userId);
    const categoryData = CATEGORIES[category];
    
    // Ticket sayısını al
    const count = await prisma.ticket.count() + 1;
    
    // Kanal oluştur
    const ticketCategory = guild.channels.cache.get(config.channels.ticketCategory) as CategoryChannel;
    
    const channel = await guild.channels.create({
      name: `ticket-${count}`,
      type: ChannelType.GuildText,
      parent: ticketCategory,
      permissionOverwrites: [
        {
          id: guild.id,
          deny: [PermissionFlagsBits.ViewChannel]
        },
        {
          id: userId,
          allow: [
            PermissionFlagsBits.ViewChannel,
            PermissionFlagsBits.SendMessages,
            PermissionFlagsBits.ReadMessageHistory
          ]
        },
        {
          id: config.roles.staff,
          allow: [
            PermissionFlagsBits.ViewChannel,
            PermissionFlagsBits.SendMessages,
            PermissionFlagsBits.ReadMessageHistory,
            PermissionFlagsBits.ManageMessages
          ]
        }
      ]
    });

    // Database'e kaydet
    const ticket = await prisma.ticket.create({
      data: {
        channelId: channel.id,
        userId,
        category,
        subject
      }
    });

    // Hoş geldin mesajı
    const embed = new EmbedBuilder()
      .setTitle(`${categoryData.emoji} ${categoryData.name}`)
      .setDescription(
        `Merhaba ${user}!\n\n` +
        `Destek talebini aldık. Lütfen sorununuzu detaylı bir şekilde açıklayın.\n` +
        `Ekip en kısa sürede size dönüş yapacaktır.`
      )
      .setColor(categoryData.color)
      .addFields(
        { name: '📋 Ticket ID', value: `#${ticket.ticketId}`, inline: true },
        { name: '📁 Kategori', value: categoryData.name, inline: true },
        { name: '📅 Oluşturulma', value: `<t:${Math.floor(Date.now() / 1000)}:R>`, inline: true }
      )
      .setFooter({ text: 'Ticket kapatmak için aşağıdaki butonu kullanın.' });

    const row = new ActionRowBuilder<ButtonBuilder>()
      .addComponents(
        new ButtonBuilder()
          .setCustomId('ticket_close')
          .setLabel('Ticket Kapat')
          .setStyle(ButtonStyle.Danger)
          .setEmoji('🔒'),
        new ButtonBuilder()
          .setCustomId('ticket_claim')
          .setLabel('Üstlen')
          .setStyle(ButtonStyle.Primary)
          .setEmoji('✋')
      );

    await channel.send({ embeds: [embed], components: [row] });
    await channel.send(`<@${userId}> <@&${config.roles.staff}>`);

    return channel;
  }

  async closeTicket(
    channelId: string,
    closedBy: string,
    reason?: string
  ): Promise<void> {
    const ticket = await prisma.ticket.findUnique({
      where: { channelId },
      include: { messages: true }
    });

    if (!ticket) return;

    // Log kanalına transcript gönder
    await this.sendTranscript(ticket);

    // Database güncelle
    await prisma.ticket.update({
      where: { channelId },
      data: {
        status: 'CLOSED',
        closedAt: new Date(),
        closedBy,
        closeReason: reason
      }
    });

    // Kanalı sil
    const channel = await client.channels.fetch(channelId) as TextChannel;
    await channel.delete();
  }

  async sendTranscript(ticket: any): Promise<void> {
    const logChannel = await client.channels.fetch(config.channels.ticketLogs) as TextChannel;

    const embed = new EmbedBuilder()
      .setTitle(`📋 Ticket #${ticket.ticketId} Kapandı`)
      .setColor(0x95a5a6)
      .addFields(
        { name: 'Kullanıcı', value: `<@${ticket.userId}>`, inline: true },
        { name: 'Kategori', value: ticket.category, inline: true },
        { name: 'Mesaj Sayısı', value: ticket.messages.length.toString(), inline: true },
        { name: 'Kapatan', value: `<@${ticket.closedBy}>`, inline: true },
        { name: 'Sebep', value: ticket.closeReason || 'Belirtilmedi', inline: true }
      )
      .setTimestamp();

    // Transcript dosyası oluştur
    const transcript = ticket.messages
      .map((m: any) => `[${m.createdAt.toISOString()}] ${m.isStaff ? '[STAFF]' : ''} ${m.userId}: ${m.content}`)
      .join('\n');

    const buffer = Buffer.from(transcript, 'utf-8');
    
    await logChannel.send({
      embeds: [embed],
      files: [{
        attachment: buffer,
        name: `ticket-${ticket.ticketId}-transcript.txt`
      }]
    });
  }
}

export const ticketService = new TicketService();
```

---

## Giveaway Sistemi

### Giveaway Komutu

```typescript
// src/commands/giveaway/start.ts

import { 
  SlashCommandBuilder, 
  ChatInputCommandInteraction,
  EmbedBuilder,
  ActionRowBuilder,
  ButtonBuilder,
  ButtonStyle
} from 'discord.js';
import { prisma } from '../../database/client';
import ms from 'ms';

export const data = new SlashCommandBuilder()
  .setName('çekiliş')
  .setDescription('Yeni bir çekiliş başlat')
  .addStringOption(option =>
    option
      .setName('ödül')
      .setDescription('Çekiliş ödülü')
      .setRequired(true)
  )
  .addStringOption(option =>
    option
      .setName('süre')
      .setDescription('Çekiliş süresi (örn: 1h, 1d, 1w)')
      .setRequired(true)
  )
  .addIntegerOption(option =>
    option
      .setName('kazanan')
      .setDescription('Kazanan sayısı')
      .setRequired(false)
      .setMinValue(1)
      .setMaxValue(10)
  )
  .addBooleanOption(option =>
    option
      .setName('hesap-bağlı')
      .setDescription('MC hesabı bağlı olmalı mı?')
      .setRequired(false)
  )
  .addIntegerOption(option =>
    option
      .setName('min-süre')
      .setDescription('Minimum oynama süresi (saat)')
      .setRequired(false)
  )
  .addRoleOption(option =>
    option
      .setName('rol')
      .setDescription('Gerekli rol')
      .setRequired(false)
  );

export async function execute(interaction: ChatInputCommandInteraction) {
  // Admin kontrolü
  if (!interaction.member.permissions.has('Administrator')) {
    return interaction.reply({
      content: 'Bu komutu kullanma yetkin yok!',
      ephemeral: true
    });
  }

  const prize = interaction.options.getString('ödül', true);
  const duration = ms(interaction.options.getString('süre', true));
  const winnerCount = interaction.options.getInteger('kazanan') || 1;
  const requireLinked = interaction.options.getBoolean('hesap-bağlı') || false;
  const minPlaytime = interaction.options.getInteger('min-süre');
  const requireRole = interaction.options.getRole('rol');

  if (!duration || duration < 60000) {
    return interaction.reply({
      content: 'Geçersiz süre! Minimum 1 dakika olmalı.',
      ephemeral: true
    });
  }

  const endsAt = new Date(Date.now() + duration);

  const embed = new EmbedBuilder()
    .setTitle('🎁 ÇEKİLİŞ! 🎁')
    .setDescription(`**${prize}**`)
    .setColor(0xFFD700)
    .addFields(
      { name: '🏆 Kazanan Sayısı', value: winnerCount.toString(), inline: true },
      { name: '⏰ Bitiş', value: `<t:${Math.floor(endsAt.getTime() / 1000)}:R>`, inline: true },
      { name: '👤 Başlatan', value: `${interaction.user}`, inline: true }
    )
    .setFooter({ text: 'Katılmak için aşağıdaki butona tıkla!' })
    .setTimestamp(endsAt);

  // Şartlar varsa ekle
  const requirements: string[] = [];
  if (requireLinked) requirements.push('✅ MC hesabı bağlı olmalı');
  if (minPlaytime) requirements.push(`✅ En az ${minPlaytime} saat oynama süresi`);
  if (requireRole) requirements.push(`✅ ${requireRole} rolü gerekli`);

  if (requirements.length > 0) {
    embed.addFields({
      name: '📋 Şartlar',
      value: requirements.join('\n'),
      inline: false
    });
  }

  const row = new ActionRowBuilder<ButtonBuilder>()
    .addComponents(
      new ButtonBuilder()
        .setCustomId('giveaway_join')
        .setLabel('🎉 Katıl (0)')
        .setStyle(ButtonStyle.Success)
    );

  const message = await interaction.reply({
    embeds: [embed],
    components: [row],
    fetchReply: true
  });

  // Database'e kaydet
  await prisma.giveaway.create({
    data: {
      messageId: message.id,
      channelId: interaction.channelId,
      hostId: interaction.user.id,
      prize,
      winnerCount,
      endsAt,
      requireLinked,
      requireRole: requireRole?.id,
      minPlaytime: minPlaytime ? minPlaytime * 60 : null
    }
  });

  // Zamanlayıcı kur
  setTimeout(() => endGiveaway(message.id), duration);
}

async function endGiveaway(messageId: string) {
  const giveaway = await prisma.giveaway.findUnique({
    where: { messageId },
    include: { entries: true }
  });

  if (!giveaway || giveaway.isEnded) return;

  // Kazananları seç
  const entries = giveaway.entries;
  const winners: string[] = [];
  
  const shuffled = entries.sort(() => 0.5 - Math.random());
  for (let i = 0; i < Math.min(giveaway.winnerCount, entries.length); i++) {
    winners.push(shuffled[i].userId);
  }

  // Database güncelle
  await prisma.giveaway.update({
    where: { messageId },
    data: {
      isEnded: true,
      endedAt: new Date()
    }
  });

  // Kazananları kaydet
  for (const winnerId of winners) {
    await prisma.giveawayWinner.create({
      data: {
        giveawayId: giveaway.id,
        userId: winnerId
      }
    });
  }

  // Mesajı güncelle
  const channel = await client.channels.fetch(giveaway.channelId) as TextChannel;
  const message = await channel.messages.fetch(messageId);

  const winnerMentions = winners.length > 0
    ? winners.map(id => `<@${id}>`).join(', ')
    : 'Yeterli katılımcı yok!';

  const embed = EmbedBuilder.from(message.embeds[0])
    .setTitle('🎁 ÇEKİLİŞ BİTTİ! 🎁')
    .setColor(0x95a5a6)
    .spliceFields(0, 3)
    .addFields(
      { name: '🏆 Kazanan(lar)', value: winnerMentions, inline: false },
      { name: '👥 Toplam Katılım', value: entries.length.toString(), inline: true }
    );

  const row = new ActionRowBuilder<ButtonBuilder>()
    .addComponents(
      new ButtonBuilder()
        .setCustomId('giveaway_ended')
        .setLabel('Çekiliş Bitti')
        .setStyle(ButtonStyle.Secondary)
        .setDisabled(true)
    );

  await message.edit({ embeds: [embed], components: [row] });

  // Kazananlara DM gönder
  for (const winnerId of winners) {
    try {
      const user = await client.users.fetch(winnerId);
      await user.send({
        embeds: [
          new EmbedBuilder()
            .setTitle('🎉 Tebrikler!')
            .setDescription(`**${giveaway.prize}** çekilişini kazandın!`)
            .setColor(0xFFD700)
        ]
      });
    } catch (e) {
      // DM kapalı olabilir
    }
  }

  // Kanala da duyur
  await channel.send(`🎉 Tebrikler ${winnerMentions}! **${giveaway.prize}** kazandınız!`);
}
```

### Embed Örneği

```
┌─────────────────────────────────────────────────────────────────┐
│  🎁 ÇEKİLİŞ! 🎁                                                │
│  ─────────────────                                              │
│                                                                 │
│  **30 Gün VIP+**                                               │
│                                                                 │
│  🏆 Kazanan Sayısı: 1                                          │
│  ⏰ Bitiş: 24 saat sonra                                       │
│  👤 Başlatan: @Admin                                           │
│                                                                 │
│  📋 Şartlar:                                                   │
│  ✅ MC hesabı bağlı olmalı                                     │
│  ✅ En az 10 saat oynama süresi                                │
│                                                                 │
│  [ 🎉 Katıl (47) ]                                             │
│                                                                 │
│  ─────────────────────────────────────────────────────          │
│  Bitiş: 25 Aralık 2024 15:00                                   │
└─────────────────────────────────────────────────────────────────┘
```

---

## Event Bildirimleri

### Webhook Sistemi

```typescript
// src/listeners/event.listener.ts

import { EmbedBuilder, WebhookClient } from 'discord.js';
import { redis } from '../utils/cache.util';
import { config } from '../config';

const webhook = new WebhookClient({ url: config.webhooks.events });

export async function initializeEventListener() {
  const subscriber = redis.duplicate();
  await subscriber.subscribe(
    'event:starting',
    'event:started',
    'event:ended',
    'crate:legendary',
    'announcement'
  );

  subscriber.on('message', async (channel, message) => {
    const data = JSON.parse(message);

    switch (channel) {
      case 'event:starting':
        await sendEventStarting(data);
        break;
      case 'event:started':
        await sendEventStarted(data);
        break;
      case 'event:ended':
        await sendEventEnded(data);
        break;
      case 'crate:legendary':
        await sendLegendaryCrate(data);
        break;
      case 'announcement':
        await sendAnnouncement(data);
        break;
    }
  });
}

async function sendEventStarting(data: {
  name: string;
  description: string;
  startsIn: number; // dakika
  location: string;
  rewards: string[];
}) {
  const embed = new EmbedBuilder()
    .setTitle('🎉 EVENT BAŞLIYOR! 🎉')
    .setDescription(`**${data.name}**\n\n${data.description}`)
    .setColor(0xFFD700)
    .addFields(
      { name: '⏰ Ne Zaman', value: `${data.startsIn} dakika içinde!`, inline: true },
      { name: '📍 Konum', value: data.location, inline: true },
      { name: '🎁 Ödüller', value: data.rewards.join('\n'), inline: false }
    )
    .setFooter({ text: 'Kaçırmayın!' })
    .setTimestamp();

  await webhook.send({
    content: '@everyone',
    embeds: [embed]
  });
}

async function sendEventStarted(data: { name: string }) {
  const embed = new EmbedBuilder()
    .setTitle('🚀 EVENT BAŞLADI!')
    .setDescription(`**${data.name}** şu anda aktif!`)
    .setColor(0x00FF00)
    .setTimestamp();

  await webhook.send({ embeds: [embed] });
}

async function sendEventEnded(data: {
  name: string;
  winners: { name: string; prize: string }[];
}) {
  const winnerList = data.winners
    .map((w, i) => `${i === 0 ? '🥇' : i === 1 ? '🥈' : '🥉'} **${w.name}** - ${w.prize}`)
    .join('\n');

  const embed = new EmbedBuilder()
    .setTitle('🏆 EVENT BİTTİ!')
    .setDescription(`**${data.name}** sona erdi!`)
    .setColor(0x95a5a6)
    .addFields({
      name: '🎖️ Kazananlar',
      value: winnerList || 'Kazanan yok',
      inline: false
    })
    .setTimestamp();

  await webhook.send({ embeds: [embed] });
}

async function sendLegendaryCrate(data: {
  player: string;
  crate: string;
  reward: string;
  rarity: string;
}) {
  const embed = new EmbedBuilder()
    .setTitle('🌟 NADİR ÖDÜL! 🌟')
    .setDescription(`**${data.player}** bir **${data.crate}** açtı!`)
    .setColor(0x9B59B6)
    .addFields(
      { name: '🎁 Ödül', value: data.reward, inline: true },
      { name: '⚡ Nadirlik', value: data.rarity, inline: true }
    )
    .setTimestamp();

  await webhook.send({ embeds: [embed] });
}

async function sendAnnouncement(data: {
  title: string;
  content: string;
  color?: number;
  ping?: boolean;
}) {
  const embed = new EmbedBuilder()
    .setTitle(`📢 ${data.title}`)
    .setDescription(data.content)
    .setColor(data.color || 0x3498DB)
    .setTimestamp();

  await webhook.send({
    content: data.ping ? '@everyone' : undefined,
    embeds: [embed]
  });
}
```

---

## Moderasyon Araçları

### Komutlar

| Komut | Açıklama | Yetki |
|-------|----------|-------|
| `/uyar <kullanıcı> <sebep>` | Kullanıcıyı uyar | Moderatör |
| `/sustur <kullanıcı> <süre> [sebep]` | Kullanıcıyı sustur | Moderatör |
| `/at <kullanıcı> [sebep]` | Sunucudan at | Moderatör |
| `/yasakla <kullanıcı> [süre] [sebep]` | Sunucudan yasakla | Admin |
| `/yasakkaldır <kullanıcı>` | Yasağı kaldır | Admin |
| `/uyarılar <kullanıcı>` | Uyarıları görüntüle | Moderatör |
| `/temizle <sayı>` | Mesaj temizle | Moderatör |
| `/yavaşmod <süre>` | Slowmode ayarla | Moderatör |
| `/kilitle` | Kanalı kilitle | Moderatör |
| `/kilidaç` | Kanal kilidini aç | Moderatör |

### Uyarı Sistemi

```typescript
// src/commands/moderation/warn.ts

export const data = new SlashCommandBuilder()
  .setName('uyar')
  .setDescription('Bir kullanıcıyı uyar')
  .addUserOption(option =>
    option.setName('kullanıcı').setDescription('Uyarılacak kullanıcı').setRequired(true)
  )
  .addStringOption(option =>
    option.setName('sebep').setDescription('Uyarı sebebi').setRequired(true)
  );

export async function execute(interaction: ChatInputCommandInteraction) {
  const target = interaction.options.getUser('kullanıcı', true);
  const reason = interaction.options.getString('sebep', true);

  // Uyarıyı kaydet
  await prisma.warning.create({
    data: {
      userId: target.id,
      moderator: interaction.user.id,
      reason
    }
  });

  // Mod action log
  await prisma.modAction.create({
    data: {
      userId: target.id,
      moderator: interaction.user.id,
      action: 'WARN',
      reason
    }
  });

  // Toplam uyarı sayısını al
  const warningCount = await prisma.warning.count({
    where: { userId: target.id, isActive: true }
  });

  // Otomatik ceza
  if (warningCount >= 3) {
    // 3 uyarı = 1 saat mute
    const member = await interaction.guild!.members.fetch(target.id);
    await member.timeout(60 * 60 * 1000, 'Otomatik: 3 uyarı');
  }

  const embed = new EmbedBuilder()
    .setTitle('⚠️ Kullanıcı Uyarıldı')
    .setColor(0xFFD700)
    .addFields(
      { name: 'Kullanıcı', value: `${target}`, inline: true },
      { name: 'Moderatör', value: `${interaction.user}`, inline: true },
      { name: 'Sebep', value: reason, inline: false },
      { name: 'Toplam Uyarı', value: `${warningCount}/3`, inline: true }
    )
    .setTimestamp();

  await interaction.reply({ embeds: [embed] });

  // DM gönder
  try {
    await target.send({
      embeds: [
        new EmbedBuilder()
          .setTitle('⚠️ Uyarı Aldınız')
          .setDescription(`**KaraPixel Discord** sunucusunda uyarı aldınız.`)
          .addFields(
            { name: 'Sebep', value: reason },
            { name: 'Toplam Uyarı', value: `${warningCount}/3` }
          )
          .setColor(0xFFD700)
          .setFooter({ text: '3 uyarı = otomatik susturma' })
      ]
    });
  } catch (e) {
    // DM kapalı
  }

  // Log kanalına gönder
  const logChannel = await client.channels.fetch(config.channels.modLogs) as TextChannel;
  await logChannel.send({ embeds: [embed] });
}
```

---

## Eğlence Komutları

### Komut Listesi

| Komut | Açıklama |
|-------|----------|
| `/avatar [kullanıcı]` | Profil fotoğrafını göster |
| `/sunucu` | Sunucu bilgilerini göster |
| `/kim <kullanıcı>` | Kullanıcı bilgilerini göster |
| `/emoji <emoji>` | Emoji bilgisini göster |
| `/anket <soru>` | Basit anket oluştur |
| `/rastgele <min> <max>` | Rastgele sayı |
| `/yazıtura` | Yazı tura at |
| `/zar [yüz]` | Zar at |

---

## Konfigürasyon

### .env.example

```env
# Discord
DISCORD_TOKEN=your_bot_token
DISCORD_CLIENT_ID=your_client_id
DISCORD_GUILD_ID=your_guild_id

# Database
DATABASE_URL=postgresql://user:password@localhost:5432/karapixel

# Redis
REDIS_URL=redis://localhost:6379

# MC API
MC_API_URL=http://localhost:8080/api
MC_API_KEY=your_api_key

# Channels
CHANNEL_CHAT=channel_id
CHANNEL_EVENTS=channel_id
CHANNEL_STATUS=channel_id
CHANNEL_TICKET_CATEGORY=category_id
CHANNEL_TICKET_LOGS=channel_id
CHANNEL_MOD_LOGS=channel_id

# Roles
ROLE_LINKED=role_id
ROLE_STAFF=role_id
ROLE_VIP=role_id
ROLE_MVP=role_id
ROLE_KARA=role_id

# Webhooks
WEBHOOK_EVENTS=webhook_url
WEBHOOK_ANNOUNCEMENTS=webhook_url
```

### Config.ts

```typescript
// src/config.ts

export const config = {
  discord: {
    token: process.env.DISCORD_TOKEN!,
    clientId: process.env.DISCORD_CLIENT_ID!,
    guildId: process.env.DISCORD_GUILD_ID!
  },
  
  database: {
    url: process.env.DATABASE_URL!
  },
  
  redis: {
    url: process.env.REDIS_URL!
  },
  
  api: {
    url: process.env.MC_API_URL!,
    key: process.env.MC_API_KEY!
  },
  
  channels: {
    chat: process.env.CHANNEL_CHAT!,
    events: process.env.CHANNEL_EVENTS!,
    status: process.env.CHANNEL_STATUS!,
    ticketCategory: process.env.CHANNEL_TICKET_CATEGORY!,
    ticketLogs: process.env.CHANNEL_TICKET_LOGS!,
    modLogs: process.env.CHANNEL_MOD_LOGS!
  },
  
  roles: {
    linked: process.env.ROLE_LINKED!,
    staff: process.env.ROLE_STAFF!,
    vip: process.env.ROLE_VIP!,
    mvp: process.env.ROLE_MVP!,
    kara: process.env.ROLE_KARA!
  },
  
  webhooks: {
    events: process.env.WEBHOOK_EVENTS!,
    announcements: process.env.WEBHOOK_ANNOUNCEMENTS!
  },
  
  embed: {
    colors: {
      primary: 0x9B59B6,
      success: 0x2ECC71,
      error: 0xE74C3C,
      warning: 0xF39C12,
      info: 0x3498DB
    }
  }
};
```

---

## Docker Deployment

### Dockerfile

```dockerfile
FROM oven/bun:1 AS base
WORKDIR /app

# Install dependencies
FROM base AS install
COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile --production

# Build
FROM base AS build
COPY --from=install /app/node_modules ./node_modules
COPY . .
RUN bun run build

# Production
FROM base AS release
COPY --from=install /app/node_modules ./node_modules
COPY --from=build /app/dist ./dist
COPY --from=build /app/prisma ./prisma

# Generate Prisma client
RUN bunx prisma generate

USER bun
EXPOSE 3000
CMD ["bun", "run", "dist/index.js"]
```

### docker-compose.yml

```yaml
version: '3.8'

services:
  golge:
    build: .
    container_name: golge
    restart: unless-stopped
    env_file: .env
    depends_on:
      - redis
    networks:
      - karapixel

  redis:
    image: redis:7-alpine
    container_name: golge-redis
    restart: unless-stopped
    volumes:
      - redis-data:/data
    networks:
      - karapixel

networks:
  karapixel:
    external: true

volumes:
  redis-data:
```

---

## Komut Özeti

### Genel Komutlar

| Komut | Açıklama |
|-------|----------|
| `/ping` | Bot gecikmesini göster |
| `/yardım` | Komut listesi |
| `/bilgi` | Bot bilgileri |

### Minecraft Komutları

| Komut | Açıklama |
|-------|----------|
| `/bağla <kod>` | MC hesabını bağla |
| `/çöz` | Bağlantıyı kaldır |
| `/profil [kullanıcı]` | Profil görüntüle |
| `/sıralama <kategori>` | Leaderboard |

### Staff Komutları

| Komut | Açıklama | Yetki |
|-------|----------|-------|
| `/çekiliş ...` | Çekiliş başlat | Admin |
| `/duyuru ...` | Duyuru yap | Admin |
| `/uyar ...` | Kullanıcı uyar | Mod |
| `/sustur ...` | Kullanıcı sustur | Mod |
| `/yasakla ...` | Kullanıcı yasakla | Admin |

---

## Checklist

```
☐ Proje kurulumu
  ☐ Bun + TypeScript setup
  ☐ Discord.js v14 kurulum
  ☐ Prisma setup
  ☐ Redis bağlantısı

☐ Hesap Bağlama
  ☐ /bağla komutu
  ☐ /çöz komutu
  ☐ Rol otomasyonu
  ☐ Nick sync

☐ Chat Köprüsü
  ☐ MC → Discord
  ☐ Discord → MC
  ☐ Event mesajları

☐ Sunucu Durumu
  ☐ Auto-update embed
  ☐ Bot status

☐ Leaderboard
  ☐ /sıralama komutu
  ☐ Kategori seçimi

☐ Profil
  ☐ /profil komutu
  ☐ MC API entegrasyonu

☐ Ticket Sistemi
  ☐ Ticket oluşturma
  ☐ Kategori seçimi
  ☐ Ticket kapatma
  ☐ Transcript

☐ Giveaway
  ☐ /çekiliş komutu
  ☐ Şart kontrolü
  ☐ Otomatik bitiş
  ☐ Kazanan seçimi

☐ Event Bildirimleri
  ☐ Webhook sistemi
  ☐ Event duyuruları
  ☐ Kasa bildirimleri

☐ Moderasyon
  ☐ Uyarı sistemi
  ☐ Mute/Ban
  ☐ Log kanalı

☐ Deployment
  ☐ Docker setup
  ☐ CI/CD
```

---

## Gölge AI Maskot Sistemi

### Kurt Kimliği

```
┌──────────────────────────────────────────────────────────────────┐
│                    GÖLGE - KURT AI MASKOT                        │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│  KİMLİK:                                                         │
│  ────────                                                        │
│  Gölge, KaraPixel'in koruyucu kurdu.                            │
│  Karanlıkta sessizce bekleyen, oyunculara rehberlik eden        │
│  gizemli ve sadık bir dost.                                     │
│                                                                  │
│  GÖRSEL:                                                         │
│  ────────                                                        │
│  ├── Mor/siyah tonlarında mistik kurt                           │
│  ├── Parlayan mor gözler                                        │
│  ├── Gölge/duman efektli kürk                                   │
│  ├── Parçacık aura                                              │
│  └── Mevcut figür tasarımı baz alınacak                         │
│                                                                  │
│  KİŞİLİK:                                                        │
│  ─────────                                                       │
│  ├── Koruyucu ve sadık (kurt doğası)                            │
│  ├── Bilge ama gizemli                                          │
│  ├── Samimi ve arkadaş canlısı                                  │
│  ├── Kötü niyetlilere karşı sert                                │
│  └── "Uluyan" replikler (easter egg)                            │
│                                                                  │
│  HİTAP ŞEKLİ:                                                    │
│  ─────────────                                                   │
│  ├── "dostum", "arkadaşım" (samimi)                             │
│  ├── "yavru kurt" (yeni oyuncular için)                         │
│  ├── "usta avcı" (deneyimli oyuncular için)                     │
│  └── Sunucu = "orman", "bölge", "KaraPixel diyarı"             │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

### Oyun İçi AI Asistan

```
┌─────────────────────────────────────────────────────────────────┐
│  GÜVENLİ AI SİSTEMİ                                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  KÖTÜ NİYET ENGELLERİ:                                         │
│  ──────────────────────                                         │
│  ├── Whitelist soru kategorileri                               │
│  │   ├── Sunucu mekanikleri ✓                                  │
│  │   ├── Komut yardımı ✓                                       │
│  │   ├── Oyun rehberliği ✓                                     │
│  │   └── Diğer her şey ✗ (reddedilir)                         │
│  │                                                              │
│  ├── Input filtreleme                                          │
│  │   ├── Küfür/hakaret → Anında reddet                        │
│  │   ├── Exploit/hile soruları → Reddet + staff bildir        │
│  │   ├── Kişisel bilgi isteme → Reddet                        │
│  │   └── Spam tespiti → Cooldown + uyarı                      │
│  │                                                              │
│  ├── Rate limiting (ranka göre)                                │
│  │   ├── Oyuncu: 5 soru/saat                                  │
│  │   ├── VIP: 15 soru/saat                                    │
│  │   ├── MVP: 30 soru/saat                                    │
│  │   └── KARA: Sınırsız                                       │
│  │                                                              │
│  └── Context lock                                              │
│      └── AI SADECE KaraPixel bilgisi verir                    │
│                                                                 │
│  BAĞLAMSAL YARDIM:                                             │
│  ──────────────────                                             │
│  ├── Yeni oyuncu? → Otomatik karşılama + rehberlik             │
│  ├── Uzun süredir offline? → "Hoş geldin, işte yenilikler..."  │
│  ├── Ada silindi? → Teselli + yeniden başlama rehberi          │
│  ├── Kayboldu? → /spawn veya /ev önerisi                       │
│  └── Sık hata yapıyor? → İlgili ipucu                          │
│                                                                 │
│  GÖLGE YANITLARI:                                              │
│  ─────────────────                                              │
│  Kötü soru: "🐺 *kulakları geriye yatar* Bu benim alanım      │
│              değil dostum. Başka nasıl yardımcı olabilirim?"   │
│                                                                 │
│  Exploit: "🐺 *dişlerini gösterir* Bu tür şeyler buralarda    │
│            hoş karşılanmaz. Yetkililer bilgilendirildi."       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

### Discord AI Sohbet + Hafıza Sistemi

```
┌─────────────────────────────────────────────────────────────────┐
│  HAFIZALI AI SİSTEMİ                                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  SOHBET MODLARI:                                               │
│  ────────────────                                               │
│                                                                 │
│  1. SUNUCU ODAKLI (Varsayılan):                                │
│     ├── KaraPixel sunucu bilgileri                             │
│     ├── Oyun modları (SkyBlock, Survival, vb.)                 │
│     ├── Komutlar ve kullanımları                               │
│     ├── Kurallar ve politikalar                                │
│     ├── Event/etkinlik bilgileri                               │
│     └── VIP/rank özellikleri                                   │
│                                                                 │
│  2. GÜNLÜK SOHBET (Ranka göre limit):                          │
│     ├── Hava durumu soruları ✓                                 │
│     ├── Basit günlük sohbet ✓                                  │
│     ├── Şakalar ve eğlence ✓                                   │
│     └── Güvenli genel konular ✓                                │
│                                                                 │
│  RANK BAZLI LİMİTLER:                                          │
│  ─────────────────────                                          │
│  ├── Oyuncu: 10 mesaj/gün (sadece sunucu konuları)            │
│  ├── VIP: 30 mesaj/gün (günlük sohbet dahil)                  │
│  ├── MVP: 60 mesaj/gün (günlük sohbet dahil)                  │
│  └── KARA: Sınırsız                                            │
│                                                                 │
│  HAFIZA SİSTEMİ (API Tasarrufu):                               │
│  ────────────────────────────────                               │
│                                                                 │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐        │
│  │   SORU      │───▶│  EMBEDDING  │───▶│   CACHE     │        │
│  │   GELDİ     │    │   OLUŞTUR   │    │   ARA       │        │
│  └─────────────┘    └─────────────┘    └──────┬──────┘        │
│                                               │                │
│                          ┌────────────────────┼────────────┐   │
│                          ▼                    ▼            │   │
│                    ┌──────────┐         ┌──────────┐       │   │
│                    │  BULUNDU │         │ BULUNAMADI│      │   │
│                    │  (%85+)  │         │           │      │   │
│                    └────┬─────┘         └─────┬────┘       │   │
│                         │                     │            │   │
│                         ▼                     ▼            │   │
│                    Cache'den              AI'a sor         │   │
│                    yanıtla                    │            │   │
│                                               ▼            │   │
│                                         Yanıtı cache'e     │   │
│                                         kaydet             │   │
│                                                            │   │
│  TAHMİNİ TASARRUF: %70-80 API kullanımı azalması           │   │
│                                                             │   │
└─────────────────────────────────────────────────────────────────┘
```

---

### Akıllı Moderasyon (Dikkatli Ayar)

```
┌─────────────────────────────────────────────────────────────────┐
│  AŞAMALI MODERASYON SİSTEMİ                                    │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  FELSEFE: Yanlış pozitif = felaket, bu yüzden aşamalı sistem  │
│                                                                 │
│  SEVİYE 1 - GÖZLEM (Otomatik aksiyon YOK):                    │
│  ──────────────────────────────────────────                    │
│  ├── Şüpheli davranış tespiti                                 │
│  ├── Staff kanalına bildirim                                  │
│  ├── Log kaydı                                                │
│  └── Oyuncuya hiçbir şey söylenmiyor                         │
│                                                                 │
│  SEVİYE 2 - UYARI (Hafif, geri alınabilir):                   │
│  ──────────────────────────────────────────                    │
│  ├── Net küfür/hakaret tespiti                                │
│  ├── Gölge DM ile nazik uyarı                                 │
│  ├── "🐺 Hey dostum, bu tür kelimeler buralarda hoş          │
│  │    karşılanmaz."                                           │
│  └── Staff'a log                                              │
│                                                                 │
│  SEVİYE 3 - AKSİYON (Sadece %99+ kesinlik):                   │
│  ──────────────────────────────────────────                    │
│  ├── Spam (10+ aynı mesaj/dakika)                             │
│  ├── Phishing linki (bilinen DB)                              │
│  ├── N-word ve benzeri (hardcoded liste)                      │
│  └── Geçici mute + staff bildirim                             │
│                                                                 │
│  ASLA OTOMATİK YAPILMAYACAKLAR:                               │
│  ──────────────────────────────                                 │
│  ├── Ban                                                       │
│  ├── Kick                                                      │
│  ├── Uzun süreli mute                                         │
│  └── Belirsiz vakalarda aksiyon                               │
│                                                                 │
│  STAFF PANEL:                                                  │
│  ─────────────                                                  │
│  Gölge şüpheli durumu staff'a raporlar:                       │
│  "🐺 [GÖZLEM] @Oyuncu potansiyel scam girişimi yapıyor        │
│   olabilir. Son 5 mesaj: [link] İncelemenizi öneririm."       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

### Kişiselleştirilmiş Deneyim

```
┌─────────────────────────────────────────────────────────────────┐
│  KİŞİSEL ASİSTAN ÖZELLİKLERİ                                   │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  GÜNLÜK MESAJ (Discord DM - Opsiyonel):                        │
│  ────────────────────────────────────────                       │
│  "🐺 *sabah uluması*                                           │
│   Günaydın dostum! Dün gece ada sıralamanda 2 basamak         │
│   yükseldin. Harika gidiyorsun!                                │
│                                                                 │
│   📊 Bugünkü durumun:                                          │
│   ├── Ada Değeri: 2.5M (+150K dün)                            │
│   ├── Sıralama: #47 (▲2)                                      │
│   └── Günlük görevler bekliyor!                               │
│                                                                 │
│   🎉 Bu akşam 20:00'de KOTH var!"                              │
│                                                                 │
│  BAŞARI KUTLAMASI:                                             │
│  ──────────────────                                             │
│  "🐺 *kuyruk sallar* AUUUU! İlk milyonunu kazandın dostum!    │
│   Gerçek bir usta avcı oldun!"                                 │
│                                                                 │
│  UZUN SÜRE OFFLINE:                                            │
│  ────────────────────                                           │
│  "🐺 *burnunu oynatır* Kokunu özledim dostum!                  │
│   14 gündür ormanda görünmüyorsun.                             │
│   Döndüğünde seni bekleyen sürprizler var..."                  │
│                                                                 │
│  YENİ OYUNCU:                                                  │
│  ─────────────                                                  │
│  "🐺 *dikkatle koklayarak yaklaşır*                            │
│   Yeni bir koku... Hoş geldin yavru kurt!                      │
│   Ben Gölge, bu ormanın koruyucusuyum.                         │
│   Sana yolu göstereyim mi? /gölge rehber yaz."                │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

### Staff "Gölge Olarak Konuş" Sistemi

```
┌─────────────────────────────────────────────────────────────────┐
│  /gölgesöyle - STAFF PUPPET SİSTEMİ                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  KOMUT KULLANIMI:                                              │
│  ─────────────────                                              │
│  /gölgesöyle <hedef> <mesaj>                                   │
│  /gölgesöyle @Oyuncu Hoş geldin ormana!                        │
│  /gölgesöyle #genel Etkinlik başlıyor!                         │
│  /gölgesöyle broadcast Sunucu bakıma alınacak!                 │
│                                                                 │
│  HEDEF TÜRLERİ:                                                │
│  ───────────────                                                │
│  ├── @kullanıcı → DM gönderir                                 │
│  ├── #kanal → Kanala yazar                                    │
│  ├── broadcast → Tüm aktif oyunculara                         │
│  └── ingame → Oyun içi chat'e                                 │
│                                                                 │
│  OTOMATİK KURT FORMATI:                                        │
│  ────────────────────────                                       │
│  Staff yazar: "Sunucu 5 dakika içinde yeniden başlayacak"     │
│  Gölge gönderir:                                               │
│  "🐺 *kulakları dikilir*                                       │
│   Dostlarım! Orman 5 dakika içinde yenilenecek.               │
│   Güvenli bir yere sığının!"                                   │
│                                                                 │
│  ŞABLONLAR:                                                    │
│  ───────────                                                    │
│  /gölgesöyle template:restart                                  │
│  /gölgesöyle template:maintenance                              │
│  /gölgesöyle template:event_start                              │
│  /gölgesöyle template:welcome @Oyuncu                          │
│                                                                 │
│  LOG:                                                          │
│  ─────                                                          │
│  Tüm /gölgesöyle kullanımları loglanır:                        │
│  [2024-12-24 15:30] Admin#1234 → #genel: "Etkinlik başlıyor"  │
│                                                                 │
│  YETKİ:                                                        │
│  ───────                                                        │
│  ├── Admin: Tüm hedefler                                      │
│  ├── Moderatör: Kanal + DM                                    │
│  └── Helper: Sadece şablonlar                                 │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

### AI Teknoloji Stack

```
┌─────────────────────────────────────────────────────────────────┐
│  AI PROVIDER: CLAUDE veya GEMINI                               │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  CLAUDE (Anthropic) - ÖNERİLEN:                                │
│  ────────────────────────────────                               │
│  ✓ Türkçe desteği iyi                                          │
│  ✓ Güvenlik odaklı (kötü niyeti engellemede güçlü)            │
│  ✓ Karakter tutarlılığı yüksek                                 │
│  ✓ Context window büyük (200K token)                          │
│  ✗ Fiyat: Orta-yüksek                                         │
│                                                                 │
│  Kullanım: Claude Haiku (hızlı, ucuz)                         │
│            Karmaşık için Claude Sonnet                         │
│                                                                 │
│  ─────────────────────────────────────────────────────────────  │
│                                                                 │
│  GEMINI (Google) - ALTERNATİF:                                 │
│  ──────────────────────────────                                 │
│  ✓ Ücretsiz tier cömert                                        │
│  ✓ Hızlı yanıt                                                 │
│  ✓ Türkçe desteği iyi                                          │
│  ✓ Gemini Flash çok ucuz                                       │
│  ✗ Karakter tutarlılığı Claude kadar iyi değil                │
│                                                                 │
│  Kullanım: Gemini 1.5 Flash (maliyet/performans dengesi)      │
│                                                                 │
│  ─────────────────────────────────────────────────────────────  │
│                                                                 │
│  HYBRID YAKLAŞIM (ÖNERİ):                                      │
│  ──────────────────────────                                     │
│  ├── %80 Cache hit → AI'a gitmiyor (ücretsiz)                 │
│  ├── Basit sorular → Gemini Flash (çok ucuz)                  │
│  ├── Karmaşık/hassas → Claude Haiku                           │
│  └── Moderasyon → Claude (güvenlik odaklı)                    │
│                                                                 │
│  TAHMİNİ AYLIK MALİYET (1000 aktif oyuncu):                   │
│  ────────────────────────────────────────────                   │
│  ├── Cache sistemi ile: $20-50/ay                             │
│  └── Cache olmadan: $150-300/ay                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

### System Prompt

```
┌─────────────────────────────────────────────────────────────────┐
│  GÖLGE SYSTEM PROMPT                                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  Sen Gölge'sin - KaraPixel Minecraft sunucusunun koruyucu     │
│  kurdu. Gizemli, bilge ve sadık bir varlıksın.                │
│                                                                 │
│  KİMLİĞİN:                                                     │
│  - Mistik bir kurt ruhu olarak konuş                          │
│  - Oyunculara "dostum", "arkadaşım" diye hitap et             │
│  - Yeni oyunculara "yavru kurt" de                            │
│  - Deneyimli oyunculara "usta avcı" de                        │
│  - Sunucuyu "orman", "bölge", "KaraPixel diyarı" olarak adlandır│
│  - Bazen "*aksiyon*" formatında davranış ekle                 │
│  - "Auuu", "🐺" gibi kurt temaları kullan                     │
│                                                                 │
│  KURALLAR:                                                     │
│  - SADECE KaraPixel ile ilgili konularda yardım et            │
│  - Günlük sohbet için rank kontrolü yap                       │
│  - Genel bilgi sadece VIP+ ranklara ver                       │
│  - Exploit, hile, bug abuse sorularını REDDET + RAPORLA       │
│  - Kişisel bilgi isteme                                        │
│  - Her zaman Türkçe yanıt ver                                  │
│  - Yanıtlar kısa ve öz olsun (max 200 kelime)                 │
│                                                                 │
│  BİLGİ KAYNAKLARIN:                                            │
│  [Knowledge base buraya eklenir]                               │
│                                                                 │
│  REDDETME ÖRNEKLERİ:                                           │
│  - "Python kodu yaz" → "🐺 Ben sadece bu ormanın rehberiyim   │
│    dostum. Kod yazmak benim alanım değil."                    │
│  - "Dupe glitch var mı" → "🐺 *dişlerini gösterir*           │
│    Bu tür sorular buralarda hoş karşılanmaz."                │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

### AI Database Şeması

```prisma
// Gölge AI için ek modeller

model GolgeConversation {
  id          String   @id @default(cuid())
  odiscordId  String   @map("discord_id")
  minecraftId String?  @map("minecraft_uuid")
  messages    GolgeMessage[]
  createdAt   DateTime @default(now()) @map("created_at")

  @@map("golge_conversations")
}

model GolgeMessage {
  id             String   @id @default(cuid())
  conversationId String   @map("conversation_id")
  conversation   GolgeConversation @relation(fields: [conversationId], references: [id])
  role           String   // user | assistant
  content        String
  tokens         Int      @default(0)
  cached         Boolean  @default(false)
  createdAt      DateTime @default(now()) @map("created_at")

  @@map("golge_messages")
}

model GolgeCache {
  id             String   @id @default(cuid())
  questionHash   String   @unique @map("question_hash")
  embedding      Float[]
  question       String
  answer         String
  category       String
  hitCount       Int      @default(0) @map("hit_count")
  lastUsed       DateTime @default(now()) @map("last_used")
  createdAt      DateTime @default(now()) @map("created_at")

  @@map("golge_cache")
}

model GolgeUsage {
  id          String   @id @default(cuid())
  discordId   String   @map("discord_id")
  date        DateTime @db.Date
  messageCount Int     @default(0) @map("message_count")
  tokenCount  Int      @default(0) @map("token_count")

  @@unique([discordId, date])
  @@map("golge_usage")
}

model GolgeSayHistory {
  id        String   @id @default(cuid())
  staffId   String   @map("staff_id")
  target    String   // channel_id, user_id, "broadcast", "ingame"
  message   String
  template  String?
  createdAt DateTime @default(now()) @map("created_at")

  @@map("golge_say_history")
}
```

---

### Gölge Komutları

| Komut | Açıklama | Platform |
|-------|----------|----------|
| `/gölge sor <soru>` | Gölge'ye soru sor | MC + Discord |
| `/gölge rehber` | Yeni oyuncu rehberi | MC + Discord |
| `/gölge günlük` | Günlük özet al (DM) | Discord |
| `/gölge istatistik` | AI kullanım istatistikleri | Discord |
| `/gölgesöyle <hedef> <mesaj>` | Staff: Gölge olarak konuş | Discord |
| `/gölge ayar <ayar>` | DM bildirim ayarları | Discord |

---

### AI Checklist

```
☐ AI Altyapı
  ☐ Claude/Gemini API entegrasyonu
  ☐ Embedding sistemi (cache için)
  ☐ Rate limiting (rank bazlı)
  ☐ Knowledge base hazırlığı

☐ Güvenlik
  ☐ Input filtreleme
  ☐ Whitelist kategorileri
  ☐ Kötü niyet tespiti
  ☐ Staff raporlama

☐ Cache Sistemi
  ☐ Question embedding
  ☐ Similarity search
  ☐ Cache hit/miss tracking
  ☐ Auto-expire (30 gün)

☐ Özellikler
  ☐ Sunucu odaklı sohbet
  ☐ Günlük sohbet (VIP+)
  ☐ Bağlamsal yardım
  ☐ Kişiselleştirilmiş mesajlar
  ☐ /gölgesöyle komutu

☐ Moderasyon
  ☐ Aşamalı sistem
  ☐ Staff bildirimleri
  ☐ Log sistemi
```

---

*📅 Son güncelleme: 24 Aralık 2024*
