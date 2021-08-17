package validators

import "regexp"

// MinLenghtValidator validate if the password has
// 9 or more characters.
// He receive the password, validate, and return the result to chain.
func MinLenghtValidator(pass string, c chan bool) {
	match, err := regexp.Match(".{9,}", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
