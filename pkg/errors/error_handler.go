package errors

import "fmt"

type CustomError struct {
	Action string
	Err    error
}

func UpdateErrorText(action string, err error) error {
	return fmt.Errorf("%s: %w", action, err)
}

func New(action string, err error) *CustomError {
	return &CustomError{
		Action: action,
		Err:    err,
	}
}
