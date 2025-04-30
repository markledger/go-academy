package handlers

import (
	"api/internal/actors"
	"api/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdRouteParam(r)
	response := make(chan actors.ResponseStruct)

	actors.RequestQueue <- actors.RequestStruct{
		Action:   "GetTask",
		Task:     models.Task{ID: id},
		Response: response,
	}
	responseData := <-response

	if responseData.Error != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	data, err := json.MarshalIndent(responseData.Data, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
