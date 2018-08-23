package tasks

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"time"
)

type TaskRepository struct {
	db *storm.DB
}

func NewRepository() *TaskRepository {
	db, err := initDb()
	if err != nil {
		panic(err)
	}
	repo := TaskRepository{db: db}
	return &repo
}

// GetTasksForHandle return tasks which must be handled
func (repo *TaskRepository) GetTasksForHandle() []Task {
	var tasks []Task
	now := time.Now()
	err := repo.db.Select(q.Lte("PerformAt", now)).Find(&tasks)
	if err != nil && err != storm.ErrNotFound {
		panic(err)
	}
	return tasks
}

// UpdateTaskPerformAtTime should set performAt time if absent or future
func (repo TaskRepository) UpdateTaskPerformAtTime(task *Task) error {
	nextTime := task.NextExecutionTime()
	err := repo.db.UpdateField(&Task{ID: task.ID}, "PerformAt", nextTime)
	return err
}

func initDb() (*storm.DB, error) {
	db, err := storm.Open("tasks.db")
	if err != nil {
		return nil, err
	}
	return db, err
}

func addInitialTasks(db *storm.DB) {
	task1 := Task{Action: "action", Schedule: "*/2 * * * *"}
	db.Save(&task1)
	task2 := Task{Action: "action", Schedule: "*/1 * * * *"}
	db.Save(&task2)
}
