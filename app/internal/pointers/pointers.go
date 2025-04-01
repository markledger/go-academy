package pointers

import "fmt"

var pl = fmt.Println

func PointersExample() {
	myInteger := 5
	pl("myInteger before:", myInteger)
	changeValueByReference(&myInteger) // point to myInteger
	pl("myInteger after:", myInteger)

	storePointer()

}

func changeValueByReference(pointerToMyInteger *int) {
	*pointerToMyInteger = 9 // write through pointerToMyInteger to myInteger
}

func storePointer() {
	pointerToStore := 10
	var storedPointer *int = &pointerToStore
	pl("pointerToStore value:", pointerToStore)
	pl("pointerToStore memory address:", storedPointer)
	pl("storedPointer value:", *storedPointer)
	*storedPointer = 55
	pl("storedPointer value (55):", *storedPointer)

	pl("pointerToStore value  before changeValueByReference:", pointerToStore)
	changeValueByReference(&pointerToStore)
	pl("pointerToStore  value aftet changeValueByReference:", pointerToStore)
	pl("storedPointer value:", *storedPointer)
}
