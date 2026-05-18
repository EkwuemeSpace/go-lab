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
