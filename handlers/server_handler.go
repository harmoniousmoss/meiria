package handlers

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

// FetchAvailableServers fetches available speed test servers
func FetchAvailableServers() ([]*speedtest.Server, error) {
	servers, err := speedtest.FetchServers()
	if err != nil {
		return nil, err
	}

	// Find the best server based on latency
	targets, err := servers.FindServer([]int{})
	if err != nil || len(targets) == 0 {
		return nil, fmt.Errorf("failed to find target servers")
	}
	return targets, nil
}
