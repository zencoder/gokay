package gokay

import (
	"fmt"
	"regexp"
	"strings"
)

// IsHex validates that the given string is a hex value
func IsHex(s *string) error {
	if s == nil {
		return nil
	}

	matches, err := regexp.MatchString("^(0x)?[0-9a-f]+$", strings.ToLower(*s))
	if err != nil {
		return err
	}
	if !matches {
		return fmt.Errorf("'%s' is not a hexadecimal string", *s)
	}

	return nil
}
