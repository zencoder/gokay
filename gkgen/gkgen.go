// Package that generates
package gkgen

import "reflect"

type Validater interface {
	GenerateValidationCode(reflect.Type, reflect.StructField, []string) (string, error)
	GetName() string
}
