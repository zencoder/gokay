package gokay

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNilString
func TestLengthNilString(t *testing.T) {
	var expected = int64(10)
	var str *string
	require.NoError(t, LengthString(expected, str))
}

// TestNotLengthNonMatch
func TestNotLengthNonMatch(t *testing.T) {
	var expected = int64(10)
	var str string = "012345678"
	require.Error(t, LengthString(expected, &str))
}

// TestNotLengthMatch
func TestNotLengthMatch(t *testing.T) {
	var expected = int64(10)
	var str string = "0123456789"
	require.NoError(t, LengthString(expected, &str))
}

// TestLengthSlice
func TestLengthSlice(t *testing.T) {
	var expected = int64(10)
	var actual = int64(99)
	require.NoError(t, LengthSlice(expected, actual))
}
