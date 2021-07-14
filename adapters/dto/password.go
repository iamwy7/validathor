package dto

import "github.com/wy7-source/iti-challenge/application"

type PasswordDto struct {
	Value string `json:"value"`
}

// Hydrate can transfer data of dto to eproduct type Password of application
func (p *PasswordDto) Hydrate(password *application.Password) error {

	password.Value = p.Value
	_, err := password.IsValid()
	if err != nil {
		return err
	}
	return nil
}
