package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_WhiteSpacesValidator(t *testing.T) {
	value := "something"
	chanTest := make(chan bool)
	go WhiteSpacesValidator(value, chanTest)
	isValid := <-chanTest
	require.True(t, isValid)

	value = "007exa mple"
	go WhiteSpacesValidator(value, chanTest)
	isValid = <-chanTest
	require.False(t, isValid)
}
