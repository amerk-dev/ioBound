package handler

import (
	"encoding/json"
	"ioBound/internal/domain"
	"ioBound/internal/worker"
	"math/rand"
	"net/http"
	"time"
)

type Server struct {
	handler http.Handler
}

func NewServer() *Server {
	return &Server{handler: http.DefaultServeMux}
}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	newTask := domain.Task{
		ID:        rand.Int(),
		Status:    domain.StatusWait,
		CreatedAt: time.Now(),
	}
	go worker.TaskWorker(&newTask)
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		panic(err)
	}

}

func (s *Server) CheckStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(taskID))
}
