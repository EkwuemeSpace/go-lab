package main

func validatorInput(first, second string) error {
	if err := validator(first); err != nil {
		return err
	}
	if err := validator(second); err != nil {
		return err
	}
	return nil
}
