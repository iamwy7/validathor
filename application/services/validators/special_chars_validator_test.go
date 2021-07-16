package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_SpecialCharsValidator(t *testing.T) {
	value := "SOM+T°hING190"
	matched := SpecialCharsValidator(value)
	require.True(t, matched)

	value = "EXAMPLE°123sss"
	matched = SpecialCharsValidator(value)

	require.False(t, matched)
}
