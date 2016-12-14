package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGenerateValidationCode_String
func TestGenerateValidationCode_String(t *testing.T) {
	lv := NewLengthValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexString")

	code, err := lv.Generate(et, field, []string{"12"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"if err := gokay.LengthString(12, &s.HexString); err != nil {\nerrorsHexString = append(errorsHexString, err)\n}",
		code,
	)
}

// TestGenerateValidationCode_StringPtr
func TestGenerateValidationCode_StringPtr(t *testing.T) {
	lv := NewLengthValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexStringPtr")
	code, err := lv.Generate(et, field, []string{"16"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"if err := gokay.LengthString(16, s.HexStringPtr); err != nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, err)\n}",
		code,
	)
}
