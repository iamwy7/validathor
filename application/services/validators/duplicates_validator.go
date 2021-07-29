package validators

import (
	"strings"
)

// DuplicatesValidator validate if the password has duplicate characters.
// He receive the password, validate, and return the result to chain.
func DuplicatesValidator(pass string, c chan bool) {
	/*  The range form of the for loop iterates over a slice or map.
		He return 2 values, the key and rune. In this case, i dont use the key, only the rune.
	See more: https://tour.golang.org/moretypes/16
	*/
	for _, r := range pass {
		// We check if the current rune value are present in password more than one time.
		// By the way, rune can hold a value, but we cant use like a value, so, i made a explicit cast.
		duplicates := strings.Count(pass, string(r))
		if duplicates > 1 {
			c <- false
		}
	}
	c <- true
}
