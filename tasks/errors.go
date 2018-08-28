package tasks

import "errors"

// Errors for tasks

var (
	// ErrTaskActionShouldBeSet is returned when no Action found in task
	ErrTaskActionShouldBeSet = errors.New("You must specify action for task")

	// ErrTaskActionShouldBeInAvailableActions is returned when try to create task with invalid Action
	ErrTaskActionShouldBeInAvailableActions = errors.New("Invalid Action for task")

	// ErrMalformedSchedule is returned when task Schedule is malformed
	ErrMalformedSchedule = errors.New("Malformed Schedule")
)
