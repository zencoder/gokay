package gkgen

import (
	"errors"
	"fmt"
	"reflect"
)

// HexValidator generates code that will verify if a field is a hex string
// 0x prefix is optional
type HexValidator struct {
	name string
}

// NewHexValidator
func NewHexValidator() *HexValidator {
	return &HexValidator{name: "Hex"}
}

// Generate
func (s *HexValidator) Generate(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
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
		return "", errors.New("HexValidator does not support nested pointer fields")
	case reflect.String:
		if isPtr {
			return fmt.Sprintf(`
				if err := gokay.IsHex(s.%[1]s); err != nil {
					errors%[1]s = append(errors%[1]s, err)
				}
				`, fieldStruct.Name), nil
		}
		return fmt.Sprintf(`
			if err := gokay.IsHex(&s.%[1]s); err != nil {
				errors%[1]s = append(errors%[1]s, err)
			}
			`, fieldStruct.Name), nil
	default:
		if isPtr {
			return "", fmt.Errorf("HexValidator does not support fields of type: '*%v'", field.Kind())
		}
		return "", fmt.Errorf("HexValidator does not support fields of type: '%v'", field.Kind())
	}
}

// Name provides access to the name field
func (s *HexValidator) Name() string {
	return s.name
}
