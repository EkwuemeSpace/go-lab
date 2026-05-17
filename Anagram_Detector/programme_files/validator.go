package main

import (
	"fmt"
	"unicode"
)

func validator(input string) error {
	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != ' ' {
			return fmt.Errorf("invalid character detected: %q", r)
		}
	}
	return nil
}
