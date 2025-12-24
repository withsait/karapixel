"use client";

import { Console } from "@/components/Console";
import { Button } from "@/components/ui/button";
import { ArrowLeft } from "lucide-react";
import Link from "next/link";
import { useParams } from "next/navigation";

export default function ConsolePage() {
  const params = useParams();
  const serverId = params.id as string;

  return (
    <div className="h-[calc(100vh-200px)] flex flex-col">
      {/* Header */}
      <div className="flex items-center gap-4 mb-4">
        <Link href="/">
          <Button variant="ghost" size="sm">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back
          </Button>
        </Link>
        <h1 className="text-2xl font-bold">Console - {serverId}</h1>
      </div>

      {/* Console */}
      <div className="flex-1 border rounded-lg overflow-hidden">
        <Console serverId={serverId} />
      </div>
    </div>
  );
}
