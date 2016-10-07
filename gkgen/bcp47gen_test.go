package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestData struct {
	BCP47String *string
}

// BCP47ValidatorTestSuite
type BCP47ValidatorTestSuite struct {
	suite.Suite
}

// TestBCP47ValidatorTestSuite
func TestBCP47ValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(BCP47ValidatorTestSuite))
}

// TestIsBCP47_ParamsLength
func (s *BCP47ValidatorTestSuite) TestIsBCP47_ParamsLength() {
	params := []string{"I'll break..."}
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47String")
	b := NewBCP47Validator()
	_, err := b.Generate(et, field, params)
	s.Require().Error(err)
}

// TestIsBCP47_FieldPtr
func (s *BCP47ValidatorTestSuite) TestIsBCP47_FieldPtr() {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonString")
	b := NewBCP47Validator()
	_, err := b.Generate(et, field, []string{})
	s.Require().Error(err)
}

// TestIsBCP47_FieldNestedPtr
func (s *BCP47ValidatorTestSuite) TestIsBCP47_FieldNestedPtr() {
	et := reflect.TypeOf(ExampleStruct{})
	field, _ := et.FieldByName("BCP47NonStringPtr")
	b := NewBCP47Validator()
	_, err := b.Generate(et, field, []string{})
	s.Require().Error(err)
}

// TestGenerateHexValidationCode_String
func (s *BCP47ValidatorTestSuite) TestGenerateHexValidationCode_String() {
	hv := NewBCP47Validator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47String")

	code, err := hv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsBCP47(&s.BCP47String); err != nil {\nerrorsBCP47String = append(errorsBCP47String, err)\n}", code)
}

// TestGenerateHexValidationCode_StringPtr
func (s *BCP47ValidatorTestSuite) TestGenerateHexValidationCode_StringPtr() {
	hv := NewBCP47Validator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47StringPtr")
	code, err := hv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsBCP47(s.BCP47StringPtr); err != nil {\nerrorsBCP47StringPtr = append(errorsBCP47StringPtr, err)\n}", code)
}
