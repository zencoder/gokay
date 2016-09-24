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
