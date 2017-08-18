package gkgen

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// MinLengthValidateGen generates code that will verify the MinLength of a String, Slice, or Map (or their respective pointers)
// It will flag nil pointers as valid, use in conjunction with NotNil validator if you don't want nil values
type MinLengthValidateGen struct {
	name string
}

// NewMinLengthValidator holds MinLengthValidator state
func NewMinLengthValidator() *MinLengthValidateGen {
	return &MinLengthValidateGen{name: "MinLength"}
}

// Generate generates validation code
func (s *MinLengthValidateGen) Generate(fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) != 1 {
		return "", errors.New("MinLength validation requires exactly 1 parameter")
	}

	expectedMinLength, err := strconv.Atoi(params[0])
	if err != nil {
		return "", err
	}
	field := fieldStruct.Type

	if field.Kind() == reflect.Ptr {
		field = field.Elem()
		switch field.Kind() {
		case reflect.String:
			return fmt.Sprintf(`
				if err := gokay.MinLengthString(%d, s.%[2]s); err != nil {
					errors%[2]s = append(errors%[2]s, err)
				}
				`, expectedMinLength, fieldStruct.Name), nil
		default:
			return "", fmt.Errorf("MinLengthValidator does not support fields of type: '*%v'", field.Kind())
		}
	}

	switch field.Kind() {
	case reflect.String:
		return fmt.Sprintf(`
			if err := gokay.MinLengthString(%d, &s.%[2]s); err != nil {
				errors%[2]s = append(errors%[2]s, err)
			}
			`, expectedMinLength, fieldStruct.Name), nil
	default:
		return "", fmt.Errorf("MinLengthValidator does not support fields of type: '%v'", field.Kind())
	}
}

// Name provides access to the name field
func (s *MinLengthValidateGen) Name() string {
	return s.name
}
