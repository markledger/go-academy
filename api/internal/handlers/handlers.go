package handlers

import (
	"api/internal/filestore"
	"api/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Data []models.Task
}
type Request struct {
	Action   string
	Request  *http.Request
	Response chan Response
}

type Response struct {
	Data  []models.Task
	Error error
}

var RequestQueue = make(chan Request)

func StartActor() {

	go func() {
		for {
			select {
			case request := <-RequestQueue:

				switch request.Action {
				case "CreateTask":
					var taskResponse []models.Task
					task, err := extractBody(request.Request)

					taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
					if err != nil {
						log.Println(err)
						response := Response{Data: taskResponse, Error: err}
						request.Response <- response
						return
					}

					task.ID = taskList[len(taskList)-1].ID + 1
					taskList = append(taskList, task)
					err = filestore.WriteFile(taskList)

					if err != nil {
						log.Println(err)
						response := Response{Data: taskResponse, Error: err}
						request.Response <- response
						return
					}
					taskResponse = append(taskResponse, task)
					response := Response{Data: taskResponse, Error: nil}
					request.Response <- response

				case "GetTask":
					//id, err := extractIdRouteParam(request.Request)
					//taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
					//if err != nil {
					//	log.Fatal(err)
					//}
					//var selectedTask []models.Task
					//for _, task := range taskList {
					//	if task.ID == id {
					//		selectedTask = append(selectedTask, task)
					//		break
					//	}
					//}
					//request.Response <- selectedTask

				case "ListAllTasks":
					//todos, err := filestore.ParseFileToSlice(filestore.FilePath)
					//if err != nil {
					//	log.Fatal("error loading todos")
					//}
					//
					//request.Response <- todos
				}
			}
		}
	}()
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	response := make(chan Response)

	RequestQueue <- Request{
		Action:   "CreateTask",
		Request:  r,
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
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func ListAllTasks(w http.ResponseWriter, r *http.Request) {
	//
	//response := make(chan []models.Task)
	//
	//RequestQueue <- Request{
	//	Action:   "ListAllTasks",
	//	Request:  r,
	//	Response: response,
	//}
	//
	//responseData := <-response
	//
	//data, err := json.MarshalIndent(responseData, "", "     ")
	//if err != nil {
	//	log.Println(err)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(data)
}

func GetTask(w http.ResponseWriter, r *http.Request) {

	//response := make(chan []models.Task)
	//
	//RequestQueue <- Request{
	//	Action:   "GetTask",
	//	Request:  r,
	//	Response: response,
	//}
	//responseData := <-response
	//
	//if len(responseData) != 1 {
	//	w.WriteHeader(http.StatusNoContent)
	//	return
	//}

	//data, err := json.MarshalIndent(responseData, "", "     ")
	//if err != nil {
	//	log.Println(err)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(data)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdRouteParam(r)
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

	id, err := extractIdRouteParam(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	patchedTask, err := extractBody(r)
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
