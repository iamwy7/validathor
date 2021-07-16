package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_LowerValidator(t *testing.T) {
	value := "SOMEThING"
	matched := DuplicatesValidator(value)
	require.True(t, matched)

	value = "EXAMPLE"
	matched = DuplicatesValidator(value)

	require.False(t, matched)
}
