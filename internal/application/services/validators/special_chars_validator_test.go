package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_SpecialCharsValidator(t *testing.T) {
	value := "SOM+T°hING190"
	chanTest := make(chan bool)
	go SpecialCharsValidator(value, chanTest)
	isValid := <-chanTest
	require.True(t, isValid)

	value = "EXAMPLE°123sss"
	go SpecialCharsValidator(value, chanTest)
	isValid = <-chanTest
	require.False(t, isValid)
}
