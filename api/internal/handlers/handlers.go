package handlers

import (
	"api/internal/filestore"
	"api/internal/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Data []models.Task
}
type Request struct {
	Ctx      context.Context
	Action   string
	Payload  any
	Response chan any
}

var RequestQueue = make(chan Request)
var ShutdownChan = make(chan struct{})

func StartActor() {

	go func() {
		for {
			select {
			case request := <-RequestQueue:
				log.Println("The request action:", request.Action)
				switch request.Action {

				case "GetTask":
					id := request.Payload
					taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
					if err != nil {
						log.Fatal(err)
					}
					var selectedTask models.Task
					for _, task := range taskList {
						if task.ID == id {
							selectedTask = task
							break
						}
					}
					request.Response <- selectedTask

				case "ListAllTasks":
					todos, err := filestore.ParseFileToSlice(filestore.FilePath)
					if err != nil {
						log.Fatal("error loading todos")
					}

					request.Response <- todos
				}
			}
		}
	}()
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := extractBody(w, r)
	if err != nil {
		log.Printf("err converting ID to integer: %+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	task.ID = taskList[len(taskList)-1].ID + 1
	taskList = append(taskList, task)
	err = filestore.WriteFile(taskList)
	if err != nil {
		log.Fatal(err)
	}

	taskResponse := &jsonResponse{
		Data: []models.Task{task},
	}

	out, err := json.MarshalIndent(taskResponse, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func ListAllTasks(w http.ResponseWriter, r *http.Request) {

	response := make(chan any)

	RequestQueue <- Request{
		Action:   "ListAllTasks",
		Payload:  nil,
		Response: response,
	}

	out := (<-response).([]models.Task)

	responseData, err := json.MarshalIndent(out, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdRouteParam(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := make(chan any)

	RequestQueue <- Request{
		Action:   "GetTask",
		Payload:  id,
		Response: response,
	}
	task := (<-response).(models.Task)

	if task.ID != id {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	responseData, err := json.MarshalIndent(task, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdRouteParam(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	var updatedTaskList []models.Task

	for _, task := range taskList {
		if task.ID == id {
			continue
		}
		updatedTaskList = append(updatedTaskList, task)
	}

	err = filestore.WriteFile(updatedTaskList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	id, err := extractIdRouteParam(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	patchedTask, err := extractBody(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	var updatedTaskList []models.Task
	var responseData models.Task
	for _, task := range taskList {
		if task.ID == id {
			task.Name = patchedTask.Name
			responseData = task
		}
		updatedTaskList = append(updatedTaskList, task)
	}

	err = filestore.WriteFile(updatedTaskList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	taskResponse := &jsonResponse{
		Data: []models.Task{responseData},
	}

	out, err := json.MarshalIndent(taskResponse, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
func extractBody(w http.ResponseWriter, r *http.Request) (models.Task, error) {
	decoder := json.NewDecoder(r.Body)
	var task models.Task
	error := decoder.Decode(&task)
	if error != nil {
		return models.Task{}, error
	}
	return task, nil
}

func extractIdRouteParam(w http.ResponseWriter, r *http.Request) (int, error) {
	idString := r.PathValue("id")
	return strconv.Atoi(idString)
}
