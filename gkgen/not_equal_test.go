package gkgen

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotEqual(t *testing.T) {
	nv := NewNotEqualValidator()
	e := NotEqualTestStruct{}
	et := reflect.TypeOf(e)

	for i := 0; i < et.NumField()-1; i++ {
		field := et.Field(i)
		expectedComparer := `0`
		if field.Name == "NotEqualString" {
			expectedComparer = `""`
		}
		code, err := nv.Generate(et, field, []string{expectedComparer})
		require.NoError(t, err)
		code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
		require.Equal(
			t,
			fmt.Sprintf("if s.%[1]s == %[2]s {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s cannot equal '%[2]s'\"))\n}", field.Name, expectedComparer),
			code,
		)
	}
}

func TestNotEqualPointer(t *testing.T) {
	nv := NewNotEqualValidator()
	e := NotEqualTestPointerStruct{}
	et := reflect.TypeOf(e)

	for i := 0; i < et.NumField()-1; i++ {
		field := et.Field(i)
		expectedComparer := `0`
		if field.Name == "NotEqualString" {
			expectedComparer = `""`
		}
		code, err := nv.Generate(et, field, []string{expectedComparer})
		require.NoError(t, err)
		code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
		require.Equal(
			t,
			fmt.Sprintf("if s.%[1]s != nil && *s.%[1]s == %[2]s {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s cannot equal '%[2]s'\"))\n}", field.Name, expectedComparer),
			code,
		)
	}
}

func TestNotEqualInvalidTypes(t *testing.T) {
	nv := NewNotEqualValidator()
	e := NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilMap")
	_, err := nv.Generate(et, field, []string{"0"})
	require.Error(t, err)
	require.Equal(
		t,
		"NotEqual does not work on type 'map'",
		err.Error(),
	)

	field, _ = et.FieldByName("NotNilSlice")
	_, err = nv.Generate(et, field, []string{"test"})
	require.Error(t, err)
	require.Equal(
		t,
		"NotEqual does not work on type 'slice'",
		err.Error(),
	)
}

func TestNotEqualInvalidParameters(t *testing.T) {
	nv := NewNotEqualValidator()
	e := NotEqualTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotEqualInt")
	_, err := nv.Generate(et, field, []string{})
	require.Error(t, err)
	require.Equal(
		t,
		"NotEqual validation requires exactly 1 parameter",
		err.Error(),
	)

	field, _ = et.FieldByName("NotEqualInt")
	_, err = nv.Generate(et, field, []string{"0", "-0"})
	require.Error(t, err)
	require.Equal(
		t,
		"NotEqual validation requires exactly 1 parameter",
		err.Error(),
	)
}
