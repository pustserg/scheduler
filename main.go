package main

import (
	"flag"
	"fmt"
	"github.com/pustserg/scheduler/daemon"
	"github.com/pustserg/scheduler/tasks"
)

func main() {
	daemonWorkers := flag.Int("daemon-workers", 1, "Integer count of daemon workers")

	flag.Parse()
	repo := tasks.NewRepository()
	startDaemon(*daemonWorkers, repo)

	var input string
	fmt.Scanln(&input)
}

func startDaemon(workersCount int, repo *tasks.TaskRepository) {
	daemonInstance := daemon.Daemon{}
	daemonInstance.Start(workersCount, repo)
}
