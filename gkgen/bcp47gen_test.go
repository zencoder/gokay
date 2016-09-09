package gkgen_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
)

type BCP47ValidatorTestSuite struct {
	suite.Suite
}

func TestBCP47ValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(BCP47ValidatorTestSuite))
}

func (s *BCP47ValidatorTestSuite) TestGenerateHexValidationCode_String() {
	hv := gkgen.NewBCP47Validator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47String")

	code, err := hv.GenerateValidationCode(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsBCP47(&s.BCP47String); err != nil {\nerrorsBCP47String = append(errorsBCP47String, err)\n}",
		code)
}

func (s *BCP47ValidatorTestSuite) TestGenerateHexValidationCode_StringPtr() {
	hv := gkgen.NewBCP47Validator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("BCP47StringPtr")
	code, err := hv.GenerateValidationCode(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if err := gokay.IsBCP47(s.BCP47StringPtr); err != nil {\nerrorsBCP47StringPtr = append(errorsBCP47StringPtr, err)\n}",
		code)
}
