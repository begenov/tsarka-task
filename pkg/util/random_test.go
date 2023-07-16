package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	str := RandomString(10)
	assert.Equal(t, 10, len(str))
}

func TestIsValid(t *testing.T) {
	valid := IsValid("991231123456")
	require.True(t, valid)

	invalid := IsValid("123456789012")
	require.False(t, invalid)
}

func TestIsValid_InvalidInput(t *testing.T) {
	invalid := IsValid("not_a_valid_iin")
	assert.False(t, invalid)
}
