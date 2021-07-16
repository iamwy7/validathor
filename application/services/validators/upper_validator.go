package validators

import "regexp"

func UpperValidator(pass string) bool {
	match, err := regexp.Match("[A-Z]", []byte(pass))

	if err != nil {
		return false
	}
	return match
}
