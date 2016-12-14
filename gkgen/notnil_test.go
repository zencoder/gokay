package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNotNil
func TestNotNil(t *testing.T) {
	nv := NewNotNilValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("HexStringPtr")
	code, err := nv.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"if s.HexStringPtr == nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, errors.New(\"is Nil\"))\n}",
		code,
	)
}

// TestNotNil_Map
func TestNotNil_Map(t *testing.T) {
	nv := NewNotNilValidator()
	e := NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilMap")
	code, err := nv.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"if s.NotNilMap == nil {\nerrorsNotNilMap = append(errorsNotNilMap, errors.New(\"is Nil\"))\n}",
		code,
	)
}

// TestNotNil_Slice
func TestNotNil_Slice(t *testing.T) {
	nv := NewNotNilValidator()
	e := NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilSlice")
	code, err := nv.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		"if s.NotNilSlice == nil {\nerrorsNotNilSlice = append(errorsNotNilSlice, errors.New(\"is Nil\"))\n}",
		code,
	)
}
