// Package domain is where the main software entities are.
package domain

import "github.com/asaskevich/govalidator"

type PasswordServiceInterface interface {
	Validate(password PasswordInterface) bool
}
type PasswordInterface interface {
	GetValue() string
	SetValue(value string)
	IsValid() (bool, error)
}

// Password is the principal entity.
type Password struct {
	value string `valid:"required"`
}

func NewPassword(value string) *Password {
	return &Password{value: value}
}

func (p *Password) GetValue() string {
	return p.value
}

func (p *Password) SetValue(value string) {
	p.value = value
}

// IsValid is used to validate the entity.
// It returns if the password is valid and any validation error encountered.
func (p *Password) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}
