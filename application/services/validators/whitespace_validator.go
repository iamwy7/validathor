package validators

import (
	"regexp"
)

func WhiteSpacesValidator(pass string, c chan bool) {
	match, err := regexp.Match("\\s", []byte(pass))

	if err != nil {
		c <- false
	}
	c <- !match
}
