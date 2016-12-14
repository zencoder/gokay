package gokay_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zencoder/gokay/gokay"
)

// NoValidate
type NoValidate struct{}

// HasValidate
type HasValidate struct{}

// Validate
func (s HasValidate) Validate() error {
	return errors.New("Validating 'HasValidate' instance")
}

// TestValidate
func TestValidate(t *testing.T) {
	a := HasValidate{}
	b := &HasValidate{}
	c := NoValidate{}

	err := gokay.Validate(a)
	require.Equal(t, errors.New("Validating 'HasValidate' instance"), err)

	err = gokay.Validate(b)
	require.Equal(t, errors.New("Validating 'HasValidate' instance"), err)

	err = gokay.Validate(c)
	require.NoError(t, err)
}

// TestErrorSliceError_Empty
func TestErrorSliceError_Empty(t *testing.T) {
	ea := gokay.ErrorSlice{}

	require.Equal(t, "[]", ea.Error())
}

// TestErrorSliceError_MultiElements
func TestErrorSliceError_MultiElements(t *testing.T) {
	ea := gokay.ErrorSlice{
		errors.New("foo"),
		errors.New("bar"),
		nil,
		gokay.ErrorSlice{errors.New("this is"), errors.New("nested")},
	}

	require.Equal(t, "[\"foo\",\"bar\",null,[\"this is\",\"nested\"]]", ea.Error())
}

// TestErrorMapError_Empty
func TestErrorMapError_Empty(t *testing.T) {
	em := gokay.ErrorMap{}
	require.Equal(t, "{}", em.Error())
}

// TestErrorMapError_NilValue
func TestErrorMapError_NilValue(t *testing.T) {
	em := gokay.ErrorMap{
		"flat":                nil,
		"nestedErrorSlice":    gokay.ErrorSlice{errors.New("this is"), errors.New("nested")},
		"nestedEmptyErrorMap": make(gokay.ErrorMap),
	}

	expectedJSONAsMap := make(map[string]interface{})
	actualJSONAsMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"flat": null,"nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)

	require.Equal(t, expectedJSONAsMap, actualJSONAsMap)
}

// TestErrorMapError_MultipleValues
func TestErrorMapError_MultipleValues(t *testing.T) {
	em := gokay.ErrorMap{
		"flat":                errors.New(`"flat" "error"`),
		"nestedErrorSlice":    gokay.ErrorSlice{errors.New("this is"), errors.New("nested")},
		"nestedEmptyErrorMap": make(gokay.ErrorMap),
	}

	expectedJSONAsMap := make(map[string]interface{})
	actualJSONAsMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"flat": "\"flat\" \"error\"","nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)

	require.Equal(t, expectedJSONAsMap, actualJSONAsMap)
}
