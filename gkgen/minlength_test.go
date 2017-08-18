package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinLengthGenerator(t *testing.T) {
	type minLengthStruct struct {
		Bool      bool
		BoolPtr   *bool
		String    string
		StringPtr *string
	}

	t.Run("RequireParam", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		s := minLengthStruct{String: "foo"}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("String")
		assert.True(found)

		_, err := validator.Generate(field, []string{})
		assert.Error(err)
	})

	t.Run("RequireParamIsANumber", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		s := minLengthStruct{String: "foo"}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("String")
		assert.True(found)

		_, err := validator.Generate(field, []string{"wat"})
		assert.Error(err)
	})

	t.Run("RejectUnsupportedType", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		s := minLengthStruct{Bool: true}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("Bool")
		assert.True(found)

		_, err := validator.Generate(field, []string{"3"})
		assert.Error(err)
	})

	t.Run("RejectUnsupportedTypePointer", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		var foo = true
		s := minLengthStruct{BoolPtr: &foo}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("BoolPtr")
		assert.True(found)

		_, err := validator.Generate(field, []string{"3"})
		assert.Error(err)
	})

	t.Run("String", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		s := minLengthStruct{String: "foo"}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("String")
		assert.True(found)

		code, err := validator.Generate(field, []string{"3"})
		assert.NoError(err)

		assert.Equal("if err := gokay.MinLengthString(3, &s.String); err != nil {errorsString = append(errorsString, err)}", collapseToOneLine(code))
	})

	t.Run("StringPtr", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		var foo = "foo"
		s := minLengthStruct{StringPtr: &foo}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("StringPtr")
		assert.True(found)

		code, err := validator.Generate(field, []string{"3"})
		assert.NoError(err)

		assert.Equal("if err := gokay.MinLengthString(3, s.StringPtr); err != nil {errorsStringPtr = append(errorsStringPtr, err)}", collapseToOneLine(code))
	})
}

func collapseToOneLine(str string) string {
	s := strings.Replace(strings.TrimSpace(str), "\t", "", -1)
	return strings.Replace(strings.TrimSpace(s), "\n", "", -1)
}
