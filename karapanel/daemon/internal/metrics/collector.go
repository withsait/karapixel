package metrics

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SystemMetrics struct {
	Timestamp   int64         `json:"timestamp"`
	CPU         CPUMetrics    `json:"cpu"`
	Memory      MemoryMetrics `json:"memory"`
	Disk        DiskMetrics   `json:"disk"`
	Network     NetworkMetrics `json:"network"`
	Uptime      int64         `json:"uptime"` // seconds
}

type CPUMetrics struct {
	UsagePercent float64   `json:"usagePercent"`
	CoreCount    int       `json:"coreCount"`
	PerCore      []float64 `json:"perCore"`
	LoadAvg      [3]float64 `json:"loadAvg"` // 1min, 5min, 15min
}

type MemoryMetrics struct {
	Total     int64   `json:"total"`     // bytes
	Used      int64   `json:"used"`      // bytes
	Free      int64   `json:"free"`      // bytes
	Available int64   `json:"available"` // bytes
	Cached    int64   `json:"cached"`    // bytes
	SwapTotal int64   `json:"swapTotal"` // bytes
	SwapUsed  int64   `json:"swapUsed"`  // bytes
	Percent   float64 `json:"percent"`
}

type DiskMetrics struct {
	Total       int64   `json:"total"`       // bytes
	Used        int64   `json:"used"`        // bytes
	Free        int64   `json:"free"`        // bytes
	Percent     float64 `json:"percent"`
	ReadBytes   int64   `json:"readBytes"`   // per second
	WriteBytes  int64   `json:"writeBytes"`  // per second
}

type NetworkMetrics struct {
	BytesRecv   int64 `json:"bytesRecv"`   // per second
	BytesSent   int64 `json:"bytesSent"`   // per second
	PacketsRecv int64 `json:"packetsRecv"` // per second
	PacketsSent int64 `json:"packetsSent"` // per second
}

type Collector struct {
	lastCPU       cpuTimes
	lastNetwork   networkStats
	lastDisk      diskStats
	lastCollect   time.Time
	mu            sync.Mutex
}

type cpuTimes struct {
	user, nice, system, idle, iowait, irq, softirq int64
}

type networkStats struct {
	bytesRecv, bytesSent, packetsRecv, packetsSent int64
}

type diskStats struct {
	readBytes, writeBytes int64
}

func NewCollector() *Collector {
	c := &Collector{
		lastCollect: time.Now(),
	}
	// Initialize baseline
	c.lastCPU = c.readCPUTimes()
	c.lastNetwork = c.readNetworkStats()
	c.lastDisk = c.readDiskStats()
	return c
}

func (c *Collector) Collect() *SystemMetrics {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(c.lastCollect).Seconds()
	if elapsed < 0.1 {
		elapsed = 1 // Avoid division by zero
	}

	metrics := &SystemMetrics{
		Timestamp: now.Unix(),
		CPU:       c.collectCPU(elapsed),
		Memory:    c.collectMemory(),
		Disk:      c.collectDisk(elapsed),
		Network:   c.collectNetwork(elapsed),
		Uptime:    c.getUptime(),
	}

	c.lastCollect = now
	return metrics
}

func (c *Collector) collectCPU(elapsed float64) CPUMetrics {
	current := c.readCPUTimes()

	totalDelta := float64(
		(current.user - c.lastCPU.user) +
		(current.nice - c.lastCPU.nice) +
		(current.system - c.lastCPU.system) +
		(current.idle - c.lastCPU.idle) +
		(current.iowait - c.lastCPU.iowait) +
		(current.irq - c.lastCPU.irq) +
		(current.softirq - c.lastCPU.softirq))

	idleDelta := float64(current.idle - c.lastCPU.idle)

	var usagePercent float64
	if totalDelta > 0 {
		usagePercent = 100 * (1 - idleDelta/totalDelta)
	}

	c.lastCPU = current

	loadAvg := c.readLoadAvg()

	return CPUMetrics{
		UsagePercent: usagePercent,
		CoreCount:    runtime.NumCPU(),
		LoadAvg:      loadAvg,
	}
}

func (c *Collector) readCPUTimes() cpuTimes {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return cpuTimes{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)
			if len(fields) >= 8 {
				user, _ := strconv.ParseInt(fields[1], 10, 64)
				nice, _ := strconv.ParseInt(fields[2], 10, 64)
				system, _ := strconv.ParseInt(fields[3], 10, 64)
				idle, _ := strconv.ParseInt(fields[4], 10, 64)
				iowait, _ := strconv.ParseInt(fields[5], 10, 64)
				irq, _ := strconv.ParseInt(fields[6], 10, 64)
				softirq, _ := strconv.ParseInt(fields[7], 10, 64)
				return cpuTimes{user, nice, system, idle, iowait, irq, softirq}
			}
		}
	}
	return cpuTimes{}
}

func (c *Collector) readLoadAvg() [3]float64 {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return [3]float64{}
	}

	fields := strings.Fields(string(data))
	if len(fields) < 3 {
		return [3]float64{}
	}

	var loadAvg [3]float64
	loadAvg[0], _ = strconv.ParseFloat(fields[0], 64)
	loadAvg[1], _ = strconv.ParseFloat(fields[1], 64)
	loadAvg[2], _ = strconv.ParseFloat(fields[2], 64)
	return loadAvg
}

func (c *Collector) collectMemory() MemoryMetrics {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return MemoryMetrics{}
	}
	defer file.Close()

	mem := make(map[string]int64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			key := strings.TrimSuffix(fields[0], ":")
			val, _ := strconv.ParseInt(fields[1], 10, 64)
			mem[key] = val * 1024 // Convert from kB to bytes
		}
	}

	total := mem["MemTotal"]
	free := mem["MemFree"]
	available := mem["MemAvailable"]
	cached := mem["Cached"] + mem["Buffers"]
	used := total - available

	return MemoryMetrics{
		Total:     total,
		Used:      used,
		Free:      free,
		Available: available,
		Cached:    cached,
		SwapTotal: mem["SwapTotal"],
		SwapUsed:  mem["SwapTotal"] - mem["SwapFree"],
		Percent:   float64(used) / float64(total) * 100,
	}
}

func (c *Collector) collectDisk(elapsed float64) DiskMetrics {
	current := c.readDiskStats()

	readPerSec := float64(current.readBytes-c.lastDisk.readBytes) / elapsed
	writePerSec := float64(current.writeBytes-c.lastDisk.writeBytes) / elapsed

	c.lastDisk = current

	// Get disk usage from statfs
	var total, used, free int64
	// Simplified - in production use syscall.Statfs
	dfOutput, err := os.ReadFile("/proc/mounts")
	if err == nil {
		// Parse first line to get root mount
		_ = dfOutput // Would need to parse this properly
	}

	return DiskMetrics{
		Total:      total,
		Used:       used,
		Free:       free,
		ReadBytes:  int64(readPerSec),
		WriteBytes: int64(writePerSec),
	}
}

func (c *Collector) readDiskStats() diskStats {
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		return diskStats{}
	}
	defer file.Close()

	var totalRead, totalWrite int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 14 {
			// Fields: major minor name reads_completed reads_merged sectors_read ms_reading writes_completed writes_merged sectors_written ms_writing ...
			// Skip loop devices and partitions
			name := fields[2]
			if strings.HasPrefix(name, "loop") || strings.HasPrefix(name, "ram") {
				continue
			}
			// Only count whole disks (sda, nvme0n1, etc)
			if len(name) > 3 && name[len(name)-1] >= '0' && name[len(name)-1] <= '9' {
				if name[len(name)-2] >= '0' && name[len(name)-2] <= '9' {
					continue // This is a partition
				}
			}

			sectorsRead, _ := strconv.ParseInt(fields[5], 10, 64)
			sectorsWritten, _ := strconv.ParseInt(fields[9], 10, 64)
			totalRead += sectorsRead * 512
			totalWrite += sectorsWritten * 512
		}
	}

	return diskStats{totalRead, totalWrite}
}

func (c *Collector) collectNetwork(elapsed float64) NetworkMetrics {
	current := c.readNetworkStats()

	recvPerSec := float64(current.bytesRecv-c.lastNetwork.bytesRecv) / elapsed
	sentPerSec := float64(current.bytesSent-c.lastNetwork.bytesSent) / elapsed
	packetsRecvPerSec := float64(current.packetsRecv-c.lastNetwork.packetsRecv) / elapsed
	packetsSentPerSec := float64(current.packetsSent-c.lastNetwork.packetsSent) / elapsed

	c.lastNetwork = current

	return NetworkMetrics{
		BytesRecv:   int64(recvPerSec),
		BytesSent:   int64(sentPerSec),
		PacketsRecv: int64(packetsRecvPerSec),
		PacketsSent: int64(packetsSentPerSec),
	}
}

func (c *Collector) readNetworkStats() networkStats {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return networkStats{}
	}
	defer file.Close()

	var stats networkStats
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		if lineNum <= 2 {
			continue // Skip headers
		}

		line := scanner.Text()
		colonIdx := strings.Index(line, ":")
		if colonIdx == -1 {
			continue
		}

		iface := strings.TrimSpace(line[:colonIdx])
		if iface == "lo" {
			continue // Skip loopback
		}

		fields := strings.Fields(line[colonIdx+1:])
		if len(fields) < 16 {
			continue
		}

		bytesRecv, _ := strconv.ParseInt(fields[0], 10, 64)
		packetsRecv, _ := strconv.ParseInt(fields[1], 10, 64)
		bytesSent, _ := strconv.ParseInt(fields[8], 10, 64)
		packetsSent, _ := strconv.ParseInt(fields[9], 10, 64)

		stats.bytesRecv += bytesRecv
		stats.bytesSent += bytesSent
		stats.packetsRecv += packetsRecv
		stats.packetsSent += packetsSent
	}

	return stats
}

func (c *Collector) getUptime() int64 {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0
	}

	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return 0
	}

	uptime, _ := strconv.ParseFloat(fields[0], 64)
	return int64(uptime)
}
