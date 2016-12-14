package gkgen

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAddValidation
func TestAddValidation(t *testing.T) {
	v := ValidateGenerator{
		Generators: make(map[string]Generater),
	}
	generator := NewNotNilValidator()

	_, ok := v.Generators[generator.Name()]
	require.False(t, ok)

	v.AddValidation(generator)
	_, ok = v.Generators[generator.Name()]
	require.True(t, ok)
}

// TestExampleStruct tests single no-param validation
func TestExampleStruct(t *testing.T) {
	out := &bytes.Buffer{}
	key := "abc123"
	e := ExampleStruct{
		HexStringPtr: &key,
	}

	v := NewValidateGenerator()

	err := v.Generate(out, e)
	require.NoError(t, err)
}

// UnknownTagStruct
type UnknownTagStruct struct {
	Field string `valid:"Length=(5),Unknown"`
}

// TestGenerateWithUnknownTag
func TestGenerateWithUnknownTag(t *testing.T) {
	out := &bytes.Buffer{}
	v := NewValidateGenerator()
	err := v.Generate(out, UnknownTagStruct{})
	require.Equal(t, errors.New("Unknown validation generator name: 'Unknown'"), err)
}

// TestGenerateMapValidationCodeNonArrayOrSlice
func TestGenerateMapValidationCodeNonArrayOrSlice(t *testing.T) {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonString")
	out := &bytes.Buffer{}
	err := generateMapValidationCode(out, field.Type, "BCP47NonString", int64(1))
	require.Error(t, err)
}

// TestGenerateSliceValidationCodeNonSlice
func TestGenerateSliceValidationCodeNonSlice(t *testing.T) {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonString")
	out := &bytes.Buffer{}
	err := generateSliceValidationCode(out, field.Type, "BCP47NonString", int64(1))
	require.Error(t, err)
}

// TestGenerateSliceValidationCodeNonSlice
func TestGenerateSliceValidationCodeSlice(t *testing.T) {
	et := reflect.TypeOf(NotNilTestStruct{})
	field, _ := et.FieldByName("NotNilSlice")
	out := &bytes.Buffer{}
	err := generateSliceValidationCode(out, field.Type, field.Name, int64(1))
	require.Error(t, err)
}
