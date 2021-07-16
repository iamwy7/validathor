package validators

import "regexp"

func SpecialCharsValidator(pass string) bool {
	match, err := regexp.Match("[!@#$%^&*()-+]", []byte(pass))

	if err != nil {
		return false
	}
	return match
}
