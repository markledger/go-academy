package handlers

import (
	"api/internal/filestore"
	"api/internal/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type jsonResponse struct {
	Data []models.Task
}

type jsonResponsePlaceholder struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var task models.Task
	err := decoder.Decode(&task)
	if err != nil {
		panic(err)
	}

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	taskList = append(taskList, models.Task{Name: task.Name})
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

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	taskResponse := &jsonResponse{
		Data: taskList,
	}

	out, err := json.MarshalIndent(taskResponse, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Printf("err converting ID to integer: %+v\n", err)
		// handle error by returning 400 - bad request, or by redirecting
		// to an error page, or rendering an error
		return
	}

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	var taskResponse *jsonResponse
	for i, task := range taskList {
		if i+1 != id {
			taskResponse = &jsonResponse{
				Data: []models.Task{task},
			}
			break
		}
	}

	out, err := json.MarshalIndent(taskResponse, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponsePlaceholder{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponsePlaceholder{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
