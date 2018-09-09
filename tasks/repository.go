package tasks

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"time"
)

// TaskRepository data source for tasks
type TaskRepository struct {
	db *storm.DB
}

// NewRepository returns pointer to TaskRepository
func NewRepository(dbFileName string) *TaskRepository {
	db, err := initDb(dbFileName)
	if err != nil {
		panic(err)
	}
	repo := TaskRepository{db: db}
	// addInitialTasks(repo)
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

// GetAllTasks return all tasks from db
func (repo *TaskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := repo.db.All(&tasks)
	return tasks, err
}

// AddTask creates task in db, returns task and db error
func (repo *TaskRepository) AddTask(params map[string]string) (Task, error) {
	task := Task{Action: params["action"], Schedule: params["schedule"]}
	err := task.Validate()
	if err != nil {
		return task, err
	}
	task.PerformAt = task.NextExecutionTime()
	repo.db.Save(&task)
	return task, nil
}

// UpdateTaskPerformAtTime should set performAt time if absent or future
func (repo TaskRepository) UpdateTaskPerformAtTime(task *Task) error {
	nextTime := task.NextExecutionTime()
	err := repo.db.UpdateField(&Task{ID: task.ID}, "PerformAt", nextTime)
	return err
}

// GetTask returns task with given id
func (repo TaskRepository) GetTask(id int) (*Task, error) {
	var task Task
	err := repo.db.One("ID", id, &task)
	return &task, err
}

// UpdateTask updates task with given params
func (repo TaskRepository) UpdateTask(id int, params map[string]string) (*Task, error) {
	taskForUpdate := Task{ID: id}
	present := false
	_, present = params["schedule"]
	if present {
		taskForUpdate.Schedule = params["schedule"]
	}
	present = false
	_, present = params["action"]
	if present {
		taskForUpdate.Action = params["action"]
	}
	err := repo.db.Update(&taskForUpdate)
	if err != nil {
		return nil, err
	}
	task, err := repo.GetTask(id)
	return task, err
}

//DeleteTask deletes task with given id
func (repo TaskRepository) DeleteTask(id int) (*Task, error) {
	task, err := repo.GetTask(id)
	if err != nil {
		return nil, err
	}
	err = repo.db.DeleteStruct(task)
	return task, err
}

func initDb(dbFileName string) (*storm.DB, error) {
	db, err := storm.Open(dbFileName)
	if err != nil {
		return nil, err
	}
	return db, err
}

func addInitialTasks(repo TaskRepository) {
	task := map[string]string{"Action": AvailableActions[0], "Schedule": "*/2 * * * *"}
	_, err := repo.AddTask(task)
	if err != nil {
		panic(err)
	}
}
