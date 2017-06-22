package gokay

import (
	"fmt"
	"regexp"
)

var (
	uuidRegexp = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")
)

// IsUUID validates that the given string is a UUID value
func IsUUID(s *string) error {
	if s == nil {
		return nil
	}

	if !uuidRegexp.MatchString(*s) {
		return fmt.Errorf("'%s' is not a UUID", *s)
	}

	return nil
}
