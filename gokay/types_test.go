package gokay_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gokay"
)

type NoValidate struct{}

type HasValidate struct{}

func (s HasValidate) Validate() error {
	return errors.New("Validating 'HasValidate' instance")
}

type GokayTypesTestSuite struct {
	suite.Suite
}

func TestGokayTypesTestSuite(t *testing.T) {
	suite.Run(t, new(GokayTypesTestSuite))
}

func (s *GokayTypesTestSuite) TestValidate() {
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

func (s *GokayTypesTestSuite) TestErrorArrayError_Empty() {
	ea := gokay.ErrorArray{}

	s.Equal("[]", ea.Error())
}

func (s *GokayTypesTestSuite) TestErrorArrayError_MultiElements() {
	ea := gokay.ErrorArray{
		errors.New("foo"),
		errors.New("bar"),
		nil,
		gokay.ErrorArray{errors.New("this is"), errors.New("nested")},
	}

	s.Equal("[\"foo\",\"bar\",null,[\"this is\",\"nested\"]]", ea.Error())
}

func (s *GokayTypesTestSuite) TestErrorMapError_Empty() {
	em := gokay.ErrorMap{}

	s.Equal("{}", em.Error())
}

func (s *GokayTypesTestSuite) TestErrorMapError_MultipleValues() {
	em := gokay.ErrorMap{
		"flat":                errors.New(`"flat" "error"`),
		"nestedErrorArray":    gokay.ErrorArray{errors.New("this is"), errors.New("nested")},
		"nestedEmptyErrorMap": make(gokay.ErrorMap),
	}

	expectedJSONAsMap := make(map[string]interface{})
	actualJSONAsMap := make(map[string]interface{})
	json.Unmarshal([]byte(`{"flat": "\"flat\" \"error\"","nestedErrorArray": ["this is","nested"],"nestedEmptyErrorMap": {}}`), &expectedJSONAsMap)
	json.Unmarshal([]byte(em.Error()), &actualJSONAsMap)

	s.Equal(expectedJSONAsMap, actualJSONAsMap)
}
