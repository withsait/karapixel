'use client';

import { useState, useEffect } from 'react';
import { Header } from '@/components/Header';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { api, DiscordLink, DiscordSettings } from '@/lib/api';

export default function DiscordPage() {
  const [links, setLinks] = useState<DiscordLink[]>([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<'links' | 'settings'>('links');
  const [page, setPage] = useState(0);
  const limit = 20;

  const fetchLinks = async () => {
    try {
      setLoading(true);
      const data = await api.getDiscordLinks({ limit, offset: page * limit });
      setLinks(data.links || []);
      setTotal(data.total);
    } catch (error) {
      console.error('Failed to fetch links:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchLinks();
  }, [page]);

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
        <h1 className="text-2xl font-bold text-white mb-6">Discord Yonetimi</h1>

        {/* Tabs */}
        <div className="flex gap-2 mb-6">
          <Button
            variant={activeTab === 'links' ? 'default' : 'outline'}
            onClick={() => setActiveTab('links')}
          >
            Bagli Hesaplar
          </Button>
          <Button
            variant={activeTab === 'settings' ? 'default' : 'outline'}
            onClick={() => setActiveTab('settings')}
          >
            Bot Ayarlari
          </Button>
        </div>

        {activeTab === 'links' && (
          <Card className="bg-obsidian-900 border-obsidian-800">
            <CardHeader>
              <CardTitle className="text-white">Bagli Discord Hesaplari ({total})</CardTitle>
            </CardHeader>
            <CardContent>
              {loading ? (
                <div className="text-center py-8 text-gray-400">Yukleniyor...</div>
              ) : links.length === 0 ? (
                <div className="text-center py-8 text-gray-400">Bagli hesap bulunamadi</div>
              ) : (
                <div className="overflow-x-auto">
                  <table className="w-full">
                    <thead>
                      <tr className="border-b border-obsidian-700">
                        <th className="text-left py-3 px-4 text-gray-400">Minecraft</th>
                        <th className="text-left py-3 px-4 text-gray-400">Discord</th>
                        <th className="text-left py-3 px-4 text-gray-400">Discord ID</th>
                        <th className="text-left py-3 px-4 text-gray-400">Baglanti Tarihi</th>
                        <th className="text-left py-3 px-4 text-gray-400">Durum</th>
                      </tr>
                    </thead>
                    <tbody>
                      {links.map((link) => (
                        <tr key={link.id} className="border-b border-obsidian-800 hover:bg-obsidian-800/50">
                          <td className="py-3 px-4">
                            <div className="flex items-center gap-2">
                              <img
                                src={`https://mc-heads.net/avatar/${link.playerUUID}/24`}
                                alt=""
                                className="w-6 h-6 rounded"
                              />
                              <span className="text-white">{link.playerUsername || link.playerUUID.slice(0, 8) + '...'}</span>
                            </div>
                          </td>
                          <td className="py-3 px-4 text-white">{link.discordUsername || 'Bilinmiyor'}</td>
                          <td className="py-3 px-4 text-gray-400 font-mono">{link.discordId}</td>
                          <td className="py-3 px-4 text-gray-300">{formatDate(link.linkedAt)}</td>
                          <td className="py-3 px-4">
                            <span className={`px-2 py-1 rounded text-xs ${
                              link.isVerified
                                ? 'bg-green-500/20 text-green-400'
                                : 'bg-yellow-500/20 text-yellow-400'
                            }`}>
                              {link.isVerified ? 'Dogrulanmis' : 'Beklemede'}
                            </span>
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
        )}

        {activeTab === 'settings' && <DiscordSettingsPanel />}
      </main>
    </div>
  );
}

function DiscordSettingsPanel() {
  const [settings, setSettings] = useState<Partial<DiscordSettings>>({
    prefix: '!',
    welcomeChannelId: '',
    logChannelId: '',
    linkedRoleId: '',
    staffRoleId: '',
  });
  const [guildId, setGuildId] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState({ type: '', text: '' });

  const handleSave = async () => {
    if (!guildId) {
      setMessage({ type: 'error', text: 'Guild ID gerekli' });
      return;
    }

    try {
      setLoading(true);
      await api.updateDiscordSettings(guildId, settings);
      setMessage({ type: 'success', text: 'Ayarlar kaydedildi' });
    } catch (error: any) {
      setMessage({ type: 'error', text: error.message || 'Ayarlar kaydedilemedi' });
    } finally {
      setLoading(false);
    }
  };

  const handleLoad = async () => {
    if (!guildId) return;

    try {
      setLoading(true);
      const data = await api.getDiscordSettings(guildId);
      setSettings(data);
      setMessage({ type: 'success', text: 'Ayarlar yuklendi' });
    } catch (error) {
      setMessage({ type: 'error', text: 'Ayarlar bulunamadi' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card className="bg-obsidian-900 border-obsidian-800">
      <CardHeader>
        <CardTitle className="text-white">Bot Ayarlari</CardTitle>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="flex gap-4">
          <div className="flex-1">
            <label className="block text-sm text-gray-400 mb-1">Guild ID</label>
            <input
              type="text"
              value={guildId}
              onChange={(e) => setGuildId(e.target.value)}
              className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
              placeholder="Discord sunucu ID'si"
            />
          </div>
          <div className="flex items-end">
            <Button variant="outline" onClick={handleLoad} disabled={loading || !guildId}>
              Yukle
            </Button>
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm text-gray-400 mb-1">Prefix</label>
            <input
              type="text"
              value={settings.prefix || ''}
              onChange={(e) => setSettings(s => ({ ...s, prefix: e.target.value }))}
              className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
              placeholder="!"
            />
          </div>
          <div>
            <label className="block text-sm text-gray-400 mb-1">Linked Role ID</label>
            <input
              type="text"
              value={settings.linkedRoleId || ''}
              onChange={(e) => setSettings(s => ({ ...s, linkedRoleId: e.target.value }))}
              className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
              placeholder="MC bagli hesaplara verilecek rol ID'si"
            />
          </div>
          <div>
            <label className="block text-sm text-gray-400 mb-1">Welcome Channel ID</label>
            <input
              type="text"
              value={settings.welcomeChannelId || ''}
              onChange={(e) => setSettings(s => ({ ...s, welcomeChannelId: e.target.value }))}
              className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
              placeholder="Hos geldin mesaji kanali"
            />
          </div>
          <div>
            <label className="block text-sm text-gray-400 mb-1">Log Channel ID</label>
            <input
              type="text"
              value={settings.logChannelId || ''}
              onChange={(e) => setSettings(s => ({ ...s, logChannelId: e.target.value }))}
              className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
              placeholder="Genel log kanali"
            />
          </div>
          <div>
            <label className="block text-sm text-gray-400 mb-1">Staff Role ID</label>
            <input
              type="text"
              value={settings.staffRoleId || ''}
              onChange={(e) => setSettings(s => ({ ...s, staffRoleId: e.target.value }))}
              className="w-full px-4 py-2 bg-obsidian-800 border border-obsidian-700 rounded-lg text-white"
              placeholder="Staff rol ID'si"
            />
          </div>
          <div className="flex items-center gap-2">
            <input
              type="checkbox"
              id="syncRoles"
              checked={settings.syncRoles || false}
              onChange={(e) => setSettings(s => ({ ...s, syncRoles: e.target.checked }))}
              className="w-4 h-4"
            />
            <label htmlFor="syncRoles" className="text-sm text-gray-400">Rolleri Senkronize Et</label>
          </div>
        </div>

        {message.text && (
          <p className={`text-sm ${message.type === 'error' ? 'text-red-400' : 'text-green-400'}`}>
            {message.text}
          </p>
        )}

        <div className="flex justify-end">
          <Button onClick={handleSave} disabled={loading || !guildId}>
            {loading ? 'Kaydediliyor...' : 'Kaydet'}
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}
