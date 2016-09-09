package gokay

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BCP47TextSuite struct {
	suite.Suite
}

func TestBCP47TextSuite(t *testing.T) {
	suite.Run(t, new(BCP47TextSuite))
}

func (s *BCP47TextSuite) TestIsBCP47_Nil() {
	s.Nil(IsBCP47(nil))
}

func (s *BCP47TextSuite) TestIsBCP47_English() {
	str := "English"
	s.NotNil(IsBCP47(&str))
}

func (s *BCP47TextSuite) TestIsBCP47_en() {
	str := "en"
	s.Nil(IsBCP47(&str))
}
