"use client";

import { useRouter, usePathname } from "next/navigation";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { LogOut } from "lucide-react";

export function Header() {
  const router = useRouter();
  const pathname = usePathname();

  // Don't show header on login page
  if (pathname === '/login') {
    return null;
  }

  const handleLogout = () => {
    api.logout();
    router.push('/login');
    router.refresh();
  };

  return (
    <header className="border-b bg-card">
      <div className="container mx-auto px-4 py-3 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 bg-primary rounded-lg flex items-center justify-center">
            <span className="text-primary-foreground font-bold">K</span>
          </div>
          <span className="font-semibold text-lg">KaraPanel</span>
        </div>
        <nav className="flex items-center gap-4">
          <a href="/" className="text-sm hover:text-primary transition-colors">
            Dashboard
          </a>
          <a href="/metrics" className="text-sm hover:text-primary transition-colors">
            Metrics
          </a>
          <Button variant="ghost" size="sm" onClick={handleLogout}>
            <LogOut className="h-4 w-4 mr-2" />
            Logout
          </Button>
        </nav>
      </div>
    </header>
  );
}
