package actors

import (
	"api/internal/filestore"
	"api/internal/models"
	"errors"
)

func GetTask(requestedTask models.Task) ResponseStruct {

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		return ResponseStruct{Data: []models.Task{}, Error: err}
	}
	var selectedTask models.Task
	for _, task := range taskList {
		if task.ID == requestedTask.ID {
			selectedTask = task
			break
		}
	}

	if selectedTask.ID != requestedTask.ID {
		return ResponseStruct{Data: []models.Task{}, Error: errors.New("Task not found")}
	}
	return ResponseStruct{Data: []models.Task{selectedTask}, Error: nil}
}
