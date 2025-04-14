package handlers

import (
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
	var t models.Task
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	t.Id = 3
	taskResponse := &jsonResponse{
		Data: []models.Task{t},
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

	task1 := models.Task{Id: 1, Name: "Make it so"}
	task2 := models.Task{Id: 2, Name: "Destroy The Dominion"}

	taskResponse := &jsonResponse{
		Data: []models.Task{task1, task2},
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
