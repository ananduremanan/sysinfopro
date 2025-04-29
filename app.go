package main

import (
	"context"
	"encoding/json"
	"fmt"

	"myproject/functions"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	functions.StartCPUMonitoring(ctx)
	functions.StartMemoryMonitoring(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetSysInfo() (string, error) {
	sysInfo, err := functions.GetSystemInfo()
	if err != nil {
		return "Error Occured", fmt.Errorf("error gathering system info: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(sysInfo, "", "  ")
	if err != nil {
		return "Error Occured", fmt.Errorf("error marshalling JSON: %w", err)
	}

	return string(jsonBytes), nil
}

func (a *App) GetInterNetSpeed() (string, error) {
	internetSpeedStat, err := functions.GetInterNetSpeed()
	if err != nil {
		return "Error Occured", fmt.Errorf("error gathering internet speed: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(internetSpeedStat, "", "  ")

	if err != nil {
		return "Error Occured", fmt.Errorf("error marshalling JSON: %w", err)
	}

	return string(jsonBytes), nil
}

func (a *App) GetCPUStats() (*functions.CPUStats, error) {
	return functions.GetCPUStats()
}

func (a *App) GetMemoryStats() (*functions.MemoryStats, error) {
	return functions.GetMemoryStats()
}

func (a *App) CheckInternetConnection() bool {
	return functions.IsConnectedToInternet()
}

func (a *App) GetBatteryDetails() (string, error) {
	batteryDetails, err := functions.GetBatteryDetails()
	if err != nil {
		return "Error Occured", fmt.Errorf("error gathering battery details: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(batteryDetails, "", "  ")
	if err != nil {
		return "Error Occured", fmt.Errorf("error marshalling JSON: %w", err)
	}

	return string(jsonBytes), nil
}

func (a *App) ScanCleanableFiles() (functions.CleanerResult, error) {
	result := functions.GetCleanableFiles()
	return result, nil
}

// ScanSafeCleanableFiles scans only user directories (no admin required)
func (a *App) ScanSafeCleanableFiles() (functions.CleanerResult, error) {
	result := functions.SafeClean()
	return result, nil
}

// CleanSelectedFiles cleans the selected files and returns results
func (a *App) CleanSelectedFiles(files []functions.FileInfo) map[string]interface{} {
	count, size, failures := functions.CleanFiles(files)

	return map[string]interface{}{
		"cleanedCount":  count,
		"cleanedSize":   size,
		"formattedSize": functions.GetFormattedSize(size),
		"failures":      failures,
	}
}

// CheckCleanerPermissions checks if we have elevated permissions
func (a *App) CheckCleanerPermissions() functions.PermissionStatus {
	return functions.CheckPermissions()
}

// GetFormattedSize returns a human-readable size
func (a *App) GetFormattedSize(size int64) string {
	return functions.GetFormattedSize(size)
}

// CheckUpdate checks for updates
func (a *App) CheckUpdate() (string, error) {
	updateStatus, err := functions.CheckUpdate()

	if err != nil {
		return "Error Occured", fmt.Errorf("error checking for update: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(updateStatus, "", "  ")
	if err != nil {
		return "Error Occured", fmt.Errorf("error marshalling JSON: %w", err)
	}

	return string(jsonBytes), nil
}
