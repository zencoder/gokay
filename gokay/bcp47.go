package gokay

import "golang.org/x/text/language"

func IsBCP47(s *string) error {
	if s == nil {
		return nil
	}

	_, err := language.Parse(*s)

	return err
}
