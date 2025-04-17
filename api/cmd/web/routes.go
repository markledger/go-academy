package main

import (
	"api/internal/handlers"
	"net/http"
)

// Use "http.FileServer" to serve a static page to a new "/about" endpoint
//
// Use "html/template" to serve a dynamic page containing a list of all to-do items to a new "/list" endpoint
func routes(mux *http.ServeMux) http.Handler {

	mux.HandleFunc("GET /list-tasks", handlers.FrontendTasks)
	mux.HandleFunc("GET /api/tasks", handlers.ListAllTasksHandler)
	mux.HandleFunc("POST /api/task", handlers.CreateTaskHandler)
	mux.HandleFunc("GET /api/task/{id}", handlers.GetTaskHandler)
	mux.HandleFunc("DELETE /api/task/{id}", handlers.DeleteTaskHandler)
	mux.HandleFunc("PATCH /api/task/{id}", handlers.UpdateTaskHandler)

	mux.HandleFunc("GET /about/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/about.html")
	})
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return mux
}
