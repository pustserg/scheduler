package daemon

import (
	"github.com/pustserg/scheduler/tasks"
	"time"
)

// sleepInterval in seconds
const sleepInterval = 20

// Daemon is a struct with daemon info
type Daemon struct {
	State       string
	StopChannel chan bool
}

// Start daemon
func (d *Daemon) Start(workersCount int) {
	d.State = "started"
	go startInfiniteHandler(workersCount)
}

// Stop daemon
func (d *Daemon) Stop() {
	d.State = "stopped"
}

func startInfiniteHandler(workersCount int) {
	for {
		tasksToHandle := tasks.GetTasksForHandle()
		for i := 0; i < workersCount; i++ {
			workerTasks := tasksToHandle
			go tasks.HandleTasks(workerTasks)
		}
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
