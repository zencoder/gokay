// Package gkgen generates validation code
package gkgen

import "reflect"

// Generater defines the behavior of types that generate validation code
type Generater interface {
	Generate(reflect.Type, reflect.StructField, []string) (string, error)
	Name() string
}
