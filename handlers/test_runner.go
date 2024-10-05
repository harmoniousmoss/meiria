package handlers

import (
	"fmt"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

// RunSpeedTest runs the download, upload, and ping tests on the provided server
func RunSpeedTest(s *speedtest.Server) (*SpeedTestResult, error) {
	// Perform ping test with a callback for logging latency
	s.PingTest(func(latency time.Duration) {
		fmt.Printf("Ping: %v\n", latency)
	})

	// Perform download and upload tests
	err := s.DownloadTest()
	if err != nil {
		return nil, fmt.Errorf("failed to perform download test")
	}

	err = s.UploadTest()
	if err != nil {
		return nil, fmt.Errorf("failed to perform upload test")
	}

	// Create result structure
	return &SpeedTestResult{
		DownloadSpeed: fmt.Sprintf("%.2f Mbps", s.DLSpeed),
		UploadSpeed:   fmt.Sprintf("%.2f Mbps", s.ULSpeed),
	}, nil
}
