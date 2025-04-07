package filestore

import (
	"testing"
)

func TestFileIsParsedToSlice(t *testing.T) {
	data, err := ParseFileToSlice("./internal/filestore/test_data/database.txt")

	if err != nil {
		t.Error(err)
	}

	if len(data) != 3 {
		t.Error("")
	}

}
