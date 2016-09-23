# gokay [![CircleCI](https://circleci.com/gh/zencoder/gokay.svg?style=svg&circle-token=90f42bc5cbb6fe74834f7649d67298130431d88d)](https://circleci.com/gh/zencoder/gokay) [![Coverage Status](https://coveralls.io/repos/github/zencoder/gokay/badge.svg?branch=circle-fixes&t=A2kWWv)](https://coveralls.io/github/zencoder/gokay?branch=circle-fixes)
Codegenerated struct validation tool for go.

## How it works
gokay parses a struct and generates a `Validate` function so that the struct implements the `Validateable` interface. It does so by parsing the `valid` tags in that struct's fields.

gokay generated Validate functions will return an ErrorMap that implements the Error interface. The ErrorMap contains a slice of failed validations for each invalid field in a struct.

## Installing gokay

This project uses [Glide](https://github.com/Masterminds/glide) to manage it's dependencies. Please refer to the glide docs to see how to install and use glide.

This project is tested on go 1.7.1 and glide 0.12.1

```bash
mkdir -p $GOPATH/github.com/zencoder
cd $GOPATH/src/github.com/zencoder
git clone https://github.com/zencoder/gokay
cd gokay
glide install
go install ./...
```

## Running gokay
### Usage
```	sh
gokay {file_name} ({custom generator package} {custom generator contructor})
```

### Example

```sh
gokay file.go gkcustom NewCustomGKGenerator
```

It relies on goimports tool to resolve import path of custom generator

## Using gokay
- Add validations to `valid` tag in struct def:

```go
type ExampleStruct struct {
	HexStringPtr            *string `valid:"Length=(16),NotNil,Hex"`
	HexString               string  `valid:"Length=(12),Hex"`
	CanBeNilWithConstraints *string `valid:"Length=(12)"`
}
```

- Run gokay command

### Tag syntax
Validations tags are comma separated, with Validation parameters delineated by open and closed parentheses.

`valid:"ValidationName1,ValidationName2=(vn2paramA)(vn2paramB)"`

In the above example, the `Hex` and `NotNil` Validations are parameterless, whereas length requires 1 parameter.

### Built-in Validations*
Name | Params | Allowed Field Types | Description
---- | ------------------- | ------ | -----------
Hex  | N/A | `(*)string` | Checks if a string field is a valid hexadecimal format number (0x prefix optional)
NotNil | N/A | pointers | Checks and fails if a pointer is nil
Length | 1 | `(*)string` | Checks if a string's length matches the tag's parameter
UUID | N/A | `(*)string` | Checks and fails if a string is not a valid UUID

### Implicitly generated validations*
These sections of code will be added to the generated `Validate()` function regardless of a field's `valid` tag's contents.
If a struct does not have any `valid` tags and no fields with implicit validation rules, then no Validate method will be generated.

- Struct fields: generated code will call static Validate function on any field that implements Validateable interface
- Slice/Map fields: Static Validate will be called on each element of a slice or map of structs or struct pointers (one level of indirection). Only supports maps with string indices.


*Note on built-in and implicit validations: With the obvious exception of NotNil, nil pointers fields are considered to be valid in order to allow support for optional fields.

### Writing your own Validations
gokay was built to allow developers to write attach their own Validations to the Validate generator.

1. Write function that validates a field. E.g:
   
	```go
// LengthString checks if the value of a string pointer has a length of exactly 'expected'
	func LengthString(expected int64, str *string) error {
		if str == nil {
			return nil // Covers the case where a value can be Nil OR has a length constraint
		}
	
		if expected != int64(len(*str)) {
			return fmt.Errorf("Length was '%d', needs to be '%d'", len(*str), expected)
		}
		return nil
	}
```

1. Write a struct that implements the `Validater` interface

    ```go
    type Validater interface {
		GenerateValidationCode(reflect.Type, reflect.StructField, []string) (string, error)
		GetName() string
	}
	```
   - GetName returns the string that will be used as a validation tag

1. GenerateValidationCode should generate a block will leverage the function defined in step 1. This block will be inserted into the generated `Validate` function. GenerateValidationCode output example:
    
    ```go
    // ValidationName
	if err := somepackage.ValidationFunction(someparamA, someparamB, s.Field); err != nil {
		errorsField = append(errorsField, err)
	}
	```

1. Write a function that constructs a `ValidateGenerator`. This function should live in a different package from your data model. Example:

	```go
	package gkcustom
	
	// To run: `gokay gkexample NewCustomValidator`
	func NewCustomGKGenerator() *gkgen.ValidateGenerator {
		v := gkgen.NewValidator()
		v.AddValidation(NewCustomValidator())
		return v
	}
	```
1. Write tests for your struct's constraints
1. Add `valid` tags to your struct fields
1. Run gokay: `gokay file.go gkcustom NewCustomGKGenerator`

[More Examples](internal/gkexample/)

## Development

### Dependencies

Tested on go 1.7.1.

### Build and run unit tests

    make test
    
### CI

[This library builds on Circle CI, here.](https://circleci.com/gh/zencoder/gokay/)

## License

[Apache License Version 2.0](LICENSE)

