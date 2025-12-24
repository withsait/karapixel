package server

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getProcessMetrics reads CPU and memory usage from /proc for a given PID
func getProcessMetrics(pid int) (cpuPercent float64, memoryBytes int64) {
	// Read /proc/[pid]/stat for CPU
	statPath := fmt.Sprintf("/proc/%d/stat", pid)
	statData, err := os.ReadFile(statPath)
	if err != nil {
		return 0, 0
	}

	// Parse stat file - format is complex, we need fields after the comm (process name in parens)
	statStr := string(statData)
	// Find the closing paren of comm field
	closeParenIdx := strings.LastIndex(statStr, ")")
	if closeParenIdx == -1 {
		return 0, 0
	}
	fields := strings.Fields(statStr[closeParenIdx+2:])
	if len(fields) < 22 {
		return 0, 0
	}

	// Field indices (0-based after comm):
	// 11 = utime, 12 = stime, 13 = cutime, 14 = cstime
	// 20 = starttime, 21 = vsize, 22 = rss

	// Read RSS (field 21, 0-indexed after extracting)
	if len(fields) > 21 {
		rss, _ := strconv.ParseInt(fields[21], 10, 64)
		// RSS is in pages, multiply by page size (usually 4096)
		memoryBytes = rss * 4096
	}

	// For CPU percent, we'd need to track over time
	// For now, we'll read from /proc/[pid]/status which has VmRSS
	statusPath := fmt.Sprintf("/proc/%d/status", pid)
	statusData, err := os.ReadFile(statusPath)
	if err == nil {
		for _, line := range strings.Split(string(statusData), "\n") {
			if strings.HasPrefix(line, "VmRSS:") {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					mem, _ := strconv.ParseInt(parts[1], 10, 64)
					// VmRSS is in kB
					memoryBytes = mem * 1024
				}
			}
		}
	}

	// CPU calculation would require sampling over time
	// For simplicity, return 0 for now - real implementation would use gopsutil
	return 0, memoryBytes
}

// GetJavaHeapUsage parses JVM memory from /proc/[pid]/cmdline or uses jstat
func GetJavaHeapUsage(pid int) (heapUsed, heapMax int64, err error) {
	// Read cmdline to get max heap from -Xmx
	cmdlinePath := fmt.Sprintf("/proc/%d/cmdline", pid)
	cmdlineData, err := os.ReadFile(cmdlinePath)
	if err != nil {
		return 0, 0, err
	}

	cmdline := string(cmdlineData)
	// Arguments are null-separated
	args := strings.Split(cmdline, "\x00")

	for _, arg := range args {
		if strings.HasPrefix(arg, "-Xmx") {
			heapMax = parseMemoryArg(arg[4:])
		}
	}

	// For actual heap usage, we'd need to use JMX or jstat
	// This is a simplified version
	return 0, heapMax, nil
}

// parseMemoryArg parses JVM memory arguments like "6G", "512M", "1024K"
func parseMemoryArg(s string) int64 {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	multiplier := int64(1)
	suffix := s[len(s)-1]

	switch suffix {
	case 'g', 'G':
		multiplier = 1024 * 1024 * 1024
		s = s[:len(s)-1]
	case 'm', 'M':
		multiplier = 1024 * 1024
		s = s[:len(s)-1]
	case 'k', 'K':
		multiplier = 1024
		s = s[:len(s)-1]
	}

	val, _ := strconv.ParseInt(s, 10, 64)
	return val * multiplier
}
