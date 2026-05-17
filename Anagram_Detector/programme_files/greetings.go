package main

import "fmt"

func greetings() {
	fmt.Println(cyan + "=====================================" + reset)
	fmt.Println(bold + blue + "        ANAGRAM DETECTOR CLI" + reset)
	fmt.Println(cyan + "=====================================" + reset)

	fmt.Println()
	fmt.Println(green + "Welcome 😊!" + reset)
	fmt.Println()

	fmt.Println(yellow + "This CLI tool checks whether two words" + reset)
	fmt.Println(yellow + "or sentences are anagrams." + reset)

	fmt.Println()
	fmt.Println("Two strings are anagrams if they contain")
	fmt.Println("the same characters in different orders.")

	fmt.Println()
	fmt.Println(bold + "Examples:" + reset)
	fmt.Println(green + `- "listen"  -> "silent"` + reset)
	fmt.Println(green + `- "evil"    -> "vile"` + reset)
	fmt.Println(green + `- "night"   -> "thing"` + reset)

	fmt.Println()
	fmt.Println(red + "ENTER 0 TO EXIT" + reset)
	fmt.Println()
}
