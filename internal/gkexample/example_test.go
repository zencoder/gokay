package gkexample

import (
	"errors"
	"testing"

	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/suite"
	"github.com/zencoder/gokay/gokay"
)

// ExampleTestSuite: Run this test suite AFTER running gokay against example.go
type ExampleTestSuite struct {
	suite.Suite
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (s *ExampleTestSuite) TestNotAValidateable() {
	underTest := NoImplicitValidate{}

	tut := reflect.TypeOf(underTest)
	validateable := reflect.TypeOf((*gokay.Validateable)(nil))

	s.False(tut.Implements(validateable.Elem()))
}

func (s *ExampleTestSuite) TestHasValidateImplicit() {
	underTest := HasValidateImplicit{}

	tut := reflect.TypeOf(underTest)
	validateable := reflect.TypeOf((*gokay.Validateable)(nil))

	s.True(tut.Implements(validateable.Elem()))
}

func (s *ExampleTestSuite) TestHasValidateImplicit_Validate() {
	underTest := HasValidateImplicit{
		InvalidStruct: &TestValidate{},
	}

	err := underTest.Validate()

	em, ok := err.(gokay.ErrorMap)
	s.True(ok)
	s.Equal(1, len(em))
	_, ok = em["ValidStruct"]
	s.False(ok)

	ea, ok := em["InvalidStruct"].(gokay.ErrorArray)
	s.True(ok)
	s.Equal(gokay.ErrorArray{
		gokay.ErrorMap{
			"Field": gokay.ErrorArray{errors.New("invalid when false")},
		},
	}, ea)
}

func (s *ExampleTestSuite) TestHasValidateImplicit_NilInvalidStruct() {
	underTest := HasValidateImplicit{
		InvalidStruct: nil,
	}

	err := underTest.Validate()
	s.Nil(err)
}

func (s *ExampleTestSuite) TestValidateImplicit_MapOfStruct() {
	underTest := HasValidateImplicit{
		InvalidStruct: nil,
		MapOfStruct: map[string]TestValidate{
			"a": {},
			"b": {},
		},
	}

	err := underTest.Validate()
	expected := gokay.ErrorMap{
		"MapOfStruct": gokay.ErrorArray{
			gokay.ErrorMap{
				"a": gokay.ErrorMap{
					"Field": gokay.ErrorArray{errors.New("invalid when false")},
				},
				"b": gokay.ErrorMap{
					"Field": gokay.ErrorArray{errors.New("invalid when false")},
				},
			},
		},
	}
	s.Equal(expected, err)
}

func (s *ExampleTestSuite) TestValidateImplicit_MapOfStructPtrs() {
	underTest := HasValidateImplicit{
		InvalidStruct: nil,
		MapOfStructPtrs: map[string]*TestValidate{
			"a": {},
			"b": (*TestValidate)(nil),
			"c": {},
			"d": {true},
		},
	}

	err := underTest.Validate()
	expected := gokay.ErrorMap{
		"MapOfStructPtrs": gokay.ErrorArray{
			gokay.ErrorMap{
				"a": gokay.ErrorMap{
					"Field": gokay.ErrorArray{errors.New("invalid when false")},
				},
				"c": gokay.ErrorMap{
					"Field": gokay.ErrorArray{errors.New("invalid when false")},
				},
			},
		},
	}
	s.Equal(expected, err)
}

func (s *ExampleTestSuite) TestValidateImplicit_MapOfMapsOfStructs() {
	underTest := HasValidateImplicit{
		InvalidStruct: nil,
		MapOfMaps: map[string]map[string]*TestValidate{
			"invalid": {
				"invalidA": {},
				"invalidB": (*TestValidate)(nil),
				"invalidC": {},
				"invalidD": {true},
			},
			"valid": {
				"validA": {true},
				"validB": (*TestValidate)(nil),
				"validC": {true},
				"validD": {true},
			},
		},
	}

	err := underTest.Validate()
	expected := gokay.ErrorMap{
		"MapOfMaps": gokay.ErrorArray{
			gokay.ErrorMap{
				"invalid": gokay.ErrorMap{
					"invalidA": gokay.ErrorMap{
						"Field": gokay.ErrorArray{errors.New("invalid when false")},
					},
					"invalidC": gokay.ErrorMap{
						"Field": gokay.ErrorArray{errors.New("invalid when false")},
					},
				},
			},
		},
	}
	s.Equal(expected, err)
}

func (s *ExampleTestSuite) TestValidateImplicit_MapOfMapsOfSlices() {
	vErr := gokay.ErrorMap{
		"Field": gokay.ErrorArray{errors.New("invalid when false")},
	}

	underTest := HasValidateImplicit{
		InvalidStruct: nil,
		MapMapsOfSlices: map[string]map[string][]*TestValidate{
			"a": {
				"aa": {
					{},
					(*TestValidate)(nil),
					{},
					{true},
				},
				"ab": {
					{true},
					{},
					(*TestValidate)(nil),
					{true},
				},
			},
			"b": {
				"ba": {
					{},
					{},
					{},
					{},
				},
				"bb": {
					{true},
					{true},
					{true},
					{true},
				},
			},
			"c": {
				"ca": {},
			},
			"d": {},
		},
	}

	err := underTest.Validate()

	expected := gokay.ErrorMap{
		"MapMapsOfSlices": gokay.ErrorArray{
			gokay.ErrorMap{
				"a": gokay.ErrorMap{
					"aa": gokay.ErrorMap{
						"0": vErr,
						"2": vErr,
					},
					"ab": gokay.ErrorMap{
						"1": vErr,
					},
				},
				"b": gokay.ErrorMap{
					"ba": gokay.ErrorMap{
						"0": vErr,
						"1": vErr,
						"2": vErr,
						"3": vErr,
					},
				},
			},
		},
	}

	s.Equal(expected, err)
}

func (s *ExampleTestSuite) TestValidateImplicit_MapOfSlicesOfMaps() {
	vErr := gokay.ErrorMap{
		"Field": gokay.ErrorArray{errors.New("invalid when false")},
	}

	underTest := HasValidateImplicit{
		InvalidStruct: nil,
		MapOfSlicesOfMaps: map[string][]map[string]*TestValidate{
			"a": {
				{
					"a0a": {},
					"a0b": (*TestValidate)(nil),
				},
				{
					"a1a": {true},
					"a1b": (*TestValidate)(nil),
					"a1c": {},
				},
				{},
			},
			"b": {
				{
					"b0a": {true},
					"b0b": (*TestValidate)(nil),
					"b0c": (*TestValidate)(nil),
				},
				{},
				{
					"b2a": {},
					"b2b": {},
					"b2c": {},
				},
			},
			"c": {},
		},
	}

	err := underTest.Validate()

	expected := gokay.ErrorMap{
		"MapOfSlicesOfMaps": gokay.ErrorArray{
			gokay.ErrorMap{
				"a": gokay.ErrorMap{
					"0": gokay.ErrorMap{
						"a0a": vErr,
					},
					"1": gokay.ErrorMap{
						"a1c": vErr,
					},
				},
				"b": gokay.ErrorMap{
					"2": gokay.ErrorMap{
						"b2a": vErr,
						"b2b": vErr,
						"b2c": vErr,
					},
				},
			},
		},
	}

	s.Equal(expected, err)
}

func (s *ExampleTestSuite) TestValidateImplicit_MapOfInterfaces() {
	vErr := gokay.ErrorMap{
		"Field": gokay.ErrorArray{errors.New("invalid when false")},
	}

	underTest := HasValidateImplicit{
		InvalidStruct: nil,
		MapOfInterfaces: map[string]interface{}{
			"a": []map[string]*TestValidate{
				{
					"a0a": {},
					"a0b": (*TestValidate)(nil),
				},
				{
					"a1a": {true},
					"a1b": (*TestValidate)(nil),
					"a1c": {},
				},
				{},
			},
			"b": []map[string]*TestValidate{
				{
					"b0a": {true},
					"b0b": (*TestValidate)(nil),
					"b0c": (*TestValidate)(nil),
				},
				{},
				{
					"b2a": {},
					"b2b": {},
					"b2c": {},
				},
			},
			"c": &TestValidate{Valid: false},
			"d": &TestValidate{Valid: true},
		},
	}

	err := underTest.Validate()
	spew.Dump(err)

	expected := gokay.ErrorMap{
		"MapOfInterfaces": gokay.ErrorArray{
			gokay.ErrorMap{
				"c": vErr,
			},
		},
	}

	s.Equal(expected, err)
}

func (s *ExampleTestSuite) TestValidateNotNil_Slice() {
	expected := gokay.ErrorMap{
		"NotNilSlice": gokay.ErrorArray{errors.New("is Nil")},
	}

	underTest := NotNilTestStruct{
		NotNilMap: map[string]interface{}{},
	}

	err := underTest.Validate()
	s.Equal(expected, err)
}

func (s *ExampleTestSuite) TestValidateNotNil_Map() {
	expected := gokay.ErrorMap{
		"NotNilMap": gokay.ErrorArray{errors.New("is Nil")},
	}

	underTest := NotNilTestStruct{
		NotNilSlice: []string{},
	}

	err := underTest.Validate()
	s.Equal(expected, err)
}
