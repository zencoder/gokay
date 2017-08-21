package gkgen

import (
	"errors"
	"fmt"
	"reflect"
)

// NotNilValidator generates code that will verify if a pointer is nil
// Slice and Array support coming later.
// It will flag nil string pointers as valid, use in conjunction with NotNil validator if you don't want nil values
type NotNilValidator struct {
	name string
}

// NewNotNilValidator holds the NotNilValidator state
func NewNotNilValidator() *NotNilValidator {
	return &NotNilValidator{name: "NotNil"}
}

// Generate generates validation code
func (s *NotNilValidator) Generate(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) != 0 {
		return "", errors.New("NotNil takes no parameters")
	}

	field := fieldStruct.Type

	switch field.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Slice, reflect.Map:
		return fmt.Sprintf(`
			if s.%[1]s == nil {
				errors%[1]s = append(errors%[1]s, errors.New("is Nil"))
			}`, fieldStruct.Name), nil
	default:
		// TODO: Add support for nil slices
		return "", fmt.Errorf("NotNil only works on pointer type Fields. %s has type '%s'", fieldStruct.Name, field.Kind())
	}
}

// Name provides access to the name field
func (s *NotNilValidator) Name() string {
	return s.name
}
