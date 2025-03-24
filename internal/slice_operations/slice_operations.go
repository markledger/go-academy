package slice_operations

import (
	"fmt"
)

var pl = fmt.Println

// Slices are arrays that can grow

func SliceOperations() {
	thisSlice := make([]string, 3)
	thisSlice[0] = "Mark"
	thisSlice[1] = "Jude"
	thisSlice[2] = "Kit"

	pl("Slice size:", len(thisSlice))
	for _, v := range thisSlice {
		pl(v)
	}

	// A slice is a view of an underlying array e.g:

	mySliceArray := [5]int{1, 2, 3, 4, 5}
	mySlice := mySliceArray[0:2]
	pl("1st three of myArray:", mySliceArray[:3])
	pl("Values after 2nd index in myArray:", mySliceArray[2:])

	mySliceArray[0] = 10
	pl("Lets alter the underlying arraymySliceArray[0] = 10 (the slice is a pointer the array)")
	pl("mySlice : ", mySlice)

	mySlice[1] = 22
	pl("Lets alter the slice now mySlice[2] = 22")
	pl("mySlice : ", mySlice)
	pl("mySliceArray : ", mySliceArray)

	pl("We can append a slice: mySlice = append(mySlice, 11)")
	mySlice = append(mySlice, 11)
	pl("mySlice : ", mySlice)
	pl("mySliceArray : ", mySliceArray)

}
