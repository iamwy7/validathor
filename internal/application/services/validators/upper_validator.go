package validators

import "regexp"

// UpperValidator validate if the password has
// at least one upper case letter.
// He receive the password, validate, and return the result to chain.
func UpperValidator(pass string, c chan bool) {
	match, err := regexp.Match("[A-Z]", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
