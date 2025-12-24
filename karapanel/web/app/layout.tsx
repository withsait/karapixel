import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Providers } from "./providers";
import { Header } from "@/components/Header";

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
            {/* Header with auth */}
            <Header />

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
