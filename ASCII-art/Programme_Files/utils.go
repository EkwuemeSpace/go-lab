package main

func argsValidator(args []string) bool {
	return len(args) >= 1 && len(args) <= 2
}
