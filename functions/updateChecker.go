package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	CurrentTag string `json:"current_tag"`
}

// Updating the Tag struct to match the actual GitHub API response
type Tag struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     Commit `json:"commit"`
	NodeID     string `json:"node_id"`
}

type Commit struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

type UpdateStatus struct {
	IsUpdateAvailable bool   `json:"is_update_available"`
	NewTag            string `json:"new_tag"`
}

func CheckUpdate() (*UpdateStatus, error) {
	updateStatus := &UpdateStatus{
		IsUpdateAvailable: false,
	}

	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config: %v", err)
	}

	// Fetch tags from GitHub
	url := "https://api.github.com/repos/ananduremanan/sysinfopro/tags"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching tags: %v", err)
	}
	defer resp.Body.Close()

	var tags []Tag
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, fmt.Errorf("error decoding tags: %v", err)
	}

	// Check if there are tags available
	if len(tags) == 0 {
		return updateStatus, nil
	}

	// The first tag in the response is the newest one
	latestTag := tags[0].Name

	// Compare with current tag
	if latestTag != config.CurrentTag {
		updateStatus.IsUpdateAvailable = true
		updateStatus.NewTag = latestTag
	}

	return updateStatus, nil
}
