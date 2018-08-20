package tasks

// GetTasksForHandle return tasks which must be handled
func GetTasksForHandle() []Task {
	task := Task{Action: "action", Schedule: "*/2 * * * *"}
	return []Task{task}
}
