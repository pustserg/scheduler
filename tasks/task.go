package tasks

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

// Task is a struct with action and schedule
type Task struct {
	ID        int `storm:"id,increment"`
	Action    string
	Schedule  string
	PerformAt time.Time `storm:"index"`
}

// Processable is an interface.
type Processable interface {
	NextExecutionTime() time.Time
	Process(*TaskRepository)
}

// NextExecutionTime returns time parsed from cron schedule
func (task Task) NextExecutionTime() time.Time {
	time := cronexpr.MustParse(task.Schedule).Next(time.Now())
	return time
}

func (task Task) Process(repo TaskRepository) {
	now := time.Now()
	if task.PerformAt.Before(now) {
		err := repo.UpdateTaskPerformAtTime(&task)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("task perform at", task.PerformAt)
}
