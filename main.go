package main

import (
	"fmt"
	"github.com/pustserg/scheduler/daemon"
)

func main() {
	daemonInstance := daemon.Daemon{}
	err := daemonInstance.Start()
	if err != nil {
		panic("Daemon not started")
	}
	if daemonInstance.State == "started" {
		fmt.Println("Instance started")
	} else {
		fmt.Println("Instance not started")
	}
	err = daemonInstance.Stop()
	if err != nil {
		panic("Daemon not stopped")
	}
	if daemonInstance.State == "started" {
		fmt.Println("Instance started")
	} else {
		fmt.Println("Instance not started")
	}
	var input string
	fmt.Scanln(&input)
}
