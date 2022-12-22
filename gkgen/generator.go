package gkgen

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
)

// Generater defines the behavior of types that generate validation code
type Generater interface {
	Generate(reflect.Type, reflect.StructField, []string) (string, error)
	Name() string
}

// ValidateGenerator holds a map of identifiers and Generator's
type ValidateGenerator struct {
	Generators map[string]Generater
}

// NewValidateGenerator creates a new pointer value of type ValidateGenerator
// - Hex: checks if a string is a valid hexadecimal format number
// - Length: takes 1 integer argument and compares the length of a string field against that
// - NotNil: Validate fails if field is nil
// - Set: Validate fails if field is not in a specific set of values
// - NotEqual: Validate fails if field is equal to specific value
// - UUID: Checks and fails if a string is not a valid UUID
func NewValidateGenerator() *ValidateGenerator {
	v := &ValidateGenerator{make(map[string]Generater)}
	v.AddValidation(NewNotNilValidator())
	v.AddValidation(NewSetValidator())
	v.AddValidation(NewNotEqualValidator())
	v.AddValidation(NewLengthValidator())
	v.AddValidation(NewHexValidator())
	v.AddValidation(NewUUIDValidator())
	v.AddValidation(NewBCP47Validator())
	v.AddValidation(NewMinLengthValidator())
	return v
}

// AddValidation adds a Validation to a ValidateGenerator, that Validation can be applied to a struct
// field using the string returned by
// validator.Name()
func (s *ValidateGenerator) AddValidation(g Generater) error {
	s.Generators[g.Name()] = g
	return nil
}

// Generate generates Validate method for a structure
// Implicitly generates code that validates Structs, Slices and Maps which can be nested. Null pointer fields are considered valid by default
// Return value of generated function is an ErrorMap of ErrorSlices, where each element of an ErrorSlice represents a failed validation
func (s *ValidateGenerator) Generate(out io.Writer, i interface{}) error {
	structValue := reflect.ValueOf(i)
	structType := reflect.TypeOf(i)

	hasValidation := false
	svBuf := &bytes.Buffer{}

	fmt.Fprintf(svBuf, "func(s %s) Validate() error {\n", structType.Name())
	fmt.Fprint(svBuf, "\tem := make(gokay.ErrorMap)\n\n")

	// Loop through fields
	for j := 0; j < structValue.NumField(); j++ {
		fieldIsValidated := false

		field := structType.Field(j)
		tag := field.Tag.Get("valid") // Everything under the 'valid' tag

		fvBuf := &bytes.Buffer{}

		fmt.Fprintf(fvBuf, `
	// BEGIN %[1]s field Validations
	errors%[1]s := make(gokay.ErrorSlice, 0, 0)
	`, field.Name)
		if tag != "" && tag != "-" {
			field := structType.Field(j)

			vcs, err := ParseTag(i, tag)
			if err != nil {
				return fmt.Errorf("Unable to parse tag: '%v'. Error: '%v'", tag, err)
			}

			for _, vc := range vcs {
				if _, ok := s.Generators[vc.name]; !ok {
					return fmt.Errorf("Unknown validation generator name: '%s'", vc.name)
				}
				code, err := s.Generators[vc.name].Generate(structType, field, vc.Params)
				fmt.Fprintf(fvBuf, "// %s", vc.name)
				fmt.Fprintln(fvBuf, code)
				if err != nil {
					return err
				}
			}
			fieldIsValidated = true
		}

		isPtr := field.Type.Kind() == reflect.Ptr

		var fieldType reflect.Type

		if isPtr {
			fieldType = field.Type.Elem()
		} else {
			fieldType = field.Type
		}

		recursiveValidate := false
		switch fieldType.Kind() {
		case reflect.Struct:
			recursiveValidate = true
			fieldIsValidated = true
		case reflect.Ptr:
			return errors.New("Nested pointers are not currently supported by gokay")
		}

		if isPtr && recursiveValidate {
			fmt.Fprintf(fvBuf, `
			if s.%s != nil {
			`, field.Name)
		}

		switch fieldType.Kind() {
		case reflect.Struct:
			fmt.Fprintf(fvBuf, `if err := gokay.Validate(s.%[1]s); err != nil {
				errors%[1]s = append(errors%[1]s, err)
			}
			`, field.Name)
		case reflect.Map:
			// TODO: Support non-string keys
			mapBuf := &bytes.Buffer{}
			err := generateMapValidationCode(mapBuf, fieldType, field.Name, 0)
			if err != nil {
				log.Printf("WARNING: Cannot generate recursive validation of map %q, %s", field.Name, err.Error())
			} else {
				io.Copy(fvBuf, mapBuf)
				recursiveValidate = true
				fieldIsValidated = true
			}
		case reflect.Slice, reflect.Array:
			slBuf := &bytes.Buffer{}
			err := generateSliceValidationCode(slBuf, fieldType, field.Name, 0)
			if err != nil {
				log.Printf("WARNING: Cannot generate recursive validation of slice %q, %s", field.Name, err.Error())
			} else {
				io.Copy(fvBuf, slBuf)
				recursiveValidate = true
				fieldIsValidated = true
			}
		}

		if isPtr && recursiveValidate {
			fmt.Fprintln(fvBuf, "}")
		}

		fmt.Fprintf(fvBuf, `
		if len(errors%[1]s) > 0 {
			em["%[1]s"] = errors%[1]s
		}
		`, field.Name)

		fmt.Fprintf(fvBuf, "// END %s field Validations\n", field.Name)

		if fieldIsValidated {
			svBuf.Write(fvBuf.Bytes())
			hasValidation = true
		}
	}
	fmt.Fprintln(svBuf, `
	if len(em) > 0 {
		return em
	} else {
		return nil
	}
	`)

	fmt.Fprintln(svBuf, "}")

	if hasValidation {
		out.Write(svBuf.Bytes())
	}

	return nil
}

// generateMapValidationCode generates validation code used with maps
func generateMapValidationCode(out io.Writer, fieldType reflect.Type, fieldName string, depth int64) error {
	if fieldType.Kind() != reflect.Map {
		return fmt.Errorf("Cannot call `generateMapValidationCode` on non-map type '%v'", fieldType)
	}

	if fieldType.Key().Kind() != reflect.String {
		return fmt.Errorf("Unsupported map Key type at depth %d of Field %s. Key Type: '%v'.", depth, fieldName, fieldType.Key().Kind())
	}

	if depth == 0 {
		fmt.Fprintf(out, "em%s := make(gokay.ErrorMap)\n", fieldName)
		fmt.Fprintf(out, "for k%d, v%d := range s.%s {\n", depth, depth, fieldName)
	} else {
		fmt.Fprintf(out, "for k%d, v%d := range v%d {\n", depth, depth, depth-1)
	}

	isPtr := fieldType.Elem().Kind() == reflect.Ptr
	if isPtr {
		fieldType = fieldType.Elem().Elem()
		fmt.Fprintf(out, "if v%d != nil {\n", depth)
	} else {
		fieldType = fieldType.Elem()
	}

	if fieldType.Kind() == reflect.Ptr {
		return fmt.Errorf("Recursive validation of nested pointers not yet supported. Field '%s' at depth '%d'", fieldName, depth)
	}

	switch fieldType.Kind() {
	case reflect.Map:
		fmt.Fprintf(out, "emv%d := make(gokay.ErrorMap)\n", depth)
		err := generateMapValidationCode(out, fieldType, fieldName, depth+1)
		if err != nil {
			return err
		}
	case reflect.Slice, reflect.Array:
		fmt.Fprintf(out, "emv%d := make(gokay.ErrorMap)\n", depth)
		err := generateSliceValidationCode(out, fieldType, fieldName, depth+1)
		if err != nil {
			return err
		}
	case reflect.Struct, reflect.Interface:
		if depth > 0 {
			fmt.Fprintf(out,
				`if err := gokay.Validate(v%d); err != nil {
						emv%d[fmt.Sprintf("%%v", k%d)] = err
			}
			`, depth, depth-1, depth)
		} else {
			fmt.Fprintf(out,
				`if err := gokay.Validate(v%d); err != nil {
						em%s[fmt.Sprintf("%%v", k%d)] = err
			}
			`, depth, fieldName, depth)
		}
	default:
		return fmt.Errorf("Unsupported map Value type at depth %d. Value Type: '%s'", depth, fieldType.Kind())
	}

	if isPtr {
		fmt.Fprint(out, "}\n")
	}

	fmt.Fprint(out, "}\n")

	if depth == 0 {
		fmt.Fprintf(out, `
			if len(em%[1]s) > 0 {
				errors%[1]s = append(errors%[1]s, em%[1]s)
			}
			`, fieldName)
	} else if depth == 1 {
		fmt.Fprintf(out, `
			if len(emv%d) > 0 {
				em%s[fmt.Sprintf("%%v", k%d)] = emv%d
			}
			`, depth-1, fieldName, depth-1, depth-1)
	} else {
		fmt.Fprintf(out, `
			if len(emv%d) > 0 {
				emv%d[fmt.Sprintf("%%v", k%d)] = emv%d
			}
			`, depth-1, depth-2, depth-1, depth-1)
	}
	return nil
}

// generateSliceValidationCode generates validation code used with slice
func generateSliceValidationCode(out io.Writer, fieldType reflect.Type, fieldName string, depth int64) error {
	if fieldType.Kind() != reflect.Slice {
		return fmt.Errorf("`generateSliceValidationCode` only supports slices and arrays. Not: '%s'", fieldType.Kind())
	}

	isPtr := fieldType.Elem().Kind() == reflect.Ptr
	// Slice of Ptrs
	if isPtr {
		fieldType = fieldType.Elem().Elem()
	} else {
		fieldType = fieldType.Elem()
	}

	if depth == 0 {
		fmt.Fprintf(out, "em%s := make(gokay.ErrorMap)\n", fieldName)
		fmt.Fprintf(out, "for k%d, v%d := range s.%s {\n", depth, depth, fieldName)
	} else {
		fmt.Fprintf(out, "for k%d, v%d := range v%d {\n", depth, depth, depth-1)
	}

	if isPtr {
		fmt.Fprintf(out, "if v%d != nil {\n", depth)
	}

	switch fieldType.Kind() {
	case reflect.Slice, reflect.Array:
		fmt.Fprintf(out, "emv%d := make(gokay.ErrorMap)\n", depth)
		err := generateSliceValidationCode(out, fieldType, fieldName, depth+1)
		if err != nil {
			return err
		}

	case reflect.Map:
		fmt.Fprintf(out, "emv%d := make(gokay.ErrorMap)\n", depth)
		err := generateMapValidationCode(out, fieldType, fieldName, depth+1)
		if err != nil {
			return err
		}

	case reflect.Struct:
		if depth > 0 {
			fmt.Fprintf(out,
				`if err := gokay.Validate(v%d); err != nil {
							emv%d[fmt.Sprintf("%%v", k%d)] = err
				}
				`, depth, depth-1, depth)
		} else {
			fmt.Fprintf(out,
				`if err := gokay.Validate(v%d); err != nil {
							em%s[fmt.Sprintf("%%v", k%d)] = err
				}
				`, depth, fieldName, depth)
		}
	default:
		return fmt.Errorf("`generateSliceValidationCode` cannot generate code to recursively validate slices of (pointers to) '%s'", fieldType.Kind())
	}

	if isPtr {
		fmt.Fprint(out, "}\n")
	}

	fmt.Fprint(out, "}\n")

	if depth == 0 {
		fmt.Fprintf(out, `
			if len(em%[1]s) > 0 {
				errors%[1]s = append(errors%[1]s, em%[1]s)
			}
			`, fieldName)
	} else if depth == 1 {
		fmt.Fprintf(out, `
			if len(emv%d) > 0 {
				em%s[fmt.Sprintf("%%v", k%d)] = emv%d
			}
			`, depth-1, fieldName, depth-1, depth-1)
	} else {
		fmt.Fprintf(out, `
			if len(emv%d) > 0 {
				emv%d[fmt.Sprintf("%%v", k%d)] = emv%d
			}
			`, depth-1, depth-2, depth-1, depth-1)
	}

	return nil
}
