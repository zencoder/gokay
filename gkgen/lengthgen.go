package gkgen

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// LengthValidateGen generates code that will verify the exact length of a String or String Pointer field.
// Slice and Array support coming later.
// It will flag nil string pointers as valid, use in conjunction with NotNil validator if you don't want nil values
type LengthValidateGen struct {
	name string
}

// NewLengthValidator holds LengthValidator state
func NewLengthValidator() *LengthValidateGen {
	return &LengthValidateGen{name: "Length"}
}

// Generate generates validation code
func (s *LengthValidateGen) Generate(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) != 1 {
		return "", errors.New("Length validation requires exactly 1 parameter")
	}

	expectedLength, err := strconv.Atoi(params[0])
	if err != nil {
		return "", err
	}
	field := fieldStruct.Type

	isPtr := false
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
		isPtr = true
	}

	switch field.Kind() {
	case reflect.Ptr:
		return "", errors.New("LengthValidator does not support nested pointer fields")
	case reflect.String:
		if isPtr {
			return fmt.Sprintf(`
				if err := gokay.LengthString(%d, s.%[2]s); err != nil {
					errors%[2]s = append(errors%[2]s, err)
				}
				`, expectedLength, fieldStruct.Name), nil
		}
		return fmt.Sprintf(`
			if err := gokay.LengthString(%d, &s.%[2]s); err != nil {
				errors%[2]s = append(errors%[2]s, err)
			}
			`, expectedLength, fieldStruct.Name), nil
	case reflect.Slice, reflect.Array:
		return "", errors.New("Length validator does not yet support Slice or Arrays")
	default:
		if isPtr {
			return "", fmt.Errorf("LengthValidator does not support fields of type: '*%v'", field.Kind())
		}
		return "", fmt.Errorf("LengthValidator does not support fields of type: '%v'", field.Kind())
	}
}

// Name provides access to the name field
func (s *LengthValidateGen) Name() string {
	return s.name
}
