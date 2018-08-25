package main

import (
	_ "github.com/pustserg/scheduler/daemon"
	"github.com/pustserg/scheduler/tasks"
	"testing"
)

func TestStartDaemon(t *testing.T) {
	repo := tasks.NewRepository()
	startDaemon(1, repo)
}
