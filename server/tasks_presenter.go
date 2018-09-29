package server

import (
	"github.com/pustserg/scheduler/tasks"
)

type apiIndexTaskItem struct {
	ID        int    `json:"id"`
	Action    string `json:"action"`
	Schedule  string `json:"schedule"`
	Message   string `json:"message"`
	PerformAt int    `json:"perform_at"`
}

func apiIndexTask(task tasks.Task) apiIndexTaskItem {
	return apiIndexTaskItem{
		ID:        task.ID,
		Action:    task.Action,
		Schedule:  task.Schedule,
		Message:   task.Message,
		PerformAt: int(task.PerformAt.Unix()),
	}
}

func apiIndexTasks(tasks []tasks.Task) []apiIndexTaskItem {
	presented := make([]apiIndexTaskItem, 0, len(tasks))
	for _, task := range tasks {
		presented = append(presented, apiIndexTask(task))
	}
	return presented
}
