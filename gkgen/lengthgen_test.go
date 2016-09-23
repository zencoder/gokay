package gkgen_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
)

// LengthValidatorTestSuite
type LengthValidatorTestSuite struct {
	suite.Suite
}

// TestLengthValidatorTestSuite
func TestLengthValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(LengthValidatorTestSuite))
}

// TestGenerateValidationCode_String
func (s *LengthValidatorTestSuite) TestGenerateValidationCode_String() {
	lv := gkgen.NewLengthValidator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexString")

	code, err := lv.Generate(et, field, []string{"12"})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.LengthString(12, &s.HexString); err != nil {\nerrorsHexString = append(errorsHexString, err)\n}",
		code)
}

// TestGenerateValidationCode_StringPtr
func (s *LengthValidatorTestSuite) TestGenerateValidationCode_StringPtr() {
	lv := gkgen.NewLengthValidator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexStringPtr")
	code, err := lv.Generate(et, field, []string{"16"})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.LengthString(16, s.HexStringPtr); err != nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, err)\n}",
		code)
}
