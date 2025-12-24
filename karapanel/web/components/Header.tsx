"use client";

import { useState } from "react";
import { useRouter, usePathname } from "next/navigation";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { LogOut, Server, HardDrive, ChevronDown } from "lucide-react";

export function Header() {
  const router = useRouter();
  const pathname = usePathname();
  const [adminOpen, setAdminOpen] = useState(false);

  if (pathname === '/login') {
    return null;
  }

  const handleLogout = () => {
    api.logout();
    router.push('/login');
    router.refresh();
  };

  const isActive = (path: string, exact: boolean = true) => {
    if (exact) return pathname === path;
    return pathname.startsWith(path);
  };

  const isAdminActive = pathname.startsWith('/admin');

  return (
    <header className="border-b border-obsidian-800 bg-obsidian-900">
      <div className="container mx-auto px-4 py-3 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 bg-indigo-600 rounded-lg flex items-center justify-center">
            <span className="text-white font-bold">K</span>
          </div>
          <span className="font-semibold text-lg text-white">KaraPanel</span>
        </div>
        <nav className="flex items-center gap-4">
          <a
            href="/"
            className={`text-sm transition-colors ${isActive('/') ? 'text-indigo-400 font-medium' : 'text-gray-400 hover:text-white'}`}
          >
            Dashboard
          </a>

          {/* Admin Dropdown */}
          <div className="relative">
            <button
              onClick={() => setAdminOpen(!adminOpen)}
              className={`flex items-center gap-1 text-sm transition-colors ${isAdminActive ? 'text-indigo-400 font-medium' : 'text-gray-400 hover:text-white'}`}
            >
              <Server className="h-4 w-4" />
              Sunucular
              <ChevronDown className={`h-3 w-3 transition-transform ${adminOpen ? 'rotate-180' : ''}`} />
            </button>

            {adminOpen && (
              <>
                <div
                  className="fixed inset-0 z-40"
                  onClick={() => setAdminOpen(false)}
                />
                <div className="absolute top-full left-0 mt-2 w-48 bg-obsidian-800 border border-obsidian-700 rounded-lg shadow-xl z-50 overflow-hidden">
                  <a
                    href="/admin/dedicated-servers"
                    onClick={() => setAdminOpen(false)}
                    className={`flex items-center gap-2 px-4 py-2.5 text-sm transition-colors ${
                      isActive('/admin/dedicated-servers', false)
                        ? 'bg-indigo-600/20 text-indigo-400'
                        : 'text-gray-300 hover:bg-obsidian-700'
                    }`}
                  >
                    <Server className="h-4 w-4" />
                    Dedicated Servers
                  </a>
                  <a
                    href="/admin/nodes"
                    onClick={() => setAdminOpen(false)}
                    className={`flex items-center gap-2 px-4 py-2.5 text-sm transition-colors ${
                      isActive('/admin/nodes', false)
                        ? 'bg-indigo-600/20 text-indigo-400'
                        : 'text-gray-300 hover:bg-obsidian-700'
                    }`}
                  >
                    <HardDrive className="h-4 w-4" />
                    Nodes
                  </a>
                </div>
              </>
            )}
          </div>

          <a
            href="/players"
            className={`text-sm transition-colors ${isActive('/players', false) ? 'text-indigo-400 font-medium' : 'text-gray-400 hover:text-white'}`}
          >
            Oyuncular
          </a>
          <a
            href="/moderation"
            className={`text-sm transition-colors ${isActive('/moderation') ? 'text-indigo-400 font-medium' : 'text-gray-400 hover:text-white'}`}
          >
            Moderasyon
          </a>
          <a
            href="/discord"
            className={`text-sm transition-colors ${isActive('/discord') ? 'text-indigo-400 font-medium' : 'text-gray-400 hover:text-white'}`}
          >
            Discord
          </a>
          <a
            href="/analytics"
            className={`text-sm transition-colors ${isActive('/analytics') ? 'text-indigo-400 font-medium' : 'text-gray-400 hover:text-white'}`}
          >
            Istatistik
          </a>
          <Button variant="ghost" size="sm" onClick={handleLogout} className="text-gray-400 hover:text-white">
            <LogOut className="h-4 w-4 mr-2" />
            Cikis
          </Button>
        </nav>
      </div>
    </header>
  );
}
