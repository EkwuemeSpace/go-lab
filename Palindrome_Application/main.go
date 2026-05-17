package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	greetings()
initial:
	fmt.Print("Please enter a word or sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "0" {
		fmt.Println()
		goodbye()
		os.Exit(0)
	}

	err := validator(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		goto initial
	}
	cleaner := cleanInput(input)
	reverse := reverseInput(cleaner)
	fmt.Println(isPalindrome(cleaner, reverse))
	goto initial
}
