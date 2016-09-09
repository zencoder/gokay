package gkgen_test

func Intptr(v int) *int {
	p := new(int)
	*p = v
	return p
}

func Float64ptr(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

func Boolptr(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}
