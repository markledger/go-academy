package recursion

import "fmt"

var pl = fmt.Println

func FactorialRecursionExample() {
	pl("Factorial 4 = ", factorialRecursion(4))
}
func factorialRecursion(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorialRecursion(num-1)
}
