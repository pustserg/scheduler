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

func TestValidate(t *testing.T) {
	schedule := "*/2 * * * * *"
	task1 := Task{Action: "", Schedule: schedule}
	err1 := task1.Validate()
	if err1 == nil {
		t.Error("Empty action is valid")
	}
	if err1 != ErrTaskActionShouldBeSet {
		t.Error("Incorrect Error for empty acton")
	}

	task2 := Task{Action: "not in list", Schedule: schedule}
	err2 := task2.Validate()
	if err2 == nil {
		t.Error("Not in list action is valid")
	}
	if err2 != ErrTaskActionShouldBeInAvailableActions {
		t.Error("Incorrect Error for not in list action")
	}

	task3 := Task{Action: AvailableActions[0], Schedule: schedule}
	err3 := task3.Validate()
	if err3 != nil {
		t.Error("Valid task is not valid")
	}

	task4 := Task{Action: AvailableActions[0], Schedule: "malformed"}
	err4 := task4.Validate()
	if err4 == nil {
		t.Error("Malformed schedule is valid")
	}

	if err4 != ErrMalformedSchedule {
		t.Error("Incorrect Errr for malformed schedule")
	}
}
