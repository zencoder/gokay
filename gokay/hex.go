package gokay

import (
	"fmt"
	"regexp"
)

var (
	hexRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]+$")
)

// IsHex validates that the given string is a hex value
func IsHex(s *string) error {
	if s == nil {
		return nil
	}

	if !hexRegexp.MatchString(*s) {
		return fmt.Errorf("'%s' is not a hexadecimal string", *s)
	}

	return nil
}
