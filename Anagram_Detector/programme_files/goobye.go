package main

import "fmt"

func goodbye() {
	fmt.Println(cyan + "=====================================" + reset)
	fmt.Println(bold + blue + "    Thank you for using the" + reset)
	fmt.Println(bold + blue + "       Anagram Detector CLI" + reset)
	fmt.Println(cyan + "=====================================" + reset)

	fmt.Println()
	fmt.Println(green + "Session terminated successfully 😊" + reset)
	fmt.Println(yellow + "Goodbye buddy!" + reset)
	fmt.Println()
}
