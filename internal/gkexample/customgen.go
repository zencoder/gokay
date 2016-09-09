package gkexample

import (
	"fmt"

	"github.com/zencoder/gokay/gkgen"
)

// To run: `gokay gkexample NewCustomValidator`
func NewCustomValidator() *gkgen.ValidateGenerator {
	fmt.Println("Generating code with a custom validator that's the same as the default validator")
	return gkgen.NewValidator()
}
