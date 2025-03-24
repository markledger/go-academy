package learn_maps

import "fmt"

var pl = fmt.Println

// Maps are basically collections of key value pairs
// and the key can be a different data type to the value.
func LearningMaps() {
	var heroes map[string]string
	heroes = make(map[string]string)
	villians := make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["DogMan"] = "A man and a dog"

	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}
	pl(superPets[1])
	checkIfKeyHasValue(superPets)
	loopOverMap(heroes)

	deleteByKey(heroes, "Batman")

}

func checkIfKeyHasValue(superPets map[int]string) {
	_, ok := superPets[3]
	pl("Is there a 3rd pet?", ok)
}

func loopOverMap(heroesArr map[string]string) {
	for k, v := range heroesArr {
		fmt.Printf("%s is %s\n", k, v)
	}
}

func deleteByKey(heroesArr map[string]string, key string) {
	delete(heroesArr, key)
	loopOverMap(heroesArr)
}
