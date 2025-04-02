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

	lines := strings.Split(string(fileData), "\n")
	for _, line := range lines {
		data = append(data, line)
	}
	return data
}

func createFile(filePath string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}
func main() {
	var todoList []string
	taskPrt := flag.String("task", "Consume double espresso", "a string")
	flag.Parse()

	todoList = parseFileToSlice(filePath)
	todoList = append(todoList, *taskPrt)
	for _, v := range todoList {
		pl(v)
	}

	//file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//if _, err := file.WriteString(*taskPrt + "\n"); err != nil {
	//	log.Fatal(err)
	//}

	//for _, i := range args {
	//	val, err := strconv.Atoi(i)
	//	if err != nil {
	//		panic(err)
	//	}
	//	iArgs = append(iArgs, val)
	//}
	//
	//max := 0
	//for _, val := range iArgs {
	//	if val > max {
	//		max = val
	//	}
	//}
	//pl("Max Value :", max)
}
