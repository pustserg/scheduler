package main

import (
	_ "github.com/pustserg/scheduler/daemon"
	"testing"
)

func TestStartDaemon(t *testing.T) {
	startDaemon(1)
}
