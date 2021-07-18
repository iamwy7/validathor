package validators

import "regexp"

func LowerValidator(pass string, c chan bool) {
	match, err := regexp.Match("[a-z]", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
