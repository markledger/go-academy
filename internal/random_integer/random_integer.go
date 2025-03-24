package random_integer

import (
	"math/rand"
	"time"
)

func GetRandomInteger(max int) int {

	seedSeconds := time.Now().Unix()
	rand.NewSource(seedSeconds)

	return rand.Intn(max)

}
