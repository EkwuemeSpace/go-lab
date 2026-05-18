package main

import (
	"fmt"
	"os"
)

func main() {
	greetings()

	for {
		firstWord, secondWord, shouldExit, err := readInput()

		if shouldExit {
			fmt.Println()
			goodbye()
			os.Exit(0)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, red+err.Error()+reset)
			continue
		}

		if err := validatorInput(firstWord, secondWord); err != nil {
			fmt.Fprintln(os.Stderr, red+err.Error()+reset)
			continue
		}

		Output, note := isAnagram(firstWord, secondWord)

		if Output {
			fmt.Printf(green+"%v %s\n"+reset, Output, note)
		} else {
			fmt.Printf(red+"%v %s\n"+reset, Output, note)
		}
	}
}
