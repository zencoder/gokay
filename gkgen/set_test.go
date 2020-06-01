package gkgen

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	nv := NewSetValidator()
	e := SetTestStruct{}
	et := reflect.TypeOf(e)

	field, ok := et.FieldByName("SetString")
	require.True(t, ok)
	code, err := nv.Generate(et, field, []string{"cat"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != \"\" && !(s.%[1]s == \"cat\") {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal cat\"))\n}", field.Name),
		code,
	)

	code, err = nv.Generate(et, field, []string{"cat", "dog", "mouse"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != \"\" && !(s.%[1]s == \"cat\" || s.%[1]s == \"dog\" || s.%[1]s == \"mouse\") {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal cat or dog or mouse\"))\n}", field.Name),
		code,
	)
}

func TestSetInts(t *testing.T) {
	nv := NewSetValidator()
	e := SetTestStruct{}
	et := reflect.TypeOf(e)

	field, ok := et.FieldByName("SetInt")
	require.True(t, ok)
	code, err := nv.Generate(et, field, []string{"3"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != \"\" && !(s.%[1]s == 3) {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal 3\"))\n}", field.Name),
		code,
	)

	code, err = nv.Generate(et, field, []string{"1", "3", "7"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != \"\" && !(s.%[1]s == 1 || s.%[1]s == 3 || s.%[1]s == 7) {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal 1 or 3 or 7\"))\n}", field.Name),
		code,
	)
}

func TestSetPointer(t *testing.T) {
	nv := NewSetValidator()
	e := SetTestStruct{}
	et := reflect.TypeOf(e)

	field, ok := et.FieldByName("SetStringPtr")
	require.True(t, ok)
	code, err := nv.Generate(et, field, []string{"cat"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != nil && !(*s.%[1]s == \"cat\") {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal cat\"))\n}", field.Name),
		code,
	)

	code, err = nv.Generate(et, field, []string{"cat", "dog", "mouse"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != nil && !(*s.%[1]s == \"cat\" || *s.%[1]s == \"dog\" || *s.%[1]s == \"mouse\") {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal cat or dog or mouse\"))\n}", field.Name),
		code,
	)
}

func TestSetIntPointer(t *testing.T) {
	nv := NewSetValidator()
	e := SetTestStruct{}
	et := reflect.TypeOf(e)

	field, ok := et.FieldByName("SetIntPtr")
	require.True(t, ok)
	code, err := nv.Generate(et, field, []string{"1"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != nil && !(*s.%[1]s == 1) {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal 1\"))\n}", field.Name),
		code,
	)

	code, err = nv.Generate(et, field, []string{"1", "3", "7"})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(
		t,
		fmt.Sprintf("if s.%[1]s != nil && !(*s.%[1]s == 1 || *s.%[1]s == 3 || *s.%[1]s == 7) {\nerrors%[1]s = append(errors%[1]s, errors.New(\"%[1]s must equal 1 or 3 or 7\"))\n}", field.Name),
		code,
	)
}

func TestSetInvalidTypes(t *testing.T) {
	nv := NewSetValidator()
	e := NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilMap")
	_, err := nv.Generate(et, field, []string{"0"})
	require.Error(t, err)
	require.Equal(
		t,
		"Set does not work on type 'map'",
		err.Error(),
	)

	field, _ = et.FieldByName("NotNilSlice")
	_, err = nv.Generate(et, field, []string{"test"})
	require.Error(t, err)
	require.Equal(
		t,
		"Set does not work on type 'slice'",
		err.Error(),
	)
}

func TestSetInvalidParameters(t *testing.T) {
	nv := NewSetValidator()
	e := SetTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("SetString")
	_, err := nv.Generate(et, field, []string{})
	require.Error(t, err)
	require.Equal(
		t,
		"Set validation requires at least 1 parameter",
		err.Error(),
	)
}
