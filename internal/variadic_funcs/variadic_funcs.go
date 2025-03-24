package variadic_funcs

import (
	"fmt"
)

var pl = fmt.Println

func SumUp() {
	pl("The sum is:", getSumOfMany(1, 2, 4))
}

func getSumOfMany(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
