package main

import (
	_ "github.com/pustserg/scheduler/daemon"
	"testing"
)

func TestStartDaemon(t *testing.T) {
	err := startDaemon(1)
	if err != nil {
		t.Error("It returns error")
	}
}
