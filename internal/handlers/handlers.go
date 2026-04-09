package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/internal/database"
)

type Handlers struct {
	store *database.TasksStore
}

func NewHandlers(store *database.TasksStore) *Handlers {
	return &Handlers{store: store}
}

func response(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(payload)
}

func responseWithError(w http.ResponseWriter, statusCode int, message string) {
	response(w, statusCode, map[string]string{"error": message})
}

func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.GetAll()

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Error fetch")
		return
	}

	response(w, http.StatusOK, tasks)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Incorrect ID")
		return
	}

	task, err := h.store.GetById(id)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response(w, http.StatusOK, task)
}

// func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
// 	json.Unmarshal()
// }
