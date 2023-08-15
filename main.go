package main

import (
	"net/http"
	"os/exec"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuUsageMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage_percent",
			Help: "CPU Usage in Percentage",
		},
		[]string{"core"},
	)

	memoryUsageMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_usage_bytes",
			Help: "Memory Usage in Bytes",
		},
	)

	// Define load average metric
	loadAverageMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "load_average",
			Help: "System Load Average",
		},
		[]string{"type"},
	)

	// Define disk usage metric
	diskUsageMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "disk_usage_bytes",
			Help: "Disk Usage in Bytes",
		},
	)

	// Define swap usage metric
	swapUsageMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "swap_usage_bytes",
			Help: "Swap Usage in Bytes",
		},
	)
)

func init() {
	prometheus.MustRegister(cpuUsageMetric)
	prometheus.MustRegister(memoryUsageMetric)
	prometheus.MustRegister(loadAverageMetric) // Register load average metric
	prometheus.MustRegister(diskUsageMetric)   // Register disk usage metric
	prometheus.MustRegister(swapUsageMetric)   // Register swap usage metric
	// Register other metrics...
}

func executeCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func updateMetrics() {
	// Update CPU metrics
	cpuUsageOutput, _ := executeCommand("top -bn 1 | grep Cpu")
	// Parse cpuUsageOutput and update cpuUsageMetric for each core

	// Update Memory metrics
	memUsageOutput, _ := executeCommand("free -h")
	// Parse memUsageOutput and update memoryUsageMetric

	// Update Load Average metric
	loadAvgOutput, _ := executeCommand("uptime")
	// Parse loadAvgOutput and update loadAverageMetric

	// Update Disk Usage metric
	diskUsageOutput, _ := executeCommand("df -h /")
	// Parse diskUsageOutput and update diskUsageMetric

	// Update Swap Usage metric
	swapUsageOutput, _ := executeCommand("free -h")
	// Parse swapUsageOutput and update swapUsageMetric

	// Update other metrics...
}

func main() {
	// Expose metrics to Prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Start HTTP server in a goroutine
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	// Update and export metrics periodically
	for {
		updateMetrics()
		time.Sleep(2 * time.Second)
	}
}
