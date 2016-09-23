package gokay

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HexTestSuite struct {
	suite.Suite
}

func TestHexTestSuite(t *testing.T) {
	suite.Run(t, new(HexTestSuite))
}

// TestNilString
func (s *HexTestSuite) TestNilString() {
	var str *string
	s.Require().Nil(IsHex(str))
}

// TestIsHex_No0x
func (s *HexTestSuite) TestIsHex_No0x() {
	str := "1a3F"
	s.Nil(IsHex(&str))
}

// TestIsHex_0x
func (s *HexTestSuite) TestIsHex_0x() {
	str := "0x1a3F"
	s.Nil(IsHex(&str))
}

// TestIsHex_NotHex
func (s *HexTestSuite) TestIsHex_NotHex() {
	str := "0x1Gbcq"
	s.Equal(errors.New("'0x1Gbcq' is not a hexadecimal string"), IsHex(&str))
}
