import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Providers } from "./providers";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "KaraPanel - Minecraft Server Control",
  description: "KaraPixel Minecraft Server Control Panel",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="tr" className="dark">
      <body className={inter.className}>
        <Providers>
          <div className="min-h-screen flex flex-col">
            {/* Header */}
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
                </nav>
              </div>
            </header>

            {/* Main content */}
            <main className="flex-1 container mx-auto px-4 py-6">
              {children}
            </main>

            {/* Footer */}
            <footer className="border-t py-4">
              <div className="container mx-auto px-4 text-center text-sm text-muted-foreground">
                KaraPanel v0.1.0 - KaraPixel Minecraft Network
              </div>
            </footer>
          </div>
        </Providers>
      </body>
    </html>
  );
}
