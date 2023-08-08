# System Monitoring CLI

This is a simple command-line tool written in Go that monitors system resources on a Windows computer and logs the information to a file.

## Features

- Monitors CPU usage, memory usage, and disk usage
- Logs information to a customizable log file
- Supports customization of the monitoring interval

## Prerequisites

- Go 1.11 or later installed
- Windows operating system

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/system-monitor.git
   cd system-monitor
   ```

2. Install the required Go packages
    
    ```bash
    go get github.com/shirou/gopsutil/cpu
    go get github.com/shirou/gopsutil/mem
    go get github.com/shirou/gopsutil/disk
    ```

3. Build the executable
    
    ```bash
    go build -o system_monitor.exe main.go
    ```

## Usage

```bash
./system_monitor.exe -log=system_monitor.log -interval=5s
