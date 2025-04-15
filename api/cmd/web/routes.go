package main

import (
	"api/internal/handlers"
	"net/http"
)

func routes(mux *http.ServeMux) http.Handler {

	mux.HandleFunc("GET /tasks", handlers.ListAllTasks)
	mux.HandleFunc("POST /task", handlers.CreateTask)
	mux.HandleFunc("GET /task/{id}", handlers.GetTask)
	mux.HandleFunc("DELETE /task/{id}", handlers.DeleteTask)
	mux.HandleFunc("PATCH /task/{id}", handlers.UpdateTask)

	return mux
}
