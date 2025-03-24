package slice_operations

import "fmt"

func SliceOperations() {

	myString := "Jude and Kit"
	myRune := []rune(myString)
	for _, v := range myRune {
		fmt.Printf("Rune array: %d\n", v)
	}

}
