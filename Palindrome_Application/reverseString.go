package main

import "strings"

func reverseInput(input string) string {
	var builder strings.Builder

	for i := len(input) - 1; i >= 0; i-- {
		if string(input[i]) == " " {
			continue
		}
		builder.WriteString(string(input[i]))
	}
	return builder.String()
}
