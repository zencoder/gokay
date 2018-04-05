package gkexample

import (
	"errors"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"github.com/zencoder/gokay/gokay"
)

func TestNotAValidateable(t *testing.T) {
	underTest := NoImplicitValidate{}

	tut := reflect.TypeOf(underTest)
	validateable := reflect.TypeOf((*gokay.Validateable)(nil))

	require.False(t, tut.Implements(validateable.Elem()))
}

func TestHasValidateImplicit(t *testing.T) {
	underTest := HasValidateImplicit{}

	tut := reflect.TypeOf(underTest)
	validateable := reflect.TypeOf((*gokay.Validateable)(nil))

	require.True(t, tut.Implements(validateable.Elem()))
}

func TestHasValidateImplicit_Validate(t *testing.T) {
	underTest := HasValidateImplicit{
		InvalidStruct: &TestValidate{},
	}

	err := underTest.Validate()

	em, ok := err.(gokay.ErrorMap)
	require.True(t, ok)
	require.Equal(t, 1, len(em))
	_, ok = em["ValidStruct"]
	require.False(t, ok)

	ea, ok := em["InvalidStruct"].(gokay.ErrorSlice)
	require.True(t, ok)
	require.Equal(t, gokay.ErrorSlice{
		gokay.ErrorMap{
			"Field": gokay.ErrorSlice{errors.New("invalid when false")},
		},
	}, ea)
}

func TestHasValidateImplicit_NilInvalidStruct(t *testing.T) {
	underTest := HasValidateImplicit{
		InvalidStruct: nil,
	}

	err := underTest.Validate()
	require.NoError(t, err)
}

func TestValidateImplicit_MapOfStruct(t *testing.T) {
	underTest := HasValidateImplicit{
		InvalidStruct: nil,
		MapOfStruct: map[string]TestValidate{
			"a": {},
			"b": {},
		},
	}

	err := underTest.Validate()
	expected := gokay.ErrorMap{
		"MapOfStruct": gokay.ErrorSlice{
			gokay.ErrorMap{
				"a": gokay.ErrorMap{
					"Field": gokay.ErrorSlice{errors.New("invalid when false")},
				},
				"b": gokay.ErrorMap{
					"Field": gokay.ErrorSlice{errors.New("invalid when false")},
				},
			},
		},
	}
	require.Equal(t, expected, err)
}

func TestValidateImplicit_MapOfStructPtrs(t *testing.T) {
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
		"MapOfStructPtrs": gokay.ErrorSlice{
			gokay.ErrorMap{
				"a": gokay.ErrorMap{
					"Field": gokay.ErrorSlice{errors.New("invalid when false")},
				},
				"c": gokay.ErrorMap{
					"Field": gokay.ErrorSlice{errors.New("invalid when false")},
				},
			},
		},
	}
	require.Equal(t, expected, err)
}

func TestValidateImplicit_MapOfMapsOfStructs(t *testing.T) {
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
		"MapOfMaps": gokay.ErrorSlice{
			gokay.ErrorMap{
				"invalid": gokay.ErrorMap{
					"invalidA": gokay.ErrorMap{
						"Field": gokay.ErrorSlice{errors.New("invalid when false")},
					},
					"invalidC": gokay.ErrorMap{
						"Field": gokay.ErrorSlice{errors.New("invalid when false")},
					},
				},
			},
		},
	}
	require.Equal(t, expected, err)
}

func TestValidateImplicit_MapOfMapsOfSlices(t *testing.T) {
	vErr := gokay.ErrorMap{
		"Field": gokay.ErrorSlice{errors.New("invalid when false")},
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
		"MapMapsOfSlices": gokay.ErrorSlice{
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

	require.Equal(t, expected, err)
}

func TestValidateImplicit_MapOfSlicesOfMaps(t *testing.T) {
	vErr := gokay.ErrorMap{
		"Field": gokay.ErrorSlice{errors.New("invalid when false")},
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
		"MapOfSlicesOfMaps": gokay.ErrorSlice{
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

	require.Equal(t, expected, err)
}

func TestValidateImplicit_MapOfInterfaces(t *testing.T) {
	vErr := gokay.ErrorMap{
		"Field": gokay.ErrorSlice{errors.New("invalid when false")},
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
		"MapOfInterfaces": gokay.ErrorSlice{
			gokay.ErrorMap{
				"c": vErr,
			},
		},
	}

	require.Equal(t, expected, err)
}

func TestValidateNotNil_Slice(t *testing.T) {
	expected := gokay.ErrorMap{
		"NotNilSlice": gokay.ErrorSlice{errors.New("is Nil")},
	}

	underTest := NotNilTestStruct{
		NotNilMap: map[string]interface{}{},
	}

	err := underTest.Validate()
	require.Equal(t, expected, err)
}

func TestValidateNotNil_Map(t *testing.T) {
	expected := gokay.ErrorMap{
		"NotNilMap": gokay.ErrorSlice{errors.New("is Nil")},
	}

	underTest := NotNilTestStruct{
		NotNilSlice: []string{},
	}

	err := underTest.Validate()
	require.Equal(t, expected, err)
}

func TestValidateNotZero_Valid(t *testing.T) {
	one := int64(1)
	underTest := NotZeroTestStruct{
		NotZeroInt64:    1,
		NotZeroInt64Ptr: &one,
	}

	err := underTest.Validate()
	require.Nil(t, err)
}
func TestValidateNotZero_NilValid(t *testing.T) {
	underTest := NotZeroTestStruct{
		NotZeroInt64: 1,
	}

	err := underTest.Validate()
	require.Nil(t, err)
}

func TestValidateNotZero_Inalid(t *testing.T) {
	expected := gokay.ErrorMap{
		"NotZeroInt64":    gokay.ErrorSlice{errors.New("is Zero")},
		"NotZeroInt64Ptr": gokay.ErrorSlice{errors.New("is Zero")},
	}

	zero := int64(0)
	underTest := NotZeroTestStruct{
		NotZeroInt64:    0,
		NotZeroInt64Ptr: &zero,
	}

	err := underTest.Validate()
	require.Equal(t, expected, err)
}
