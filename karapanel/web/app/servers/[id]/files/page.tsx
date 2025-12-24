"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { api, FileInfo } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { formatBytes, formatDate } from "@/lib/utils";
import {
  ArrowLeft,
  File,
  Folder,
  FolderUp,
  Download,
  Trash2,
  Save,
  RefreshCw
} from "lucide-react";
import Link from "next/link";

export default function FilesPage() {
  const params = useParams();
  const serverId = params.id as string;
  const queryClient = useQueryClient();

  const [currentPath, setCurrentPath] = useState("/");
  const [selectedFile, setSelectedFile] = useState<string | null>(null);
  const [fileContent, setFileContent] = useState<string>("");
  const [isEditing, setIsEditing] = useState(false);

  const { data: filesData, refetch } = useQuery({
    queryKey: ["files", serverId, currentPath],
    queryFn: () => api.listFiles(serverId, currentPath),
  });

  const { data: fileData, isLoading: loadingFile } = useQuery({
    queryKey: ["file", serverId, selectedFile],
    queryFn: () => api.getFile(serverId, selectedFile!),
    enabled: !!selectedFile,
  });

  const saveMutation = useMutation({
    mutationFn: () => api.saveFile(serverId, selectedFile!, fileContent),
    onSuccess: () => {
      setIsEditing(false);
      queryClient.invalidateQueries({ queryKey: ["file", serverId, selectedFile] });
    },
  });

  const deleteMutation = useMutation({
    mutationFn: (path: string) => api.deleteFile(serverId, path),
    onSuccess: () => {
      refetch();
      if (selectedFile) {
        setSelectedFile(null);
      }
    },
  });

  const files = filesData?.files || [];

  const navigateUp = () => {
    const parts = currentPath.split("/").filter(Boolean);
    parts.pop();
    setCurrentPath("/" + parts.join("/"));
    setSelectedFile(null);
  };

  const navigateToFolder = (folder: FileInfo) => {
    setCurrentPath(folder.path);
    setSelectedFile(null);
  };

  const openFile = (file: FileInfo) => {
    setSelectedFile(file.path);
    setIsEditing(false);
  };

  return (
    <div className="space-y-4">
      {/* Header */}
      <div className="flex items-center gap-4">
        <Link href="/">
          <Button variant="ghost" size="sm">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back
          </Button>
        </Link>
        <h1 className="text-2xl font-bold">Files - {serverId}</h1>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {/* File Browser */}
        <Card>
          <CardHeader className="pb-2">
            <div className="flex items-center justify-between">
              <CardTitle className="text-lg">
                {currentPath}
              </CardTitle>
              <Button variant="ghost" size="sm" onClick={() => refetch()}>
                <RefreshCw className="h-4 w-4" />
              </Button>
            </div>
          </CardHeader>
          <CardContent>
            <div className="space-y-1">
              {/* Navigate up */}
              {currentPath !== "/" && (
                <button
                  onClick={navigateUp}
                  className="w-full flex items-center gap-3 p-2 rounded hover:bg-secondary/50 transition-colors"
                >
                  <FolderUp className="h-4 w-4 text-muted-foreground" />
                  <span>..</span>
                </button>
              )}

              {/* Files list */}
              {files.map((file) => (
                <div
                  key={file.path}
                  className="flex items-center justify-between p-2 rounded hover:bg-secondary/50 transition-colors group"
                >
                  <button
                    onClick={() => file.isDir ? navigateToFolder(file) : openFile(file)}
                    className="flex items-center gap-3 flex-1 text-left"
                  >
                    {file.isDir ? (
                      <Folder className="h-4 w-4 text-blue-400" />
                    ) : (
                      <File className="h-4 w-4 text-muted-foreground" />
                    )}
                    <span className="truncate">{file.name}</span>
                  </button>

                  <div className="flex items-center gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                    {!file.isDir && (
                      <>
                        <span className="text-xs text-muted-foreground">
                          {formatBytes(file.size)}
                        </span>
                        <Button
                          variant="ghost"
                          size="icon"
                          className="h-7 w-7"
                          onClick={() => window.open(`${process.env.NEXT_PUBLIC_API_URL}/api/servers/${serverId}/files/download?path=${encodeURIComponent(file.path)}`)}
                        >
                          <Download className="h-3 w-3" />
                        </Button>
                      </>
                    )}
                    <Button
                      variant="ghost"
                      size="icon"
                      className="h-7 w-7 text-destructive"
                      onClick={() => deleteMutation.mutate(file.path)}
                    >
                      <Trash2 className="h-3 w-3" />
                    </Button>
                  </div>
                </div>
              ))}

              {files.length === 0 && (
                <div className="text-center py-8 text-muted-foreground">
                  Empty directory
                </div>
              )}
            </div>
          </CardContent>
        </Card>

        {/* File Editor */}
        <Card>
          <CardHeader className="pb-2">
            <div className="flex items-center justify-between">
              <CardTitle className="text-lg truncate">
                {selectedFile || "Select a file"}
              </CardTitle>
              {selectedFile && fileData && (
                <div className="flex gap-2">
                  {isEditing ? (
                    <>
                      <Button
                        size="sm"
                        onClick={() => saveMutation.mutate()}
                        disabled={saveMutation.isPending}
                      >
                        <Save className="h-4 w-4 mr-1" />
                        Save
                      </Button>
                      <Button
                        size="sm"
                        variant="ghost"
                        onClick={() => {
                          setIsEditing(false);
                          setFileContent(fileData.content);
                        }}
                      >
                        Cancel
                      </Button>
                    </>
                  ) : (
                    <Button size="sm" onClick={() => {
                      setFileContent(fileData.content);
                      setIsEditing(true);
                    }}>
                      Edit
                    </Button>
                  )}
                </div>
              )}
            </div>
          </CardHeader>
          <CardContent>
            {loadingFile ? (
              <div className="h-[500px] flex items-center justify-center">
                <div className="animate-spin h-8 w-8 border-2 border-primary border-t-transparent rounded-full" />
              </div>
            ) : selectedFile && fileData ? (
              <div className="relative">
                {isEditing ? (
                  <textarea
                    value={fileContent}
                    onChange={(e) => setFileContent(e.target.value)}
                    className="w-full h-[500px] p-4 font-mono text-sm bg-secondary/30 rounded border resize-none focus:outline-none focus:ring-2 focus:ring-primary"
                    spellCheck={false}
                  />
                ) : (
                  <pre className="w-full h-[500px] p-4 font-mono text-sm bg-secondary/30 rounded overflow-auto">
                    {fileData.content}
                  </pre>
                )}
              </div>
            ) : (
              <div className="h-[500px] flex items-center justify-center text-muted-foreground">
                Select a file to view its content
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
