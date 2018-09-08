package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pustserg/scheduler/tasks"
	"log"
	"net/http"
	"sync"
)

// Server struct wrapper for http server
type Server struct {
	router *mux.Router
	Status string
	repo   *tasks.TaskRepository
}

var serverInstance *Server
var once sync.Once

// NewServer returns Server
func NewServer(env string, repo *tasks.TaskRepository) *Server {
	once.Do(func() {
		router := mux.NewRouter()
		// Handle index tasks
		router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
			GetTasks(w, r, repo)
		}).Methods("GET")
		// Handle create task
		router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
			CreateTask(w, r, repo)
		}).Methods("POST")
		// TODO: add this methods
		// Handle show task
		// Handle update task
		// Handle delete task
		server := Server{Status: "initiated", router: router, repo: repo}
		log.Println("Server initiated once")
		serverInstance = &server
	})
	return serverInstance
}

// Start starts server
func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(":8000", server.router))
}

// GetTasks index for tasks
func GetTasks(w http.ResponseWriter, r *http.Request, repo *tasks.TaskRepository) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := repo.GetAllTasks()
	if err != nil {
		panic(err)
	}
	presentedTasks := apiIndexTasks(tasks)
	json.NewEncoder(w).Encode(presentedTasks)
}

// CreateTask creates task
func CreateTask(w http.ResponseWriter, r *http.Request, repo *tasks.TaskRepository) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}
	task, err := repo.AddTask(params)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(apiIndexTask(task))
}
