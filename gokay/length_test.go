package gokay

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// LengthTestSuite
type LengthTestSuite struct {
	suite.Suite
	expected int64
	actual   int64
}

// SetupSuite
func (s *LengthTestSuite) SetupSuite() {
	s.expected = int64(10)
	s.actual = int64(99)
}

// TestLengthTestSuite
func TestLengthTestSuite(t *testing.T) {
	suite.Run(t, new(LengthTestSuite))
}

// TestNilString
func (s *LengthTestSuite) TestNilString() {
	var str *string
	s.Require().Nil(LengthString(s.expected, str))
}

// TestNotLengthNonMatch
func (s *LengthTestSuite) TestNotLengthNonMatch() {
	var str string
	str = "012345678"
	s.Require().Error(LengthString(s.expected, &str))
}

// TestNotLengthMatch
func (s *LengthTestSuite) TestNotLengthMatch() {
	var str string
	str = "0123456789"
	s.Require().Nil(LengthString(s.expected, &str))
}

// TestLengthSlice
func (s *LengthTestSuite) TestLengthSlice() {
	s.Require().Nil(LengthSlice(s.expected, s.actual))
}
