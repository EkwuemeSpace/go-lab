package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	greetings()
	for {
		fmt.Print("Please enter a word or sentence: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input)

		if input == "0" {
			fmt.Println()
			goodbye()
			os.Exit(0)
		}

		err = validator(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		cleaner := cleanInput(input)
		reverse := reverseInput(cleaner)

		if isPalindrome(cleaner, reverse) {
			fmt.Println("Palindrome ✔️")
		} else {
			fmt.Println("Not a palindrome❌")
		}
	}
}
