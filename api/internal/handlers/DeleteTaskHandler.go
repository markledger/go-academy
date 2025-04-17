package handlers

import (
	"api/internal/actors"
	"api/internal/models"
	"net/http"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdRouteParam(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := make(chan actors.ResponseStruct)

	actors.RequestQueue <- actors.RequestStruct{
		Action:   "DeleteTask",
		Task:     models.Task{ID: id},
		Response: response,
	}

	responseData := <-response

	if responseData.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
