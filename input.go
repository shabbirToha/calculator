package main

import (
	"errors"
	"fmt"
)

// getUserInput prompts the user for two numbers and an operator.
func getUserInput() (float64, float64, string, error) {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err1 := fmt.Scanln(&num1)
	if err1 != nil {
		return 0, 0, "", errors.New("invalid input for first number")
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	_, err2 := fmt.Scanln(&operator)
	if err2 != nil {
		return 0, 0, "", errors.New("invalid input for operator")
	}

	fmt.Print("Enter second number: ")
	_, err3 := fmt.Scanln(&num2)
	if err3 != nil {
		return 0, 0, "", errors.New("invalid input for second number")
	}

	return num1, num2, operator, nil
}
