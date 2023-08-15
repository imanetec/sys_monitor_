package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"
)

func executeCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func displayMetrics() {
	clearScreen()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// CPU Metrics
	cpuUsageOutput, _ := executeCommand("top -bn 1 | grep Cpu")

	// Load Average
	loadAverageOutput, _ := executeCommand("uptime")

	// Memory Metrics
	memUsageOutput, _ := executeCommand("free -h")

	// Swap Usage
	swapUsageOutput, _ := executeCommand("swapon --show")

	// Disk Usage
	diskUsageOutput, _ := executeCommand("df -h")

	// Print Metrics
	fmt.Fprintln(w, "=== CPU Usage ===")
	fmt.Fprintln(w, cpuUsageOutput)
	fmt.Fprintln(w, "=== Load Average ===")
	fmt.Fprintln(w, loadAverageOutput)
	fmt.Fprintln(w, "=== Memory Usage ===")
	fmt.Fprintln(w, memUsageOutput)
	fmt.Fprintln(w, "=== Swap Usage ===")
	fmt.Fprintln(w, swapUsageOutput)
	fmt.Fprintln(w, "=== Disk Usage ===")
	fmt.Fprintln(w, diskUsageOutput)

	w.Flush()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		displayMetrics()
		time.Sleep(2 * time.Second)
	}
}
