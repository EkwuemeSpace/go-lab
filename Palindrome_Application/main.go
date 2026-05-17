// Package main implements a palindrome checker CLI application.
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
		fmt.Print(cyan + "Please enter a word or sentence: " + reset)

		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, red+err.Error()+reset)
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
			fmt.Fprintln(os.Stderr, red+err.Error()+reset)
			continue
		}

		cleaner := cleanInput(input)
		reverse := reverseInput(cleaner)

		if isPalindrome(cleaner, reverse) {
			fmt.Println(green + "Palindrome ✔️" + reset)
		} else {
			fmt.Println(red + "Not a palindrome ❌" + reset)
		}
	}
}
