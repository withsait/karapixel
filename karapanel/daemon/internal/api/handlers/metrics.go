package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/withsait/karapixel/karapanel/internal/metrics"
)

type MetricsHandler struct {
	collector *metrics.Collector
}

func NewMetricsHandler(collector *metrics.Collector) *MetricsHandler {
	return &MetricsHandler{collector: collector}
}

// GetSystemMetrics returns current system metrics
func (h *MetricsHandler) GetSystemMetrics(c *fiber.Ctx) error {
	m := h.collector.Collect()
	return c.JSON(m)
}

// GetCPUMetrics returns CPU metrics
func (h *MetricsHandler) GetCPUMetrics(c *fiber.Ctx) error {
	m := h.collector.Collect()
	return c.JSON(m.CPU)
}

// GetMemoryMetrics returns memory metrics
func (h *MetricsHandler) GetMemoryMetrics(c *fiber.Ctx) error {
	m := h.collector.Collect()
	return c.JSON(m.Memory)
}

// GetDiskMetrics returns disk metrics
func (h *MetricsHandler) GetDiskMetrics(c *fiber.Ctx) error {
	m := h.collector.Collect()
	return c.JSON(m.Disk)
}

// GetNetworkMetrics returns network metrics
func (h *MetricsHandler) GetNetworkMetrics(c *fiber.Ctx) error {
	m := h.collector.Collect()
	return c.JSON(m.Network)
}
