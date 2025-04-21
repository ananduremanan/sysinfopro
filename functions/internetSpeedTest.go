package functions

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Constants for the speed test
const (
	testDurationSeconds = 10
	downloadURL         = "https://speed.cloudflare.com/__down?bytes=100000000" // 100MB test file
	uploadURL           = "https://speed.cloudflare.com/__up"                   // Cloudflare upload endpoint
	bufferSize          = 4096                                                  // 4KB buffer
)

func GetInterNetSpeed() (*InternetSpeedStat, error) {
	speedStat := &InternetSpeedStat{}

	downloadSpeed := testDownloadSpeed()
	speedStat.DownloadSpeed = downloadSpeed

	uploadSpeed := testUploadSpeedOptimized()
	speedStat.UploadSpeed = uploadSpeed

	return speedStat, nil
}

// testDownloadSpeed measures download speed by downloading a large file
func testDownloadSpeed() float64 {
	fmt.Println("Testing download speed...")

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Duration(testDurationSeconds+5) * time.Second,
	}

	startTime := time.Now()

	// Make the request
	resp, err := client.Get(downloadURL)
	if err != nil {
		fmt.Printf("Error during download test: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	// Read the response
	buffer := make([]byte, bufferSize)
	var bytesRead int64

	// Use a timer to limit test duration
	timer := time.NewTimer(time.Duration(testDurationSeconds) * time.Second)
	done := make(chan bool)

	// Start reading in a goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				n, err := resp.Body.Read(buffer)
				if err != nil && err != io.EOF {
					fmt.Printf("Error reading response: %v\n", err)
					return
				}
				bytesRead += int64(n)
				if err == io.EOF {
					done <- true
					return
				}
			}
		}
	}()

	// Wait for either timeout or EOF
	select {
	case <-timer.C:
		done <- true
	case <-done:
		// Reading complete before timeout
	}

	wg.Wait()

	elapsedTime := time.Since(startTime).Seconds()
	if elapsedTime > float64(testDurationSeconds) {
		elapsedTime = float64(testDurationSeconds)
	}

	// Calculate speed in Mbps (megabits per second)
	// 8 bits per byte, divide by 1,000,000 for mega, divide by elapsed time for per second
	speedMbps := float64(bytesRead) * 8 / 1000000 / elapsedTime

	return speedMbps
}

// testUploadSpeed measures upload speed by uploading random data
func testUploadSpeedOptimized() float64 {
	fmt.Println("Testing upload speed...")

	// Create a smaller payload (5MB instead of 10MB)
	dataSize := 5000000 // 5MB
	reader := strings.NewReader(strings.Repeat("X", dataSize))

	// Set a shorter test duration (5 seconds)
	testDuration := 5 * time.Second

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: testDuration + (2 * time.Second),
	}

	startTime := time.Now()
	bytesUploaded := int64(0)

	// Use a context with timeout to control the overall test duration
	ctx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	// Only do a single upload or a small fixed number of uploads
	for i := 0; i < 3 && ctx.Err() == nil; i++ {
		// Reset reader position
		reader.Seek(0, io.SeekStart)

		// Create a request with context
		req, err := http.NewRequestWithContext(ctx, "POST", uploadURL, reader)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}
		req.Header.Set("Content-Type", "application/octet-stream")

		// Make upload request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error during upload test: %v\n", err)
			continue
		}

		// Read and discard response
		_, err = io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Error reading upload response: %v\n", err)
			continue
		}

		// Count uploaded bytes
		bytesUploaded += int64(dataSize)
	}

	elapsedTime := time.Since(startTime).Seconds()
	if elapsedTime < 0.1 {
		return 0 // Avoid division by zero or unrealistically short times
	}

	// Calculate speed in Mbps (megabits per second)
	speedMbps := float64(bytesUploaded) * 8 / 1000000 / elapsedTime

	return speedMbps
}
