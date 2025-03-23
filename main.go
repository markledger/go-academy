package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	pl("many eyes")

	pl(stringToFloat("3.14"))

	pl(floatToString(3.142))

	conditional()
	stringOperations("Mark Ledger")

	runeOperations()

	timeOperations()

	createRandomValues(5)

	mathOperations()

	convertDegreesToRadians(90)
	printOperations()
}

func printOperations() {
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value
	fmt.Printf("%s %d %c %f %t %o %x\n",
		"stuff", 1, 'A', 3.14, true, 1, 1)

	//output width of 9 spaces
	fmt.Printf("%9f\n", 3.14)

	//print to 2 decimal places
	fmt.Printf("%.2f\n", 3.14)
	fmt.Printf("%9.1f\n", 3.14)
}

func convertDegreesToRadians(myDegrees float64) {
	radians := myDegrees * math.Pi / 180
	degrees := radians * (180 / math.Pi)

	fmt.Printf("%.2f radians = %.2f degrees\n", radians, degrees)
}

func mathOperations() {

	pl("Abs (-10) = ", math.Abs(-10))
	pl("Pow(4,2) = ", math.Pow(4, 2))
	pl("Sqrt(16) = ", math.Sqrt(16))
	pl("Cbrt(8)= ", math.Cbrt(8))
	pl("Ceil(8.8)= ", math.Ceil(8.8))
	pl("Floor(9.8)= ", math.Floor(9.8))
	pl("Max(5,6) = ", math.Max(5, 6))
	pl("Min(7,9) = ", math.Min(7, 9))
}

func createRandomValues(howMany int) {
	seedSeconds := time.Now().Unix()
	rand.Seed(seedSeconds)

	for i := 0; i < howMany; i++ {
		pl("Random Number:", rand.Intn(25))
	}

}

func timeOperations() {
	now := time.Now()
	pl(fmt.Sprintf("%d-%d-%d", now.Day(), now.Month(), now.Year()))
	pl(fmt.Sprintf("%d:%d:%d", now.Hour(), now.Minute(), now.Second()))
}

func runeOperations() {
	runeString := "abcdefg"

	pl("Rune Count:", utf8.RuneCountInString(runeString))

	for i, runeVal := range runeString {
		fmt.Printf("index %d : %#U : %c\n", i, runeVal, runeVal)
	}
}

func stringOperations(myString string) {

	//Replace
	replacer := strings.NewReplacer("a", "aaaaaa")
	pl(replacer.Replace(myString))

	//Contains

	pl("Contains letter 'M'? ", strings.Contains(myString, "M"))

	//Length
	pl("Length: ", len(myString))

	//first index match
	pl("k index: ", strings.Index(myString, "k"))

	//replace all instances
	pl("Replace all instances of e: ", strings.ReplaceAll(myString, "e", "u"))

	//remove leading and trailing whitespace
	var newLines = "\r\n\tNew Lines Too Much\n"
	pl(newLines)
	newLines = strings.TrimSpace(newLines)
	pl("Remove the new lines:", newLines)

	//split string
	pl("Split:", strings.Split("M-a-r-K", "-"))

	//upper or lowercase conversion
	pl("Lowercase:", strings.ToUpper(myString))
	pl("Uppercase:", strings.ToLower(myString))

	//prefix
	pl("Does", myString, "have T prefix? ", strings.HasPrefix(myString, "T"))

	//suffix
	pl("Does", myString, "have r suffix? ", strings.HasSuffix(myString, "r"))

}

func floatToString(myFloat float64) string {
	return fmt.Sprintf("The float as string: %f", myFloat)

}

func stringToFloat(myString string) float64 {
	converted, err := strconv.ParseFloat(myString, 64)
	if err == nil {
		return converted
	}
	return 0.0
}

func conditional() {

	age := 8

	if (age >= 1) && (age <= 18) {
		pl("Important age")
	}

	if (age == 21) || (age == 50) {
		pl("kinda important")
	}
	pl("!true = ", !true)

}
