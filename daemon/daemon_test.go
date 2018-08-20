package daemon

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	fmt.Println("testing daemon.Start")
	daemonInstance := Daemon{}
	daemonInstance.Start(1)

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
