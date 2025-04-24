package functions

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// CPUStats holds the CPU utilization data
type CPUStats struct {
	Usage     float64 `json:"usage"`
	Timestamp int64   `json:"timestamp"`
}

// GetCPUStats returns the current CPU utilization percentage
func GetCPUStats() (*CPUStats, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	// Just get the overall CPU usage (not per core)
	return &CPUStats{
		Usage:     percentages[0],
		Timestamp: time.Now().Unix(),
	}, nil
}

// StartCPUMonitoring begins periodic monitoring of CPU usage and emits events
// Context is passed from the main App
func StartCPUMonitoring(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			stats, err := GetCPUStats()
			if err == nil {
				// Emit event to frontend with current CPU stats
				runtime.EventsEmit(ctx, "cpu-stats-update", stats)
			}
		}
	}()
}
