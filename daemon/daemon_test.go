package daemon

import (
	"fmt"
	"github.com/pustserg/scheduler/tasks"
	"testing"
)

func TestStart(t *testing.T) {
	fmt.Println("testing daemon.Start")
	repo := tasks.NewRepository("test.db")
	daemonInstance := Daemon{}
	daemonInstance.Start(1, repo)

	if daemonInstance.State != "started" {
		t.Error("Expected daemon to be started")
	}
}

func TestStop(t *testing.T) {
	daemonInstance := Daemon{}
	daemonInstance.Stop()

	if daemonInstance.State != "stopped" {
		t.Error("Expected daemon to be stopped")
	}
}
