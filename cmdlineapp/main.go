package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

// - Create a command line application that uses flags to accept a to-do item adds it to an empty list of to-do items and prints the list to console
// - After printing the list of to-do items, save them to a file on disk
// - When the application starts, load all to-do items from disk before adding new item
// - Allow the user to update the description of a to-do item or delete it
var pl = fmt.Println
var filePath = "./todo-list.txt"

const CREATE = "create"
const EDIT = "edit"
const DELETE = "delete"

var id int
var action string
var task string
var validActions = []string{EDIT, CREATE, DELETE}

func init() {
	flag.IntVar(&id, "id", 0, "use in combination with the -action flag to select task to be modified")
	flag.StringVar(&action, "action", CREATE, "use in combination with -id. Select action from: "+EDIT+"|"+DELETE)
	flag.StringVar(&task, "task", "example task to complete", "the task you want to create, or the new task if editing")
	flag.Parse()
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
	todoList = parseFileToSlice(filePath)
	if id < 1 || id > len(todoList) && action != CREATE {
		pl(fmt.Errorf("Invalid id selected. Please select an id between 1 and %d", len(todoList)))
		os.Exit(1)
	}
	if !slices.Contains(validActions, action) {
		pl(fmt.Errorf("Invalid action selected. Please select from: create, edit or delete"))
		os.Exit(1)
	}

	todoList = append(todoList, task)
	pl("id | Task")

	if id > 0 && id < len(todoList) && action == "edit" {

	}
	if id > 0 && id < len(todoList) && action == "delete" {

	}
	for i, todo := range todoList {
		pl(fmt.Sprintf("[%d]: %s", i+1, todo))
	}

	if task == "create" {
		appendFile(filePath, task)
	}
}
