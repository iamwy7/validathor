package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDomain_NewPassword(t *testing.T) {
	value := "something"
	pass := NewPassword(value)
	require.NotEmpty(t, pass)
	require.Equal(t, pass.GetValue(), "something")
}
