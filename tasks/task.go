package tasks

import (
	"github.com/gorhill/cronexpr"
	"log"
	"time"
)

var (
	// AvailableActions array with valid actions
	AvailableActions = [...]string{
		"send_telegram_message",
		"send_rabbit_message",
	}
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

// Process processes task
func (task Task) Process(repo TaskRepository) {
	now := time.Now()
	if task.PerformAt.Before(now) {
		log.Println("task perform is set to", task.PerformAt, "process task", task.ID)
		err := repo.UpdateTaskPerformAtTime(&task)
		if err != nil {
			panic(err)
		}
	}
}

// Validate task before save to db
func (task Task) Validate() error {
	actionErr := validateAction(task)
	if actionErr != nil {
		return actionErr
	}
	scheduleErr := validateSchedule(task)
	if scheduleErr != nil {
		return scheduleErr
	}
	return nil
}

func validateAction(task Task) error {
	if task.Action == "" {
		return ErrTaskActionShouldBeSet
	}
	found := false
	for _, action := range AvailableActions {
		if action == task.Action {
			found = true
			break
		}
	}
	if !found {
		return ErrTaskActionShouldBeInAvailableActions
	}
	return nil
}

func validateSchedule(task Task) error {
	_, err := cronexpr.Parse(task.Schedule)
	if err != nil {
		return ErrMalformedSchedule
	}
	return nil
}
