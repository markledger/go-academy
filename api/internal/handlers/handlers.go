package handlers

import (
	"api/internal/actors"
	"api/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Data []models.Task
}

func StartActor() {

	go func() {
		for {
			select {
			case request := <-actors.RequestQueue:
				switch request.Action {
				case "CreateTask":
					request.Response <- actors.CreateTask(request.Task)
				case "GetTask":
					request.Response <- actors.GetTask(request.Task)
				case "ListAllTasks":
					request.Response <- actors.ListAllTasks()
				case "DeleteTask":
					request.Response <- actors.DeleteTask(request.Task)
				case "UpdateTask":
					request.Response <- actors.UpdateTask(request.Task)
				}
			}
		}
	}()
}

func extractBody(r *http.Request) (models.Task, error) {
	decoder := json.NewDecoder(r.Body)
	var task models.Task
	error := decoder.Decode(&task)
	if error != nil {
		return models.Task{}, error
	}
	return task, nil
}

func extractIdRouteParam(r *http.Request) (int, error) {
	idString := r.PathValue("id")
	return strconv.Atoi(idString)
}
