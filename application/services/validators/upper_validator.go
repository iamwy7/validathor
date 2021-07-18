package validators

import "regexp"

func UpperValidator(pass string, c chan bool) {
	match, err := regexp.Match("[A-Z]", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
