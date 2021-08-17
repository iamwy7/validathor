package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_MinLenghtValidator(t *testing.T) {
	value := "SOMEThING123890"
	chanTest := make(chan bool)
	go MinLenghtValidator(value, chanTest)
	isValid := <-chanTest
	require.True(t, isValid)

	value = "EXAMPLE!"
	go MinLenghtValidator(value, chanTest)
	isValid = <-chanTest

	require.False(t, isValid)
}
