package validators

import "regexp"

func DigitValidator(pass string) bool {
	match, err := regexp.Match("\\d", []byte(pass))

	if err != nil {
		return false
	}
	return match
}
