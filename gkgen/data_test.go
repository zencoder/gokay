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

// NotZeroTestStruct is used in the NotZero unit tests
type NotZeroTestStruct struct {
	NotZeroInt        int        `valid:"NotZero"`
	NotZeroInt8       int8       `valid:"NotZero"`
	NotZeroInt16      int16      `valid:"NotZero"`
	NotZeroInt32      int32      `valid:"NotZero"`
	NotZeroInt64      int64      `valid:"NotZero"`
	NotZeroUint       uint       `valid:"NotZero"`
	NotZeroUint8      uint8      `valid:"NotZero"`
	NotZeroUint16     uint16     `valid:"NotZero"`
	NotZeroUint32     uint32     `valid:"NotZero"`
	NotZeroUint64     uint64     `valid:"NotZero"`
	NotZeroUintptr    uintptr    `valid:"NotZero"`
	NotZeroFloat32    float32    `valid:"NotZero"`
	NotZeroFloat64    float64    `valid:"NotZero"`
	NotZeroComplex64  complex64  `valid:"NotZero"`
	NotZeroComplex128 complex128 `valid:"NotZero"`
}

// NotZeroTestPointerStruct is used in the NotZero unit tests
type NotZeroTestPointerStruct struct {
	NotZeroInt        *int        `valid:"NotZero"`
	NotZeroInt8       *int8       `valid:"NotZero"`
	NotZeroInt16      *int16      `valid:"NotZero"`
	NotZeroInt32      *int32      `valid:"NotZero"`
	NotZeroInt64      *int64      `valid:"NotZero"`
	NotZeroUint       *uint       `valid:"NotZero"`
	NotZeroUint8      *uint8      `valid:"NotZero"`
	NotZeroUint16     *uint16     `valid:"NotZero"`
	NotZeroUint32     *uint32     `valid:"NotZero"`
	NotZeroUint64     *uint64     `valid:"NotZero"`
	NotZeroUintptr    *uintptr    `valid:"NotZero"`
	NotZeroFloat32    *float32    `valid:"NotZero"`
	NotZeroFloat64    *float64    `valid:"NotZero"`
	NotZeroComplex64  *complex64  `valid:"NotZero"`
	NotZeroComplex128 *complex128 `valid:"NotZero"`
}
