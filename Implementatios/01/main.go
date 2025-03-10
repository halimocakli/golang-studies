package main

import "fmt"

func main() {
	var a, b int
	fmt.Printf("Insert two numbers: ")
	_, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Printf("Error scanning input: %v", err)
		return
	}

	fmt.Println(getAddSubstractMultiply(a, b))
}

func getAddSubstractMultiply(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}
