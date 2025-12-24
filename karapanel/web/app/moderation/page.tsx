'use client';

import { useState, useEffect } from 'react';
import { Header } from '@/components/Header';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { api, Punishment, PunishmentStats } from '@/lib/api';

export default function ModerationPage() {
  const [punishments, setPunishments] = useState<Punishment[]>([]);
  const [stats, setStats] = useState<PunishmentStats | null>(null);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(true);
  const [filter, setFilter] = useState({ type: '', status: '' });
  const [page, setPage] = useState(0);
  const [showModal, setShowModal] = useState(false);
  const limit = 20;

  const fetchPunishments = async () => {
    try {
      setLoading(true);
      const active = filter.status === 'active' ? true : filter.status === 'expired' ? false : undefined;
      const data = await api.getPunishments({
        limit,
        offset: page * limit,
        type: filter.type || undefined,
        active,
      });
      setPunishments(data.punishments || []);
      setTotal(data.total);
    } catch (error) {
      console.error('Failed to fetch punishments:', error);
    } finally {
      setLoading(false);
    }
  };

  const fetchStats = async () => {
    try {
      const data = await api.getPunishmentStats();
      setStats(data);
    } catch (error) {
      console.error('Failed to fetch stats:', error);
    }
  };

  useEffect(() => {
    fetchPunishments();
    fetchStats();
  }, [page, filter]);

  const handleRevoke = async (id: number) => {
    if (!confirm('Bu cezayi kaldirmak istediginize emin misiniz?')) return;
    try {
      await api.revokePunishment(id, 'Admin');
      fetchPunishments();
      fetchStats();
    } catch (error) {
      console.error('Failed to revoke punishment:', error);
    }
  };

  const formatDate = (dateStr: string) => {
    return new Date(dateStr).toLocaleDateString('tr-TR', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const formatDuration = (seconds?: number) => {
    if (!seconds) return 'Kalici';
    const days = Math.floor(seconds / 86400);
    const hours = Math.floor((seconds % 86400) / 3600);
    if (days > 0) return `${days} gun`;
    return `${hours} saat`;
  };

  const getTypeColor = (type: string) => {
    switch (type) {
      case 'ban': return 'bg-red-500/20 text-red-400';
      case 'mute': return 'bg-orange-500/20 text-orange-400';
      case 'warn': return 'bg-yellow-500/20 text-yellow-400';
      case 'kick': return 'bg-blue-500/20 text-blue-400';
      default: return 'bg-gray-500/20 text-gray-400';
    }
  };

  return (
    <div className="min-h-screen bg-obsidian-950">
      <Header />
      <main className="container mx-auto px-4 py-8">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold text-white">Moderasyon</h1>
          <Button onClick={() => setShowModal(true)}>Yeni Ceza</Button>
        </div>

        {/* Stats */}
        {stats && (
          <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4 mb-6">
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardContent className="pt-4 pb-4">
                <p className="text-sm text-gray-400">Toplam Ban</p>
                <p className="text-2xl font-bold text-red-400">{stats.totalBans}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardContent className="pt-4 pb-4">
                <p className="text-sm text-gray-400">Aktif Ban</p>
                <p className="text-2xl font-bold text-red-400">{stats.activeBans}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardContent className="pt-4 pb-4">
                <p className="text-sm text-gray-400">Toplam Mute</p>
                <p className="text-2xl font-bold text-orange-400">{stats.totalMutes}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardContent className="pt-4 pb-4">
                <p className="text-sm text-gray-400">Aktif Mute</p>
                <p className="text-2xl font-bold text-orange-400">{stats.activeMutes}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardContent className="pt-4 pb-4">
                <p className="text-sm text-gray-400">Toplam Warn</p>
                <p className="text-2xl font-bold text-yellow-400">{stats.totalWarns}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardContent className="pt-4 pb-4">
                <p className="text-sm text-gray-400">Toplam Kick</p>
                <p className="text-2xl font-bold text-blue-400">{stats.totalKicks}</p>
              </CardContent>
            </Card>
          </div>
        )}

        {/* Filters */}
        <div className="flex gap-4 mb-6">
          <select
            value={filter.type}
            onChange={(e) => setFilter(f => ({ ...f, type: e.target.value }))}
            className="px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
          >
            <option value="">Tum Turler</option>
            <option value="ban">Ban</option>
            <option value="mute">Mute</option>
            <option value="warn">Warn</option>
            <option value="kick">Kick</option>
          </select>
          <select
            value={filter.status}
            onChange={(e) => setFilter(f => ({ ...f, status: e.target.value }))}
            className="px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
          >
            <option value="">Tum Durumlar</option>
            <option value="active">Aktif</option>
            <option value="expired">Bitmis</option>
            <option value="appealed">Itiraz Edilmis</option>
          </select>
        </div>

        {/* Punishments Table */}
        <Card className="bg-obsidian-900 border-obsidian-800">
          <CardHeader>
            <CardTitle className="text-white">Ceza Kayitlari ({total})</CardTitle>
          </CardHeader>
          <CardContent>
            {loading ? (
              <div className="text-center py-8 text-gray-400">Yukleniyor...</div>
            ) : punishments.length === 0 ? (
              <div className="text-center py-8 text-gray-400">Ceza kaydi bulunamadi</div>
            ) : (
              <div className="overflow-x-auto">
                <table className="w-full">
                  <thead>
                    <tr className="border-b border-obsidian-700">
                      <th className="text-left py-3 px-4 text-gray-400">Oyuncu</th>
                      <th className="text-left py-3 px-4 text-gray-400">Tur</th>
                      <th className="text-left py-3 px-4 text-gray-400">Sebep</th>
                      <th className="text-left py-3 px-4 text-gray-400">Yetkili</th>
                      <th className="text-left py-3 px-4 text-gray-400">Sure</th>
                      <th className="text-left py-3 px-4 text-gray-400">Tarih</th>
                      <th className="text-left py-3 px-4 text-gray-400">Durum</th>
                      <th className="text-right py-3 px-4 text-gray-400">Islemler</th>
                    </tr>
                  </thead>
                  <tbody>
                    {punishments.map((p) => (
                      <tr key={p.id} className="border-b border-obsidian-800 hover:bg-obsidian-800/50">
                        <td className="py-3 px-4">
                          <div className="flex items-center gap-2">
                            <img
                              src={`https://mc-heads.net/avatar/${p.playerUUID}/24`}
                              alt=""
                              className="w-6 h-6 rounded"
                            />
                            <span className="text-white">{p.playerUsername || p.playerUUID.slice(0, 8)}</span>
                          </div>
                        </td>
                        <td className="py-3 px-4">
                          <span className={`px-2 py-1 rounded text-xs uppercase ${getTypeColor(p.type)}`}>
                            {p.type}
                          </span>
                        </td>
                        <td className="py-3 px-4 text-gray-300 max-w-xs truncate">{p.reason}</td>
                        <td className="py-3 px-4 text-gray-300">{p.issuedBy || 'Sistem'}</td>
                        <td className="py-3 px-4 text-gray-300">{formatDuration(p.duration)}</td>
                        <td className="py-3 px-4 text-gray-300">{formatDate(p.issuedAt)}</td>
                        <td className="py-3 px-4">
                          <span className={`px-2 py-1 rounded text-xs ${
                            p.isActive ? 'bg-red-500/20 text-red-400' : 'bg-green-500/20 text-green-400'
                          }`}>
                            {p.isActive ? 'Aktif' : 'Bitmis'}
                          </span>
                        </td>
                        <td className="py-3 px-4 text-right">
                          {p.isActive && (
                            <Button
                              variant="outline"
                              size="sm"
                              onClick={() => handleRevoke(p.id)}
                              className="text-red-400"
                            >
                              Kaldir
                            </Button>
                          )}
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}

            {total > limit && (
              <div className="flex justify-center gap-2 mt-4">
                <Button variant="outline" disabled={page === 0} onClick={() => setPage(p => p - 1)}>
                  Onceki
                </Button>
                <span className="px-4 py-2 text-gray-400">
                  Sayfa {page + 1} / {Math.ceil(total / limit)}
                </span>
                <Button variant="outline" disabled={(page + 1) * limit >= total} onClick={() => setPage(p => p + 1)}>
                  Sonraki
                </Button>
              </div>
            )}
          </CardContent>
        </Card>

        {/* New Punishment Modal */}
        {showModal && (
          <PunishmentModal onClose={() => setShowModal(false)} onSuccess={() => { fetchPunishments(); fetchStats(); }} />
        )}
      </main>
    </div>
  );
}

function PunishmentModal({ onClose, onSuccess }: { onClose: () => void; onSuccess: () => void }) {
  const [form, setForm] = useState({
    playerName: '',
    playerUUID: '',
    type: 'ban' as 'ban' | 'kick' | 'mute' | 'warn',
    reason: '',
    duration: '',
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');

    let playerUUID = form.playerUUID;
    if (!playerUUID) {
      // Try to find player by name
      try {
        const player = await api.searchPlayer(form.playerName);
        playerUUID = player.uuid;
      } catch {
        setError('Oyuncu bulunamadi');
        return;
      }
    }

    try {
      setLoading(true);
      const durationSeconds = form.duration ? parseInt(form.duration) * 3600 : undefined;
      await api.createPunishment({
        playerUUID,
        type: form.type,
        reason: form.reason,
        duration: durationSeconds,
        isPermanent: !durationSeconds,
      });
      onSuccess();
      onClose();
    } catch (err: any) {
      setError(err.message || 'Ceza olusturulamadi');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <Card className="bg-obsidian-900 border-obsidian-800 w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-white">Yeni Ceza</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-sm text-gray-400 mb-1">Oyuncu Adi</label>
              <input
                type="text"
                value={form.playerName}
                onChange={(e) => setForm(f => ({ ...f, playerName: e.target.value }))}
                className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
                required
              />
            </div>
            <div>
              <label className="block text-sm text-gray-400 mb-1">Ceza Turu</label>
              <select
                value={form.type}
                onChange={(e) => setForm(f => ({ ...f, type: e.target.value as 'ban' | 'kick' | 'mute' | 'warn' }))}
                className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
              >
                <option value="ban">Ban</option>
                <option value="mute">Mute</option>
                <option value="warn">Warn</option>
                <option value="kick">Kick</option>
              </select>
            </div>
            <div>
              <label className="block text-sm text-gray-400 mb-1">Sebep</label>
              <textarea
                value={form.reason}
                onChange={(e) => setForm(f => ({ ...f, reason: e.target.value }))}
                className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
                rows={3}
                required
              />
            </div>
            {form.type !== 'kick' && form.type !== 'warn' && (
              <div>
                <label className="block text-sm text-gray-400 mb-1">Sure (saat, bos = kalici)</label>
                <input
                  type="number"
                  value={form.duration}
                  onChange={(e) => setForm(f => ({ ...f, duration: e.target.value }))}
                  className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
                  placeholder="Bos birakirsaniz kalici olur"
                />
              </div>
            )}
            {error && <p className="text-red-400 text-sm">{error}</p>}
            <div className="flex gap-2 justify-end">
              <Button type="button" variant="outline" onClick={onClose}>Iptal</Button>
              <Button type="submit" disabled={loading}>
                {loading ? 'Olusturuluyor...' : 'Olustur'}
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
