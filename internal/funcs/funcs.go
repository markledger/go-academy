package funcs

import (
	"fmt"
	"log"
)

var pl = fmt.Println

func FuncExamples() {
	pl(getTwoNumbers(1, 2))
	nums := [3]float64{22, 11, 0}
	answer, err := getQuotient(nums[0], nums[1])
	if err != nil {
		log.Fatalln(err)
	}
	pl(fmt.Sprintf("The result of: %.1f / %.1f == %.1f", nums[0], nums[1], answer))

	// nextAnswer, err := getQuotient(nums[0], nums[2])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// pl(fmt.Sprintf("The result of: %.1f / %.1f == %.1f", nums[0], nums[1], answer))
}

func getTwoNumbers(x int, y int) (int, int) {
	return x + 1, y + 1
}

func getQuotient(x float64, y float64) (answer float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	}

	return x / y, nil

}
