package main

import (
	"fmt"
	"math"
)

func main() {
	sqrtNeg := math.Sqrt(-1)
	fmt.Printf("math.Sqrt(-1) = %v\n", sqrtNeg) // NaN

	if math.IsNaN(sqrtNeg) {
		fmt.Println("sqrtNeg is NaN (Not-a-Number)!")
	}

	inf := math.Inf(1) // +Infinity
	infMinusInf := inf - inf
	fmt.Printf("+Infinity - +Infinity = %v\n", infMinusInf) // NaN

	if math.IsNaN(infMinusInf) {
		fmt.Println("infMinusInf is NaN (Not-a-Number)!")
	}

	logNeg := math.Log(-10)                    // Negatif bir sayının logaritması
	fmt.Printf("math.Log(-10) = %v\n", logNeg) // NaN

	if math.IsNaN(logNeg) {
		fmt.Println("logNeg is NaN (Not-a-Number)!")
	}

	logZero := math.Log(0)
	fmt.Printf("math.Log(0) = %v\n", logZero)

	if math.IsInf(logZero, 0) {
		fmt.Println("logZero is negative Infinity!")
	}
}
