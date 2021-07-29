package validators

import "regexp"

// SpecialCharsValidator validate if the password has
// at least one special character, which can be: !@#$%^&*()-+
// He receive the password, validate, and return the result to chain.
func SpecialCharsValidator(pass string, c chan bool) {
	match, err := regexp.Match("[!@#$%^&*()-+]", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
