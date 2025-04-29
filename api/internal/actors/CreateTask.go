package actors

import (
	"api/internal/filestore"
	"api/internal/models"
	"log"
)

func CreateTaskActor(task models.Task) ResponseStruct {
	var taskResponse []models.Task

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Println(err)
		return ResponseStruct{Data: taskResponse, Error: err}
	}

	task.ID = taskList[len(taskList)-1].ID + 1
	taskList = append(taskList, task)
	err = filestore.WriteFile(taskList)

	if err != nil {
		log.Println(err)
		return ResponseStruct{Data: taskResponse, Error: err}
	}
	taskResponse = append(taskResponse, task)
	return ResponseStruct{Data: taskResponse, Error: nil}
}
