package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_UpperValidator(t *testing.T) {
	value := "Something$"
	matched := UpperValidator((value))
	require.True(t, matched)

	value = "someth1ng$3ls3"
	matched = UpperValidator((value))

	require.False(t, matched)
}
