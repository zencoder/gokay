package gokay

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthNilString(t *testing.T) {
	expected := int64(10)
	var str *string
	require.NoError(t, LengthString(expected, str))
}

func TestNotLengthNonMatch(t *testing.T) {
	expected := int64(10)
	var str string = "012345678"
	require.Error(t, LengthString(expected, &str))
}

func TestNotLengthMatch(t *testing.T) {
	expected := int64(10)
	var str string = "0123456789"
	require.NoError(t, LengthString(expected, &str))
}

func TestLengthSlice(t *testing.T) {
	expected := int64(10)
	actual := int64(99)
	require.NoError(t, LengthSlice(expected, actual))
}

func TestMinLengthString(t *testing.T) {
	validCases := []struct {
		String    string
		MinLength int64
	}{
		{String: "foo", MinLength: 1},
		{String: "foo", MinLength: 3},
		{String: "", MinLength: 0},
		{String: "a", MinLength: 1},
	}
	for _, tc := range validCases {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			require.NoError(t, MinLengthString(tc.MinLength, &tc.String))
		})
	}

	invalidCases := []struct {
		String    string
		MinLength int64
	}{
		{String: "1", MinLength: 2},
		{String: "a", MinLength: 100},
		{String: "", MinLength: 1},
	}
	for _, tc := range invalidCases {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			require.Error(t, MinLengthString(tc.MinLength, &tc.String))
		})
	}
}
