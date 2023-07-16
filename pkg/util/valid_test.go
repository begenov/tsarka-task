package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxLengthSubstring(t *testing.T) {
	str := MaxLengthSubstring("abcabcbb")
	require.Equal(t, "abc", str)

	str = MaxLengthSubstring("bbbbb")
	require.Equal(t, "b", str)

	str = MaxLengthSubstring("pwwkew")
	require.Equal(t, "wke", str)

	str = MaxLengthSubstring("a")
	require.Equal(t, "a", str)

	str = MaxLengthSubstring("")
	require.Equal(t, "", str)
}

func TestEmailsCheck(t *testing.T) {
	input := []string{
		"test@example.com",
		"invalid_email",
		"another@example.com",
		"not_an_email",
	}

	expected := []string{
		"test@example.com",
		"another@example.com",
	}

	result := EmailsCheck(input)
	require.Equal(t, expected, result)
}

func TestInnCheck(t *testing.T) {
	input := []string{
		"011203432542",
		"987654321098",
		"invalid_inn",
		"031103432335",
	}

	expected := []string{
		"011203432542",
		"031103432335",
	}

	result := InnCheck(input)
	require.Equal(t, expected, result)
}
