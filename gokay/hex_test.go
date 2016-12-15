package gokay

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNilString(t *testing.T) {
	var str *string
	require.NoError(t, IsHex(str))
}

func TestIsHex_No0x(t *testing.T) {
	str := "1a3F"
	require.NoError(t, IsHex(&str))
}

func TestIsHex_0x(t *testing.T) {
	str := "0x1a3F"
	require.NoError(t, IsHex(&str))
}

func TestIsHex_NotHex(t *testing.T) {
	str := "0x1Gbcq"
	require.EqualError(t, IsHex(&str), "'0x1Gbcq' is not a hexadecimal string")
}
