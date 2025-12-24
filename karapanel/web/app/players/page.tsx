'use client';

import { useState, useEffect } from 'react';
import { Header } from '@/components/Header';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { api, Player } from '@/lib/api';
import Link from 'next/link';

export default function PlayersPage() {
  const [players, setPlayers] = useState<Player[]>([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');
  const [onlineOnly, setOnlineOnly] = useState(false);
  const [page, setPage] = useState(0);
  const limit = 20;

  const fetchPlayers = async () => {
    try {
      setLoading(true);
      const data = await api.getPlayers({
        limit,
        offset: page * limit,
        search: search || undefined,
        online: onlineOnly,
      });
      setPlayers(data.players || []);
      setTotal(data.total);
    } catch (error) {
      console.error('Failed to fetch players:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPlayers();
  }, [page, onlineOnly]);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setPage(0);
    fetchPlayers();
  };

  const formatPlaytime = (seconds: number) => {
    const hours = Math.floor(seconds / 3600);
    const days = Math.floor(hours / 24);
    if (days > 0) return `${days}d ${hours % 24}h`;
    return `${hours}h`;
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

  return (
    <div className="min-h-screen bg-obsidian-950">
      <Header />
      <main className="container mx-auto px-4 py-8">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold text-white">Oyuncu Yonetimi</h1>
          <div className="flex gap-4">
            <form onSubmit={handleSearch} className="flex gap-2">
              <input
                type="text"
                placeholder="Oyuncu ara..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                className="px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-indigo-500"
              />
              <Button type="submit">Ara</Button>
            </form>
            <Button
              variant={onlineOnly ? 'default' : 'outline'}
              onClick={() => setOnlineOnly(!onlineOnly)}
            >
              {onlineOnly ? 'Online' : 'Tumu'}
            </Button>
          </div>
        </div>

        <Card className="bg-obsidian-900 border-obsidian-800">
          <CardHeader>
            <CardTitle className="text-white flex justify-between">
              <span>Oyuncular ({total})</span>
            </CardTitle>
          </CardHeader>
          <CardContent>
            {loading ? (
              <div className="text-center py-8 text-gray-400">Yukleniyor...</div>
            ) : players.length === 0 ? (
              <div className="text-center py-8 text-gray-400">Oyuncu bulunamadi</div>
            ) : (
              <div className="overflow-x-auto">
                <table className="w-full">
                  <thead>
                    <tr className="border-b border-obsidian-700">
                      <th className="text-left py-3 px-4 text-gray-400 font-medium">Oyuncu</th>
                      <th className="text-left py-3 px-4 text-gray-400 font-medium">Durum</th>
                      <th className="text-left py-3 px-4 text-gray-400 font-medium">Oynama Suresi</th>
                      <th className="text-left py-3 px-4 text-gray-400 font-medium">Son Giris</th>
                      <th className="text-left py-3 px-4 text-gray-400 font-medium">Ilk Giris</th>
                      <th className="text-right py-3 px-4 text-gray-400 font-medium">Islemler</th>
                    </tr>
                  </thead>
                  <tbody>
                    {players.map((player) => (
                      <tr key={player.uuid} className="border-b border-obsidian-800 hover:bg-obsidian-800/50">
                        <td className="py-3 px-4">
                          <div className="flex items-center gap-3">
                            <img
                              src={`https://mc-heads.net/avatar/${player.uuid}/32`}
                              alt={player.username}
                              className="w-8 h-8 rounded"
                            />
                            <span className="text-white font-medium">{player.username}</span>
                          </div>
                        </td>
                        <td className="py-3 px-4">
                          <span className={`px-2 py-1 rounded text-xs ${
                            player.isOnline
                              ? 'bg-green-500/20 text-green-400'
                              : 'bg-gray-500/20 text-gray-400'
                          }`}>
                            {player.isOnline ? 'Online' : 'Offline'}
                          </span>
                        </td>
                        <td className="py-3 px-4 text-gray-300">
                          {formatPlaytime(player.totalPlaytime)}
                        </td>
                        <td className="py-3 px-4 text-gray-300">
                          {formatDate(player.lastSeen)}
                        </td>
                        <td className="py-3 px-4 text-gray-300">
                          {formatDate(player.firstJoin)}
                        </td>
                        <td className="py-3 px-4 text-right">
                          <Link href={`/players/${player.uuid}`}>
                            <Button variant="outline" size="sm">Detay</Button>
                          </Link>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}

            {/* Pagination */}
            {total > limit && (
              <div className="flex justify-center gap-2 mt-4">
                <Button
                  variant="outline"
                  disabled={page === 0}
                  onClick={() => setPage(p => p - 1)}
                >
                  Onceki
                </Button>
                <span className="px-4 py-2 text-gray-400">
                  Sayfa {page + 1} / {Math.ceil(total / limit)}
                </span>
                <Button
                  variant="outline"
                  disabled={(page + 1) * limit >= total}
                  onClick={() => setPage(p => p + 1)}
                >
                  Sonraki
                </Button>
              </div>
            )}
          </CardContent>
        </Card>
      </main>
    </div>
  );
}
