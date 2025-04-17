package handlers

import (
	"api/internal/filestore"
	"html/template"
	"log"
	"net/http"
)

func FrontendTasks(w http.ResponseWriter, r *http.Request) {
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
		log.Println(task.ID, task.Name)

	}
	err = template.Execute(w, taskList)
	if err != nil {
		log.Fatal("that not good")
	}

}
