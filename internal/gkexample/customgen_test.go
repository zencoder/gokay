package gkexample

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// CustomGeneratorTestSuite: Run this test suite AFTER running gokay against example.go
type CustomGeneratorTestSuite struct {
	suite.Suite
}

// TestCustomGeneratorTestSuite
func TestCustomGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(CustomGeneratorTestSuite))
}

// TestNewCustomGenerator
func (s *CustomGeneratorTestSuite) TestNewCustomGenerator() {
	s.Require().NotNil(NewCustomGenerator())
}
