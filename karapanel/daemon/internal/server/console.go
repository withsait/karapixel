package server

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// ConsoleManager manages console connections for servers
type ConsoleManager struct {
	sessions map[string]*ConsoleSession
	mu       sync.RWMutex
}

type ConsoleSession struct {
	ServerID    string
	LogFile     string
	subscribers map[string]chan string
	cancel      context.CancelFunc
	mu          sync.RWMutex
}

func NewConsoleManager() *ConsoleManager {
	return &ConsoleManager{
		sessions: make(map[string]*ConsoleSession),
	}
}

// Subscribe creates a new log subscription for a server
func (cm *ConsoleManager) Subscribe(serverID, subscriberID string, workDir string) (<-chan string, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	session, exists := cm.sessions[serverID]
	if !exists {
		// Create new session
		ctx, cancel := context.WithCancel(context.Background())
		session = &ConsoleSession{
			ServerID:    serverID,
			LogFile:     filepath.Join(workDir, "logs", "latest.log"),
			subscribers: make(map[string]chan string),
			cancel:      cancel,
		}
		cm.sessions[serverID] = session
		go session.tailLog(ctx)
	}

	// Create channel for this subscriber
	ch := make(chan string, 100)
	session.mu.Lock()
	session.subscribers[subscriberID] = ch
	session.mu.Unlock()

	return ch, nil
}

// Unsubscribe removes a log subscription
func (cm *ConsoleManager) Unsubscribe(serverID, subscriberID string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	session, exists := cm.sessions[serverID]
	if !exists {
		return
	}

	session.mu.Lock()
	if ch, ok := session.subscribers[subscriberID]; ok {
		close(ch)
		delete(session.subscribers, subscriberID)
	}
	session.mu.Unlock()

	// If no more subscribers, close the session
	if len(session.subscribers) == 0 {
		session.cancel()
		delete(cm.sessions, serverID)
	}
}

// tailLog continuously reads from the log file and broadcasts to subscribers
func (s *ConsoleSession) tailLog(ctx context.Context) {
	// Use journalctl for systemd services
	cmd := exec.CommandContext(ctx, "journalctl", "-u", s.ServerID, "-f", "-n", "100", "--no-pager", "-o", "cat")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		s.broadcast(fmt.Sprintf("[ERROR] Failed to start log stream: %v", err))
		return
	}

	if err := cmd.Start(); err != nil {
		// Fallback to file tailing if journalctl fails
		s.tailFile(ctx)
		return
	}

	reader := bufio.NewReader(stdout)
	for {
		select {
		case <-ctx.Done():
			cmd.Process.Kill()
			return
		default:
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					s.broadcast(fmt.Sprintf("[ERROR] Log read error: %v", err))
				}
				continue
			}
			s.broadcast(line)
		}
	}
}

// tailFile falls back to tailing the actual log file
func (s *ConsoleSession) tailFile(ctx context.Context) {
	file, err := os.Open(s.LogFile)
	if err != nil {
		s.broadcast(fmt.Sprintf("[ERROR] Cannot open log file: %v", err))
		return
	}
	defer file.Close()

	// Seek to end
	file.Seek(0, io.SeekEnd)

	reader := bufio.NewReader(file)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					// Wait for new content
					continue
				}
				return
			}
			s.broadcast(line)
		}
	}
}

// broadcast sends a message to all subscribers
func (s *ConsoleSession) broadcast(msg string) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, ch := range s.subscribers {
		select {
		case ch <- msg:
		default:
			// Drop if channel is full
		}
	}
}

// GetRecentLogs returns recent log entries
func (cm *ConsoleManager) GetRecentLogs(serverID string, lines int) ([]string, error) {
	cmd := exec.Command("journalctl", "-u", serverID, "-n", fmt.Sprintf("%d", lines), "--no-pager", "-o", "cat")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var logs []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	return logs, nil
}
