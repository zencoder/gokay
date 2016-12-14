package gokay

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsUUID_No0x(t *testing.T) {
	str := "603c9a2a-38dB-4987-932a-2f57733a29f1"
	require.NoError(t, IsUUID(&str))
}

// TestNilUUID tests to see if the string passed in is nil
func TestNilUUID(t *testing.T) {
	var str *string
	require.NoError(t, IsUUID(str))
}

// TestIsUUID_NotMatch
func TestIsUUID_NotMatch(t *testing.T) {
	str := "603c9a2a-38db-4987-932a-2f57733a29fQ"
	require.EqualError(t, IsUUID(&str), "'603c9a2a-38db-4987-932a-2f57733a29fQ' is not a UUID")
}

func TestIsUUID_NotUUIDTooLong(t *testing.T) {
	str := "AB603c9a2a-38db-4987-932a-2f57733a29fQ"
	require.EqualError(t, IsUUID(&str), "'AB603c9a2a-38db-4987-932a-2f57733a29fQ' is not a UUID")
}
