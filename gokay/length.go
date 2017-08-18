package gokay

import "fmt"

// LengthString checks if the value of a string pointer has a length of exactly 'expected'
func LengthString(expected int64, str *string) error {
	if str == nil {
		return nil // Covers the case where a value can be Nil OR has a length constraint
	}

	if expected != int64(len(*str)) {
		return fmt.Errorf("Length was '%d', needs to be '%d'", len(*str), expected)
	}
	return nil
}

// LengthSlice not yet implemented
func LengthSlice(expected int64, actual int64) error {
	return nil
}

// MinLengthString checks if the value of a string pointer has a length of at least 'expected'
func MinLengthString(expected int64, str *string) error {
	if str == nil {
		return nil // Covers the case where a value can be Nil OR has a length constraint
	}

	if expected < int64(len(*str)) {
		return fmt.Errorf("Length was '%d', needs to be at least '%d'", len(*str), expected)
	}
	return nil
}
