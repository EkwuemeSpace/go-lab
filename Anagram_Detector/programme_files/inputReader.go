package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() (string, string, bool, error) {
	fmt.Print("Please enter a word:")
	reader := bufio.NewReader(os.Stdin)

	firstInput, err := reader.ReadString('\n')
	if err != nil {
		return "", "", false, err
	}

	firstInput = strings.TrimSpace(firstInput)
	if firstInput == "0" {
		return "", "", true, nil
	}

	fmt.Print("Please enter a second word:")
	secondInput, err := reader.ReadString('\n')
	if err != nil {
		return "", "", false, err
	}

	secondInput = strings.TrimSpace(secondInput)
	if secondInput == "0" {
		return "", "", true, nil
	}

	return firstInput, secondInput, false, nil
}
