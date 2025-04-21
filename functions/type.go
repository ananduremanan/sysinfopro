package functions

type SystemInfo struct {
	CPUModel        string  `json:"cpu_model"`
	Manufacturer    string  `json:"cpu_vendor"`
	Cores           int     `json:"cpu_cores"`
	LoggedInUser    string  `json:"logged_in_user"`
	TotalRAMMB      uint64  `json:"total_ram_mb"`
	FreeRAMMB       uint64  `json:"free_ram_mb"`
	DiskType        string  `json:"disk_type"`
	TotalDiskGB     uint64  `json:"total_disk_gb"`
	FreeDiskGB      uint64  `json:"free_disk_gb"`
	DiskUsagePct    float64 `json:"disk_usage_percent"`
	NetworkStatus   string  `json:"network_status"`
	OS              string  `json:"host_os"`
	Platform        string  `json:"host_platform"`
	PlatformVersion string  `json:"host_platform_version"`
}

type InternetSpeedStat struct {
	DownloadSpeed float64 `json:"download_speed"`
	UploadSpeed   float64 `json:"upload_speed"`
}
