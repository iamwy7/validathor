package validators

import (
	"regexp"
)

func WhiteSpacesValidator(pass string) bool {
	match, err := regexp.Match("\\s", []byte(pass))

	if err != nil {
		return false
	}
	return !match
}
