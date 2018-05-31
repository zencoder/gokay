package gkgen

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// EqualsValidator generates code that will verify a fields does equals one of an allowed set of values
// The EqualsValidator will look at the field or the dereferenced value of the field
// nil values for a field are not considered invalid
type EqualsValidator struct {
	name string
}

// NewEqualsValidator holds the EqualsValidator state
func NewEqualsValidator() *EqualsValidator {
	return &EqualsValidator{name: "Equals"}
}

// Generate generates validation code
func (s *EqualsValidator) Generate(sType reflect.Type, fieldStruct reflect.StructField, params []string) (string, error) {
	if len(params) == 0 {
		return "", errors.New("Equals validation requires at least 1 parameter")
	}

	field := fieldStruct.Type

	switch field.Kind() {
	case reflect.String:
		conditions := make([]string, len(params))
		for i, param := range params {
			conditions[i] = fmt.Sprintf(`s.%[1]s == "%[2]s"`, fieldStruct.Name, param)
		}
		condition := strings.Join(conditions, " || ")
		return fmt.Sprintf(`
			if s.%[1]s != "" && !(%[2]s) {
					errors%[1]s = append(errors%[1]s, errors.New("%[1]s must equal %[3]s"))
			}`, fieldStruct.Name, condition, strings.Join(params, " or ")), nil
	case reflect.Ptr:
		field = field.Elem()
		switch field.Kind() {
		case reflect.String:
			conditions := make([]string, len(params))
			for i, param := range params {
				conditions[i] = fmt.Sprintf(`*s.%[1]s == "%[2]s"`, fieldStruct.Name, param)
			}
			condition := strings.Join(conditions, " || ")
			return fmt.Sprintf(`
				if s.%[1]s != nil && !(%[2]s) {
						errors%[1]s = append(errors%[1]s, errors.New("%[1]s must equal %[3]s"))
				}`, fieldStruct.Name, condition, strings.Join(params, " or ")), nil
		default:
			return "", fmt.Errorf("Equals does not work on type '%s'", field.Kind())
		}
	default:
		return "", fmt.Errorf("Equals does not work on type '%s'", field.Kind())
	}
}

// Name provides access to the name field
func (s *EqualsValidator) Name() string {
	return s.name
}
