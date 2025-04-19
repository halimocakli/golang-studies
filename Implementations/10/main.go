package main

import (
	"fmt"
)

func main() {
	performIntegerOperations()
	performComplexOperations()
	showInvalidOperations()
}

func performIntegerOperations() {
	var (
		num1 int
		num2 int
	)

	fmt.Print("Input two integers (e.g., 3 5): ")
	if _, err := fmt.Scan(&num1, &num2); err != nil {
		fmt.Println("Invalid input. Please enter valid integers.")
		return
	}

	fmt.Printf("\nInteger Operations:\n")
	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)

	if num2 != 0 {
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	} else {
		fmt.Printf("Division by zero is not allowed.\n")
	}

	fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)
}

func performComplexOperations() {
	var complex1, complex2 complex128
	fmt.Print("\nInput two complex numbers (e.g., 3+4i 5-2i): ")
	if _, err := fmt.Scan(&complex1, &complex2); err != nil {
		fmt.Println("Invalid input. Please enter valid complex numbers.")
		return
	}
	fmt.Printf("\nComplex Operations:\n")
	fmt.Printf("%v + %v = %v\n", complex1, complex2, complex1+complex2)
	fmt.Printf("%v - %v = %v\n", complex1, complex2, complex1-complex2)
}

func showInvalidOperations() {
	fmt.Println("\nReminder: %% [Modulo Operator] cannot be used with Floats or Complex numbers.")
}
