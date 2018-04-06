package gkexample

import (
	"errors"

	"github.com/zencoder/gokay/gokay"
)

type NoImplicitValidate struct {
	StringField string

	NonStringKeyMap             map[int]TestValidate
	NestedNonStringKeyMap       map[string]map[int]map[string]*TestValidate
	NestedMapWithNonStructValue map[string]map[string]map[string]*int64

	SliceOfBuiltins       []string
	NestedSliceofBuiltins [][][]*int
}

type HasValidateImplicit struct {
	InvalidStruct *TestValidate
	ValidStruct   AlwaysValid

	MapOfStruct     map[string]TestValidate
	MapOfStructPtrs map[string]*TestValidate
	MapOfMaps       map[string]map[string]*TestValidate
	MapMapsOfSlices map[string]map[string][]*TestValidate
	MapOfInterfaces map[string]interface{}

	SimpleSlice           []*TestValidate
	SliceOfSlicesOfSlices [][][]*TestValidate

	MapOfSlicesOfMaps map[string][]map[string]*TestValidate
}

type NotNilTestStruct struct {
	NotNilMap       map[string]interface{} `valid:"NotNil"`
	NotNilSlice     []string               `valid:"NotNil"`
	NotNilInterface interface{}
}

// Example struct definition with tags
type ExampleStruct struct {
	HexStringPtr            *string `valid:"Length=(16),NotNil,Hex"`
	HexString               string  `valid:"Length=(12),Hex"`
	BCP47StringPtr          *string `valid:"NotNil,BCP47"`
	BCP47String             string  `valid:"BCP47"`
	CanBeNilWithConstraints *string `valid:"Length=(12)"`
	BCP47NonString          int
	BCP47NonStringPtr       *int
}

type NotEqualTestStruct struct {
	NotEqualString    string  `valid:"NotEqual=()"`
	NotEqualStringPtr *string `valid:"NotEqual=()"`
	NotEqualInt64     int64   `valid:"NotEqual=(0)"`
	NotEqualInt64Ptr  *int64  `valid:"NotEqual=(0)"`
}

type TestValidate struct {
	Valid bool
}

func (s TestValidate) Validate() error {
	if !s.Valid {
		return gokay.ErrorMap{
			"Field": gokay.ErrorSlice{errors.New("invalid when false")},
		}
	}
	return nil
}

type AlwaysValid struct{}

func (s AlwaysValid) Validate() error {
	return nil
}

type Example struct {
	MapOfInterfaces map[string]interface{} `valid:"NotNil"`
}
