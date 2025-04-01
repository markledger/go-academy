package regex_examples

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func RegexExamples() {
	reStr := "Cat rat mat fat tap hat"

	r, _ := regexp.Compile("([crmfp]at)")
	pl("MatchString: ", r.MatchString(reStr))
	pl("FindString: ", r.FindString(reStr))
	pl("Indexg: ", r.FindStringIndex(reStr))
	pl("All String: ", r.FindAllString(reStr, -1))
	pl("1st Two Strings: ", r.FindAllString(reStr, 2))
	pl("All submatch index: ", r.FindAllStringSubmatchIndex(reStr, -1))
	pl("Replace all matches with DOG: ", r.ReplaceAllString(reStr, "DOG"))
}
