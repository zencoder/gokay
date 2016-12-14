package gkgen

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

// EmptyStruct
type EmptyStruct struct{}

// TestParseTagSingleNoParamValidation tests single no-param validation
func TestParseTagSingleNoParamValidation(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar"
	vcs, err := ParseTag(s, tag)
	require.NoError(t, err)

	expectedCommand := NewValidationCommand("bar")
	require.Equal(t, expectedCommand, vcs[0])
	require.Equal(t, 1, len(vcs))
}

// Test single no-param validation
func TestExampleValidStruct(t *testing.T) {
	key := "abc123"
	s := ExampleStruct{
		HexStringPtr: &key,
	}

	_, err := ParseTag(s, "valid")
	require.Nil(t, err)
}

// Test multiple no-param validaitons
func TestParseTagMultipleNoParamValidations(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar,biz,buz"
	vcs, err := ParseTag(s, tag)

	require.Nil(t, err)
	barCommand := NewValidationCommand("bar")
	bizCommand := NewValidationCommand("biz")
	buzCommand := NewValidationCommand("buz")

	expectedVcs := []ValidationCommand{barCommand, bizCommand, buzCommand}

	require.Equal(t, expectedVcs, vcs)
}

// Test leading comma
func TestParseTagLeadingComma(t *testing.T) {
	s := new(EmptyStruct)
	tag := ",bar"
	_, err := ParseTag(s, tag)
	require.Error(t, err)
}

// Test trailing commas
func TestParseTagTrailingCommas(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar,"
	vcs, err := ParseTag(s, tag)
	require.Nil(t, err)
	expectedVcs := []ValidationCommand{NewValidationCommand("bar")}
	require.Equal(t, expectedVcs, vcs)

	tag = "two_commas,,"
	_, err = ParseTag(s, tag)
	require.Error(t, err)
}

// Test validation with multiple parameters
func TestParseTagWithConstParam(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar=(hello world,\\)How are you?)"
	vcs, err := ParseTag(s, tag)
	require.NoError(t, err)
	require.Equal(t, 1, len(vcs))
	require.Equal(t, "bar", vcs[0].Name())
	require.Equal(t, 1, len(vcs[0].Params))
	require.Equal(t, "hello world,)How are you?", vcs[0].Params[0])
}

func TestParseTagWithConstParamSyntaxError(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar=(?foo\\)[biz]"
	_, err := ParseTag(s, tag)
	require.Error(t, err)
}

func TestParseTagMissingParamSyntaxError(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar=,foo"
	_, err := ParseTag(s, tag)
	require.Error(t, err)

	tag = "bar="
	_, err = ParseTag(s, tag)
	require.Equal(t, io.EOF, err)
}

func TestParseTagLeadingEquals(t *testing.T) {
	s := new(EmptyStruct)
	tag := "="
	_, err := ParseTag(s, tag)
	require.Error(t, err)
}

func TestParseTagWithMultipleParams(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar=(bar0)(bar1)"
	vcs, err := ParseTag(s, tag)
	require.Nil(t, err)
	require.Equal(t, 1, len(vcs))
	require.Equal(t, "bar", vcs[0].Name())
	require.Equal(t, 2, len(vcs[0].Params))
	require.Equal(t, "bar0", vcs[0].Params[0])
	require.Equal(t, "bar1", vcs[0].Params[1])
}

func TestParseTag2ValidationsWith1ParamEach(t *testing.T) {
	s := new(EmptyStruct)
	tag := "bar=(bar0)(bar1),foo=(foo0)"
	vcs, err := ParseTag(s, tag)
	require.Nil(t, err)
	require.Equal(t, 2, len(vcs))

	require.Equal(t, "bar", vcs[0].Name())
	require.Equal(t, 2, len(vcs[0].Params))
	require.Equal(t, "bar0", vcs[0].Params[0])
	require.Equal(t, "bar1", vcs[0].Params[1])

	require.Equal(t, "foo", vcs[1].Name())
	require.Equal(t, 1, len(vcs[1].Params))
	require.Equal(t, "foo0", vcs[1].Params[0])
}
