package daemon

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	fmt.Println("testing daemon.Start")
	daemonInstance := Daemon{}
	err := daemonInstance.Start()
	if err != nil {
		t.Error("daemon start returns error")
	}

	if daemonInstance.State != "started" {
		t.Error("Expected daemon to be started")
	}
}

func TestStop(t *testing.T) {
	daemonInstance := Daemon{}
	err := daemonInstance.Stop()
	if err != nil {
		t.Error("daemon stop returns error")
	}

	if daemonInstance.State != "stopped" {
		t.Error("Expected daemon to be stopped")
	}
}
