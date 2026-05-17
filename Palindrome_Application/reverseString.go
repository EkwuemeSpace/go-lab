package main

import "strings"

func reverseInput(input string) string {
	runes := []rune(input)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[i], runes[i]
	}
	return strings.ReplaceAll(string(runes), " ", "")
}
