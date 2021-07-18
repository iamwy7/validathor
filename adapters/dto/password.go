package dto

import (
	"github.com/wy7-source/iti-challenge/application/domain"
)

type PasswordDto struct {
	Value string `json:"value"`
}

// Hydrate can transfer data of dto to Password type.
func (p *PasswordDto) Hydrate(password *domain.Password) error {

	password.SetValue(p.Value)
	_, err := password.IsValid()
	if err != nil {
		return err
	}
	return nil
}
