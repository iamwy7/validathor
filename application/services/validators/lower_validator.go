package validators

import "regexp"

func LowerValidator(pass string) bool {
	match, err := regexp.Match("[a-z]", []byte(pass))

	if err != nil {
		return false
	}
	return match
}
