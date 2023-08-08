package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	logFilePath := flag.String("log", "system_monitor.log", "Path to the log file")
	interval := flag.Duration("interval", 5*time.Second, "Interval between monitoring updates")

	flag.Parse()

	logFile, err := os.OpenFile(*logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	ticker := time.Tick(*interval)
	for range ticker {
		if err := monitorAndLog(logger); err != nil {
			logger.Println("Error:", err)
		}
	}
}

func monitorAndLog(logger *log.Logger) error {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return fmt.Errorf("error getting CPU usage: %w", err)
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Errorf("error getting memory info: %w", err)
	}

	diskInfo, err := disk.Usage("/")
	if err != nil {
		return fmt.Errorf("error getting disk usage info: %w", err)
	}

	logger.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
	logger.Printf("Memory Info: %+v\n", memInfo)
	logger.Printf("Disk Usage: %+v\n", diskInfo)

	return nil
}
