package main

import (
	"testing"
)

func TestValidateFlagsRequiresTaskToCreate(t *testing.T) {
	action = "create"
	id = 0
	task = "  "
	err := validateFlags(0)
	if err == nil {
		t.Error(err)
	}

	task = ""
	err = validateFlags(0)
	if err == nil {
		t.Error(err)
	}

	task = "Make it ao"
	err = validateFlags(12)
	if err != nil {
		t.Error(err)
	}
}
func TestValidateFlagsHasValidIdToDelete(t *testing.T) {
	action = "delete"
	id = 1
	err := validateFlags(2)
	if err != nil {
		t.Error(err)
	}

	id = 0
	err = validateFlags(2)
	if err == nil {
		t.Error(err)
	}

	id = 13
	err = validateFlags(2)
	if err == nil {
		t.Error(err)
	}
}

func TestValidateFlagsRequiresValidIdAndTaskContentToEdit(t *testing.T) {
	action = "edit"
	id = 0
	task = "Make it so number one"
	err := validateFlags(2)
	if err == nil {
		t.Error(err)
	}

	task = " "
	err = validateFlags(2)
	if err == nil {
		t.Error(err)
	}

	id = 3
	task = "Make it so number one"
	err = validateFlags(2)
	if err == nil {
		t.Error(err)
	}

	task = "    "
	err = validateFlags(2)
	if err == nil {
		t.Error(err)
	}

	id = 2
	task = "Make it so number one"
	err = validateFlags(2)
	if err != nil {
		t.Error(err)
	}

	id = 2
	task = ""
	err = validateFlags(2)
	if err == nil {
		t.Error(err)
	}
}

func TestHandleDeleteAction(t *testing.T) {
	todoList := []string{"Meet the Klingons", "Resist the Borg", "Defeat the Dominion"}
	action = "delete"
	id = 2

	updatedTodoList, _ := handleAction(todoList)

	if len(updatedTodoList) != 2 {
		t.Error()
	}

	if updatedTodoList[0] != "Meet the Klingons" {
		t.Error()
	}

	if updatedTodoList[1] != "Defeat the Dominion" {
		t.Error()
	}

}

func TestHandleEditAction(t *testing.T) {
	todoList := []string{"Meet the Klingons", "Resist the Borg", "Defeat the Dominion"}
	action = "edit"
	id = 2
	task = "Beat the Borg"

	updatedTodoList, _ := handleAction(todoList)

	if len(updatedTodoList) != 3 {
		t.Error()
	}

	if updatedTodoList[0] != "Meet the Klingons" {
		t.Error()
	}

	if updatedTodoList[1] != "Beat the Borg" {
		t.Error()
	}

	if updatedTodoList[2] != "Defeat the Dominion" {
		t.Error()
	}

}

func TestHandleCreateAction(t *testing.T) {
	todoList := []string{"Meet the Klingons", "Resist the Borg", "Defeat the Dominion"}

	action = "create"
	task = "Run away"
	appendedTodoList, err := handleAction(todoList)
	if len(appendedTodoList) != 4 {
		t.Error("There should be 4 elements in the slice")
	}
	if appendedTodoList[0] != "Meet the Klingons" {
		t.Error()
	}

	if appendedTodoList[1] != "Resist the Borg" {
		t.Error()
	}

	if appendedTodoList[2] != "Defeat the Dominion" {
		t.Error()
	}

	if appendedTodoList[3] != "Run away" {
		t.Error()
	}

	if err != nil {
		t.Error()
	}
}
