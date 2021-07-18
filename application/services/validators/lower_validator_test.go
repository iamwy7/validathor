package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_LowerValidator(t *testing.T) {
	value := "SOMEThING"
	chanTest := make(chan bool)
	go LowerValidator(value, chanTest)
	isValid := <-chanTest
	require.True(t, isValid)

	value = "EXAMPLE"
	go LowerValidator(value, chanTest)
	isValid = <-chanTest
	require.False(t, isValid)
}
