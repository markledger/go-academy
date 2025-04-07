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

const FilePath = "./todo-list.txt"
const CreateAction = "create"
const EditAction = "edit"
const DeleteAction = "delete"

var id int
var action string
var task string
var validActions = []string{EditAction, CreateAction, DeleteAction}

/*
*
Setup for the program

- Create the file to store tasks if it doesn't exist
- Declare and parse the flags
*/
func init() {

	_, err := os.Stat(FilePath)
	if errors.Is(err, os.ErrNotExist) {
		createFile(FilePath)
	}

	flag.IntVar(&id, "id", 0, "use in combination with the -action flag to select task to be modified")
	flag.StringVar(&action, "action", CreateAction, "use in combination with -id. Select action from: "+EditAction+"|"+DeleteAction)
	flag.StringVar(&task, "task", "example task to complete", "the task you want to create, or the new task if editing")
	flag.Parse()
}

func main() {
	var todoList []string
	todoList = parseFileToSlice(FilePath)
	err := validateFlags(len(todoList))
	if err != nil {
		logError(err)
	}

	todoList = handleAction(todoList)
	listCurrentTasks(todoList)
	updateFile(FilePath, todoList)
}

func logError(errorMessage error) {
	f, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(errorMessage)
}

/*
*
Parse the file located at filePath and split on new lines storing each line as a
a task in the data slice.
*/
func parseFileToSlice(filePath string) []string {
	var data []string
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		logError(err)
	}

	for _, line := range strings.Split(string(fileData), "\n") {
		if line == "" {
			continue
		}
		data = append(data, line)
	}
	return data
}

/*
*
Create a file at the location provided in filePath
*/
func createFile(filePath string) {
	_, err := os.Create(filePath)
	if err != nil {
		logError(err)
	}
}

/*
*
Delete the file if a file exists at the filePath location
*/
func deleteFile(filePath string) {
	err := os.Remove(filePath) //remove the file
	if err != nil {
		logError(err)
		return
	}
	fmt.Println(FilePath + " deleted")
}

func updateFile(filePath string, todoList []string) {

	if len(todoList) == 0 {
		deleteFile(FilePath)
		return
	}
	createFile(filePath)
	f, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, todo := range todoList {
		if _, err := f.WriteString(todo + "\n"); err != nil {
			log.Fatal(err)
		}
	}

}

func listCurrentTasks(todoList []string) {
	fmt.Println("id | Task")

	for i, todo := range todoList {
		fmt.Println(fmt.Sprintf("[%d]: %s", i+1, todo))
	}
}

func validateFlags(numberOfTasks int) error {
	var errorMsg string
	invalidId := (id < 1 || id > numberOfTasks) && action != CreateAction
	invalidAction := !slices.Contains(validActions, action)

	if invalidId {
		errorMsg = fmt.Sprintf("Invalid id selected. Please select an id between 1 and %d", numberOfTasks)
		log.Fatal(errorMsg)
		return errors.New(errorMsg)
	}
	if invalidAction {
		errorMsg = "Invalid action selected. Please select from: " + CreateAction + ", " + EditAction + "or " + DeleteAction
		log.Fatal(errorMsg)
		return errors.New(errorMsg)
	}

	return nil
}

/*
Handle either creating, editing or deleting a task and return the updated todoList
*/
func handleAction(todoList []string) []string {
	if action == "edit" {
		todoList[id-1] = task
	}

	if action == "delete" {
		todoList = append(todoList[:id-1], todoList[id:]...)
	}

	if action == "create" {
		todoList = append(todoList, task)
	}
	return todoList
}
