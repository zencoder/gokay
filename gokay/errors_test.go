package gokay_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
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

// GokayErrorsTestSuite
type GokayErrorsTestSuite struct {
	suite.Suite
}

// TestGokayTypesTestSuite
func TestGokayTypesTestSuite(t *testing.T) {
	suite.Run(t, new(GokayErrorsTestSuite))
}

// TestValidate
func (s *GokayErrorsTestSuite) TestValidate() {
	a := HasValidate{}
	b := &HasValidate{}
	c := NoValidate{}

	err := gokay.Validate(a)
	s.Equal(errors.New("Validating 'HasValidate' instance"), err)

	err = gokay.Validate(b)
	s.Equal(errors.New("Validating 'HasValidate' instance"), err)

	err = gokay.Validate(c)
	s.Nil(err)
}

// TestErrorSliceError_Empty
func (s *GokayErrorsTestSuite) TestErrorSliceError_Empty() {
	ea := gokay.ErrorSlice{}

	s.Equal("[]", ea.Error())
}

// TestErrorSliceError_MultiElements
func (s *GokayErrorsTestSuite) TestErrorSliceError_MultiElements() {
	ea := gokay.ErrorSlice{
		errors.New("foo"),
		errors.New("bar"),
		nil,
		gokay.ErrorSlice{errors.New("this is"), errors.New("nested")},
	}

	s.Equal("[\"foo\",\"bar\",null,[\"this is\",\"nested\"]]", ea.Error())
}

// TestErrorMapError_Empty
func (s *GokayErrorsTestSuite) TestErrorMapError_Empty() {
	em := gokay.ErrorMap{}
	s.Equal("{}", em.Error())
}

// TestErrorMapError_NilValue
func (s *GokayErrorsTestSuite) TestErrorMapError_NilValue() {
	em := gokay.ErrorMap{
		"flat":                nil,
		"nestedErrorSlice":    gokay.ErrorSlice{errors.New("this is"), errors.New("nested")},
		"nestedEmptyErrorMap": make(gokay.ErrorMap),
	}

	expectedJSONAsMap := make(map[string]interface{})
	actualJSONAsMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"flat": null,"nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)

	s.Equal(expectedJSONAsMap, actualJSONAsMap)
}

// TestErrorMapError_MultipleValues
func (s *GokayErrorsTestSuite) TestErrorMapError_MultipleValues() {
	em := gokay.ErrorMap{
		"flat":                errors.New(`"flat" "error"`),
		"nestedErrorSlice":    gokay.ErrorSlice{errors.New("this is"), errors.New("nested")},
		"nestedEmptyErrorMap": make(gokay.ErrorMap),
	}

	expectedJSONAsMap := make(map[string]interface{})
	actualJSONAsMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"flat": "\"flat\" \"error\"","nestedErrorSlice": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)

	s.Equal(expectedJSONAsMap, actualJSONAsMap)
}
