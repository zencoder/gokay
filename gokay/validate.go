package gokay

// Validateable specifies a generic error return type instead of an
// ErrorMap return type in order to allow for handwritten Validate
// methods to work in tandem with gokay generated Validate methods.
type Validateable interface {
	Validate() error
}

// Validate calls validate on structs that implement the Validateable
// interface. If they do not, then that struct is valid.
func Validate(i interface{}) error {
	if v, ok := i.(Validateable); ok {
		return v.Validate()
	}
	return nil
}
