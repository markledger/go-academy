package handlers

import (
	"api/internal/actors"
	"api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func ListAllTasksHandler(w http.ResponseWriter, r *http.Request) {

	response := make(chan actors.ResponseStruct)

	actors.RequestQueue <- actors.RequestStruct{
		Action:   "ListAllTasks",
		Task:     models.Task{},
		Response: response,
	}

	responseData := <-response

	data, err := json.MarshalIndent(responseData.Data, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
