package gkgen_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
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
	hv := gkgen.NewHexValidator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexString")

	code, err := hv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsHex(&s.HexString); err != nil {\nerrorsHexString = append(errorsHexString, err)\n}", code)
}

// TestGenerateHexValidationCode_StringPtr
func (s *HexValidatorTestSuite) TestGenerateHexValidationCode_StringPtr() {
	hv := gkgen.NewHexValidator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexStringPtr")
	code, err := hv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsHex(s.HexStringPtr); err != nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, err)\n}", code)
}
