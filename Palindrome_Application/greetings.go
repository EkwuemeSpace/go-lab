package main

import "fmt"

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
	bold   = "\033[1m"
)

func greetings() {
	fmt.Println(cyan + "=====================================" + reset)
	fmt.Println(bold + blue + "      PALINDROME CHECKER TOOL" + reset)
	fmt.Println(cyan + "=====================================" + reset)

	fmt.Println()
	fmt.Println(green + "Welcome 😊!" + reset)
	fmt.Println()

	fmt.Println(yellow + "This CLI tool checks whether a word or sentence" + reset)
	fmt.Println(yellow + "is a palindrome." + reset)

	fmt.Println()
	fmt.Println("A palindrome is a word or phrase that reads the")
	fmt.Println("same forward and backward.")

	fmt.Println()
	fmt.Println(bold + "Examples:" + reset)
	fmt.Println(green + "- level" + reset)
	fmt.Println(green + "- madam" + reset)
	fmt.Println(green + "- A man a plan a canal Panama" + reset)

	fmt.Println()
	fmt.Println(red + "ENTER 0 TO EXIT" + reset)
	fmt.Println()
}
