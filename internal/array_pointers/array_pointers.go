package array_pointers

import "fmt"

var pl = fmt.Println

func ArrayPointers() {
	myArray := [4]int{4, 8, 12, 1e6}
	pl("Array before doubling:", myArray)
	doubleArrayValues(&myArray)
	pl("Array after doubling:", myArray)

	sliceByReference()
}

func doubleArrayValues(arr *[4]int) {
	for x := 0; x < len(arr); x++ {
		arr[x] *= 2
	}
}

func sliceByReference() {
	mySlice := []float64{11, 14, 17, 465}
	fmt.Printf("Average: %.3f \n", getAverage(mySlice...))
}

func getAverage(sliceToAverage ...float64) float64 {
	var sum float64 = 0.0
	var NumSize float64 = float64(len(sliceToAverage))
	for _, num := range sliceToAverage {
		sum += num
	}
	return sum / NumSize

}
