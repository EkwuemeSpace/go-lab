package main

import "strings"

func cleanInput(input string) string {
	var newInput strings.Builder

	input = strings.ToLower(input)
	for _, ch := range input {
		if ch == ' ' {
			continue
		}
		newInput.WriteString(string(ch))
	}
	return newInput.String()
}
