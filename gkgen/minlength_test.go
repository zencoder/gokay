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
		Slice     []string
		SlicePtr  *[]string
		Map       map[string]string
		MapPtr    *map[string]string
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

	t.Run("Slice", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		s := minLengthStruct{Slice: []string{"foo"}}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("Slice")
		assert.True(found)

		code, err := validator.Generate(field, []string{"3"})
		assert.NoError(err)

		assert.Equal("if err := gokay.MinLengthSlice(3, &s.Slice); err != nil {errorsSlice = append(errorsSlice, err)}", collapseToOneLine(code))
	})

	t.Run("SlicePtr", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		var foo = []string{"foo"}
		s := minLengthStruct{SlicePtr: &foo}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("SlicePtr")
		assert.True(found)

		code, err := validator.Generate(field, []string{"3"})
		assert.NoError(err)

		assert.Equal("if err := gokay.MinLengthSlice(3, s.SlicePtr); err != nil {errorsSlicePtr = append(errorsSlicePtr, err)}", collapseToOneLine(code))
	})

	t.Run("Map", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		s := minLengthStruct{Map: map[string]string{"foo": "foo"}}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("Map")
		assert.True(found)

		code, err := validator.Generate(field, []string{"3"})
		assert.NoError(err)

		assert.Equal("if err := gokay.MinLengthMap(3, &s.Map); err != nil {errorsMap = append(errorsMap, err)}", collapseToOneLine(code))
	})

	t.Run("MapPtr", func(t *testing.T) {
		t.Parallel()
		assert := require.New(t)
		validator := NewMinLengthValidator()
		var foo = map[string]string{"foo": "foo"}
		s := minLengthStruct{MapPtr: &foo}
		st := reflect.TypeOf(s)
		field, found := st.FieldByName("MapPtr")
		assert.True(found)

		code, err := validator.Generate(field, []string{"3"})
		assert.NoError(err)

		assert.Equal("if err := gokay.MinLengthMap(3, s.MapPtr); err != nil {errorsMapPtr = append(errorsMapPtr, err)}", collapseToOneLine(code))
	})
}

func collapseToOneLine(str string) string {
	s := strings.Replace(strings.TrimSpace(str), "\t", "", -1)
	return strings.Replace(strings.TrimSpace(s), "\n", "", -1)
}
