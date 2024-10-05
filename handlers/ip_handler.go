package handlers

import (
	"fmt"
	"net/http"
)

// GetPublicIP fetches the public IP address of the user
func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ip string
	_, err = fmt.Fscanf(resp.Body, "%s", &ip)
	if err != nil {
		return "", err
	}

	return ip, nil
}
