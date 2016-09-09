package gokay

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UUIDTestSuite struct {
	suite.Suite
}

func TestUUIDTestSuite(t *testing.T) {
	suite.Run(t, new(UUIDTestSuite))
}

func (s *UUIDTestSuite) TestIsUUID_No0x() {
	str := "603c9a2a-38dB-4987-932a-2f57733a29f1"
	s.Nil(IsUUID(&str))
}

func (s *UUIDTestSuite) TestIsUUID_NotUUID() {
	str := "603c9a2a-38db-4987-932a-2f57733a29fQ"
	s.Equal(errors.New("'603c9a2a-38db-4987-932a-2f57733a29fQ' is not a UUID"), IsUUID(&str))
}

func (s *UUIDTestSuite) TestIsUUID_NotUUIDTooLong() {
	str := "AB603c9a2a-38db-4987-932a-2f57733a29fQ"
	s.Equal(errors.New("'AB603c9a2a-38db-4987-932a-2f57733a29fQ' is not a UUID"), IsUUID(&str))
}
