// Package validators provides the validators to service use to work.
package validators

import "regexp"

// DigitValidator validate if the password has at least one number.
// He receive the password, validate, and return the result to chain.
func DigitValidator(pass string, c chan bool) {
	match, err := regexp.Match("\\d", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
