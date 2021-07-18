package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_UpperValidator(t *testing.T) {
	value := "Something$"
	chanTest := make(chan bool)
	go UpperValidator(value, chanTest)
	isValid := <-chanTest
	require.True(t, isValid)

	value = "someth1ng$3ls3"
	go UpperValidator(value, chanTest)
	isValid = <-chanTest
	require.False(t, isValid)
}
