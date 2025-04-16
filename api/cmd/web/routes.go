package main

import (
	"api/internal/handlers"
	"net/http"
)

// Use "http.FileServer" to serve a static page to a new "/about" endpoint
//
// Use "html/template" to serve a dynamic page containing a list of all to-do items to a new "/list" endpoint
func routes(mux *http.ServeMux) http.Handler {

	mux.HandleFunc("GET /tasks", handlers.ListAllTasks)
	mux.HandleFunc("POST /task", handlers.CreateTask)
	mux.HandleFunc("GET /task/{id}", handlers.GetTask)
	mux.HandleFunc("DELETE /task/{id}", handlers.DeleteTask)
	mux.HandleFunc("PATCH /task/{id}", handlers.UpdateTask)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/about", http.StripPrefix("/static", fileServer))

	return mux
}
