package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedTestResult struct {
	DownloadSpeed string `json:"download_speed"`
	UploadSpeed   string `json:"upload_speed"`
	IPAddress     string `json:"ip_address"`
}

// getPublicIP fetches the public IP address of the user
func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip := ""
	_, err = fmt.Fscanf(resp.Body, "%s", &ip)
	if err != nil {
		return "", err
	}

	return ip, nil
}

// SpeedTestHandler performs the actual speed test and returns the results
func SpeedTestHandler(c *fiber.Ctx) error {
	// Get public IP (using an external service)
	ipAddress, err := getPublicIP()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to get public IP")
	}

	// Fetch available speed test servers
	servers, err := speedtest.FetchServers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to fetch servers")
	}

	// Select the best server based on latency
	targets, err := servers.FindServer([]int{})
	if err != nil || len(targets) == 0 {
		return c.Status(http.StatusInternalServerError).SendString("Failed to find target servers")
	}

	// Perform speed tests on each target server
	for _, s := range targets {
		// Perform ping test with a callback for logging latency
		s.PingTest(func(latency time.Duration) {
			fmt.Printf("Ping: %v\n", latency)
		})

		// Perform download and upload tests
		s.DownloadTest()
		s.UploadTest()

		// Create result structure
		result := SpeedTestResult{
			DownloadSpeed: fmt.Sprintf("%.2f Mbps", s.DLSpeed),
			UploadSpeed:   fmt.Sprintf("%.2f Mbps", s.ULSpeed),
			IPAddress:     ipAddress,
		}

		// Return JSON result to the client
		return c.JSON(result)
	}

	// If no server tests were run, return an error
	return c.Status(http.StatusInternalServerError).SendString("Failed to perform speed test")
}
