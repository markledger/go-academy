package main

import (
	"cmdlineapp/internal/filestore"
	"errors"
	"flag"
	"fmt"
	"log"
	"slices"
)

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

	err := filestore.CreateFile()
	if err != nil {
		log.Fatal("unable to create file to persist todo tasks")
	}

	flag.IntVar(&id, "id", 0, "use in combination with the -action flag to select task to be modified")
	flag.StringVar(&action, "action", CreateAction, "use in combination with -id. Select action from: "+EditAction+"|"+DeleteAction)
	flag.StringVar(&task, "task", "example task to complete", "the task you want to create, or the new task if editing")
	flag.Parse()
}

func main() {
	var todoList []string

	todoList, err := filestore.ParseFileToSlice()
	if err != nil {
		log.Fatal(err)
	}

	err = validateFlags(len(todoList))
	if err != nil {
		log.Fatal(err)
	}

	todoList = handleAction(todoList)
	listCurrentTasks(todoList)
	err = filestore.WriteFile(todoList)
	if err != nil {
		log.Fatal(err)
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
		return errors.New(errorMsg)
	}
	if invalidAction {
		errorMsg = "Invalid action selected. Please select from: " + CreateAction + ", " + EditAction + "or " + DeleteAction
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
