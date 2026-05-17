package main

import "strings"

// cleanInput converts input to lowercase and removes spaces.
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
