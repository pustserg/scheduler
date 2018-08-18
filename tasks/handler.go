package tasks

import (
	"fmt"
	"time"
)

// HandleTasks receives int sleepInterval. It runs handle func and sleep given interval
func HandleTasks(sleepInterval int64) {
	for {
		handle()
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}

func handle() {
	fmt.Println("handle")
}
