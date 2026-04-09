package database

import (
	"database/sql"
	"fmt"
	"time"
	"todo/internal/models"

	"github.com/jmoiron/sqlx"
)

type TasksStore struct {
	db *sqlx.DB
}

func NewTasksStore(db *sqlx.DB) *TasksStore {
	return &TasksStore{db: db}
}

func (s *TasksStore) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	query := `
		SELECT (id, title, description, completed, idcreated_at, updated_at) 
			FROM tasks
			ORDER BY created_at DESC
	`

	if err := s.db.Select(&tasks, query); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TasksStore) GetById(id int) (*models.Task, error) {
	var task models.Task

	query := `
		SELECT (id, title, description, completed, idcreated_at, updated_at) 
			FROM tasks
				WHERE id = $1
	`

	err := s.db.Get(&task, query, id)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task with id %d not found", id)
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *TasksStore) Create(form models.CreateTaskForm) (*models.Task, error) {
	var task models.Task

	now := time.Now()

	query := `
		INSERT INTO tasks (title, description, completed, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
				RETURNING id, title, description, completed, created_at, updated_at
	`

	err := s.db.QueryRowx(query, form.Title, form.Description, form.Completed, now, now).StructScan(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *TasksStore) Update(id int, form models.UpdateTaskForm) (*models.Task, error) {
	task, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	if form.Title != nil {
		task.Title = *form.Title
	}

	if form.Description != nil {
		task.Description = *form.Description
	}

	if form.Completed != nil {
		task.Completed = *form.Completed
	}

	task.UpdatedAt = time.Now()

	query := `
		UPDATE tasks
			SET
				title = $1
				description = $2
				completed = $3
				updated_at = $4
			RETURNING id, title, description, completed, created_at, updated_at
			
		
	`

	var updatedTask models.Task

	err = s.db.QueryRowx(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.UpdatedAt,
		id,
	).StructScan(&updatedTask)

	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

func (s *TasksStore) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id=$1`

	_, err := s.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
