package main

import "errors"

// calculate selects the correct operation based on the operator provided.
func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return add(num1, num2), nil
	case "-":
		return subtract(num1, num2), nil
	case "*":
		return multiply(num1, num2), nil
	case "/":
		return divide(num1, num2)
	default:
		return 0, errors.New("invalid operator")
	}
}
