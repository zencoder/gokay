package gokay

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIsBCP47_Nil
func TestIsBCP47_Nil(t *testing.T) {
	require.NoError(t, IsBCP47(nil))
}

// TestIsBCP47_English
func TestIsBCP47_English(t *testing.T) {
	str := "English"
	require.EqualError(t, IsBCP47(&str), "language: tag is not well-formed")
}

// TestIsBCP47_en
func TestIsBCP47_en(t *testing.T) {
	str := "en"
	require.NoError(t, IsBCP47(&str))
}
