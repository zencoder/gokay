package gkgen_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
)

// UUIDValidatorTestSuite
type UUIDValidatorTestSuite struct {
	suite.Suite
}

// TestUUIDValidatorTestSuite
func TestUUIDValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(UUIDValidatorTestSuite))
}

// TestGenerateUUIDValidationCode_String
func (s *UUIDValidatorTestSuite) TestGenerateUUIDValidationCode_String() {
	v := gkgen.NewUUIDValidator()
	e := UUIDTestStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("UUIDString")

	code, err := v.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsUUID(&s.UUIDString); err != nil {\nerrorsUUIDString = append(errorsUUIDString, err)\n}", code)
}

// TestGenerateUUIDValidationCode_StringPtr
func (s *UUIDValidatorTestSuite) TestGenerateUUIDValidationCode_StringPtr() {
	v := gkgen.NewUUIDValidator()
	e := UUIDTestStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("UUIDStringPtr")
	code, err := v.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsUUID(s.UUIDStringPtr); err != nil {\nerrorsUUIDStringPtr = append(errorsUUIDStringPtr, err)\n}", code)
}

// TestGenerateUUIDValidationCode_NonString
func (s *UUIDValidatorTestSuite) TestGenerateUUIDValidationCode_NonString() {
	v := gkgen.NewUUIDValidator()
	e := UUIDTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("UUIDNonString")
	_, err := v.Generate(et, field, []string{})
	s.Equal(errors.New("UUIDValidator does not support fields of type: 'int'"), err)
}

// UUIDTestStruct
type UUIDTestStruct struct {
	UUIDString    string  `valid:"UUID"`
	UUIDStringPtr *string `valid:"UUID"`
	UUIDNonString int     `valid:"UUID"`
}
