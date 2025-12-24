'use client';

import { useState, useEffect } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { Header } from '@/components/Header';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { api, PlayerWithStats, Punishment } from '@/lib/api';

export default function PlayerDetailPage() {
  const params = useParams();
  const router = useRouter();
  const uuid = params.uuid as string;

  const [player, setPlayer] = useState<PlayerWithStats | null>(null);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<'info' | 'stats' | 'punishments' | 'ips'>('info');

  useEffect(() => {
    const fetchPlayer = async () => {
      try {
        const data = await api.getPlayer(uuid);
        setPlayer(data);
      } catch (error) {
        console.error('Failed to fetch player:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchPlayer();
  }, [uuid]);

  const formatPlaytime = (seconds: number) => {
    const hours = Math.floor(seconds / 3600);
    const days = Math.floor(hours / 24);
    const remainingHours = hours % 24;
    const minutes = Math.floor((seconds % 3600) / 60);
    if (days > 0) return `${days} gun ${remainingHours} saat`;
    if (hours > 0) return `${hours} saat ${minutes} dakika`;
    return `${minutes} dakika`;
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

  const formatNumber = (num: number) => {
    return num.toLocaleString('tr-TR');
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-obsidian-950">
        <Header />
        <main className="container mx-auto px-4 py-8">
          <div className="text-center text-gray-400">Yukleniyor...</div>
        </main>
      </div>
    );
  }

  if (!player) {
    return (
      <div className="min-h-screen bg-obsidian-950">
        <Header />
        <main className="container mx-auto px-4 py-8">
          <div className="text-center text-gray-400">Oyuncu bulunamadi</div>
        </main>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-obsidian-950">
      <Header />
      <main className="container mx-auto px-4 py-8">
        <Button variant="outline" onClick={() => router.back()} className="mb-4">
          Geri
        </Button>

        {/* Player Header */}
        <Card className="bg-obsidian-900 border-obsidian-800 mb-6">
          <CardContent className="pt-6">
            <div className="flex items-center gap-6">
              <img
                src={`https://mc-heads.net/body/${player.uuid}/100`}
                alt={player.username}
                className="w-24 h-auto"
              />
              <div className="flex-1">
                <h1 className="text-3xl font-bold text-white mb-2">{player.username}</h1>
                <div className="flex items-center gap-4 text-gray-400">
                  <span className={`px-3 py-1 rounded text-sm ${
                    player.isOnline
                      ? 'bg-green-500/20 text-green-400'
                      : 'bg-gray-500/20 text-gray-400'
                  }`}>
                    {player.isOnline ? 'Online' : 'Offline'}
                  </span>
                  <span>UUID: {player.uuid}</span>
                </div>
              </div>
              <div className="flex gap-2">
                <Button variant="outline" className="text-yellow-400 border-yellow-400/50">
                  Uyar
                </Button>
                <Button variant="outline" className="text-orange-400 border-orange-400/50">
                  Sustur
                </Button>
                <Button variant="outline" className="text-red-400 border-red-400/50">
                  Yasakla
                </Button>
              </div>
            </div>
          </CardContent>
        </Card>

        {/* Tabs */}
        <div className="flex gap-2 mb-6">
          {(['info', 'stats', 'punishments', 'ips'] as const).map((tab) => (
            <Button
              key={tab}
              variant={activeTab === tab ? 'default' : 'outline'}
              onClick={() => setActiveTab(tab)}
            >
              {tab === 'info' && 'Bilgiler'}
              {tab === 'stats' && 'Istatistikler'}
              {tab === 'punishments' && 'Cezalar'}
              {tab === 'ips' && 'IP Gecmisi'}
            </Button>
          ))}
        </div>

        {/* Tab Content */}
        {activeTab === 'info' && (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Toplam Oynama Suresi</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-white">{formatPlaytime(player.totalPlaytime)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Ilk Giris</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-xl font-bold text-white">{formatDate(player.firstJoin)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Son Giris</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-xl font-bold text-white">{formatDate(player.lastSeen)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Son IP</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-xl font-bold text-white">{player.lastIP || 'Bilinmiyor'}</p>
              </CardContent>
            </Card>
          </div>
        )}

        {activeTab === 'stats' && player.stats && (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Oldurmeler</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-green-400">{formatNumber(player.stats.kills)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Olumler</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-red-400">{formatNumber(player.stats.deaths)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">K/D Orani</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-indigo-400">
                  {player.stats.deaths > 0
                    ? (player.stats.kills / player.stats.deaths).toFixed(2)
                    : player.stats.kills}
                </p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Mob Oldurmeler</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-white">{formatNumber(player.stats.mobKills)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Kirilan Blok</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-white">{formatNumber(player.stats.blocksBroken)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Yerlestirilen Blok</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-white">{formatNumber(player.stats.blocksPlaced)}</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Yurunen Mesafe</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-white">{formatNumber(Math.floor(player.stats.distanceWalked / 1000))} km</p>
              </CardContent>
            </Card>
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader className="pb-2">
                <CardTitle className="text-sm text-gray-400">Oynama Suresi</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-2xl font-bold text-white">{formatPlaytime(player.stats.timePlayed)}</p>
              </CardContent>
            </Card>
          </div>
        )}

        {activeTab === 'punishments' && (
          <Card className="bg-obsidian-900 border-obsidian-800">
            <CardContent className="pt-6">
              {player.punishments && player.punishments.length > 0 ? (
                <table className="w-full">
                  <thead>
                    <tr className="border-b border-obsidian-700">
                      <th className="text-left py-3 px-4 text-gray-400">Tur</th>
                      <th className="text-left py-3 px-4 text-gray-400">Sebep</th>
                      <th className="text-left py-3 px-4 text-gray-400">Yetkili</th>
                      <th className="text-left py-3 px-4 text-gray-400">Tarih</th>
                      <th className="text-left py-3 px-4 text-gray-400">Durum</th>
                    </tr>
                  </thead>
                  <tbody>
                    {player.punishments.map((p) => (
                      <tr key={p.id} className="border-b border-obsidian-800">
                        <td className="py-3 px-4">
                          <span className={`px-2 py-1 rounded text-xs ${
                            p.type === 'ban' ? 'bg-red-500/20 text-red-400' :
                            p.type === 'mute' ? 'bg-orange-500/20 text-orange-400' :
                            p.type === 'warn' ? 'bg-yellow-500/20 text-yellow-400' :
                            'bg-gray-500/20 text-gray-400'
                          }`}>
                            {p.type.toUpperCase()}
                          </span>
                        </td>
                        <td className="py-3 px-4 text-white">{p.reason}</td>
                        <td className="py-3 px-4 text-gray-300">{p.issuedBy || 'Sistem'}</td>
                        <td className="py-3 px-4 text-gray-300">{formatDate(p.issuedAt)}</td>
                        <td className="py-3 px-4">
                          <span className={`px-2 py-1 rounded text-xs ${
                            p.isActive ? 'bg-red-500/20 text-red-400' : 'bg-green-500/20 text-green-400'
                          }`}>
                            {p.isActive ? 'Aktif' : 'Bitmis'}
                          </span>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              ) : (
                <div className="text-center py-8 text-gray-400">Ceza kaydı bulunamadi</div>
              )}
            </CardContent>
          </Card>
        )}

        {activeTab === 'ips' && (
          <Card className="bg-obsidian-900 border-obsidian-800">
            <CardContent className="pt-6">
              {player.ips && player.ips.length > 0 ? (
                <table className="w-full">
                  <thead>
                    <tr className="border-b border-obsidian-700">
                      <th className="text-left py-3 px-4 text-gray-400">IP Adresi</th>
                      <th className="text-left py-3 px-4 text-gray-400">Ilk Kullanim</th>
                      <th className="text-left py-3 px-4 text-gray-400">Son Kullanim</th>
                      <th className="text-left py-3 px-4 text-gray-400">Kullanim Sayisi</th>
                    </tr>
                  </thead>
                  <tbody>
                    {player.ips.map((ip, idx) => (
                      <tr key={idx} className="border-b border-obsidian-800">
                        <td className="py-3 px-4 text-white font-mono">{ip.ip}</td>
                        <td className="py-3 px-4 text-gray-300">{formatDate(ip.firstUsed)}</td>
                        <td className="py-3 px-4 text-gray-300">{formatDate(ip.lastUsed)}</td>
                        <td className="py-3 px-4 text-gray-300">{ip.usageCount}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              ) : (
                <div className="text-center py-8 text-gray-400">IP gecmisi bulunamadi</div>
              )}
            </CardContent>
          </Card>
        )}
      </main>
    </div>
  );
}
