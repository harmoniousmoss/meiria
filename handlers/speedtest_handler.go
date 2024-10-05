package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// SpeedTestHandler performs the actual speed test and returns the results
func SpeedTestHandler(c *fiber.Ctx) error {
	// Get public IP (using an external service)
	ipAddress, err := GetPublicIP()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to get public IP")
	}

	// Fetch available speed test servers
	servers, err := FetchAvailableServers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to fetch speed test servers")
	}

	// Perform speed test on the first server (remove the loop)
	if len(servers) > 0 {
		s := servers[0]
		result, err := RunSpeedTest(s)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Failed to perform speed test")
		}

		// Add public IP to the result
		result.IPAddress = ipAddress

		// Return the JSON result
		return c.JSON(result)
	}

	// If no servers are available
	return c.Status(http.StatusInternalServerError).SendString("Failed to perform speed test")
}
