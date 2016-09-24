package gkgen

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

// GeneratorTestSuite
type GeneratorTestSuite struct {
	suite.Suite
}

// TestGeneratorTestSuite
func TestGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}

// SetupTest
func (s *GeneratorTestSuite) SetupTest() {}

// TestAddValidation
func (s *GeneratorTestSuite) TestAddValidation() {
	v := ValidateGenerator{
		Generators: make(map[string]Generater),
	}
	generator := NewNotNilValidator()

	_, ok := v.Generators[generator.Name()]
	s.False(ok)

	v.AddValidation(generator)
	_, ok = v.Generators[generator.Name()]
	s.True(ok)
}

// TestExampleStruct tests single no-param validation
func (s *GeneratorTestSuite) TestExampleStruct() {
	out := &bytes.Buffer{}
	key := "abc123"
	e := ExampleStruct{
		HexStringPtr: &key,
	}

	v := NewValidateGenerator()

	err := v.Generate(out, e)
	s.Nil(err)
}

// UnknownTagStruct
type UnknownTagStruct struct {
	Field string `valid:"Length=(5),Unknown"`
}

// TestGenerateWithUnknownTag
func (s *GeneratorTestSuite) TestGenerateWithUnknownTag() {
	out := &bytes.Buffer{}
	v := NewValidateGenerator()
	err := v.Generate(out, UnknownTagStruct{})
	s.Equal(errors.New("Unknown validation generator name: 'Unknown'"), err)
}

// TestGenerateMapValidationCode
func (s *GeneratorTestSuite) TestGenerateMapValidationCode() {
	out := &bytes.Buffer{}
	var fieldType reflect.Type
	err := generateMapValidationCode(out, fieldType, "", int64(1))
	s.Require().Error(err)
}
