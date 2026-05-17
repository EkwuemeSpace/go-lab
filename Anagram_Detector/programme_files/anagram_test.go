package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		first  string
		second string
		expect bool
	}{
		{"listen", "silent", true},
		{"evil", "vile", true},
		{"night", "thing", true},
		{"hello", "world", false},
		{"ab", "abc", false},
		{"Innocent", "nnoectIn", true},
		{"aabbcc", "abcabc", true},
	}

	for _, tc := range tests {
		result, _ := isAnagram(tc.first, tc.second)

		if result != tc.expect {
			t.Errorf(
				"FAILED: (%s, %s) expected %v but got %v",
				tc.first,
				tc.second,
				tc.expect,
				result,
			)
		}
	}
}
