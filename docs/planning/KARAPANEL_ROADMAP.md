# KaraPanel Geliştirme Planı

## Mevcut Durum (v0.1.0)
- [x] Go Fiber daemon backend
- [x] Next.js frontend (shadcn/ui)
- [x] JWT authentication
- [x] Dashboard (sunucu listesi, sistem metrikleri)
- [x] Sunucu start/stop/restart
- [x] Dosya yöneticisi (temel)
- [x] WebSocket konsol (temel)
- [x] SSL/HTTPS (Let's Encrypt)
- [x] Route protection middleware

---

## Faz 1: Temel İyileştirmeler (Öncelikli)

### 1.1 Konsol Geliştirmeleri
- [ ] Komut gönderme özelliği (input field)
- [ ] Konsol geçmişi (scrollback)
- [ ] ANSI renk desteği iyileştirme
- [ ] Konsol temizleme butonu
- [ ] Log filtreleme (ERROR, WARN, INFO)
- [ ] Log indirme (son X satır)

### 1.2 Dosya Yöneticisi Geliştirmeleri
- [ ] Drag & drop dosya yükleme
- [ ] Çoklu dosya seçimi ve silme
- [ ] Dosya/klasör kopyalama, taşıma
- [ ] Yeni klasör oluşturma
- [ ] Dosya arama
- [ ] Kod editörü syntax highlighting (Monaco Editor)
- [ ] Dosya izinleri görüntüleme

### 1.3 Dashboard Geliştirmeleri
- [ ] Gerçek zamanlı metrik güncelleme (WebSocket)
- [ ] Sunucu TPS gösterimi
- [ ] Son 24 saat grafikleri (CPU, RAM, Players)
- [ ] Hızlı eylemler (tüm sunucuları başlat/durdur)
- [ ] Sunucu sağlık durumu (health check)

---

## Faz 2: Sunucu Yönetimi

### 2.1 Sunucu Konfigürasyonu
- [ ] server.properties editörü (GUI)
- [ ] spigot.yml / paper.yml editörü
- [ ] velocity.toml editörü
- [ ] JVM argümanları ayarlama
- [ ] Port yönetimi

### 2.2 Plugin Yönetimi
- [ ] Yüklü plugin listesi
- [ ] Plugin aktif/pasif yapma
- [ ] Plugin silme
- [ ] Plugin güncelleme kontrolü
- [ ] Modrinth/Hangar entegrasyonu (plugin arama & indirme)

### 2.3 Dünya Yönetimi
- [ ] Dünya listesi
- [ ] Dünya yedekleme
- [ ] Dünya silme
- [ ] Dünya indirme (zip)
- [ ] Dünya yükleme

---

## Faz 3: Yedekleme Sistemi

### 3.1 Otomatik Yedekleme
- [ ] Zamanlanmış yedeklemeler (cron)
- [ ] Yedekleme rotasyonu (son X yedek)
- [ ] Sıkıştırma seçenekleri (tar.gz, zip)
- [ ] Yedekleme öncesi sunucu durdurma opsiyonu

### 3.2 Manuel Yedekleme
- [ ] Tek tıkla yedekleme
- [ ] Seçici yedekleme (sadece dünyalar, config, vb.)
- [ ] Yedekleme listesi ve geri yükleme
- [ ] Yedekleme indirme

### 3.3 Uzak Yedekleme
- [ ] S3/MinIO entegrasyonu
- [ ] SFTP yedekleme
- [ ] Google Drive / Dropbox (opsiyonel)

---

## Faz 4: Oyuncu Yönetimi

### 4.1 Oyuncu Listesi
- [ ] Online oyuncu listesi (gerçek zamanlı)
- [ ] Oyuncu bilgileri (UUID, IP, oyun süresi)
- [ ] Kick / Ban / Mute işlemleri
- [ ] Oyuncu arama

### 4.2 Whitelist & Banlist
- [ ] Whitelist yönetimi (ekle/çıkar)
- [ ] Banlist yönetimi
- [ ] IP ban listesi
- [ ] Ban geçmişi

### 4.3 Oyuncu İstatistikleri
- [ ] Oyun süresi
- [ ] Giriş/çıkış geçmişi
- [ ] Oyuncu sayısı grafikleri

---

## Faz 5: Güvenlik & Kullanıcı Yönetimi

### 5.1 Çoklu Kullanıcı
- [ ] Kullanıcı ekleme/silme (panel üzerinden)
- [ ] Rol tabanlı yetkilendirme (RBAC)
  - Admin: Tam yetki
  - Moderator: Sunucu kontrol, oyuncu yönetimi
  - Viewer: Sadece görüntüleme
- [ ] Kullanıcı bazlı sunucu erişimi

### 5.2 Güvenlik
- [ ] 2FA (TOTP) desteği
- [ ] Oturum yönetimi (aktif oturumlar)
- [ ] IP whitelist (panel erişimi)
- [ ] Rate limiting
- [ ] Audit log (kim ne yaptı)
- [ ] Şifre politikası

### 5.3 API Güvenliği
- [ ] API key sistemi
- [ ] Webhook desteği (sunucu olayları)
- [ ] API rate limiting

---

## Faz 6: Bildirimler & Entegrasyonlar

### 6.1 Bildirim Sistemi
- [ ] Discord webhook entegrasyonu
  - Sunucu durumu değişikliği
  - Oyuncu giriş/çıkış
  - Yüksek CPU/RAM uyarıları
  - Yedekleme tamamlandı
- [ ] Email bildirimleri
- [ ] Panel içi bildirimler

### 6.2 Entegrasyonlar
- [ ] Discord bot entegrasyonu
- [ ] Pterodactyl import/export
- [ ] RCON protokolü desteği

---

## Faz 7: Gelişmiş Özellikler

### 7.1 Çoklu Sunucu (Multi-Node)
- [ ] Uzak sunucu ekleme (SSH/Agent)
- [ ] Merkezi yönetim paneli
- [ ] Sunucu grupları

### 7.2 Otomatik Görevler
- [ ] Zamanlanmış komutlar
- [ ] Otomatik restart (belirli saatlerde)
- [ ] Crash algılama ve otomatik restart
- [ ] Düşük TPS uyarısı ve restart

### 7.3 Performans İzleme
- [ ] Detaylı performans grafikleri
- [ ] Timings analizi
- [ ] Spark profiler entegrasyonu
- [ ] Garbage collection izleme

### 7.4 Şablon Sistemi
- [ ] Sunucu şablonları (Paper, Velocity, vb.)
- [ ] Tek tıkla yeni sunucu oluşturma
- [ ] Şablon paylaşımı

---

## Faz 8: UI/UX İyileştirmeleri

### 8.1 Tema & Özelleştirme
- [ ] Light/Dark tema geçişi
- [ ] Özel renk şemaları
- [ ] Dashboard widget düzenleme
- [ ] Dil desteği (TR/EN)

### 8.2 Mobil Uyumluluk
- [ ] Responsive tasarım iyileştirme
- [ ] PWA desteği (Progressive Web App)
- [ ] Mobil bildirimler

### 8.3 Erişilebilirlik
- [ ] Klavye navigasyonu
- [ ] Screen reader uyumluluğu
- [ ] Yüksek kontrast modu

---

## Teknik Borç & İyileştirmeler

### Daemon (Go)
- [ ] Unit testler
- [ ] Graceful shutdown iyileştirme
- [ ] Metrics caching (Redis opsiyonel)
- [ ] Structured logging (zerolog)
- [ ] Config hot-reload

### Frontend (Next.js)
- [ ] React Query ile data fetching
- [ ] Zustand/Jotai state management
- [ ] Error boundary iyileştirme
- [ ] Loading states & skeletons
- [ ] Unit & E2E testler (Vitest, Playwright)

### DevOps
- [ ] Docker compose setup
- [ ] GitHub Actions CI/CD
- [ ] Otomatik deployment
- [ ] Monitoring (Prometheus + Grafana)

---

## Öncelik Sıralaması

| Öncelik | Özellik | Tahmini Zorluk |
|---------|---------|----------------|
| 1 | Konsol komut gönderme | Kolay |
| 2 | Gerçek zamanlı metrikler | Orta |
| 3 | Plugin yönetimi | Orta |
| 4 | Otomatik yedekleme | Orta |
| 5 | Discord webhook | Kolay |
| 6 | Çoklu kullanıcı & RBAC | Zor |
| 7 | 2FA desteği | Orta |
| 8 | Oyuncu yönetimi | Orta |
| 9 | Zamanlanmış görevler | Orta |
| 10 | Multi-node desteği | Zor |

---

## Notlar

- Her faz bağımsız olarak geliştirilebilir
- Kullanıcı geri bildirimlerine göre öncelikler değişebilir
- Güvenlik her zaman öncelikli olmalı
- Her özellik için test yazılmalı
