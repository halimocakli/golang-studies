package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Print("Input a float value (e.g., 3.14): ")
	var number float64
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	fmt.Printf("\nOriginal: %.2f\n", number)
	fmt.Printf("Rounded: %.2f\n", math.Round(number))
	fmt.Printf("Floored: %.2f\n", math.Floor(number))
	fmt.Printf("Ceiled: %.2f\n", math.Ceil(number))
}
