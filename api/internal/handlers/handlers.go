package handlers

import (
	"api/internal/filestore"
	"api/internal/models"
	"encoding/json"
	"log"
	"net/http"
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
