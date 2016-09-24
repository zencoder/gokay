package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

// HexValidatorTestSuite
type HexValidatorTestSuite struct {
	suite.Suite
}

// TestHexValidatorTestSuite
func TestHexValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(HexValidatorTestSuite))
}

// TestGenerateHexValidationCode_String
func (s *HexValidatorTestSuite) TestGenerateHexValidationCode_String() {
	hv := NewHexValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexString")

	code, err := hv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsHex(&s.HexString); err != nil {\nerrorsHexString = append(errorsHexString, err)\n}", code)
}

// TestGenerateHexValidationCode_StringPtr
func (s *HexValidatorTestSuite) TestGenerateHexValidationCode_StringPtr() {
	hv := NewHexValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexStringPtr")
	code, err := hv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsHex(s.HexStringPtr); err != nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, err)\n}", code)
}
