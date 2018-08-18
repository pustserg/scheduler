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
	// fmt.Println("handle")
	task := Task{Action: "action", Schedule: "*/2 * * * *"}
	// currentTime := time.Now()
	nextTime := task.NextExecutionTime()
	fmt.Println("Next execution time for schedule", task.Schedule, "is a", nextTime)
}
