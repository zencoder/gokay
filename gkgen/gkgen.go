// Package gkgen generates validation code
package gkgen

import "reflect"

// Validater
type Validater interface {
	GenerateValidationCode(reflect.Type, reflect.StructField, []string) (string, error)
	Name() string
}

// Generater defines the behavior of types that generate validation code
type Generater interface {
	Generate(reflect.Type, reflect.StructField, []string) (string, error)
	Name() string
}
