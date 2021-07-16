package services

import (
	"github.com/wy7-source/iti-challenge/application/domain"
	"github.com/wy7-source/iti-challenge/application/services/validators"
)

type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (p *PasswordService) Validate(password domain.PasswordInterface) bool {
	isValid := validators.DigitValidator(password.GetValue())
	if !isValid {
		return false
	}
	isValid = validators.DuplicatesValidator(password.GetValue())
	if !isValid {
		return false
	}
	isValid = validators.LowerValidator(password.GetValue())
	if !isValid {
		return false
	}
	isValid = validators.SpecialCharsValidator(password.GetValue())
	if !isValid {
		return false
	}
	isValid = validators.UpperValidator(password.GetValue())
	if !isValid {
		return false
	}
	isValid = validators.WhiteSpacesValidator(password.GetValue())
	if !isValid {
		return false
	}
	isValid = validators.MinLenghtValidator(password.GetValue())
	if !isValid {
		return false
	}
	return true
}
