package handlers

import (
	"api/internal/actors"
	"encoding/json"
	"log"
	"net/http"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdRouteParam(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	task, err := extractBody(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	task.ID = id

	response := make(chan actors.ResponseStruct)

	actors.RequestQueue <- actors.RequestStruct{
		Action:   "UpdateTask",
		Task:     task,
		Response: response,
	}

	taskResponse := <-response

	if taskResponse.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	data, err := json.MarshalIndent(taskResponse.Data, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
