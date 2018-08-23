package daemon

import (
	"fmt"
	"github.com/pustserg/scheduler/tasks"
	"time"
)

// sleepInterval in seconds
const sleepInterval = 5

// Daemon is a struct with daemon info
type Daemon struct {
	State       string
	StopChannel chan bool
}

// Start daemon
func (d *Daemon) Start(workersCount int, repo *tasks.TaskRepository) {
	d.State = "started"
	go startInfiniteHandler(workersCount, repo)
}

// Stop daemon
func (d *Daemon) Stop() {
	d.State = "stopped"
}

func startInfiniteHandler(workersCount int, repo *tasks.TaskRepository) {
	for {
		tasksToHandle := repo.GetTasksForHandle()
		fmt.Println("In daemon got tasks to handle", len(tasksToHandle))
		for _, task := range tasksToHandle {
			go task.Process(*repo)
		}
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
