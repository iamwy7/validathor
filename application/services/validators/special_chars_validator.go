package validators

import "regexp"

func SpecialCharsValidator(pass string, c chan bool) {
	match, err := regexp.Match("[!@#$%^&*()-+]", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
