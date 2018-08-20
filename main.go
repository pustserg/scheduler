package main

import (
	"flag"
	"fmt"
	"github.com/pustserg/scheduler/daemon"
)

func main() {
	daemonWorkers := flag.Int("daemon-workers", 1, "Integer count of daemon workers")

	flag.Parse()
	startDaemon(*daemonWorkers)

	var input string
	fmt.Scanln(&input)
}

func startDaemon(workersCount int) {
	daemonInstance := daemon.Daemon{}
	daemonInstance.Start(workersCount)
}
