package main

import (
	"api/internal/filestore"
	"api/internal/handlers"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Use "http.FileServer" to serve a static page to a new "/about" endpoint
//
// Use "html/template" to serve a dynamic page containing a list of all to-do items to a new "/list" endpoint
func routes(mux *http.ServeMux) http.Handler {

	mux.HandleFunc("GET /list-tasks", func(w http.ResponseWriter, r *http.Request) {
		log.Println("YO")

		template, err := template.New("tasks").ParseFiles("./templates/tasks.page.tmpl")
		if err != nil {
			log.Fatal("cannot fwr tmpalte")
		}
		taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
		if err != nil {
			log.Fatal("oh dear")
		}
		for _, task := range taskList {
			log.Println(task.Name)
		}

		err = template.Execute(os.Stdout, taskList)
		if err != nil {
			log.Fatal("WTF")
		}
		//_ = myCache["list-tasks"].Execute(buf, taskList)
		//
		//_, err = buf.WriteTo(w)
		//if err != nil {
		//	fmt.Println("error writing template to browser", err)
		//}

	})
	mux.HandleFunc("GET /tasks", handlers.ListAllTasks)
	mux.HandleFunc("POST /task", handlers.CreateTask)
	mux.HandleFunc("GET /task/{id}", handlers.GetTask)
	mux.HandleFunc("DELETE /task/{id}", handlers.DeleteTask)
	mux.HandleFunc("PATCH /task/{id}", handlers.UpdateTask)

	mux.HandleFunc("GET /about/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/about.html")
	})
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return mux
}
