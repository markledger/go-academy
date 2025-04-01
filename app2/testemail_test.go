package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("Hello")
	if err != nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("mark@mark.com")
	if err == nil {
		t.Error("mark@mark.com is an email")
	}

	// and a test created to demonstrate a fail
	//_, err = IsEmail("mark@mark")
	//if err != nil {
	//	t.Error("mark@mark is not an email")
	//}
}
