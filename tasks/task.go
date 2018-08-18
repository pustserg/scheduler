package tasks

import (
	"github.com/gorhill/cronexpr"
	"time"
)

// Task is a struct with action and schedule
type Task struct {
	Action   string
	Schedule string
}

// NextExecutionTime returns time parsed from cron schedule
func (task *Task) NextExecutionTime() time.Time {
	time := cronexpr.MustParse(task.Schedule).Next(time.Now())
	return time
}
