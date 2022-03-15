package gkgen

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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
// 	Supports the "except" parameter with comma-separated values to allow language exceptions
// 	Example Field: "FieldName string `valid:BCP47=(except=en-AB,en-WL)`"
func (s *BCP47Validator) Generate(_ reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) > 1 {
		return "", errors.New("BCP47 takes up to 1 parameter")
	}

	field := fieldStruct.Type

	exceptions := []string{}
	for _, param := range params {
		if strings.HasPrefix(param, "except=") {
			exceptions = strings.Split(strings.TrimPrefix(param, "except="), ",")
		}
	}

	exceptionSetOut := fmt.Sprintf("bcp47Exceptions%s := map[string]bool {\n", fieldStruct.Name)
	if len(exceptions) > 0 {
		for _, exception := range exceptions {
			exceptionSetOut += fmt.Sprintf("\"%s\": true,\n", exception)
		}
	}
	exceptionSetOut += "}"

	isPtr := false
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
		isPtr = true
	}

	switch field.Kind() {
	case reflect.Ptr:
		return "", errors.New("BCP47Validator does not support nested pointer fields")
	case reflect.String:
		fieldPtr := "s." + fieldStruct.Name
		fieldVal := "s." + fieldStruct.Name
		if isPtr {
			fieldVal = "*" + fieldVal
		} else {
			fieldPtr = "&" + fieldPtr
		}
		return fmt.Sprintf(`
			%[1]s
			if err := gokay.IsBCP47(%[2]s); err != nil && !bcp47Exceptions%[4]s[%[3]s] {
				errors%[4]s = append(errors%[4]s, err)
			}
			`, exceptionSetOut, fieldPtr, fieldVal, fieldStruct.Name), nil
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
