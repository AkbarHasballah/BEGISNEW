package musik

import (
	"fmt"
	"strings"
)

func IsMatch(str string, subs ...string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	fmt.Printf("String: \"%s\", Substrings: %s\n", str, subs)

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches += 1
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}
