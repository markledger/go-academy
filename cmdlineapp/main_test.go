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
