package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo/internal/database"
	"todo/internal/models"
)

type Handlers struct {
	store *database.TasksStore
}

func NewHandlers(store *database.TasksStore) *Handlers {
	return &Handlers{store: store}
}

func response(w http.ResponseWriter, statusCode int, payload any) {
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
		responseWithError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("Fetch failed %v", err).Error(),
		)
		return
	}

	response(w, http.StatusOK, tasks)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = h.store.Delete(id)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response(w, http.StatusNoContent, nil)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	task, err := h.store.GetById(id)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response(w, http.StatusOK, task)
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var createForm models.CreateTaskForm

	err := json.NewDecoder(r.Body).Decode(&createForm)

	if err != nil {
		responseWithError(
			w,
			http.StatusBadRequest,
			fmt.Errorf("Bad request %v", err).Error(),
		)
		return
	}

	task, err := h.store.Create(createForm)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, task)
}

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateForm models.UpdateTaskForm

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updateForm)

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	task, err := h.store.Update(id, updateForm)

	if err != nil {
		responseWithError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("Internal server error %v", err).Error(),
		)
		return
	}

	response(w, http.StatusOK, task)
}
