package main

import "time"

type SpeedTestResultClient struct {
	IP      string `json:"ip";`
	ISP     string `json:"isp";`
	Lat     string `json:"lat";`
	Lon     string `json:"lon";`
	Country string `json:"country";`
}

type SpeedTestResultServer struct {
	Url     string  `json:"url";`
	Name    string  `json:"name";`
	Country string  `json:"country";`
	CC      string  `json:"cc";`
	ID      string  `json:"id";`
	Host    string  `json:"host";`
	Latency float64 `json:"latency";`
}

type SpeedTestResult struct {
	Ping float64 `json:"ping";`

	Download      float64   `json:"download";`
	Upload        float64   `json:"upload";`
	BytesSent     int64     `json:"bytes_sent";`
	BytesReceived int64     `json:"bytes_received";`
	Timestamp     time.Time `json:"timestamp";`
	Server        SpeedTestResultServer
	Client        SpeedTestResultClient

	Duration time.Duration
}
