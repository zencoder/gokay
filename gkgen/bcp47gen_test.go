package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBCP47_ParamsLengthSuccess(t *testing.T) {
	params := []string{"I'll work..."}
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47String")
	b := NewBCP47Validator()
	_, err := b.Generate(et, field, params)
	require.NoError(t, err)
}

func TestIsBCP47_ParamsLengthFail(t *testing.T) {
	params := []string{"I'll work...", "I'll fail..."}
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47String")
	b := NewBCP47Validator()
	_, err := b.Generate(et, field, params)
	require.Error(t, err)
}

func TestIsBCP47_FieldPtr(t *testing.T) {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonString")
	b := NewBCP47Validator()
	_, err := b.Generate(et, field, []string{})
	require.Error(t, err)
}

func TestIsBCP47_FieldNestedPtr(t *testing.T) {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonStringPtr")
	b := NewBCP47Validator()
	_, err := b.Generate(et, field, []string{})
	require.Error(t, err)
}

func TestGenerateBCP47ValidationCode_String(t *testing.T) {
	hv := NewBCP47Validator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47String")

	code, err := hv.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"bcp47Exceptions := map[string]bool {\n}\nif err := gokay.IsBCP47(&s.BCP47String); err != nil && !bcp47Exceptions[s.BCP47String] {\nerrorsBCP47String = append(errorsBCP47String, err)\n}",
		code,
	)
}

func TestGenerateBCP47ValidationCode_WithExceptions(t *testing.T) {
	hv := NewBCP47Validator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47String")

	code, err := hv.Generate(et, field, []string{"except=en-AB,en-WL"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"bcp47Exceptions := map[string]bool {\n\"en-AB\": true,\n\"en-WL\": true,\n}\nif err := gokay.IsBCP47(&s.BCP47String); err != nil && !bcp47Exceptions[s.BCP47String] {\nerrorsBCP47String = append(errorsBCP47String, err)\n}",
		code,
	)
}

func TestGenerateBCP47ValidationCode_StringPtr(t *testing.T) {
	hv := NewBCP47Validator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47StringPtr")
	code, err := hv.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"bcp47Exceptions := map[string]bool {\n}\nif err := gokay.IsBCP47(s.BCP47StringPtr); err != nil && !bcp47Exceptions[*s.BCP47StringPtr] {\nerrorsBCP47StringPtr = append(errorsBCP47StringPtr, err)\n}",
		code,
	)
}
