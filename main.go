package main

import (
	"flag"
	"fmt"
	"github.com/pustserg/scheduler/daemon"
)

func main() {
	daemonWorkers := flag.Int("daemon-workers", 1, "Integer count of daemon workers")

	flag.Parse()

	fmt.Println(*daemonWorkers)

	err := startDaemon(*daemonWorkers)
	if err != nil {
		panic("Daemon not started")
	}

	var input string
	fmt.Scanln(&input)
}

func startDaemon(workersCount int) error {
	daemonInstance := daemon.Daemon{}
	err := daemonInstance.Start(workersCount)
	return err
}
