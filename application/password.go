package application

import "errors"

type Password struct {
	Value string `json:"value"`
}

func NewPassword(value string) *Password {
	return &Password{Value: value}
}
func (p *Password) IsValid() (bool, error) {
	// Regex Here
	return true, errors.New("example")
}
