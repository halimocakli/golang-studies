package main

import (
	"fmt"  // Formatlı giriş/çıkış işlemleri için kullanılan paket (Input/Output operations)
	"math" // Matematiksel işlemler ve özel durum kontrolleri için kullanılan paket (Mathematical operations and special case handling)
)

func main() {
	// Bölme işlemleri için kullanılan pay (numerators) dizisi.
	numerators := []float64{10.0, 0.0, -5.0, 10.0, 0} // Paylar (Numerators)
	// Bölme işlemleri için kullanılan payda (denominators) dizisi.
	denominators := []float64{2.0, 0.0, 0.0, 0.0, 10.0} // Paydalar (Denominators)

	// Her bir pay ve payda çifti üzerinde bölme işlemi gerçekleştirilmek üzere for döngüsü kullanılır.
	for i := 0; i < len(numerators); i++ {
		num := numerators[i]
		denom := denominators[i]

		// Float bölme işlemi yapılır.
		result := num / denom

		// Hesaplanan sonucun özel durumlarını kontrol ediyoruz:
		// Eğer sonuç NaN (Not a Number) ise:
		if math.IsNaN(result) {
			fmt.Printf("Floating-point division: %.2f / %.2f = NaN (Not a Number)\n", num, denom)
			// Eğer sonuç pozitif sonsuzluk (+Inf) ise:
		} else if math.IsInf(result, 1) {
			fmt.Printf("Floating-point division: %.2f / %.2f = +Inf (Positive Infinity)\n", num, denom)
			// Eğer sonuç negatif sonsuzluk (-Inf) ise:
		} else if math.IsInf(result, -1) {
			fmt.Printf("Floating-point division: %.2f / %.2f = -Inf (Negative Infinity)\n", num, denom)
			// Diğer durumlarda, sonuç normal bir değer olarak kabul edilir:
		} else {
			fmt.Printf("Floating-point division: %.2f / %.2f = %.2f\n", num, denom, result)
		}
	}

	/*
		Aşağıdaki kod parçası direkt olarak "RUNTIME ERROR" (çalışma zamanı hatası) üretecektir:

				c := 10 / 0
				fmt.Println("c =", c)
	*/
}
