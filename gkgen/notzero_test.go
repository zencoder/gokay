package gkgen

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotZero(t *testing.T) {
	nv := NewNotZeroValidator()
	e := NotZeroTestStruct{}
	et := reflect.TypeOf(e)

	for i := 0; i < et.NumField(); i++ {
		field := et.Field(i)
		code, err := nv.Generate(et, field, []string{})
		require.NoError(t, err)
		code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
		require.Equal(
			t,
			fmt.Sprintf("if s.%[1]s == 0 {\nerrors%[1]s = append(errors%[1]s, errors.New(\"is Zero\"))\n}", field.Name),
			code,
		)
	}
}

func TestNotZeroPointer(t *testing.T) {
	nv := NewNotZeroValidator()
	e := NotZeroTestPointerStruct{}
	et := reflect.TypeOf(e)

	for i := 0; i < et.NumField(); i++ {
		field := et.Field(i)
		code, err := nv.Generate(et, field, []string{})
		require.NoError(t, err)
		code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
		require.Equal(
			t,
			fmt.Sprintf("if s.%[1]s != nil && *s.%[1]s == 0 {\nerrors%[1]s = append(errors%[1]s, errors.New(\"is Zero\"))\n}", field.Name),
			code,
		)
	}
}

func TestNotZeroInvalidTypes(t *testing.T) {
	nv := NewNotZeroValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("HexStringPtr")
	_, err := nv.Generate(et, field, []string{})
	require.Error(t, err)
	require.Equal(
		t,
		"NotZero only works on number and number pointer type Fields. HexStringPtr has type 'string'",
		err.Error(),
	)

	field, _ = et.FieldByName("HexString")
	_, err = nv.Generate(et, field, []string{})
	require.Error(t, err)
	require.Equal(
		t,
		"NotZero only works on number and number pointer type Fields. HexString has type 'string'",
		err.Error(),
	)
}

func TestNotZeroInvalidParameters(t *testing.T) {
	nv := NewNotZeroValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("HexStringPtr")
	_, err := nv.Generate(et, field, []string{"42"})
	require.Error(t, err)
	require.Equal(
		t,
		"NotZero takes no parameters",
		err.Error(),
	)
}
