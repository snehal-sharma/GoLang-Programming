package main

import (
	"errors"
	"fmt"
)

type AgeValidationError struct {
	age     int
	message string
}

func (av *AgeValidationError) Error() string {
	return fmt.Sprintf("Age validation failed for age %d with message %s", av.age, av.message)
}

func ValidateAge(age int) error {
	if age < 18 {
		return &AgeValidationError{
			age:     age,
			message: "Person is below 18"}
	}
	return nil
}

func main() {
	var ave *AgeValidationError
	err := ValidateAge(16)
	if err != nil {
		if errors.As(err, &ave) {
			fmt.Println(err)
		}
	}

}
