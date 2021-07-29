// Package dto provides Data Transfer Objects to receive requests
// over API.
package dto

import (
	"github.com/wy7-source/iti-challenge/application/domain"
)

// PasswordDto is the value received by API.
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
