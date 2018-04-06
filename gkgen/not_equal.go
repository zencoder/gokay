package gkgen

import (
	"errors"
	"fmt"
	"reflect"
)

// NotEqualValidator generates code that will verify a fields does not equal a set value
// The validator will look at the field or the dereferenced value of the field
// nil values for a field are not considered invalid
type NotEqualValidator struct {
	name string
}

// NewNotEqualValidator holds the NotEqualValidator state
func NewNotEqualValidator() *NotEqualValidator {
	return &NotEqualValidator{name: "NotEqual"}
}

// Generate generates validation code
func (s *NotEqualValidator) Generate(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) != 1 {
		return "", errors.New("NotEqual validation requires exactly 1 parameter")
	}

	restrictedValue := params[0]
	field := fieldStruct.Type

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf(`
			if s.%[1]s == %[2]s {
				errors%[1]s = append(errors%[1]s, errors.New("%[1]s cannot equal '%[2]s'"))
			}`, fieldStruct.Name, restrictedValue), nil
	case reflect.String:
		return fmt.Sprintf(`
			if s.%[1]s == "%[2]s" {
					errors%[1]s = append(errors%[1]s, errors.New("%[1]s cannot equal '%[2]s'"))
			}`, fieldStruct.Name, restrictedValue), nil
	case reflect.Ptr:
		field = field.Elem()
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
			reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
			return fmt.Sprintf(`
				if s.%[1]s != nil && *s.%[1]s == %[2]s {
					errors%[1]s = append(errors%[1]s, errors.New("%[1]s cannot equal '%[2]s'"))
				}`, fieldStruct.Name, restrictedValue), nil
		case reflect.String:
			return fmt.Sprintf(`
				if s.%[1]s != nil && *s.%[1]s == "%[2]s" {
						errors%[1]s = append(errors%[1]s, errors.New("%[1]s cannot equal '%[2]s'"))
				}`, fieldStruct.Name, restrictedValue), nil
		default:
			return "", fmt.Errorf("NotEqual does not work on type '%s'", field.Kind())
		}
	default:
		return "", fmt.Errorf("NotEqual does not work on type '%s'", field.Kind())
	}
}

// Name provides access to the name field
func (s *NotEqualValidator) Name() string {
	return s.name
}
