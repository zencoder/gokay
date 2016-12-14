package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGenerateHexValidationCode_String
func TestGenerateHexValidationCode_String(t *testing.T) {
	hv := NewHexValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexString")

	code, err := hv.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(t, "if err := gokay.IsHex(&s.HexString); err != nil {\nerrorsHexString = append(errorsHexString, err)\n}", code)
}

// TestGenerateHexValidationCode_StringPtr
func TestGenerateHexValidationCode_StringPtr(t *testing.T) {
	hv := NewHexValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexStringPtr")
	code, err := hv.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(t, "if err := gokay.IsHex(s.HexStringPtr); err != nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, err)\n}", code)
}
