package server

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type SystemdClient struct {
	// We use exec instead of dbus for simplicity and portability
}

func NewSystemdClient() (*SystemdClient, error) {
	// Verify systemctl is available
	if _, err := exec.LookPath("systemctl"); err != nil {
		return nil, fmt.Errorf("systemctl not found: %w", err)
	}
	return &SystemdClient{}, nil
}

func (s *SystemdClient) GetUnitStatus(unitName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "systemctl", "is-active", unitName)
	output, err := cmd.Output()
	if err != nil {
		// is-active returns exit code 3 for inactive units
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 3 {
				return "inactive", nil
			}
		}
		return "unknown", nil
	}

	return strings.TrimSpace(string(output)), nil
}

func (s *SystemdClient) GetUnitPID(unitName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "systemctl", "show", unitName, "--property=MainPID", "--value")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func (s *SystemdClient) StartUnit(unitName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "systemctl", "start", unitName)
	return cmd.Run()
}

func (s *SystemdClient) StopUnit(unitName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "systemctl", "stop", unitName)
	return cmd.Run()
}

func (s *SystemdClient) RestartUnit(unitName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "systemctl", "restart", unitName)
	return cmd.Run()
}

func (s *SystemdClient) KillUnit(unitName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "systemctl", "kill", "-s", "SIGKILL", unitName)
	return cmd.Run()
}

func (s *SystemdClient) Close() error {
	return nil
}

// StreamJournalLogs streams logs from journalctl for a unit
func (s *SystemdClient) StreamJournalLogs(ctx context.Context, unitName string, lines int) (<-chan string, error) {
	cmd := exec.CommandContext(ctx, "journalctl", "-u", unitName, "-f", "-n", strconv.Itoa(lines), "--no-pager", "-o", "cat")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	logChan := make(chan string, 100)

	go func() {
		defer close(logChan)
		defer cmd.Wait()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			case logChan <- scanner.Text():
			default:
				// Drop log if channel is full
			}
		}
	}()

	return logChan, nil
}

// GetRecentLogs gets recent logs without streaming
func (s *SystemdClient) GetRecentLogs(unitName string, lines int) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "journalctl", "-u", unitName, "-n", strconv.Itoa(lines), "--no-pager", "-o", "cat")
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
