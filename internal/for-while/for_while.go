package for_while

import (
	"fmt"
	"go_academy_course/internal/random_integer"
)

var pl = fmt.Println

func LoopOperations() {

	var cycles int = random_integer.GetRandomInteger(10)
	pl(fmt.Sprintf("Number of cycles to loop: %d", cycles))
	for i := 0; i <= cycles; i++ {
		pl("Cycle #", i)
	}

}
