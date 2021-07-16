package services

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wy7-source/iti-challenge/application/domain"
)

var passwordService domain.PasswordServiceInterface

func setUp() {
	passwordService = NewPasswordService()
}
func TestValidators_PasswordService(t *testing.T) {
	setUp()
	passwordCandidate := domain.NewPassword("AbTp9!fok")
	isValid := passwordService.Validate(passwordCandidate)
	require.True(t, isValid)

	passwordCandidate.SetValue("007example")
	isValid = passwordService.Validate(passwordCandidate)

	require.False(t, isValid)
}
