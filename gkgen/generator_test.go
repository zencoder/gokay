package gkgen_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
)

type ValidatorTestSuite struct {
	suite.Suite
}

func TestValidatorSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}

func (s *ValidatorTestSuite) SetupTest() {
}

func (s *ValidatorTestSuite) TestAddValidation() {
	v := gkgen.ValidateGenerator{make(map[string]gkgen.Validater)}
	validator := gkgen.NewNotNilValidator()

	_, ok := v.Validaters[validator.GetName()]
	s.False(ok)

	v.AddValidation(validator)
	_, ok = v.Validaters[validator.GetName()]
	s.True(ok)
}

// Test single no-param validation
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

type UnknownTagStruct struct {
	Field string `valid:"Length=(5),Unknown"`
}

func (s *ValidatorTestSuite) TestGenerateWithUnknownTag() {
	out := &bytes.Buffer{}
	v := gkgen.NewValidator()
	err := v.Generate(out, UnknownTagStruct{})
	s.Equal(errors.New("Unknown validation generator name: 'Unknown'"), err)
}
