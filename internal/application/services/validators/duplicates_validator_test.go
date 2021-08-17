package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators_DuplicatesValidator(t *testing.T) {
	value := "something"
	chanTest := make(chan bool)
	go DuplicatesValidator(value, chanTest)
	isValid := <-chanTest
	require.True(t, isValid)

	value = "007example"
	go DuplicatesValidator(value, chanTest)
	isValid = <-chanTest
	require.False(t, isValid)
}
