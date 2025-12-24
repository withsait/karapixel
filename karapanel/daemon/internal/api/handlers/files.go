package handlers

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/withsait/karapixel/karapanel/internal/server"
)

type FilesHandler struct {
	manager *server.Manager
}

func NewFilesHandler(manager *server.Manager) *FilesHandler {
	return &FilesHandler{manager: manager}
}

type FileInfo struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	IsDir   bool   `json:"isDir"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"modTime"`
}

// ListFiles lists files in a server's directory
func (h *FilesHandler) ListFiles(c *fiber.Ctx) error {
	serverID := c.Params("id")
	relativePath := c.Query("path", "/")

	srv, err := h.manager.GetServerConfig(serverID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sanitize path to prevent directory traversal
	cleanPath := filepath.Clean(relativePath)
	if strings.Contains(cleanPath, "..") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid path",
		})
	}

	fullPath := filepath.Join(srv.WorkDir, cleanPath)

	// Ensure we're still within workDir
	if !strings.HasPrefix(fullPath, srv.WorkDir) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path outside server directory",
		})
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	files := make([]FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		files = append(files, FileInfo{
			Name:    entry.Name(),
			Path:    filepath.Join(relativePath, entry.Name()),
			IsDir:   entry.IsDir(),
			Size:    info.Size(),
			ModTime: info.ModTime().Unix(),
		})
	}

	return c.JSON(fiber.Map{
		"path":  relativePath,
		"files": files,
	})
}

// GetFile returns a file's content
func (h *FilesHandler) GetFile(c *fiber.Ctx) error {
	serverID := c.Params("id")
	relativePath := c.Query("path")

	if relativePath == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path required",
		})
	}

	srv, err := h.manager.GetServerConfig(serverID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sanitize path
	cleanPath := filepath.Clean(relativePath)
	if strings.Contains(cleanPath, "..") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid path",
		})
	}

	fullPath := filepath.Join(srv.WorkDir, cleanPath)

	if !strings.HasPrefix(fullPath, srv.WorkDir) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path outside server directory",
		})
	}

	// Check file size (limit to 1MB for text preview)
	info, err := os.Stat(fullPath)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "File not found",
		})
	}

	if info.IsDir() {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot read directory",
		})
	}

	if info.Size() > 1024*1024 {
		return c.Status(400).JSON(fiber.Map{
			"error": "File too large for preview",
		})
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"path":    relativePath,
		"content": string(content),
		"size":    info.Size(),
	})
}

// SaveFile saves content to a file
func (h *FilesHandler) SaveFile(c *fiber.Ctx) error {
	serverID := c.Params("id")

	var input struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	srv, err := h.manager.GetServerConfig(serverID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sanitize path
	cleanPath := filepath.Clean(input.Path)
	if strings.Contains(cleanPath, "..") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid path",
		})
	}

	fullPath := filepath.Join(srv.WorkDir, cleanPath)

	if !strings.HasPrefix(fullPath, srv.WorkDir) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path outside server directory",
		})
	}

	// Create backup of existing file
	if _, err := os.Stat(fullPath); err == nil {
		backupPath := fullPath + ".bak"
		os.Rename(fullPath, backupPath)
	}

	// Write new content
	if err := os.WriteFile(fullPath, []byte(input.Content), 0644); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File saved",
		"path":    input.Path,
	})
}

// UploadFile handles file uploads
func (h *FilesHandler) UploadFile(c *fiber.Ctx) error {
	serverID := c.Params("id")
	relativePath := c.FormValue("path", "/")

	srv, err := h.manager.GetServerConfig(serverID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sanitize path
	cleanPath := filepath.Clean(relativePath)
	if strings.Contains(cleanPath, "..") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid path",
		})
	}

	destDir := filepath.Join(srv.WorkDir, cleanPath)

	if !strings.HasPrefix(destDir, srv.WorkDir) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path outside server directory",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "No file provided",
		})
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer src.Close()

	// Create destination file
	destPath := filepath.Join(destDir, file.Filename)
	dst, err := os.Create(destPath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer dst.Close()

	// Copy file
	if _, err := io.Copy(dst, src); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":  "File uploaded",
		"filename": file.Filename,
		"path":     filepath.Join(relativePath, file.Filename),
	})
}

// DeleteFile deletes a file or directory
func (h *FilesHandler) DeleteFile(c *fiber.Ctx) error {
	serverID := c.Params("id")
	relativePath := c.Query("path")

	if relativePath == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path required",
		})
	}

	srv, err := h.manager.GetServerConfig(serverID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sanitize path
	cleanPath := filepath.Clean(relativePath)
	if strings.Contains(cleanPath, "..") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid path",
		})
	}

	fullPath := filepath.Join(srv.WorkDir, cleanPath)

	if !strings.HasPrefix(fullPath, srv.WorkDir) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path outside server directory",
		})
	}

	// Prevent deleting root
	if fullPath == srv.WorkDir {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot delete server root directory",
		})
	}

	if err := os.RemoveAll(fullPath); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Deleted",
		"path":    relativePath,
	})
}

// DownloadFile serves a file for download
func (h *FilesHandler) DownloadFile(c *fiber.Ctx) error {
	serverID := c.Params("id")
	relativePath := c.Query("path")

	if relativePath == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path required",
		})
	}

	srv, err := h.manager.GetServerConfig(serverID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sanitize path
	cleanPath := filepath.Clean(relativePath)
	if strings.Contains(cleanPath, "..") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid path",
		})
	}

	fullPath := filepath.Join(srv.WorkDir, cleanPath)

	if !strings.HasPrefix(fullPath, srv.WorkDir) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Path outside server directory",
		})
	}

	return c.SendFile(fullPath)
}
