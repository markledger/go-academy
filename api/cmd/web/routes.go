package main

import (
	"api/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handlers.ListAllTasks)
	mux.HandleFunc("POST /task", handlers.CreateTask)
	mux.HandleFunc("GET /task/{id}", handlers.GetTask)
	mux.HandleFunc("DELETE /task", handlers.DeleteTask)
	mux.HandleFunc("PATCH /task", handlers.UpdateTask)

	return mux
}
