package validators

import "regexp"

func MinLenghtValidator(pass string) bool {
	match, err := regexp.Match(".{9,}", []byte(pass))

	if err != nil {
		return false
	}
	return match
}
