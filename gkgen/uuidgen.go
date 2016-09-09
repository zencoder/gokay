package gkgen

import (
	"errors"
	"fmt"
	"reflect"
)

// UUIDValidator generates code that will verify if a field is a UUID string
type UUIDValidator struct {
	Name string
}

func NewUUIDValidator() *UUIDValidator {
	return &UUIDValidator{Name: "UUID"}
}

func (s *UUIDValidator) GenerateValidationCode(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) != 0 {
		return "", errors.New("Hex takes no parameters")
	}

	field := fieldStruct.Type

	isPtr := false
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
		isPtr = true
	}

	switch field.Kind() {
	case reflect.Ptr:
		return "", errors.New("UUIDValidator does not support nested pointer fields.")
	case reflect.String:
		if isPtr {
			return fmt.Sprintf(`
				if err := gokay.IsUUID(s.%[1]s); err != nil {
					errors%[1]s = append(errors%[1]s, err)
				}
				`, fieldStruct.Name), nil
		} else {
			return fmt.Sprintf(`
				if err := gokay.IsUUID(&s.%[1]s); err != nil {
					errors%[1]s = append(errors%[1]s, err)
				}
				`, fieldStruct.Name), nil
		}
	default:
		if isPtr {
			return "", fmt.Errorf("UUIDValidator does not support fields of type: '*%v'", field.Kind())
		} else {
			return "", fmt.Errorf("UUIDValidator does not support fields of type: '%v'", field.Kind())
		}
	}
}

func (s *UUIDValidator) GetName() string {
	return s.Name
}
