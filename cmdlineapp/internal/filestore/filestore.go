package filestore

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const FilePath = "./internal/filestore/database.txt"

/*
*
Parse the file located at filePath and split on new lines storing each line as a
a task in the data slice.
*/
func ParseFileToSlice() ([]string, error) {
	var data []string
	fileData, err := os.ReadFile(FilePath)
	if err != nil {
		return data, err
	}

	for _, line := range strings.Split(string(fileData), "\n") {
		if line == "" {
			continue
		}
		data = append(data, line)
	}
	return data, nil
}

/*
*
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

func WriteFile(todoList []string) error {
	emptyTodoList := len(todoList) == 0
	if emptyTodoList {
		err := deleteFile()
		if err != nil {
			return err
		}
		return nil
	}

	err := CreateFile()
	if err != nil {
		return err
	}

	f, openFileError := os.OpenFile(FilePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if openFileError != nil {
		return openFileError
	}

	defer f.Close()
	for _, todo := range todoList {
		if _, err := f.WriteString(todo + "\n"); err != nil {
			return err
		}
	}
	return nil
}
