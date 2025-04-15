package filestore

import (
	"api/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

const FilePath = "database.json"

/*
*
Parse the file located at filePath and split on new lines storing each line as a
a task in the data slice.
*/
func ParseFileToSlice(filePath string) ([]models.Task, error) {
	var data []models.Task
	file, err := os.Open(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			err = CreateFile()
			if err != nil {
				return data, err
			}
		}
		return data, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

/*
Create a file at the location provided in filePath
*/
func CreateFile() error {
	_, err := os.Stat(FilePath)
	if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(FilePath)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
Delete the file if a file exists at the filePath location
*/
func deleteFile() error {
	err := os.Remove(FilePath) //remove the file
	if err != nil {
		return err
	}
	fmt.Println(FilePath + " deleted")
	return nil
}

func WriteFile(todoList []models.Task) error {
	err := CreateFile()
	if err != nil {
		return err
	}

	emptyTodoList := len(todoList) == 0
	if emptyTodoList {
		err := deleteFile()
		if err != nil {
			return err
		}
		return nil
	}

	file, err := os.Create(FilePath)
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "	")

	if err := encoder.Encode(todoList); err != nil {
		return err
	}

	log.Println("todoList saved to file: " + FilePath)
	return nil
}
