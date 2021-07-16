package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApi_responseFormater(t *testing.T) {
	msg := "Hello Json"
	result := responseFormater(msg)
	require.Equal(t, string([]byte(`{"message":"Hello Json"}`)), string(result))
}
