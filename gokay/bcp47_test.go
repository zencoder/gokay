package gokay

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// BCP47TextSuite
type BCP47TextSuite struct {
	suite.Suite
}

// TestBCP47TextSuite
func TestBCP47TextSuite(t *testing.T) {
	suite.Run(t, new(BCP47TextSuite))
}

// TestIsBCP47_Nil
func (s *BCP47TextSuite) TestIsBCP47_Nil() {
	s.Nil(IsBCP47(nil))
}

// TestIsBCP47_English
func (s *BCP47TextSuite) TestIsBCP47_English() {
	str := "English"
	s.NotNil(IsBCP47(&str))
}

// TestIsBCP47_en
func (s *BCP47TextSuite) TestIsBCP47_en() {
	str := "en"
	s.Nil(IsBCP47(&str))
}
