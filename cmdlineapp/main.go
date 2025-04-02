package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// - Create a command line application that uses flags to accept a to-do item adds it to an empty list of to-do items and prints the list to console
// - After printing the list of to-do items, save them to a file on disk
// - When the application starts, load all to-do items from disk before adding new item
// - Allow the user to update the description of a to-do item or delete it
var pl = fmt.Println
var filePath = "./todo-list.txt"

var id int
var action string
var task string

func init() {
	flag.IntVar(&id, "id", 0, "use in combination with the -action flag to select task to be modified")
	flag.StringVar(&action, "action", "create", "use in combination with -id. Select action from: edit|delete ")
	flag.StringVar(&task, "task", "example task to complete", "the task you want to create, or the new task if editing")
}

func parseFileToSlice(filePath string) []string {
	var data []string

	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		pl("File doesn't exist - It will be created.")
		createFile(filePath)
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	removeNewLine := strings.NewReplacer("\n", "")

	lines := strings.Split(string(fileData), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		data = append(data, removeNewLine.Replace(line))
	}
	return data
}

func createFile(filePath string) {
	_, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func appendFile(filePath string, todo string) {

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(todo + "\n"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var todoList []string
	flag.Parse()

	todoList = parseFileToSlice(filePath)
	todoList = append(todoList, task)
	pl(action)
	pl(id)
	pl(task)
	pl("id | Task")

	for i, todo := range todoList {
		pl(fmt.Sprintf("[%d]: %s", i+1, todo))
	}
	appendFile(filePath, task)
}
