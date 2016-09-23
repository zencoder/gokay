package gkgen_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
)

// ValidatorTestSuite
type ValidatorTestSuite struct {
	suite.Suite
}

// TestValidatorSuite
func TestValidatorSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}

// SetupTest
func (s *ValidatorTestSuite) SetupTest() {}

// TestAddValidation
func (s *ValidatorTestSuite) TestAddValidation() {
	v := gkgen.ValidateGenerator{make(map[string]gkgen.Validater)}
	validator := gkgen.NewNotNilValidator()

	_, ok := v.Validaters[validator.Name()]
	s.False(ok)

	v.AddValidation(validator)
	_, ok = v.Validaters[validator.Name()]
	s.True(ok)
}

// TestExampleStruct tests single no-param validation
func (s *ValidatorTestSuite) TestExampleStruct() {
	out := &bytes.Buffer{}
	key := "abc123"
	e := gkexample.ExampleStruct{
		HexStringPtr: &key,
	}

	v := gkgen.NewValidator()

	err := v.Generate(out, e)
	s.Nil(err)
}

// UnknownTagStruct
type UnknownTagStruct struct {
	Field string `valid:"Length=(5),Unknown"`
}

// TestGenerateWithUnknownTag
func (s *ValidatorTestSuite) TestGenerateWithUnknownTag() {
	out := &bytes.Buffer{}
	v := gkgen.NewValidator()
	err := v.Generate(out, UnknownTagStruct{})
	s.Equal(errors.New("Unknown validation generator name: 'Unknown'"), err)
}
