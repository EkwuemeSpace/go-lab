package main

import "fmt"

func goodbye() {
	fmt.Println(cyan + "=====================================" + reset)
	fmt.Println(bold + blue + "   Thank you for using the" + reset)
	fmt.Println(bold + blue + "      Palindrome Checker" + reset)
	fmt.Println(cyan + "=====================================" + reset)

	fmt.Println()
	fmt.Println(green + "Session terminated successfully 😍" + reset)
	fmt.Println(yellow + "Goodbye and happy coding!" + reset)
	fmt.Println()
}
