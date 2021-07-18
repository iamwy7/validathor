package validators

import "regexp"

func DigitValidator(pass string, c chan bool) {
	match, err := regexp.Match("\\d", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
