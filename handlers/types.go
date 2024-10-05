package handlers

// SpeedTestResult represents the result of a speed test
type SpeedTestResult struct {
	DownloadSpeed string `json:"download_speed"`
	UploadSpeed   string `json:"upload_speed"`
	IPAddress     string `json:"ip_address"`
}
