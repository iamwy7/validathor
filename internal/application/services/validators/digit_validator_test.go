package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_DigitValidator(t *testing.T) {
	value := "something"
	chanTest := make(chan bool)
	go DigitValidator(value, chanTest)
	isValid := <-chanTest
	require.False(t, isValid)

	value = "007example"
	go DigitValidator(value, chanTest)
	isValid = <-chanTest
	require.True(t, isValid)
}
