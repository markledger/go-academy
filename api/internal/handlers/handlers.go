package handlers

import (
	"api/internal/filestore"
	"api/internal/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Data []models.Task
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
	taskId = taskList[len(taskList)-1].Id
	jsonTask, _ := json.Marshal(person1) fmt.Println(string(jsonResult)) }
	//taskList = append(taskList, models.Task{Name: task.Name})
	//err = filestore.WriteFile(taskList)
	//if err != nil {
	//	log.Fatal(err)
	//}


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

//func ListAllTasks(w http.ResponseWriter, r *http.Request) {
//
//	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	taskResponse := &jsonResponse{
//		Data: taskList,
//	}
//
//	out, err := json.MarshalIndent(taskResponse, "", "     ")
//	if err != nil {
//		log.Println(err)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(out)
//}
//
//func GetTask(w http.ResponseWriter, r *http.Request) {
//	id, err := extractIdRouteParam(w, r)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var taskResponse *jsonResponse
//	for i, task := range taskList {
//		if i+1 != id {
//			taskResponse = &jsonResponse{
//				Data: []models.Task{task},
//			}
//			break
//		}
//	}
//
//	out, err := json.MarshalIndent(taskResponse, "", "     ")
//	if err != nil {
//		log.Println(err)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(out)
//}
//
//func DeleteTask(w http.ResponseWriter, r *http.Request) {
//	id, err := extractIdRouteParam(w, r)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	taskList = append(taskList[:id-1], taskList[id:]...)
//
//	err = filestore.WriteFile(taskList)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusNoContent)
//
//}
//
//func UpdateTask(w http.ResponseWriter, r *http.Request) {
//
//	id, err := extractIdRouteParam(w, r)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	task, err := extractBody(w, r)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	taskList[id-1] = task
//	taskResponse := &jsonResponse{
//		Data: []models.Task{task},
//	}
//
//	err = filestore.WriteFile(taskList)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	out, err := json.MarshalIndent(taskResponse, "", "     ")
//	if err != nil {
//		log.Println(err)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(out)
//}
//
//func extractBody(w http.ResponseWriter, r *http.Request) (models.Task, error) {
//	decoder := json.NewDecoder(r.Body)
//	var task models.Task
//	err := decoder.Decode(&task)
//	if err != nil {
//		return models.Task{}, err
//	}
//	return models.Task{Name: task.Name}, nil
//}
//
//func extractIdRouteParam(w http.ResponseWriter, r *http.Request) (int, error) {
//	idString := chi.URLParam(r, "id")
//	return strconv.Atoi(idString)
//}
