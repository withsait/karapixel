import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatBytes(bytes: number, decimals = 2): string {
  if (bytes === 0) return '0 B';

  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(decimals))} ${sizes[i]}`;
}

export function formatUptime(seconds: number): string {
  const days = Math.floor(seconds / 86400);
  const hours = Math.floor((seconds % 86400) / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);

  if (days > 0) {
    return `${days}d ${hours}h ${minutes}m`;
  }
  if (hours > 0) {
    return `${hours}h ${minutes}m`;
  }
  return `${minutes}m`;
}

export function formatDate(timestamp: number): string {
  return new Date(timestamp * 1000).toLocaleString('tr-TR');
}

export function getStatusColor(status: string): string {
  switch (status) {
    case 'online':
      return 'text-green-500';
    case 'offline':
      return 'text-red-500';
    case 'starting':
    case 'stopping':
      return 'text-yellow-500';
    default:
      return 'text-gray-500';
  }
}

export function getStatusBgColor(status: string): string {
  switch (status) {
    case 'online':
      return 'bg-green-500/20';
    case 'offline':
      return 'bg-red-500/20';
    case 'starting':
    case 'stopping':
      return 'bg-yellow-500/20';
    default:
      return 'bg-gray-500/20';
  }
}
