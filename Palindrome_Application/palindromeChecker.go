package main

func isPalindrome(input, reversed string) (bool, string) {
	if input == reversed {
		return true, "Palindrome"
	}
	return false, "not a palindrome"
}
