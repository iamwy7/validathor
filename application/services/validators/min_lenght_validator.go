package validators

import "regexp"

func MinLenghtValidator(pass string, c chan bool) {
	match, err := regexp.Match(".{9,}", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- match
}
