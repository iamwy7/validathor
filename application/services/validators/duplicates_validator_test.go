package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_DuplicatesValidator(t *testing.T) {
	value := "something"
	matched := DuplicatesValidator(value)
	require.True(t, matched)

	value = "007example"
	matched = DuplicatesValidator(value)

	require.False(t, matched)
}
