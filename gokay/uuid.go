package gokay

import (
	"fmt"
	"regexp"
	"strings"
)

// IsUUID validates that the given string is a UUID value
func IsUUID(s *string) error {
	if s == nil {
		return nil
	}

	matches, err := regexp.MatchString("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", strings.ToLower(*s))
	if err != nil {
		return err
	}
	if !matches {
		return fmt.Errorf("'%s' is not a UUID", *s)
	}

	return nil
}
