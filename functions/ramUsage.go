package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func GetMemoryStats() (*MemoryStats, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error getting memory: %v", err)
	}

	return &MemoryStats{
		TotalRAMMB:   vmStat.Total / 1024 / 1024,
		FreeRAMMB:    vmStat.Free / 1024 / 1024,
		UsedRAMMB:    (vmStat.Total - vmStat.Free) / 1024 / 1024,
		UsagePercent: vmStat.UsedPercent,
		Timestamp:    time.Now().Unix(),
	}, nil
}

// StartMemoryMonitoring begins periodic monitoring of RAM usage and emits events
func StartMemoryMonitoring(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			stats, err := GetMemoryStats()
			if err == nil {
				// Emit event to frontend with current RAM stats
				runtime.EventsEmit(ctx, "ram-stats-update", stats)
			}
		}
	}()
}
