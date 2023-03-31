package gokay

import (
	"bytes"
	"fmt"
)

// ErrorMap contains an entry for each invalid field in a struct. Values can be
// any struct that implements the go Error interface, including nested ErrorMaps.
type ErrorMap map[string]error

// Error returns a JSON formatted representation of the ErrorMap.
func (em ErrorMap) Error() string {
	out := &bytes.Buffer{}
	out.WriteRune('{')
	for k, v := range em {
		if out.Len() > 1 {
			out.WriteRune(',')
		}
		if v != nil {
			switch v.(type) {
			case ErrorSlice, ErrorMap:
				fmt.Fprintf(out, `%q: %s`, k, v.Error())
			default:
				fmt.Fprintf(out, `%q: %q`, k, v.Error())
			}
		} else {
			fmt.Fprintf(out, `%q: null`, k)
		}
	}
	out.WriteRune('}')
	return out.String()
}

// ErrorSlice is a slice of errors. Typically an ErrorSlice will be an entry
// in the ErrorMap outputted by a generated Validate function, each element
// of the array represents a failed validation on that field.
type ErrorSlice []error

// Returns a JSON formatted representation of the ErrorSlice
func (ea ErrorSlice) Error() string {
	out := &bytes.Buffer{}
	out.WriteRune('[')
	for i := range ea {
		if i != 0 {
			out.WriteRune(',')
		}
		if ea[i] != nil {
			switch ea[i].(type) {
			case ErrorSlice, ErrorMap:
				fmt.Fprintf(out, `%s`, ea[i].Error())
			default:
				fmt.Fprintf(out, `%q`, ea[i].Error())
			}
		} else {
			out.WriteString("null")
		}
	}
	out.WriteRune(']')
	return out.String()
}
