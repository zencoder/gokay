package gkgen_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gkgen"
	"github.com/zencoder/gokay/internal/gkexample"
)

// EmptyStruct
type EmptyStruct struct{}

// TagParserTestSuite
type TagParserTestSuite struct {
	suite.Suite
}

// SetupTest
func (suite *TagParserTestSuite) SetupTest() {}

// TestTagParserTestSuite
func TestTagParserTestSuite(t *testing.T) {
	suite.Run(t, new(TagParserTestSuite))
}

// TestParseTagSingleNoParamValidation tests single no-param validation
func (suite *TagParserTestSuite) TestParseTagSingleNoParamValidation() {
	s := new(EmptyStruct)
	tag := "bar"
	vcs, err := gkgen.ParseTag(s, tag)
	assert.Nil(suite.T(), err)

	expectedCommand := gkgen.NewValidationCommand("bar")
	assert.Equal(suite.T(), expectedCommand, vcs[0])
	assert.Equal(suite.T(), 1, len(vcs))
}

// Test single no-param validation
func (suite *TagParserTestSuite) TestExampleStruct() {
	key := "abc123"
	s := gkexample.ExampleStruct{
		HexStringPtr: &key,
	}

	_, err := gkgen.ParseTag(s, "valid")
	assert.Nil(suite.T(), err)
}

// Test multiple no-param validaitons
func (suite *TagParserTestSuite) TestParseTagMultipleNoParamValidations() {
	s := new(EmptyStruct)
	tag := "bar,biz,buz"
	vcs, err := gkgen.ParseTag(s, tag)

	assert.Nil(suite.T(), err)
	barCommand := gkgen.NewValidationCommand("bar")
	bizCommand := gkgen.NewValidationCommand("biz")
	buzCommand := gkgen.NewValidationCommand("buz")

	expectedVcs := []gkgen.ValidationCommand{barCommand, bizCommand, buzCommand}

	assert.Equal(suite.T(), expectedVcs, vcs)
}

// Test leading comma
func (suite *TagParserTestSuite) TestParseTagLeadingComma() {
	s := new(EmptyStruct)
	tag := ",bar"
	_, err := gkgen.ParseTag(s, tag)
	suite.NotNil(err)
}

// Test trailing commas
func (suite *TagParserTestSuite) TestParseTagTrailingCommas() {
	s := new(EmptyStruct)
	tag := "bar,"
	vcs, err := gkgen.ParseTag(s, tag)
	assert.Nil(suite.T(), err)
	expectedVcs := []gkgen.ValidationCommand{gkgen.NewValidationCommand("bar")}
	assert.Equal(suite.T(), expectedVcs, vcs)

	tag = "two_commas,,"
	_, err = gkgen.ParseTag(s, tag)
	suite.NotNil(err)
}

// Test validation with multiple parameters
func (suite *TagParserTestSuite) TestParseTagWithConstParam() {
	s := new(EmptyStruct)
	tag := "bar=(hello world,\\)How are you?)"
	vcs, err := gkgen.ParseTag(s, tag)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(vcs))
	assert.Equal(suite.T(), "bar", vcs[0].Name())
	assert.Equal(suite.T(), 1, len(vcs[0].Params))
	assert.Equal(suite.T(), "hello world,)How are you?", vcs[0].Params[0])
}

func (suite *TagParserTestSuite) TestParseTagWithConstParamSyntaxError() {
	s := new(EmptyStruct)
	tag := "bar=(?foo\\)[biz]"
	_, err := gkgen.ParseTag(s, tag)
	suite.NotNil(err)
}

func (suite *TagParserTestSuite) TestParseTagMissingParamSyntaxError() {
	s := new(EmptyStruct)
	tag := "bar=,foo"
	_, err := gkgen.ParseTag(s, tag)
	suite.NotNil(err)

	tag = "bar="
	_, err = gkgen.ParseTag(s, tag)
	assert.Equal(suite.T(), io.EOF, err)
}

func (suite *TagParserTestSuite) TestParseTagLeadingEquals() {
	s := new(EmptyStruct)
	tag := "="
	_, err := gkgen.ParseTag(s, tag)
	suite.NotNil(err)
}

func (suite *TagParserTestSuite) TestParseTagWithMultipleParams() {
	s := new(EmptyStruct)
	tag := "bar=(bar0)(bar1)"
	vcs, err := gkgen.ParseTag(s, tag)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(vcs))
	assert.Equal(suite.T(), "bar", vcs[0].Name())
	assert.Equal(suite.T(), 2, len(vcs[0].Params))
	assert.Equal(suite.T(), "bar0", vcs[0].Params[0])
	assert.Equal(suite.T(), "bar1", vcs[0].Params[1])
}

func (suite *TagParserTestSuite) TestParseTag2ValidationsWith1ParamEach() {
	s := new(EmptyStruct)
	tag := "bar=(bar0)(bar1),foo=(foo0)"
	vcs, err := gkgen.ParseTag(s, tag)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(vcs))

	assert.Equal(suite.T(), "bar", vcs[0].Name())
	assert.Equal(suite.T(), 2, len(vcs[0].Params))
	assert.Equal(suite.T(), "bar0", vcs[0].Params[0])
	assert.Equal(suite.T(), "bar1", vcs[0].Params[1])

	assert.Equal(suite.T(), "foo", vcs[1].Name())
	assert.Equal(suite.T(), 1, len(vcs[1].Params))
	assert.Equal(suite.T(), "foo0", vcs[1].Params[0])
}

// Intptr provides an int typed pointer
func Intptr(v int) *int {
	p := new(int)
	*p = v
	return p
}

// Float64ptr provides a float64 typed pointer
func Float64ptr(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

// Boolptr provides a boolean typed pointer
func Boolptr(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}
