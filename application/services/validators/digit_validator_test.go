package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_DigitValidator(t *testing.T) {
	value := "something"
	matched := DigitValidator(value)
	require.False(t, matched)

	value = "007example"
	matched = DigitValidator(value)

	require.True(t, matched)
}
