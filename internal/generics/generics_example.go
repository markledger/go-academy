package generics_example

import "fmt"

var pl = fmt.Println

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func GenericsExample() {
	pl("55+55 = ", getSumGen(55, 55))
}
