package main

import "errors"

// add returns the sum of two numbers.
func add(a, b float64) float64 {
	return a + b
}

// subtract returns the difference of two numbers.
func subtract(a, b float64) float64 {
	return a - b
}

// multiply returns the product of two numbers.
func multiply(a, b float64) float64 {
	return a * b
}

// divide returns the quotient of two numbers or an error if dividing by zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
