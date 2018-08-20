package tasks

import (
	"fmt"
)

// HandleTasks receives int sleepInterval. It runs handle func and sleep given interval
func HandleTasks(tasks []Task) {
	for _, task := range tasks {
		handle(task)
	}
}

func handle(task Task) {
	// fmt.Println("handle")
	// task := Task{Action: "action", Schedule: "*/2 * * * *"}
	// currentTime := time.Now()
	nextTime := task.NextExecutionTime()
	fmt.Println("Next execution time for schedule", task.Schedule, "is a", nextTime)
}
