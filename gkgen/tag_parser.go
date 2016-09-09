package gkgen

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

type ValidationCommand struct {
	Name   string
	Params []string
}

func ParseTag(interf interface{}, tag string) ([]ValidationCommand, error) {
	inr := strings.NewReader(tag)
	out := new(bytes.Buffer)
	vcs := make([]ValidationCommand, 0, 8)

	ch, _, readerErr := inr.ReadRune()
	for readerErr == nil {
		var name string
		switch ch {
		case '=':
			if len(out.Bytes()) == 0 {
				return nil, errors.New("Error in 'valid' tag syntax")
			}
			name = out.String()
			out.Reset()
			params, err := parseParams(interf, inr)
			if err != nil {
				return nil, err
			}
			vc := ValidationCommand{
				Name:   name,
				Params: params}
			vcs = append(vcs, vc)
		case ',':
			if len(out.Bytes()) == 0 {
				return nil, errors.New("Error in 'valid' tag syntax")
			}
			name = out.String()
			vc := ValidationCommand{
				Name: name}
			vcs = append(vcs, vc)
			out.Reset()
		default:
			out.WriteRune(ch)
		}

		ch, _, readerErr = inr.ReadRune()
		if readerErr == io.EOF && len(out.Bytes()) > 0 {
			name = out.String()
			vc := ValidationCommand{
				Name: name}
			vcs = append(vcs, vc)
		}
	}

	if readerErr == io.EOF {
		return vcs, nil
	} else {
		return nil, readerErr
	}
}

func parseParams(interf interface{}, inr *strings.Reader) ([]string, error) {
	ch, _, readerErr := inr.ReadRune()
	params := make([]string, 0, 8)
	if readerErr != nil {
		return nil, readerErr
	}

	for readerErr == nil {
		switch ch {
		case '(':
			param, err := parseParam(inr)
			if err != nil {
				return nil, err
			}
			params = append(params, param)
			ch, _, readerErr = inr.ReadRune()

		case ',':
			if len(params) == 0 {
				return nil, errors.New("Error in 'valid' tag syntax")
			}
			return params, nil
		default:
			if len(params) == 0 {
				return nil, errors.New("Error in 'valid' tag syntax")
			}
			ch, _, readerErr = inr.ReadRune()
		}
	}
	return params, nil
}

func parseParam(inr *strings.Reader) (string, error) {
	ch, _, readerErr := inr.ReadRune()
	out := new(bytes.Buffer)

	if readerErr != nil {
		return "", readerErr
	}

	for readerErr == nil {
		switch ch {
		case ')':
			return out.String(), nil
		case '\\':
			ch, _, readerErr = inr.ReadRune()
			if readerErr != nil {
				return "", readerErr
			}
			out.WriteRune(ch)
		default:
			out.WriteRune(ch)
		}
		ch, _, readerErr = inr.ReadRune()
	}

	// Reader should reach ')' before EOF
	if readerErr == io.EOF {
		return "", errors.New("Error in 'valid' tag syntax")
	} else {
		return "", readerErr
	}
}
