package array_operations

import "fmt"

var pl = fmt.Println

func ArrayOperations() {
	var intArray [4]int
	intArray[0] = 0
	intArray[1] = 1
	intArray[2] = 2
	// intArray[3] == 0 defaults to zero for int array

	arrayLoopLength(intArray)
	arrayLoopRange(intArray)
	arrayLoopMultiDimensional()
}

func arrayLoopLength(intArray [4]int) {

	for i := 0; i < len(intArray); i++ {
		pl("Value:", intArray[i])
	}

}

func arrayLoopRange(intArray [4]int) {
	for index, value := range intArray {
		pl("index:", index, "value:", value)
	}
}

func arrayLoopMultiDimensional() {

	multiDimensionalArray := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i, _ := range multiDimensionalArray {
		for ii := 0; ii < len(multiDimensionalArray[i]); ii++ {
			pl(fmt.Sprintf("multiDimensionalArray[%d][%d] == %d", i, ii, multiDimensionalArray[i][ii]))
		}
	}
}
