package gkgen

import (
	"errors"
	"fmt"
	"reflect"
)

// NotZeroValidator generates code that will verify if a number or dereferenced number pointer is zero
type NotZeroValidator struct {
	name string
}

// NewNotZeroValidator holds the NotZeroValidator state
func NewNotZeroValidator() *NotZeroValidator {
	return &NotZeroValidator{name: "NotZero"}
}

// Generate generates validation code
func (s *NotZeroValidator) Generate(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) != 0 {
		return "", errors.New("NotZero takes no parameters")
	}

	field := fieldStruct.Type

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf(`
			if s.%[1]s == 0 {
				errors%[1]s = append(errors%[1]s, errors.New("is Zero"))
			}`, fieldStruct.Name), nil
	case reflect.Ptr:
		field = field.Elem()
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
			reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
			return fmt.Sprintf(`
				if s.%[1]s != nil && *s.%[1]s == 0 {
					errors%[1]s = append(errors%[1]s, errors.New("is Zero"))
				}`, fieldStruct.Name), nil
		default:
			return "", fmt.Errorf("NotZero only works on number and number pointer type Fields. %s has type '%s'", fieldStruct.Name, field.Kind())
		}
	default:
		return "", fmt.Errorf("NotZero only works on number and number pointer type Fields. %s has type '%s'", fieldStruct.Name, field.Kind())
	}
}

// Name provides access to the name field
func (s *NotZeroValidator) Name() string {
	return s.name
}
