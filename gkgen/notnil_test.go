package gkgen

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

// NotNilTestSuite
type NotNilTestSuite struct {
	suite.Suite
}

//TestNotNilSuite
func TestNotNilSuite(t *testing.T) {
	suite.Run(t, new(NotNilTestSuite))
}

// SetupTest
func (s *NotNilTestSuite) SetupTest() {}

// TestNotNil
func (s *NotNilTestSuite) TestNotNil() {
	nv := NewNotNilValidator()
	e := ExampleStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("HexStringPtr")
	code, err := nv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if s.HexStringPtr == nil {\nerrorsHexStringPtr = append(errorsHexStringPtr, errors.New(\"is Nil\"))\n}", code)
}

// TestNotNil_Map
func (s *NotNilTestSuite) TestNotNil_Map() {
	nv := NewNotNilValidator()
	e := NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilMap")
	code, err := nv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if s.NotNilMap == nil {\nerrorsNotNilMap = append(errorsNotNilMap, errors.New(\"is Nil\"))\n}", code)
}

// TestNotNil_Slice
func (s *NotNilTestSuite) TestNotNil_Slice() {
	nv := NewNotNilValidator()
	e := NotNilTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("NotNilSlice")
	code, err := nv.Generate(et, field, []string{})
	s.Nil(err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	s.Equal("if s.NotNilSlice == nil {\nerrorsNotNilSlice = append(errorsNotNilSlice, errors.New(\"is Nil\"))\n}", code)
}
