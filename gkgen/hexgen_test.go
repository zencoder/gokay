package gkgen_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
)

type HexValidatorTestSuite struct {
	suite.Suite
}

func TestHexValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(HexValidatorTestSuite))
}

func (s *HexValidatorTestSuite) TestGenerateHexValidationCode_String() {
	hv := gkgen.NewHexValidator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexString")

	code, err := hv.GenerateValidationCode(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsHex(&s.HexString); err != nil {\nerrorsHexString = append(errorsHexString, err)\n}",
		code)
}

func (s *HexValidatorTestSuite) TestGenerateHexValidationCode_StringPtr() {
	hv := gkgen.NewHexValidator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("HexStringPtr")
	code, err := hv.GenerateValidationCode(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsHex(s.HexStringPtr); err != nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, err)\n}",
		code)
}
