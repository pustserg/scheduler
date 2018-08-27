package tasks

import (
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	repo := NewRepository("test.db")
	defer os.Remove("test.db")

	// test success
	task := map[string]string{"Action": AvailableActions[0], "Schedule": "*/2 * * * *"}
	createdTask, err := repo.AddTask(task)

	if err != nil {
		t.Error("Add task returned error")
	}
	if createdTask.Action != task["Action"] {
		t.Error("Add task created tasks with wrong Action")
	}
	if createdTask.Schedule != task["Schedule"] {
		t.Error("Add task created task with wrong Schedule")
	}

	// test validation for invalid actions
	invalidTask := map[string]string{"Action": "invalid action", "Schedule": "* * * * * *"}
	_, elseErr := repo.AddTask(invalidTask)

	if elseErr == nil {
		t.Error("Add task did not returns error for invalid task")
	}
	if elseErr != ErrTaskActionShouldBeInAvailableActions {
		t.Error("Add task return incorrect error")
	}

	// test validation for empty action
	taskWithEmptyAction := map[string]string{"Action": "", "Schedule": "* * * * * *"}
	_, elseOneErr := repo.AddTask(taskWithEmptyAction)

	if elseOneErr == nil {
		t.Error("Add task did not returns error for invalid task")
	}
	if elseOneErr != ErrTaskActionShouldBeSet {
		t.Error("Add task return incorrect error")
	}
}
