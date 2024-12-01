package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int

	fmt.Printf("Insert two numbers: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Fatalf("Error scanning input: %v", err)
	}

	sum, sub, mul, div, err := calculate(a, b)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Sum: %d, Subtraction: %d, Multiplication: %d, Division: %.2f\n", sum, sub, mul, div)
}

func calculate(a, b int) (sum, sub, mul int, div float64, err error) {
	sum = a + b
	sub = a - b
	mul = a * b

	if b == 0 {
		err = fmt.Errorf("division by zero is undefined")
	} else {
		div = float64(a) / float64(b)
	}
	return
}
