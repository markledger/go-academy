package actors

import (
	"api/internal/filestore"
)

func ListAllTasks() ResponseStruct {
	tasks, err := filestore.ParseFileToSlice(filestore.FilePath)
	return ResponseStruct{Data: tasks, Error: err}
}
