package main

import (
	"flag"
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
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err != nil {
			logger.Println("Error getting CPU usage:", err)
		}

		memInfo, err := mem.VirtualMemory()
		if err != nil {
			logger.Println("Error getting memory info:", err)
		}

		diskInfo, err := disk.Usage("/")
		if err != nil {
			logger.Println("Error getting disk usage info:", err)
		}

		logger.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
		logger.Printf("Memory Info: %+v\n", memInfo)
		logger.Printf("Disk Usage: %+v\n", diskInfo)
	}
}
