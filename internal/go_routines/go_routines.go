package go_routines

import (
	"fmt"
)

var pl = fmt.Println

func printTo15() {
	for i := 0; i <= 15; i++ {
		pl("Fun 1: ", i)
	}
}

func printTo10() {
	for i := 0; i <= 10; i++ {
		pl("Fun 2: ", i)
	}
}

// Go Routines can communicate with each other using channels
// The sending go routine is gonna make sure the recieveing
// go routine will receive values beofre it attempts to use them,

func nums1(channel chan int) {
	channel <- 1
	channel <- 244
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func GoRoutines() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)

	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)
}
