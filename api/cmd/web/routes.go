package main

import (
	"api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/tasks", handlers.ListAllTasks)
	mux.Post("/task", handlers.CreateTask)

	mux.Get("/task/{id}", handlers.GetTask)
	mux.Patch("/task/{id}", handlers.UpdateTask)
	mux.Delete("/task/{id}", handlers.DeleteTask)

	return mux
}
