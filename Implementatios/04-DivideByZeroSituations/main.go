package main

import (
	"fmt"
	"math"
)

func main() {
	numerators := []float64{10.0, 0.0, -5.0, 10.0, 0}   // Paylar (numerators)
	denominators := []float64{2.0, 0.0, 0.0, 0.0, 10.0} // Paydalar (denominators)

	for i := 0; i < len(numerators); i++ {
		num := numerators[i]
		denom := denominators[i]

		// Float bölme işlemi (floating-point division)
		result := num / denom

		// Özel durumları kontrol edelim
		if math.IsNaN(result) {
			fmt.Printf("Floating-point division: %.2f / %.2f = NaN (Not a Number)\n", num, denom)
		} else if math.IsInf(result, 1) {
			fmt.Printf("Floating-point division: %.2f / %.2f = +Inf (Positive Infinity)\n", num, denom)
		} else if math.IsInf(result, -1) {
			fmt.Printf("Floating-point division: %.2f / %.2f = -Inf (Negative Infinity)\n", num, denom)
		} else {
			fmt.Printf("Floating-point division: %.2f / %.2f = %.2f\n", num, denom, result)
		}
	}

	/*
		Aşağıdaki kod parçası direkt olarak "RUNTIME ERROR" çıktısı üretecektir:

				c := 10 / 0
				fmt.Println("c =", c)
	*/
}
