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

func (s *HexTestSuite) TestIsHex_No0x() {
	str := "1a3F"
	s.Nil(IsHex(&str))
}

func (s *HexTestSuite) TestIsHex_0x() {
	str := "0x1a3F"
	s.Nil(IsHex(&str))
}

func (s *HexTestSuite) TestIsHex_NotHex() {
	str := "0x1Gbcq"
	s.Equal(errors.New("'0x1Gbcq' is not a hexadecimal string"), IsHex(&str))
}
