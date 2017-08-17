package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestData struct {
	BCP47String *string
}

func TestIsBCP47_ParamsLength(t *testing.T) {
	params := []string{"I'll break..."}
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47String")
	b := NewBCP47Validator()
	_, err := b.Generate(field, params)
	require.Error(t, err)
}

func TestIsBCP47_FieldPtr(t *testing.T) {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonString")
	b := NewBCP47Validator()
	_, err := b.Generate(field, []string{})
	require.Error(t, err)
}

func TestIsBCP47_FieldNestedPtr(t *testing.T) {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonStringPtr")
	b := NewBCP47Validator()
	_, err := b.Generate(field, []string{})
	require.Error(t, err)
}

func TestGenerateBCP47ValidationCode_String(t *testing.T) {
	hv := NewBCP47Validator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47String")

	code, err := hv.Generate(field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"if err := gokay.IsBCP47(&s.BCP47String); err != nil {\nerrorsBCP47String = append(errorsBCP47String, err)\n}",
		code,
	)
}

func TestGenerateBCP47ValidationCode_StringPtr(t *testing.T) {
	hv := NewBCP47Validator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47StringPtr")
	code, err := hv.Generate(field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"if err := gokay.IsBCP47(s.BCP47StringPtr); err != nil {\nerrorsBCP47StringPtr = append(errorsBCP47StringPtr, err)\n}",
		code,
	)
}
