package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_MinLenghtValidator(t *testing.T) {
	value := "SOMEThING123890"
	matched := MinLenghtValidator(value)
	require.True(t, matched)

	value = "EXAMPLE!"
	matched = MinLenghtValidator(value)

	require.False(t, matched)
}
