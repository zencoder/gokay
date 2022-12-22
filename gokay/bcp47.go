package gokay

import "golang.org/x/text/language"

func IsBCP47(s *string) error {
	if s == nil || *s == "" {
		return nil
	}

	_, err := language.Parse(*s)

	// Pass tags that are well-formed, but not in the spec
	if _, ok := err.(language.ValueError); ok {
		return nil
	}

	return err
}
