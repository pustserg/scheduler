package tasks

import (
	"github.com/gorhill/cronexpr"
	"testing"
	"time"
)

func TestNextExecutionTime(t *testing.T) {
	task := Task{Action: "action", Schedule: "*/2 * * * *"}
	expectedTime := cronexpr.MustParse(task.Schedule).Next(time.Now())
	gotTime := task.NextExecutionTime()
	if expectedTime != gotTime {
		t.Error("times does not match")
	}
}
