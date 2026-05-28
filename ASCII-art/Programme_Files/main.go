package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if !argsValidator(args) {
		fmt.Fprintln(os.Stderr, "usage: go run . [STRING] [BANNER]")
		os.Exit(1)
	}

	input := args[0]
	banner := "standard"
	if len(args) == 2 {
		banner = strings.ToLower(args[1])
	}

	validateBanner := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !validateBanner[banner] {
		fmt.Fprintf(os.Stderr, "invalid banner: %q expected[standard, shadow, thinkertoy]\n", banner)
		os.Exit(1)
	}

	path := "banners/" + banner + ".txt"
	data, err := readBannerFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	art, err := parseBannerFile(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = renderString(os.Stdout, art, input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
