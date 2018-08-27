package main

import (
	_ "github.com/pustserg/scheduler/daemon"
	"github.com/pustserg/scheduler/tasks"
	"os"
	"testing"
)

func TestStartDaemon(t *testing.T) {
	repo := tasks.NewRepository("test.db")
	defer os.Remove("test.db")
	startDaemon(1, repo)
}
