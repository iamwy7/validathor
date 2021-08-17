package validators

import (
	"regexp"
)

// WhiteSpacesValidator validate if the password has
// whitespaces.
// He receive the password, validate, and return the result to chain.
func WhiteSpacesValidator(pass string, c chan bool) {
	match, err := regexp.Match("\\s", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- !match
}
