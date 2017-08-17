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

// MinLengthSlice checks if the value of a slice pointer has a length of at least 'expected'
func MinLengthSlice(expected int64, s *[]interface{}) error {
	if s == nil {
		return nil // Covers the case where a value can be Nil OR has a length constraint
	}

	if expected < int64(len(*s)) {
		return fmt.Errorf("Length was '%d', needs to be at least '%d'", len(*s), expected)
	}
	return nil
}

// MinLengthMap checks if the value of a slice pointer has a length of at least 'expected'
func MinLengthMap(expected int64, m *map[interface{}]interface{}) error {
	if m == nil {
		return nil // Covers the case where a value can be Nil OR has a length constraint
	}

	if expected < int64(len(*m)) {
		return fmt.Errorf("Length was '%d', needs to be at least '%d'", len(*m), expected)
	}
	return nil
}
