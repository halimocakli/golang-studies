package main

import (
	"fmt"
	"log"
)

func main() {
	complex1 := scanComplexNumber("Enter the first complex number [like 3.4+4.5i or (3.4+4.5i)]: ")
	complex2 := scanComplexNumber("Enter the second complex number [like 3.4+4.5i or (3.4+4.5i)]: ")

	fmt.Print("\n-----------------------------------\n")

	printComplexNumber("First complex number", complex1)
	printComplexNumber("Second complex number", complex2)

	fmt.Print("-----------------------------------\n")
	printRealAndImagParts("Complex1", complex1)
	fmt.Print("-----------------------------------\n")
	printRealAndImagParts("Complex2", complex2)
	fmt.Print("-----------------------------------\n")

	performArithmeticOperations("Complex1", "Complex2", complex1, complex2)
}

func scanComplexNumber(prompt string) complex128 {
	var c complex128

	fmt.Print(prompt)

	_, err := fmt.Scan(&c)
	if err != nil {
		log.Fatalf("Wrong input: %v", err)
	}
	return c
}

func printComplexNumber(label string, c complex128) {
	fmt.Printf("%s: %f\n", label, c)
}

func printRealAndImagParts(label string, c complex128) {
	fmt.Printf("%s Real part: %f\n", label, real(c))
	fmt.Printf("%s Imaginary part: %f\n", label, imag(c))
}

func performArithmeticOperations(label1, label2 string, c1, c2 complex128) {
	fmt.Printf("%s: %f, %s: %f\n", label1, c1, label2, c2)
	fmt.Printf("Sum: %f\n", c1+c2)
	fmt.Printf("Difference: %f\n", c1-c2)
	fmt.Printf("Product: %f\n", c1*c2)

	if c2 == 0 {
		fmt.Println("Cannot divide by zero")
		return
	}

	fmt.Printf("Quotient: %f\n", c1/c2)
}
