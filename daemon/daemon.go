package daemon

import (
	// _ "errors"
	"github.com/pustserg/scheduler/tasks"
)

const sleepInterval = 2

// Daemon is a struct with daemon info
type Daemon struct {
	State string
}

// Start daemon
func (d *Daemon) Start(workersCount int) error {
	d.State = "started"
	for i := 0; i < workersCount; i++ {
		go tasks.HandleTasks(sleepInterval)
	}
	return nil
}

// Stop daemon
func (d *Daemon) Stop() error {
	d.State = "stopped"
	return nil
}
