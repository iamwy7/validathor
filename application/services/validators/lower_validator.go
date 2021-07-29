package validators

import "regexp"

// LowerValidator validate if the password has at least
// one lower case letter.
// He receive the password, validate, and return the result to chain.
func LowerValidator(pass string, c chan bool) {
	match, err := regexp.Match("[a-z]", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
