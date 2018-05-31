package gkgen

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEquals(t *testing.T) {
	nv := NewEqualsValidator()
	e := EqualsTestStruct{}
	et := reflect.TypeOf(e)

	field, ok := et.FieldByName("EqualsString")
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

func TestEqualsPointer(t *testing.T) {
	nv := NewEqualsValidator()
	e := EqualsTestStruct{}
	et := reflect.TypeOf(e)

	field, ok := et.FieldByName("EqualsStringPtr")
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

func TestEqualsInvalidTypes(t *testing.T) {
	nv := NewEqualsValidator()
	e := NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilMap")
	_, err := nv.Generate(et, field, []string{"0"})
	require.Error(t, err)
	require.Equal(
		t,
		"Equals does not work on type 'map'",
		err.Error(),
	)

	field, _ = et.FieldByName("NotNilSlice")
	_, err = nv.Generate(et, field, []string{"test"})
	require.Error(t, err)
	require.Equal(
		t,
		"Equals does not work on type 'slice'",
		err.Error(),
	)
}

func TestEqualsInvalidParameters(t *testing.T) {
	nv := NewEqualsValidator()
	e := EqualsTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("EqualsString")
	_, err := nv.Generate(et, field, []string{})
	require.Error(t, err)
	require.Equal(
		t,
		"Equals validation requires at least 1 parameter",
		err.Error(),
	)
}
