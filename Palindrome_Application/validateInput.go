package main

import (
	"fmt"
	"unicode"
)

func validator(input string) error {
	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != ' ' {
			return fmt.Errorf("error: you have enter an invalid character: %q", r)
		}
	}
	return nil
}
