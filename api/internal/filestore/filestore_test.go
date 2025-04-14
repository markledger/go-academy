package filestore

import (
	"testing"
)

func TestFileIsParsedToSlice(t *testing.T) {
	data, err := ParseFileToSlice("./test_data/database.txt")

	if err != nil {
		t.Error(err)
	}

	if len(data) != 3 {
		t.Error("Incorrect length")
	}

	if data[0] != "Fish and Chips" {
		t.Error("Slice element 0 incorrect")
	}

	if data[1] != "Jerk Chicken" {
		t.Error("Slice element 1 incorrect")
	}

	if data[2] != "Fish and Crab" {
		t.Error("Slice element 2 incorrect")
	}
}
