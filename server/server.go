package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pustserg/scheduler/tasks"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
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
		router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
			ShowTask(w, r, repo)
		}).Methods("GET")
		// Handle update task
		router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
			UpdateTask(w, r, repo)
		}).Methods("PATCH")
		// Handle delete task
		router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
			DeleteTask(w, r, repo)
		}).Methods("DELETE")
		server := Server{Status: "initiated", router: router, repo: repo}
		log.Println("Server initiated once")
		serverInstance = &server
	})
	return serverInstance
}

// Start starts server
func (server *Server) Start() {
	c := cors.New(
		cors.Options{
			Debug:          true,
			AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		},
	)
	handler := c.Handler(server.router)
	log.Fatal(http.ListenAndServe(":8000", handler))
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

// ShowTask render one task
func ShowTask(w http.ResponseWriter, r *http.Request, repo *tasks.TaskRepository) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	task, err := repo.GetTask(taskID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(apiIndexTask(task))
}

// UpdateTask updates task in db
func UpdateTask(w http.ResponseWriter, r *http.Request, repo *tasks.TaskRepository) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	var updateParams map[string]string
	err = decoder.Decode(&updateParams)
	if err != nil {
		panic(err)
	}
	task, err := repo.UpdateTask(taskID, updateParams)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(apiIndexTask(task))
}

// DeleteTask deletes task from db
func DeleteTask(w http.ResponseWriter, r *http.Request, repo *tasks.TaskRepository) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	task, err := repo.DeleteTask(taskID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(apiIndexTask(task))
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
