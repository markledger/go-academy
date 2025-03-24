package for_while

import (
	"bufio"
	"fmt"
	"go_academy_course/internal/random_integer"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func LoopOperations() {

	var cycles int = random_integer.GetRandomInteger(10)
	pl(fmt.Sprintf("Number of cycles to loop: %d", cycles))
	for i := 0; i < cycles; i++ {
		pl("Cycle #", i)
	}

}

func LoopFixedArray(arraySize int) {
	sized := []int{1, 2, 44}

	for index, number := range sized {
		pl("index:", index, "Value:", number)
	}

}

func GuessNumberGame() {
	var randomNumber int = random_integer.GetRandomInteger(10)
	for true {
		pl("Guess a number between 0 and 10:")
		pl("Random Number is:", randomNumber)
		reader := bufio.NewReader((os.Stdin))
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randomNumber {
			pl("Pick a lower value...")
		}

		if iGuess < randomNumber {
			pl("Pick a higher value...")
		}

		if iGuess == randomNumber {
			pl("You got it! It was:", iGuess)
			break
		}
	}
}
