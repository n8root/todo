package main

import (
	"log"
	"net/http"
	"todo/internal/database"
	"todo/internal/handlers"
)

func main() {
	db, err := database.Connect("dbname=tasksdb user=tasksuser password=taskspass port=5000 sslmode=disable")

	if err != nil {
		log.Fatalf("connection to database failed %v", err)
	}

	storage := database.NewTasksStore(db)
	handlers := handlers.NewHandlers(storage)

	http.HandleFunc("GET /tasks", handlers.GetAllTasks)
	http.HandleFunc("GET /tasks/{id}", handlers.GetTask)
	http.HandleFunc("POST /tasks", handlers.CreateTask)
	http.HandleFunc("PATCH /tasks/{id}", handlers.UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", handlers.DeleteTask)

	http.ListenAndServe(":8080", nil)
}
