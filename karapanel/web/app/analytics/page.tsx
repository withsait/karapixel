'use client';

import { useState, useEffect } from 'react';
import { Header } from '@/components/Header';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { api, DashboardStats, ActivityLog, PlayerHistoryPoint } from '@/lib/api';

export default function AnalyticsPage() {
  const [stats, setStats] = useState<DashboardStats | null>(null);
  const [logs, setLogs] = useState<ActivityLog[]>([]);
  const [history, setHistory] = useState<PlayerHistoryPoint[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const [statsData, logsData, historyData] = await Promise.all([
          api.getDashboardStats(),
          api.getActivityLogs({ limit: 20 }),
          api.getPlayerHistory(24),
        ]);
        setStats(statsData);
        setLogs(logsData.logs || []);
        setHistory(historyData.history || []);
      } catch (error) {
        console.error('Failed to fetch analytics:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const formatDate = (dateStr: string) => {
    return new Date(dateStr).toLocaleDateString('tr-TR', {
      day: '2-digit',
      month: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  return (
    <div className="min-h-screen bg-obsidian-950">
      <Header />
      <main className="container mx-auto px-4 py-8">
        <h1 className="text-2xl font-bold text-white mb-6">Istatistikler & Analiz</h1>

        {loading ? (
          <div className="text-center py-8 text-gray-400">Yukleniyor...</div>
        ) : (
          <>
            {/* Stats Cards */}
            {stats && (
              <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4 mb-6">
                <Card className="bg-obsidian-900 border-obsidian-800">
                  <CardContent className="pt-4 pb-4">
                    <p className="text-sm text-gray-400">Toplam Oyuncu</p>
                    <p className="text-2xl font-bold text-white">{stats.totalPlayers}</p>
                  </CardContent>
                </Card>
                <Card className="bg-obsidian-900 border-obsidian-800">
                  <CardContent className="pt-4 pb-4">
                    <p className="text-sm text-gray-400">Online Oyuncu</p>
                    <p className="text-2xl font-bold text-green-400">{stats.onlinePlayers}</p>
                  </CardContent>
                </Card>
                <Card className="bg-obsidian-900 border-obsidian-800">
                  <CardContent className="pt-4 pb-4">
                    <p className="text-sm text-gray-400">Bugun Katilan</p>
                    <p className="text-2xl font-bold text-indigo-400">{stats.newPlayersToday}</p>
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
                    <p className="text-sm text-gray-400">Aktif Mute</p>
                    <p className="text-2xl font-bold text-orange-400">{stats.activeMutes}</p>
                  </CardContent>
                </Card>
                <Card className="bg-obsidian-900 border-obsidian-800">
                  <CardContent className="pt-4 pb-4">
                    <p className="text-sm text-gray-400">Discord Bagli</p>
                    <p className="text-2xl font-bold text-blue-400">{stats.linkedDiscordAccounts}</p>
                  </CardContent>
                </Card>
              </div>
            )}

            {/* Player History Chart */}
            <Card className="bg-obsidian-900 border-obsidian-800 mb-6">
              <CardHeader>
                <CardTitle className="text-white">Oyuncu Sayisi (Son 24 Saat)</CardTitle>
              </CardHeader>
              <CardContent>
                {history.length === 0 ? (
                  <div className="text-center py-8 text-gray-400">Veri bulunamadi</div>
                ) : (
                  <div className="h-64 flex items-end gap-1">
                    {history.map((point, idx) => {
                      const maxPeak = Math.max(...history.map(h => h.peakPlayers), 1);
                      const heightPercent = (point.peakPlayers / maxPeak) * 100;
                      return (
                        <div
                          key={idx}
                          className="flex-1 bg-indigo-500/50 hover:bg-indigo-500/70 transition-colors rounded-t relative group"
                          style={{ height: `${Math.max(heightPercent, 5)}%` }}
                        >
                          <div className="absolute bottom-full left-1/2 -translate-x-1/2 mb-2 px-2 py-1 bg-obsidian-800 rounded text-xs text-white opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                            Peak: {point.peakPlayers} | Avg: {point.avgPlayers}
                            <br />
                            {new Date(point.hour).toLocaleTimeString('tr-TR', { hour: '2-digit', minute: '2-digit' })}
                          </div>
                        </div>
                      );
                    })}
                  </div>
                )}
              </CardContent>
            </Card>

            {/* Activity Logs */}
            <Card className="bg-obsidian-900 border-obsidian-800">
              <CardHeader>
                <CardTitle className="text-white">Son Aktiviteler</CardTitle>
              </CardHeader>
              <CardContent>
                {logs.length === 0 ? (
                  <div className="text-center py-8 text-gray-400">Aktivite bulunamadi</div>
                ) : (
                  <div className="space-y-3">
                    {logs.map((log) => (
                      <div
                        key={log.id}
                        className="flex items-center gap-4 p-3 bg-obsidian-800/50 rounded-lg"
                      >
                        <div className={`w-2 h-2 rounded-full ${
                          log.action.includes('ban') ? 'bg-red-400' :
                          log.action.includes('login') ? 'bg-green-400' :
                          log.action.includes('logout') ? 'bg-yellow-400' :
                          'bg-blue-400'
                        }`} />
                        <div className="flex-1">
                          <p className="text-white">
                            <span className="font-medium">{log.actorName}</span>
                            <span className="text-gray-400"> {log.action} </span>
                            {log.targetName && (
                              <span className="font-medium">{log.targetName}</span>
                            )}
                          </p>
                          {log.details && (
                            <p className="text-sm text-gray-500">{log.details}</p>
                          )}
                        </div>
                        <span className="text-sm text-gray-500">{formatDate(log.createdAt)}</span>
                      </div>
                    ))}
                  </div>
                )}
              </CardContent>
            </Card>
          </>
        )}
      </main>
    </div>
  );
}
