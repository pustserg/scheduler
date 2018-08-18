package daemon

import (
	_ "errors"
	"github.com/pustserg/scheduler/tasks"
)

// Daemon is a struct with daemon info
type Daemon struct {
	State string
}

// Start daemon
func (d *Daemon) Start() error {
	d.State = "started"
	go tasks.HandleTasks(5)
	return nil
}

// Stop daemon
func (d *Daemon) Stop() error {
	d.State = "stopped"
	return nil
}
