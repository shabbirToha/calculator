package main

import "fmt"

func main() {
	num1, num2, operator, inputErr := getUserInput()
	if inputErr != nil {
		fmt.Println("Error:", inputErr)
		return
	}

	result, calcErr := calculate(num1, num2, operator)
	if calcErr != nil {
		fmt.Println("Error:", calcErr)
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
