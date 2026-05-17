package main

import "strings"

func isAnagram(first, second string) (bool, string) {
	if len(first) != len(second) {
		return false, "Not an Anagram❌"
	}

	first = strings.ToLower(first)
	second = strings.ToLower(second)

	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	for _, ch := range first {
		freq1[ch]++
	}
	for _, ch := range second {
		freq2[ch]++
	}

	for key, value := range freq1 {
		if freq2[key] != value {
			return false, "Not an Anagram❌"
		}
	}

	return true, "Anagram✔️"
}
