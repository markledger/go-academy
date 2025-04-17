package handlers

import (
	"api/internal/actors"
	"encoding/json"
	"log"
	"net/http"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := extractBody(r)

	// Oliver... Is it safe to respond here if there
	// is an error? I would also like to include validation
	// at this point so I'm not putting duff data on the channel
	// but this, I fear is removing responsibility from the actor
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	response := make(chan actors.ResponseStruct)

	actors.RequestQueue <- actors.RequestStruct{
		Action:   "CreateTask",
		Task:     task,
		Response: response,
	}

	taskResponse := <-response

	if taskResponse.Error != nil {
		log.Println(taskResponse.Error)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	data, err := json.MarshalIndent(taskResponse.Data, "", "     ")
	if err != nil {
		log.Println(err)
		//handle this error for the user
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
