package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/tabwriter"
)

func executeCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// CPU Metrics
	cpuUsageOutput, _ := executeCommand("top -bn 1 | grep Cpu")
	loadAverageOutput, _ := executeCommand("uptime")

	// Memory Metrics
	memUsageOutput, _ := executeCommand("free -h")
	swapUsageOutput, _ := executeCommand("swapon --show")

	// Disk Metrics
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
