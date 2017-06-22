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

var (
	testHex = "0x1234567890ABCDEF"
)

func devNull(i interface{}) {}

func BenchmarkHexV1(b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		err = IsHex(&testHex)
	}
	devNull(err)
}

func BenchmarkHexV2(b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		err = IsHexV2(&testHex)
	}
	devNull(err)
}

func BenchmarkHexV3(b *testing.B) {
	var err error
	for n := 0; n < b.N; n++ {
		err = IsHexV3(&testHex)
	}
	devNull(err)
}
