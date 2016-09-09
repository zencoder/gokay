package gkgen_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
)

type NotNilTestSuite struct {
	suite.Suite
}

func TestNotNilSuite(t *testing.T) {
	suite.Run(t, new(NotNilTestSuite))
}

func (s *NotNilTestSuite) SetupTest() {
}

func (s *NotNilTestSuite) TestNotNil() {
	nv := gkgen.NewNotNilValidator()
	e := gkexample.ExampleStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("HexStringPtr")
	code, err := nv.GenerateValidationCode(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if s.HexStringPtr == nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, errors.New(\"is Nil\"))\n}", code)
}

func (s *NotNilTestSuite) TestNotNil_Map() {
	nv := gkgen.NewNotNilValidator()
	e := gkexample.NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilMap")
	code, err := nv.GenerateValidationCode(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if s.NotNilMap == nil {\nerrorsNotNilMap = append(errorsNotNilMap, errors.New(\"is Nil\"))\n}", code)
}

func (s *NotNilTestSuite) TestNotNil_Slice() {
	nv := gkgen.NewNotNilValidator()
	e := gkexample.NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilSlice")
	code, err := nv.GenerateValidationCode(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if s.NotNilSlice == nil {\nerrorsNotNilSlice = append(errorsNotNilSlice, errors.New(\"is Nil\"))\n}", code)
}
