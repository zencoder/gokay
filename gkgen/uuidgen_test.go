package gkgen

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateUUIDValidationCode_String(t *testing.T) {
	v := NewUUIDValidator()
	e := UUIDTestStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("UUIDString")

	code, err := v.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(t, "if err := gokay.IsUUID(&s.UUIDString); err != nil {\nerrorsUUIDString = append(errorsUUIDString, err)\n}", code)
}

func TestGenerateUUIDValidationCode_StringPtr(t *testing.T) {
	v := NewUUIDValidator()
	e := UUIDTestStruct{}
	et := reflect.TypeOf(e)
	field, _ := et.FieldByName("UUIDStringPtr")
	code, err := v.Generate(et, field, []string{})
	require.NoError(t, err)
	code = strings.Replace(strings.TrimSpace(code), "\t", "", -1)
	require.Equal(t, "if err := gokay.IsUUID(s.UUIDStringPtr); err != nil {\nerrorsUUIDStringPtr = append(errorsUUIDStringPtr, err)\n}", code)
}

func TestGenerateUUIDValidationCode_NonString(t *testing.T) {
	v := NewUUIDValidator()
	e := UUIDTestStruct{}
	et := reflect.TypeOf(e)

	field, _ := et.FieldByName("UUIDNonString")
	_, err := v.Generate(et, field, []string{})
	require.Equal(t, errors.New("UUIDValidator does not support fields of type: 'int'"), err)
}

type UUIDTestStruct struct {
	UUIDString    string  `valid:"UUID"`
	UUIDStringPtr *string `valid:"UUID"`
	UUIDNonString int     `valid:"UUID"`
}
