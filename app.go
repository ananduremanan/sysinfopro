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
