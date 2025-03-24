package file_io

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

var filePath = "./internal/file_io/data.txt"

func WritePrimeNumbersToFile() {

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	iPrimeNumbersArray := []int{2, 3, 5, 7, 11}
	var sPrimeNumsArr []string
	for _, integer := range iPrimeNumbersArray {
		sPrimeNumsArr = append(sPrimeNumsArr, strconv.Itoa(integer))
	}
	for _, num := range sPrimeNumsArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		pl("Prime Number:", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

}

func AppendFile() {
	WritePrimeNumbersToFile()
	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		pl("File doesn't exist - It will be created.")
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString("13\n"); err != nil {
		log.Fatal(err)
	}
}
