package tasks

import (
	"os"
	"testing"
)

func TestAddTask(t *testing.T) {
	repo := NewRepository("test.db")
	defer os.Remove("test.db")

	// test success
	task := map[string]string{"action": AvailableActions[0], "schedule": "*/2 * * * *"}
	createdTask, err := repo.AddTask(task)

	if err != nil {
		t.Error("Add task returned error")
	}
	if createdTask.Action != task["action"] {
		t.Error("Add task created tasks with wrong Action")
	}
	if createdTask.Schedule != task["schedule"] {
		t.Error("Add task created task with wrong Schedule")
	}

	// test validation for invalid actions
	invalidTask := map[string]string{"action": "invalid action", "schedule": "* * * * * *"}
	_, elseErr := repo.AddTask(invalidTask)

	if elseErr == nil {
		t.Error("Add task did not returns error for invalid task")
	}
	if elseErr != ErrTaskActionShouldBeInAvailableActions {
		t.Error("Add task return incorrect error")
	}

	// test validation for empty action
	taskWithEmptyAction := map[string]string{"action": "", "schedule": "* * * * * *"}
	_, elseOneErr := repo.AddTask(taskWithEmptyAction)

	if elseOneErr == nil {
		t.Error("Add task did not returns error for invalid task")
	}
	if elseOneErr != ErrTaskActionShouldBeSet {
		t.Error("Add task return incorrect error")
	}
}

func TestGetAllTasks(t *testing.T) {
	repo := NewRepository("test.db")
	defer os.Remove("test.db")

	createdTask1, _ := repo.AddTask(map[string]string{"action": AvailableActions[0], "schedule": "*/2 * * * *"})
	createdTask2, _ := repo.AddTask(map[string]string{"action": AvailableActions[0], "schedule": "*/3 * * * *"})

	repoTasks, err := repo.GetAllTasks()
	if err != nil {
		t.Error(err)
	}

	if len(repoTasks) != 2 {
		t.Error("Got wrong tasks count")
	}

	found := false
	for _, task := range repoTasks {
		if task.Action == createdTask1.Action && task.Schedule == createdTask1.Schedule {
			found = true
		}
	}

	if !found {
		t.Error("task1 is not returned")
	}

	found = false
	for _, task := range repoTasks {
		if task.Action == createdTask2.Action && task.Schedule == createdTask2.Schedule {
			found = true
		}
	}

	if !found {
		t.Error("task2 is not returned")
	}
}

func TestGetTask(t *testing.T) {
	repo := NewRepository("test.db")
	defer os.Remove("test.db")
	task, _ := repo.AddTask(map[string]string{"action": AvailableActions[0], "schedule": "*/2 * * * *"})

	gotTask, err := repo.GetTask(task.ID)
	if err != nil {
		t.Error(err)
	}
	if *gotTask != task {
		t.Error("Got wrong task")
	}

	_, err = repo.GetTask(task.ID + 1)
	if err == nil {
		t.Error("Incorrect task ID as arg did not return error")
	}
}

func TestUpdateTask(t *testing.T) {
	repo := NewRepository("test.db")
	defer os.Remove("test.db")
	oldTask, _ := repo.AddTask(map[string]string{"action": AvailableActions[0], "schedule": "*/2 * * * *"})

	newTaskParams := map[string]string{"schedule": "*/3 * * * * * *"}
	updatedTask, err := repo.UpdateTask(oldTask.ID, newTaskParams)

	if err != nil {
		t.Error(err)
	}

	if updatedTask.Schedule != newTaskParams["schedule"] {
		t.Error("schedule is not updated")
	}

	if updatedTask.Action != oldTask.Action {
		t.Error("updated field which not present in params")
	}
}

func TestDeleteTask(t *testing.T) {
	repo := NewRepository("test.db")
	defer os.Remove("test.db")
	oldTask, _ := repo.AddTask(map[string]string{"action": AvailableActions[0], "schedule": "*/2 * * * *"})
	allTasksBeforeDelete, _ := repo.GetAllTasks()
	if len(allTasksBeforeDelete) != 1 {
		t.Error("Cannot create before tasks")
	}

	deletedTask, err := repo.DeleteTask(oldTask.ID)
	if err != nil {
		t.Error(err)
	}
	if deletedTask.ID != oldTask.ID {
		t.Error("deleted wrong task")
	}

	allTasksAfterDelete, err := repo.GetAllTasks()
	if len(allTasksAfterDelete) != 0 {
		t.Error("Task is not deleted")
	}
}
