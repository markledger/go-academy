package actors

import (
	"api/internal/filestore"
	"api/internal/models"
	"log"
)

func DeleteTask(task models.Task) ResponseStruct {
	var updatedTaskList []models.Task
	taskList, err := filestore.ParseFileToSlice(filestore.FilePath)
	if err != nil {
		return ResponseStruct{Data: []models.Task{task}, Error: err}
	}

	for _, comparatorTask := range taskList {
		if comparatorTask.ID == task.ID {
			continue
		}
		updatedTaskList = append(updatedTaskList, comparatorTask)
	}
	log.Println(updatedTaskList)
	err = filestore.WriteFile(updatedTaskList)
	if err != nil {
		return ResponseStruct{Data: []models.Task{task}, Error: err}
	}

	return ResponseStruct{[]models.Task{}, nil}
}
