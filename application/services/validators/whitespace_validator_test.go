package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_WhiteSpacesValidator(t *testing.T) {
	value := "something"
	matched := WhiteSpacesValidator(value)
	require.True(t, matched)

	value = "007exa mple"
	matched = WhiteSpacesValidator(value)

	require.False(t, matched)
}
