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
	cases := []struct {
		task     Task
		valid    bool
		err      error
		errorMsg string
	}{
		{
			task:     Task{Action: "", Schedule: schedule},
			valid:    false,
			err:      ErrTaskActionShouldBeSet,
			errorMsg: "Incorrect Error for empty action",
		},
		{
			task:     Task{Action: "not in list", Schedule: schedule},
			valid:    false,
			err:      ErrTaskActionShouldBeInAvailableActions,
			errorMsg: "Incorrect Error for not in list action",
		},
		{
			task:     Task{Action: AvailableActions[0], Schedule: schedule},
			valid:    true,
			err:      nil,
			errorMsg: "Valid task is not valid",
		},
		{
			task:     Task{Action: AvailableActions[0], Schedule: "malformed"},
			valid:    false,
			err:      ErrMalformedSchedule,
			errorMsg: "Incorrect Errr for malformed schedule",
		},
	}

	for _, testCase := range cases {
		e := testCase.task.Validate()
		if testCase.valid && e != nil {
			t.Error("Valid task returns error")
		}
		if !testCase.valid && e == nil {
			t.Error("Invalid task did not return error")
		}
		if e != testCase.err {
			t.Error(testCase.errorMsg)
		}
	}
}
