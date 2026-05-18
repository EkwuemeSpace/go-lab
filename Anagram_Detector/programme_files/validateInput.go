package main

func validatorInput(first, second string) error {
	if err := validateChar(first); err != nil {
		return err
	}
	if err := validateChar(second); err != nil {
		return err
	}
	return nil
}

func isEmptyInput(first, second string) bool {
	if !(first == "" || second == "") {
		return true
	}
	return false
}
