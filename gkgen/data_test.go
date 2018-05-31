package gkgen

// ExampleStruct is used in a number of unit tests in this package
type ExampleStruct struct {
	HexStringPtr            *string `valid:"Length=(16),NotNil,Hex"`
	HexString               string  `valid:"Length=(12),Hex"`
	BCP47StringPtr          *string `valid:"NotNil,BCP47"`
	BCP47String             string  `valid:"BCP47"`
	CanBeNilWithConstraints *string `valid:"Length=(12)"`
	BCP47NonString          int
	BCP47NonStringPtr       *int
}

// NotNilTestStruct is used in the NotNil unit tests
type NotNilTestStruct struct {
	NotNilMap       map[string]interface{} `valid:"NotNil"`
	NotNilSlice     []string               `valid:"NotNil"`
	NotNilInterface interface{}
}

// NotEqualTestStruct is used in the NotEqual unit tests
type NotEqualTestStruct struct {
	NotEqualInt        int        `valid:"NotEqual=(0)"`
	NotEqualInt8       int8       `valid:"NotEqual=(0)"`
	NotEqualInt16      int16      `valid:"NotEqual=(0)"`
	NotEqualInt32      int32      `valid:"NotEqual=(0)"`
	NotEqualInt64      int64      `valid:"NotEqual=(0)"`
	NotEqualUint       uint       `valid:"NotEqual=(0)"`
	NotEqualUint8      uint8      `valid:"NotEqual=(0)"`
	NotEqualUint16     uint16     `valid:"NotEqual=(0)"`
	NotEqualUint32     uint32     `valid:"NotEqual=(0)"`
	NotEqualUint64     uint64     `valid:"NotEqual=(0)"`
	NotEqualUintptr    uintptr    `valid:"NotEqual=(0)"`
	NotEqualFloat32    float32    `valid:"NotEqual=(0)"`
	NotEqualFloat64    float64    `valid:"NotEqual=(0)"`
	NotEqualComplex64  complex64  `valid:"NotEqual=(0)"`
	NotEqualComplex128 complex128 `valid:"NotEqual=(0)"`
	NotEqualString     string     `valid:"NotEqual=()"`
}

// NotEqualTestPointerStruct is used in the NotEqual unit tests
type NotEqualTestPointerStruct struct {
	NotEqualInt        *int        `valid:"NotEqual=(0)"`
	NotEqualInt8       *int8       `valid:"NotEqual=(0)"`
	NotEqualInt16      *int16      `valid:"NotEqual=(0)"`
	NotEqualInt32      *int32      `valid:"NotEqual=(0)"`
	NotEqualInt64      *int64      `valid:"NotEqual=(0)"`
	NotEqualUint       *uint       `valid:"NotEqual=(0)"`
	NotEqualUint8      *uint8      `valid:"NotEqual=(0)"`
	NotEqualUint16     *uint16     `valid:"NotEqual=(0)"`
	NotEqualUint32     *uint32     `valid:"NotEqual=(0)"`
	NotEqualUint64     *uint64     `valid:"NotEqual=(0)"`
	NotEqualUintptr    *uintptr    `valid:"NotEqual=(0)"`
	NotEqualFloat32    *float32    `valid:"NotEqual=(0)"`
	NotEqualFloat64    *float64    `valid:"NotEqual=(0)"`
	NotEqualComplex64  *complex64  `valid:"NotEqual=(0)"`
	NotEqualComplex128 *complex128 `valid:"NotEqual=(0)"`
	NotEqualString     string      `valid:"NotEqual=(0)"`
}

// SetTestStruct is used in the Set unit tests
type SetTestStruct struct {
	SetString    string  `valid:"Set=(cat)(dog)(mouse)"`
	SetStringPtr *string `valid:"Set=(cat)(dog)(mouse)"`
}
