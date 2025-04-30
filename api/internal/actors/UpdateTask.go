package actors

import (
	"api/internal/filestore"
	"api/internal/models"
	"log"
)

func UpdateTask(patchedTask models.Task) ResponseStruct {

	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	var updatedTaskList []models.Task
	var responseData models.Task
	for _, task := range taskList {
		if task.ID == patchedTask.ID {
			task.Name = patchedTask.Name
			responseData = task
		}
		updatedTaskList = append(updatedTaskList, task)
	}

	err = filestore.WriteFile(updatedTaskList)
	if err != nil {
		return ResponseStruct{Data: []models.Task{}, Error: err}
	}

	return ResponseStruct{Data: []models.Task{responseData}, Error: nil}

}
