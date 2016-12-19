package gokay

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type NoValidate struct{}

type HasValidate struct{}

func (s HasValidate) Validate() error {
	return errors.New("Validating 'HasValidate' instance")
}

func TestValidate(t *testing.T) {
	a := HasValidate{}
	b := &HasValidate{}
	c := NoValidate{}

	err := Validate(a)
	require.Equal(t, errors.New("Validating 'HasValidate' instance"), err)

	err = Validate(b)
	require.Equal(t, errors.New("Validating 'HasValidate' instance"), err)

	err = Validate(c)
	require.NoError(t, err)
}

func TestErrorSliceError_Empty(t *testing.T) {
	ea := ErrorSlice{}

	require.Equal(t, "[]", ea.Error())
}

func TestErrorSliceError_MultiElements(t *testing.T) {
	ea := ErrorSlice{
		errors.New("foo"),
		errors.New("bar"),
		nil,
		ErrorSlice{errors.New("this is"), errors.New("nested")},
	}

	require.Equal(t, "[\"foo\",\"bar\",null,[\"this is\",\"nested\"]]", ea.Error())
}

func TestErrorMapError_Empty(t *testing.T) {
	em := ErrorMap{}
	require.Equal(t, "{}", em.Error())
}

func TestErrorMapError_NilValue(t *testing.T) {
	em := ErrorMap{
		"flat":                nil,
		"nestedErrorSlice":    ErrorSlice{errors.New("this is"), errors.New("nested")},
		"nestedEmptyErrorMap": make(ErrorMap),
	}

	expectedJSONAsMap := make(map[string]interface{})
	actualJSONAsMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"flat": null,"nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)

	require.Equal(t, expectedJSONAsMap, actualJSONAsMap)
}

func TestErrorMapError_MultipleValues(t *testing.T) {
	em := ErrorMap{
		"flat":                errors.New(`"flat" "error"`),
		"nestedErrorSlice":    ErrorSlice{errors.New("this is"), errors.New("nested")},
		"nestedEmptyErrorMap": make(ErrorMap),
	}

	expectedJSONAsMap := make(map[string]interface{})
	actualJSONAsMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"flat": "\"flat\" \"error\"","nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)

	require.Equal(t, expectedJSONAsMap, actualJSONAsMap)
}
