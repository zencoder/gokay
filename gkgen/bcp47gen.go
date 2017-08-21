package gkgen

import (
	"errors"
	"fmt"
	"reflect"
)

// BCP47Validator generates code that will verify if a field is a BCP47 compatible string
// https://tools.ietf.org/html/bcp47
type BCP47Validator struct {
	name string
}

// NewBCP47Validator holds the BCP47Validator state
func NewBCP47Validator() *BCP47Validator {
	return &BCP47Validator{name: "BCP47"}
}

// Generate generates validation code
func (s *BCP47Validator) Generate(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) != 0 {
		return "", errors.New("BCP47 takes no parameters")
	}

	field := fieldStruct.Type

	isPtr := false
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
		isPtr = true
	}

	switch field.Kind() {
	case reflect.Ptr:
		return "", errors.New("BCP47Validator does not support nested pointer fields")
	case reflect.String:
		if isPtr {
			return fmt.Sprintf(`
				if err := gokay.IsBCP47(s.%[1]s); err != nil {
					errors%[1]s = append(errors%[1]s, err)
				}
				`, fieldStruct.Name), nil
		}
		return fmt.Sprintf(`
			if err := gokay.IsBCP47(&s.%[1]s); err != nil {
				errors%[1]s = append(errors%[1]s, err)
			}
			`, fieldStruct.Name), nil
	default:
		if isPtr {
			return "", fmt.Errorf("BCP47Validator does not support fields of type: '*%v'", field.Kind())
		}
		return "", fmt.Errorf("BCP47Validator does not support fields of type: '%v'", field.Kind())
	}
}

// Name provides access to the name field
func (s *BCP47Validator) Name() string {
	return s.name
}
